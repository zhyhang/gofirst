package json

import (
	"encoding/json"
	"log"
	"net/http"
)

const seniverseApi = "https://api.seniverse.com/v3/weather/now.json?key=A1KQAK6L4E&location=beijing&language=zh-Hans&unit=c"

type SeniverseWeather struct {
	Results []struct {
		Location struct {
			Id             string `json:"id"`
			Name           string `json:"name"`
			Country        string `json:"country"`
			Path           string `json:"path"`
			Timezone       string `json:"timezone"`
			TimezoneOffset string `json:"timezone_offset"`
		} `json:"location"`
		Now struct {
			Text        string `json:"text"`
			Code        string `json:"code"`
			Temperature string `json:"temperature"`
		} `json:"now"`
		LastUpdate string `json:"last_update"`
	} `json:"results"`
}

func FromHttpByDecoder() {
	resp, err := http.Get(seniverseApi)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var weather SeniverseWeather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err == nil {
		log.Printf("response(struct string):\n%v\n", weather)
		if bytes, err := json.MarshalIndent(weather, "", "\t"); err == nil {
			log.Printf("response(json string):\n%s\n", string(bytes))
		} else {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}

}
