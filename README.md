# slack-weather-message

네이버 날씨 크롤러 & 슬랙 WebHook 연동 프로그램

---

## Initial Setting

본 프로그램은 Github Actions를 통해 슬랙에 자동으로 메세지를 전송하도록 설정하였습니다.

슬랙 연동의 경우, slack WebHook url을 본 repo와 연동하면 됩니다.

[github action과 slack WebHook 연동 방법](https://heegs.tistory.com/95)

### How to change Schedule time?

슬랙 전송 목표 시간은 한국 시간 기준 평일 오전 8시입니다.

GitHub Actions의 예약 실행은 정각에 부하가 집중되어 지연되거나 누락될 수 있습니다.
이를 완화하기 위해 한국 시간 07:47에 실행을 예약하고, 실행 환경 준비 후 08:00까지 기다렸다가 날씨를 조회하고 전송합니다.
08:00 이후에 시작되면 추가 대기 없이 바로 진행합니다. `push` 실행에는 대기를 적용하지 않습니다.

시간을 변경하려면 `.github/workflows/action.yaml`의 cron과 `Wait until 08:00 KST` 단계의 목표 시간을 함께 변경하세요.

```yaml
on:
  schedule:
    - cron: '47 22 * * 0-4'
```
cron으로 선언되는 변수의 값은 하단과 같습니다.

> minute / hour / day(month) / month / day(week)
>
> 0 ~ 59 / 0 ~ 23 / 1 ~ 31 / 1 ~ 12 / 0 ~ 6

현재 cron은 UTC 기준으로 설정되어 있습니다.

한국 시간으로 설정하고자 할 경우, 9시간을 뺀 시간으로 계산해서 기입해야 합니다.
위 설정은 UTC 일~목 22:47, 한국 시간 월~금 07:47입니다.

이 변경은 지연 가능성을 줄이는 완화책이며, 08:00 정시 전송을 보장하지 않습니다.
정시 실행이 필수라면 상시 실행 서버의 cron 등 별도 스케줄러에서 직접 프로그램을 실행하고 Slack으로 전송해야 합니다.
외부 스케줄러로 GitHub Actions만 호출하는 경우에도 runner 대기 지연은 남습니다.
자세한 제한은 [GitHub 예약 실행 문서](https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows#schedule)를 참고하세요.

---

## Installation for Develop

Local에서 실행하려 할 경우, 다음 명령어를 입력해주시기 바랍니다.

```bash
  make run
```

Make가 없는 경우, 별도로 설치가 필요합니다.
- [Make.exe install](https://gnuwin32.sourceforge.net/packages/make.htm)

---

## Authors

- [@sorrychoe](https://www.github.com/sorrychoe)
