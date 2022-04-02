package handlers

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"strconv"
)

func WriteMetrics(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	metricMap := make(map[string]MetricValue)
	var I int64
	var m MetricValue
	q := r.URL.RequestURI()
	log.Println(q)
	reqMethod := r.Method
	log.Println(reqMethod)

	if reqMethod == "POST" {

		switch chi.URLParam(r, "metricType") {
		case "gauge":
			f, err := strconv.ParseFloat(chi.URLParam(r, "metricValue"), 64)
			if err != nil {
				log.Fatal(err)
			}
			m.val = float64ToBytes(f)
			fmt.Println(float64FromBytes(m.val[:]))
			m.isCounter = false
			metricMap[chi.URLParam(r, "metricName")] = m
			w.WriteHeader(http.StatusOK)
			r.Body.Close()

		case "counter":

			i, err := strconv.ParseInt(chi.URLParam(r, "metricValue"), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			I = I + i
			m.val = int64ToBytes(I)
			m.isCounter = true
			fmt.Println(int64FromBytes(m.val[:]))
			metricMap[chi.URLParam(r, "metricName")] = m
			w.WriteHeader(http.StatusOK)
			r.Body.Close()

		default:
			fmt.Println("Type", chi.URLParam(r, "metricType"), "wrong")
			outputMessage := "Type " + chi.URLParam(r, "metricType") + " not supported, only [counter/gauge]"
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(outputMessage))
			r.Body.Close()
		}

		//fmt.Println(metricMap)

		options := os.O_WRONLY | os.O_TRUNC | os.O_CREATE
		file, err := os.OpenFile("metrics.data", options, os.FileMode(0600))
		if err != nil {
			log.Fatal(err)
		}
		_, err = fmt.Fprintln(file, metricMap)
		if err != nil {
			log.Fatal(err)
		}
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Method is wrong")
		outputMessage := "Only POST method is alload"
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(outputMessage))

	}
}
