package pricing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	pb "github.com/textileio/bidbot/gen/v1"
	golog "github.com/textileio/go-log/v2"
)

const (
	// Use cached rules if they are loaded no earlier than this period.
	cidGravityCachePeriod = time.Minute
	// If loading rules from CID gravity API takes longer than this timeout, turn it into background and use the
	// cached rules for prices lookup.
	cidGravityLoadRulesTimeout = 5 * time.Second
)

var (
	log              = golog.Logger("bidbot/pricing")
	cidGravityAPIUrl = "https://staging-api.cidgravity.com/api/integrations/bidbot"
)

// CIDGravityRules is the format CID gravity API returns. Needs to be public to unmarshal JSON payload. Do not use.
type CIDGravityRules struct {
	Blocked         bool
	MaintenanceMode bool
	PricingRules    []struct {
		Verified    bool
		MinSize     uint64 // bytes
		MaxSize     uint64 // bytes
		MinDuration uint64 // epoches
		MaxDuration uint64 // epoches
		Price       int64  // attoFil
	}
}

type cidGravityRules struct {
	apiKey      string
	clientRules map[string]*clientRules
	lkRules     sync.Mutex
}

// NewCIDGravityRules returns PricingRules based on CID gravity configuration for the storage provider.
func NewCIDGravityRules(apiKey string) PricingRules {
	return &cidGravityRules{apiKey: apiKey, clientRules: make(map[string]*clientRules)}
}

// PricesFor looks up prices for the auction based on its client address.
func (cg *cidGravityRules) PricesFor(auction *pb.Auction) (prices ResolvedPrices, valid bool) {
	cg.lkRules.Lock()
	rules, exists := cg.clientRules[auction.ClientAddress]
	if !exists {
		rules = newClientRulesFor(cg.apiKey, auction.ClientAddress)
	}
	cg.clientRules[auction.ClientAddress] = rules
	cg.lkRules.Unlock()
	return rules.PricesFor(auction)
}

type clientRules struct {
	apiKey           string
	clientAddress    string
	rules            atomic.Value // *CIDGravityRules
	rulesLastUpdated atomic.Value // time.Time
	rulesETag        atomic.Value // string
	lkLoadRules      sync.Mutex
}

func newClientRulesFor(apiKey, clientAddress string) *clientRules {
	rules := &clientRules{
		apiKey:        apiKey,
		clientAddress: clientAddress,
	}
	rules.rulesLastUpdated.Store(time.Time{})
	return rules
}

// PricesFor checks the CID gravity rules for one specific client address and returns the resolved prices for the
// auction. The rules are cached locally for some time. It returns valid = false if the rules were never loaded from the
// CID gravity API.
func (cg *clientRules) PricesFor(auction *pb.Auction) (prices ResolvedPrices, valid bool) {
	cg.maybeReloadRules(cidGravityAPIUrl, cidGravityLoadRulesTimeout, cidGravityCachePeriod)
	rules := cg.rules.Load()
	if rules == nil {
		return
	}
	valid = true
	if rules.(*CIDGravityRules).Blocked {
		return
	}
	if rules.(*CIDGravityRules).MaintenanceMode {
		return
	}
	// rules are checked in sequence and the first match wins.
	for _, r := range rules.(*CIDGravityRules).PricingRules {
		if auction.DealSize >= r.MinSize && auction.DealSize < r.MaxSize &&
			auction.DealDuration >= r.MinDuration && auction.DealDuration < r.MaxDuration {
			if r.Verified && !prices.VerifiedPriceValid {
				prices.VerifiedPriceValid, prices.VerifiedPrice = true, r.Price
			} else if !prices.UnverifiedPriceValid {
				prices.UnverifiedPriceValid, prices.UnverifiedPrice = true, r.Price
			}
			if prices.VerifiedPriceValid && prices.UnverifiedPriceValid {
				return
			}
		}
	}
	return
}

// maybeReloadRules reloads rules from the CID gravity API if the cache is expired. It reloads only once if being called
// concurrently. When loading takes more than the timeout, reloading turns to background and the method returns.
func (cg *clientRules) maybeReloadRules(url string, timeout time.Duration, cachePeriod time.Duration) {
	cg.lkLoadRules.Lock()
	defer cg.lkLoadRules.Unlock()
	lastUpdated := cg.rulesLastUpdated.Load().(time.Time)
	if time.Since(lastUpdated) < cachePeriod {
		return
	}
	// use buffered channel to avoid blocking the goroutine when the receiver is gone.
	chErr := make(chan error, 1)
	go func() {
		chErr <- func() error {
			body := fmt.Sprintf(`{"clientAddress":"%s"}`, cg.clientAddress)
			req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(body))
			if err != nil {
				return fmt.Errorf("creating HTTP request: %v", err)
			}
			req.Header.Set("Authorization", cg.apiKey)
			etag := cg.rulesETag.Load()
			if etag != nil {
				req.Header.Set("If-None-Match", etag.(string))
			}
			start := time.Now()
			resp, err := http.DefaultClient.Do(req)
			log.Infof("loading rules from API took %v", time.Since(start))
			if err != nil {
				return fmt.Errorf("contacting CID gravity server: %v", err)
			}
			switch resp.StatusCode {
			case http.StatusOK:
				// proceed
			case http.StatusNotModified:
				return nil
			default:
				return fmt.Errorf("unexpected HTTP status '%v'", resp.Status)
			}
			cg.rulesETag.Store(resp.Header.Get("ETag"))
			defer func() { _ = resp.Body.Close() }()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("reading http response: %v", err)
			}
			var rules CIDGravityRules
			if err := json.Unmarshal(b, &rules); err != nil {
				return fmt.Errorf("unmarshalling rules: %v", err)
			}
			cg.rulesLastUpdated.Store(time.Now())
			cg.rules.Store(&rules)
			return nil
		}()
		close(chErr)
	}()
	select {
	case err := <-chErr:
		if err != nil {
			log.Errorf("loading rules from API: %v", err)
		}
	case <-time.After(timeout):
	}
}
