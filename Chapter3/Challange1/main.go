package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
    Water int `json:"water"`
    Wind  int `json:"wind"`
}

func main() {
    t := time.NewTicker(1 * time.Second)

    for {
        select {
        case <-t.C:
            rand.Seed(time.Now().UnixNano())
            valueWater := rand.Intn(100) + 1
            valueWind := rand.Intn(100) + 1

            var statusWater string
            if valueWater < 5 {
                statusWater = "aman"
            } else if valueWater >= 5 && valueWater <= 8 {
                statusWater = "siaga"
            } else {
                statusWater = "bahaya"
            }

 
            var statusWind string
            if valueWind < 6 {
                statusWind = "aman"
            } else if valueWind >= 6 && valueWind <= 15 {
                statusWind = "siaga"
            } else {
                statusWind = "bahaya"
            }

            data := Data{
                Water: valueWater,
                Wind:  valueWind,
            }

            jsonValue, err := json.Marshal(data)
            if err != nil {
                log.Fatal(err)
            }

            resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonValue))
            if err != nil {
                log.Fatal(err)
            }

            defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
				_, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}
			}

            log.Printf("Data: %s\nStatus water: %s\nStatus wind: %s\n", jsonValue, statusWater, statusWind)
        }
    }
}
