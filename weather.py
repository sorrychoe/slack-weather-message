import json

import requests
from bs4 import BeautifulSoup

weather_emoji = {
    "맑음": " :sunny:",
    "구름": " :cloud:",
    "비": " :unbrella:",
    "눈": " :snowman",
    "뇌우": " :zap:",
    "안개": " :foggy:",
    "소나기": " :cyclone:",
    "우박" : " :snowflake:",
    "갬": " :rainbow:",
}

status_emoji = {
    "좋음": " :smile:",
    "보통": " :grinning:",
    "나쁨": " :worried:",
}


def get_weather_data(area: str):

    title = f"[오늘의 {area} 날씨]"
    url = 'https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=' + area + '+날씨'

    html = requests.get(url)
    soup = BeautifulSoup(html.text, 'html.parser')

    temperature = soup.select_one("div.temperature_text").text.strip()
    status = soup.select_one("div.temperature_info > p").text.strip()
    for weather in weather_emoji:
        if weather in status:
            status += weather_emoji[weather]
    weather_etc = soup.select_one("div.temperature_info > dl").text.strip()
    weather_etc2 = [etc_info.text.strip() for etc_info in soup.select_one("div.report_card_wrap > ul")][1:-1:2]

    message = title + '\n\n' + temperature + ' ' + status + '\n\n' + weather_etc + '\n\n'

    for info in weather_etc2:
        for status in status_emoji:
            if status in info:
                info += status_emoji[status]
        message += info + '\n\n'

    return message


if __name__ == "__main__":
    message = get_weather_data("강남")
    data = {"text": message}
    with open('./weather.json','w') as f:
        json.dump(data, f, ensure_ascii=False)
    print(message)
