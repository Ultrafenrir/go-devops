package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (c *Config) InputLoop(
	ctx context.Context,
	m *Metrics,
) {
	localCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	var counter = 0
	t := time.NewTicker(c.PollInterval * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			counter++
			m.ReadStats(counter)
		case <-localCtx.Done():
			return
		}
	}
}

func (c *Config) OutputLoop(ctx context.Context, m *Metrics) {
	localCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	t := time.NewTicker(c.ReportInterval * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:

			m.PushMetrics(c.ServerAddr)

		case <-localCtx.Done():
			return
		}
	}
}

func main() {

	var m Metrics
	c := Config{
		ServerAddr:     "http://127.0.0.1" + ":" + "8080",
		Timeout:        1,
		ReportInterval: 10,
		PollInterval:   2,
	}

	ctxI, cancelI := context.WithCancel(context.Background())
	go c.InputLoop(ctxI, &m)
	log.Println("input loop started")
	ctxO, cancelO := context.WithCancel(context.Background())
	go c.OutputLoop(ctxO, &m)
	log.Println("output loop started")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Printf("caught %v", <-sig)
	cancelI() //stop inputs
	cancelO() // stop outputs
}
