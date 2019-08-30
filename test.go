package main

import (
	"encoding/json"
	"log"
)

type AutoGenerated struct {
	Order     string  `json:"order"`
	Datetime  string  `json:"datetime"`
	Font      string  `json:"font"`
	Payprice  float64 `json:"payprice"`
	Fontprice float64 `json:"fontprice"`
	Status    string  `json:"status"`
	Platform  string  `json:"platform"`
	User      string  `json:"user"`
}

func main() {
	msg := "[{\"order\": \"61fdca2e-f00f-3059-bbf2-af414b8a463a\", \"datetime\": \"2019-08-19 17:36:50\", \"font\": \"爆米花体\", \"payprice\": 9.0, \"fontprice\": 1.0, \"status\": \"200\", \"platform\": \"oppo\", \"user\": \"17118048396\"}, {\"order\": \"21300218-575a-30f5-9d7d-9ee834491234\", \"datetime\": \"2019-08-20 09:52:14\", \"font\": \"音乐体\", \"payprice\": 9.0, \"fontprice\": 1.0, \"status\": \"1006\", \"platform\": \"oppo\", \"user\": \"17118048396\"}]"
	var dat []AutoGenerated
	if err := json.Unmarshal([]byte(msg), &dat); err == nil {
		log.Println("==============json str 转map=======================")
		log.Println(dat)
	}
}
