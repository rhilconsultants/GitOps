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
		Help: "The Humidity value comming from the Sendor",
	})
)

type ReturnValue struct {
	Value int `json:"value"`
	Status string `json:"status"`
	Id int `json:"id"`
}

type myServerHandler struct {
	sensor_url string
}

func getSensor(hs myServerHandler) {

	log.Println("starting the humidity function")
	
		for {

			log.Println("in the humidity loop")
			 
			resp, err := http.Get(hs.sensor_url)	

			if err != nil {
				log.Fatalln(err)
			}
			body , err := ioutil.ReadAll(resp.Body)

			defer resp.Body.Close()
			
			sd := string(body)
			log.Println(sd)
				// JSON UnMarshal
			rValue := ReturnValue{}

			if err := json.Unmarshal([]byte(sd),&rValue); err != nil {
				log.Println("unable to Unmarshal the Body Request", err)
				return
			}

			humidityValue.Set(float64(rValue.Value))
			time.Sleep(2 * time.Second)
		}
}


func main() {

		log.Println("Starting Sensor Reading")
		sensor_url := getEnv("SENSOR_URL", "null")
		port := getEnv("PORT","9500")
		
		if sensor_url == "null" {
			log.Fatalln("missing environment variable SENSOR_URL")
			os.Exit(1)
		}

		hs := myServerHandler{}
		hs.sensor_url = sensor_url
		hs.sensor_type = sensor_type
		http.Handle("/metrics", promhttp.Handler())
		go func() {
			log.Printf("Starting HTTP Service on port %v", port)
			log.Fatal(http.ListenAndServe(":"+port, nil))
		}()
		getSensor(hs)
}
