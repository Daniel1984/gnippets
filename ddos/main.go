package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gnippets/ddos/pkg/cfg"
	"github.com/gnippets/ddos/pkg/ddos"
)

func main() {
	acfg, err := cfg.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}
	wg.Add(acfg.Attempts)

	dd := ddos.New(acfg, wg)
	start := time.Now()
	dd.Atack()
	wg.Wait()

	log.Printf("Time since beginning of atack:%s, Result:%+v", time.Since(start), dd)
}
