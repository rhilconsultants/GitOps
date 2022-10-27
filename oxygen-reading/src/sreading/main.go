package main

import (
	"net/http"
	"os"
	"log"
	"io/ioutil"
	"time"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func getEnv(key,fallback string) string {

	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}


var (
	humidityValue = promauto.NewGauge(prometheus.GaugeOpts {
		Name: "sendor_humidity_value",
		Help: "The Humidity Level comming from the Sendor",
	})
	oxygenValue = promauto.NewGauge(prometheus.GaugeOpts {
		Name: "sendor_oxygen_value",
		Help: "The Oxygen Level comming from the Sendor",
	})
)

type ReturnValue struct {
	Value int `json:"value"`
	Status string `json:"status"`
	Id int `json:"id"`
}

type myServerHandler struct {
	oxygen_url string
	humidity_url string
}

func getSensor(hs myServerHandler) {

	log.Println("starting the collector function")
	
		for {
			 
			o_resp, err := http.Get(hs.oxygen_url)	

			if err != nil {
				log.Fatalln(err)
			}

			o_body , err := ioutil.ReadAll(o_resp.Body)
			defer o_resp.Body.Close()

			h_resp, err := http.Get(hs.humidity_url)	

			if err != nil {
				log.Fatalln(err)
			}

			h_body , err := ioutil.ReadAll(h_resp.Body)

			defer h_resp.Body.Close()
			
			o_sd := string(o_body)
			h_sd := string(h_body)
			log.Println(o_sd, h_sd)
				// JSON UnMarshal
			o_rValue := ReturnValue{}
			h_rValue := ReturnValue{}

			if err := json.Unmarshal([]byte(h_sd),&h_rValue); err != nil {
				log.Println("unable to Unmarshal the Body Request", err)
				return
			}

			humidityValue.Set(float64(h_rValue.Value))

			if err := json.Unmarshal([]byte(o_sd),&o_rValue); err != nil {
				log.Println("unable to Unmarshal the Body Request", err)
				return
			}

			oxygenValue.Set(float64(o_rValue.Value))

			time.Sleep(2 * time.Second)
		}
}


func main() {

		log.Println("Starting Sensor Reading")
		humidity_url := getEnv("HUMIDITY_URL", "null")
		oxygen_url := getEnv("OXYGEN_URL", "null")
		port := getEnv("PORT","9500")
		
		if  humidity_url == "null" {
			log.Fatalln("missing environment variable HUMIDITY_URL")
			os.Exit(1)
		}
		
			if  oxygen_url == "null" {
			log.Fatalln("missing environment variable OXYGEN_URL")
			os.Exit(1)
		}

		hs := myServerHandler{}
		hs.humidity_url = humidity_url
		hs.oxygen_url = oxygen_url
		http.Handle("/metrics", promhttp.Handler())
		go func() {
			log.Printf("Starting HTTP Service on port %v", port)
			log.Fatal(http.ListenAndServe(":"+port, nil))
		}()
		getSensor(hs)
}
