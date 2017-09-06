package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type MapRequest struct {
	Query string `json:"query"`
	Key   string `json:"appkey"`
}

type Documents struct {
	Meta      map[string]interface{} `json:"meta"`
	Documents []struct {
		PlaceName   string `json:"place_name"`
		PlaceURL    string `json:"place_url"`
		Category    string `json:"category_name"`
		Address     string `json:"address_name"`
		RoadAddress string `json:"road_address_name"`
		Phone       string `json:"phone"`
	} `json:"documents"`
}

type EatingPlace struct {
	PlaceURL    string `json:"place_url"`
	Category    string `json:"category_name"`
	Address     string `json:"address_name"`
	RoadAddress string `json:"road_address_name"`
	Phone       string `json:"phone"`
}

func CallAPI() []byte {

	jsonreq, err := json.Marshal(MapRequest{
		Key: "1b5396fee020a19230beade9dbee9381",
	})

	if err != nil {
		log.Fatalln(err)
		recover()
	}

	return jsonreq

}

func ReqEatMap(jsonreq []byte, where string) map[string]EatingPlace {

	var docu Documents
	eatspot := make(map[string]EatingPlace)
	mapurl := "https://dapi.kakao.com/v2/local/search/keyword.json?query="
	myurl := mapurl + url.QueryEscape(where)

	// 1. 한번 url로 값을 전송해 본다.
	req, err := http.NewRequest("GET", myurl, bytes.NewBuffer(jsonreq))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "KakaoAK 1b5396fee020a19230beade9dbee9381")

	client := &http.Client{}

	// 3. 값을 받아온다
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	// 항상 Close함수는 defer로 까먹지 말고 만듭니다(자주 까먹음)
	defer resp.Body.Close()

	// 4. 값을 읽고 리턴해 줍니다.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	if err := json.Unmarshal(body, &docu); err != nil {
		log.Fatal(err)
	}

	for _, v := range docu.Documents {

		eatspot[v.PlaceName] = EatingPlace{
			Address:     v.Address,
			Category:    v.Category,
			Phone:       v.Phone,
			PlaceURL:    v.PlaceURL,
			RoadAddress: v.RoadAddress,
		}
	}

	return eatspot

}

func main() {

	j := CallAPI()
	ReqEatMap(j, "김밥천국 삼성")
}
