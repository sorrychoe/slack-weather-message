package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"emoji"
	"github.com/PuerkitoBio/goquery"
)

func getWeatherData(area string) (string, error) {
	title := fmt.Sprintf("*[오늘의 %s 날씨]*", area)
	url := fmt.Sprintf("https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=%s+날씨", area)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	temperature := doc.Find("div.temperature_text").Text()
	status := doc.Find("div.temperature_info > p").Text()
	for key, emoji := range statusEmoji {
		if strings.Contains(status, key) {
			status += emoji
		}
	}

	weatherEtc := doc.Find("div.temperature_info > dl").Text()
	weatherEtc2 := []string{}
	doc.Find("div.report_card_wrap > ul").Children().Each(func(i int, s *goquery.Selection) {
		if i%2 != 0 {
			weatherEtc2 = append(weatherEtc2, s.Text())
		}
	})

	message := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n", title, temperature, status, weatherEtc)

	for _, info := range weatherEtc2 {
		for key, emoji := range statusEmoji {
			if strings.Contains(info, key) {
				info += emoji
			}
		}
		message += fmt.Sprintf("%s\n\n", info)
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
