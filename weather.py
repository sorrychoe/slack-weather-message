import json
import requests
from bs4 import BeautifulSoup


def get_weather_data(area: str):

    title = f"[오늘의 {area} 날씨]"
    url = 'https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=' + area + '+날씨'

    html = requests.get(url)
    soup = BeautifulSoup(html.text, 'html.parser')

    temperature = soup.select_one("div.temperature_text").text.strip()
    status = soup.select_one("div.temperature_info > p").text.strip()
    weather_etc = soup.select_one("div.temperature_info > dl").text.strip()
    weather_etc2 = [etc_info.text.strip() for etc_info in soup.select_one("div.report_card_wrap > ul")][1:-1:2]

    message = title + '\n\n' + temperature + ' ' + status + '\n\n' + weather_etc + '\n\n'

    for info in weather_etc2:
        message += info + '\n\n'

    return message



if __name__ == "__main__":
    message = get_weather_data("강남")
    print({"text": message})