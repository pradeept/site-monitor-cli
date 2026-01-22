package call

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/pradeept/site-monitor-cli/internals/store"
)

func worker(address string, ch chan<- store.SiteStatus) {
	resp, err := http.Get(address)
	if err != nil {
		ch <- store.SiteStatus{
			Address:    address,
			StatusCode: 503,
			Status:     "Not reachable",
		}
		return
	}
	defer resp.Body.Close()

	ch <- store.SiteStatus{
		Address:    address,
		StatusCode: resp.StatusCode,
		Status:     strings.Split(resp.Status, " ")[1],
	}
}

func Request() {
	addressList := []string{"https://gobyexample.com", "https://invalidwebsiteeasdf.com", "https://pradeept.dev/notasdf", "https://facebook.com", "https://google.com"}
	ch := make(chan store.SiteStatus, len(addressList))

	var wg sync.WaitGroup

	for _, addr := range addressList {
		wg.Add(1)
		go func(a string) {
			defer wg.Done()
			worker(a, ch)
		}(addr)
	}

	// close channel once ALL workers finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		log.Println(res)
	}

	// s := store.NewStore()
	// s.Insert(entry)
	// entries := s.GetAll()
	// log.Println(entries)
}
