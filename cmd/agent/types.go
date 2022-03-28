package main

import "time"

type Metrics struct {
	Alloc,
	TotalAlloc,
	LiveObjects,
	BuckHashSys,
	Frees,
	GCSys,
	HeapAlloc,
	HeapIdle,
	HeapInuse,
	HeapObjects,
	HeapReleased,
	HeapSys,
	LastGC,
	Lookups,
	MCacheInuse,
	MCacheSys,
	MSpanInuse,
	MSpanSys,
	Mallocs,
	NextGC,
	OtherSys,
	StackInuse,
	StackSys,
	Sys,
	PauseTotalNs,
	NumGC,
	NumForcedGC,
	GCCPUFraction,
	RandomValue float64
	PollCount int64
}

type Config struct {
	ServerAddr     string
	PollInterval   time.Duration
	ReportInterval time.Duration
	Timeout        time.Duration
}
