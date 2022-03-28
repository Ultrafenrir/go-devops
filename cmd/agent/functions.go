package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
)

func (m *Metrics) ReadStats(counter int) {

	var r runtime.MemStats
	runtime.ReadMemStats(&r)
	m.Alloc = float64(r.Alloc)
	m.BuckHashSys = float64(r.BuckHashSys)
	m.Frees = float64(r.Frees)
	m.GCCPUFraction = float64(r.GCCPUFraction)
	m.GCSys = float64(r.GCSys)
	m.HeapAlloc = float64(r.HeapAlloc)
	m.HeapIdle = float64(r.HeapIdle)
	m.HeapInuse = float64(r.HeapInuse)
	m.HeapObjects = float64(r.HeapObjects)
	m.HeapReleased = float64(r.HeapReleased)
	m.HeapSys = float64(r.HeapSys)
	m.LastGC = float64(r.LastGC)
	m.Lookups = float64(r.Lookups)
	m.MCacheInuse = float64(r.MCacheInuse)
	m.MCacheSys = float64(r.MCacheSys)
	m.MSpanInuse = float64(r.MSpanInuse)
	m.MSpanSys = float64(r.MSpanSys)
	m.Mallocs = float64(r.Mallocs)
	m.NextGC = float64(r.NextGC)
	m.NumForcedGC = float64(r.NumForcedGC)
	m.NumGC = float64(r.NumGC)
	m.OtherSys = float64(r.OtherSys)
	m.PauseTotalNs = float64(r.PauseTotalNs)
	m.StackInuse = float64(r.StackInuse)
	m.StackSys = float64(r.StackSys)
	m.Sys = float64(r.Sys)
	m.TotalAlloc = float64(r.TotalAlloc)
	m.RandomValue = rand.ExpFloat64()
	m.PollCount = int64(counter)
	// Debug print
	d, _ := json.Marshal(m)
	fmt.Println(string(d))

}

func (m *Metrics) PushMetrics(serverAddr string) {

	b, _ := json.Marshal(m)
	var metrics map[string]float64
	err := json.Unmarshal(b, &metrics)
	if err != nil {
		log.Fatal(err)
	}

	for key, val := range metrics {
		var uri string
		var f interface{} = val
		switch f.(type) {
		case nil:
			fmt.Println("metric value type is nil, skipping")
		case float64:
			uri = "/update/gauge/" + key + "/" + strconv.FormatFloat(val, 'f', -1, 64)
		case int64, int, int32:
			uri = "/update/counter/" + key + "/" + strconv.FormatFloat(val, 'f', -1, 64)
		default:
			fmt.Println("metric value type is unknown, skipping")
		}

		req, err := http.Post(serverAddr+uri, "text/plain", bytes.NewReader([]byte(strconv.FormatFloat(val, 'f', -1, 64))))
		if err != nil {
			log.Fatal(err)
		}
		req.Body.Close()
		fmt.Println(serverAddr + uri)
	}

}
