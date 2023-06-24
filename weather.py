import os

import requests as re
from bs4 import BeautifulSoup

url = os.environ.get('SLACK_URL')


def get_weather_data(area: str):

    title = "[오늘의 날씨]"
    url = 'https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=' + area + '+날씨'

    html = re.get(url)
    soup = BeautifulSoup(html.text, 'html.parser')

    temperature = soup.select_one("div.temperature_text").text.strip()
    status = soup.select_one("div.temperature_info > p").text.strip()
    weather_etc = soup.select_one("div.temperature_info > dl").text.strip()
    weather_etc2 = [etc_info.text.strip() for etc_info in soup.select_one("div.report_card_wrap > ul")][1:-1:2]

    message = title + '\n\n' + temperature + ' ' + status + '\n\n' + weather_etc + '\n\n'

    for info in weather_etc2:
        message += info + '\n\n'

    return message


def slack_post_text(url, text):
    result = re.post(url, {"message": text})
    return result


if __name__ == "__main__":
    slack_post_text(url, get_weather_data("강남"))
