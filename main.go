package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var weatherEmoji = map[string]string{
	"맑음":  " :sunny:",
	"구름":  " :cloud:",
	"흐림":  " :cloud:",
	"비":   " :umbrella:",
	"눈":   " :snowman",
	"뇌우":  " :zap:",
	"안개":  " :foggy:",
	"소나기": " :cyclone:",
	"우박":  " :snowflake:",
	"갬":   " :rainbow:",
}

var statusEmoji = map[string]string{
	"좋음": " :smile:",
	"보통": " :grinning:",
	"나쁨": " :worried:",
	"높음": " :worried:",
}

func getWeatherData(area string) (string, error) {
	title := fmt.Sprintf("*[오늘의 %s 날씨]*", area)
	url := fmt.Sprintf("https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=%s+날씨", area)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	temperature := doc.Find("div.temperature_text").First().Text()
	status := doc.Find("div.temperature_info > p").First().Text()
	for key, emoji := range weatherEmoji {
		if strings.Contains(status, key) {
			status += emoji
		}
	}

	weatherEtc := doc.Find("div.temperature_info > dl").First().Text()
	weatherEtc2 := doc.Find("div.report_card_wrap > ul").First().Text()
	sliceEtc2 := strings.Split(weatherEtc2, "  ")

	message := fmt.Sprintf("%s\n\n%s\n\n%s\n\n %s\n\n%s", title, temperature, status, weatherEtc)

	for _, info := range sliceEtc2 {
		for key, emoji := range statusEmoji {
			if strings.Contains(info, key) {
				info += emoji
				message += fmt.Sprintf("\n\n%s", info)
			}
		}
	}

	return message, nil
}

func main() {
	message, err := getWeatherData("강남")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data := map[string]string{"text": message}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = ioutil.WriteFile("./weather.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
