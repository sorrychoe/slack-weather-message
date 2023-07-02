# slack-weather-message

네이버 날씨 크롤러 & 슬랙 WebHook 연동 프로그램

---

## Initial Setting

본 프로그램은 Github Actions를 통해 슬랙에 자동으로 메세지를 전송하도록 설정하였습니다.

슬랙 연동의 경우, slack WebHook url을 본 repo와 연동하면 됩니다.

[github action과 slack WebHook 연동 방법](https://heegs.tistory.com/95)

### How to change Schedule time?

시간대는 한국 시간 기준 평일 오전 7시로 설정돼 있습니다.

시간대 변경은 .github/workflows/action.yaml 파일에 설정된 시간대를 변경해주시면 됩니다.

```yaml
on:
  schedule:
    - cron: '0 22 * * 0-4'
```
cron으로 선언되는 변수의 값은 하단과 같습니다.

> minute / hour / day(month) / month / day(week)
>
> 0 ~ 59 / 0 ~ 23 / 1 ~ 31 / 1 ~ 12 / 0 ~ 6

단, 시간대는 UTC 기준으로설정해야 합니다.

한국 시간으로 설정하고자 할 경우, 9시간을 뺀 시간으로 계산해서 기입해야 합니다.

---

## Installation for Develop

Local에서 실행하거나 설정을 변경하고자 할 경우, 하단의 명령어를 통해 개발 환경을 세팅해주시기 바랍니다.

```bash
  make install
```

Make가 없는 경우, 별도로 설치가 필요합니다.
- [Make.exe install](https://gnuwin32.sourceforge.net/packages/make.htm)

---

## Authors

- [@sorrychoe](https://www.github.com/sorrychoe)
