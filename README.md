# tenka

**주의사항: 파서 만들줄도 모릅니다**

애니 음악이나 아이돌 음악 가사 파일 포맷을 위한 프로젝트 `tenka`.

새로 입문하는 애니나 아이돌(물)의 음악을 들으면, 누가 지금 부르고 있는지 알기가 어렵습니다. 게다가 일본어(또는 원어)를 할 줄 모르면,
그냥 반주랑 목소리만 듣고 있는 셈입니다.

그냥 일본어 가사는 못 알아들으니 

- 한글 발음
- 한국어 번역
- 부르고 있는 캐릭터/아이돌
- 가사 타이밍

... 등이 있는 가사 포맷이 필요해 새로운 가사 포맷을 구상하게 되었습니다.

결론: 덕질 편하려고...

이 문서의 예시는 알스트로메리아의 [Love Addiction](https://youtu.be/moVb4o6xn-k)입니다.

## 메타 데이터

가사를 작성하기 전에, 음악에 관련된 메타 데이터를 작성해야합니다.

```toml
title = "Love Addiction"

languages = [
    "ja",
    "ja-ko",
    "ko",
]

# titles translated
titles = { ja = "Love Addiction", ja-ko = "Love Addiction", ko = "Love Addiction" }

artist = "アルストロメリア"
album = "THE IDOLM@STER SHINY COLORS FR@GMENT WING 05"
composer = "Jam9,家原正樹"
genre = ""
year = "2019"
disk = 1
track = 2
duration = "03:55.00"

project = "THE iDOLM@STER SHINY COLORS"
```

`title`, `artist`, `album`, `composer`, `genre`, `year`, `disk`, `track`은 mp3 태그와 같습니다. 이 중 `title`을 제외하면
굳이 다시 적을 필요는 없습니다.
`duration`은 음악의 길이입니다. `project`는 해당 음악이 속한 프로젝트 이름을 적습니다(선택). 예를 들면, `THE iDOLM@STER`(아이마스), 
`Lovelive! Sunshine!`(러브라이브 선샤인), `Is the Order a Rabbit?`(주문은 토끼입니까?) 등이 있습니다. 프로젝트명을 로마자로 적을지,
아니면 일본어 원어 그대로 적을지는 아직 고민사항입니다.

`languages`는 원어와 번역된 언어의 목록을 적습니다. 다시 말해 가사 파일에서 사용될 언어의 목록입니다. 
`ISO 639-1 Code`로 적습니다. 첫 번째 언어는 무조건 원어(이 음악의 경우에는 일본어)가 되어야합니다. 
위 파일의 `ja-ko`는 일본어의 한글 발음입니다. `ja-en`(일본어를 로마자로)식으로도 적을 수 있습니다.

이상의 메타데이터 아래에는 `singers`를 기술합니다.

```toml
[singers]
    [singers.all] # overrides default 'all' properties
    color = "#000000"
    ja = "みんな"
    ja-ko = "all"
    ko = "모두"

    [singers.amana]
    color = "#f53c71"
    ja = "大崎甘奈"
    ja-ko = "오사키 아마나"
    ko = "오사키 아마나"
    actor = "黒木ほの香"

    [singers.chiyuki]
    color = "#fbfbfb"
    ja = "桑山千雪"
    ja-ko = "쿠와야마 치유키"
    ko = "쿠와야마 치유키"
    actor = "芝崎典子"

    [singers.tenka]
    color = "#e75bec"
    ja = "大崎甜花"
    ja-ko = "오사키 텐카"
    ko = "오사키 텐카"
    actor = "前川涼子"
```

`singers`는 노래를 부른 사람/아이돌/캐릭터 목록을 적습니다. 각 `singer`의 첫줄은 `singers.code`입니다. `code`는 앞으로 가사 파일에서
캐릭터를 부르는 변수명 같은 것입니다. 되도록이면 로마자 이름을 사용해주세요.

`color`를 지정할 수 있습니다. 이 캐릭터가 부르는 가사의 글자색입니다. 여기서는 퍼스널 컬러로 되어있습니다.

`ja`, `ja-ko`, `ko`는 언어별(그리고 발음별) 캐릭터의 이름입니다. `languages`에서 선언한 목록을 선언한 순서대로 써주세요.

`actor`는 성우/배우의 이름입니다(선택). 성우/배우의 이름은 번역할 필요가 없다고 판단했습니다. 

`singers`에는 기본 값으로 `all`이 있습니다. 다만, 이 파일처럼 기본값을 덮어씌울 수 있습니다. 

## 가사 파일

가사 파일은 각각의 `block`으로 이루어집니다. 형식은 다음과 같습니다.

```ttsl
[hh:mm:ss.ms][code]
ja
ja-ko
ko
```

첫 줄은 가사가 나타날 시간인 `timing`과 메타 데이터에서 선언한 캐릭터 이름인 `code`가 있습니다. 

### timing

`timing`은 다음 형식으로 적습니다.

```
[hh:mm:ss.ms]
```

- `hh`는 시간입니다. 필요하지 않은 경우 쓰지 않아도 됩니다.
- `mm`는 분입니다.
- `ss`는 초입니다.
- `ms`는 밀리세컨드(1/1000)의 10배입니다. 예를 들면, `ss.ms`가 `05.91`일 경우 `5910ms`와 같습니다.

### code

메타 데이터에서 선언한 캐릭터의 코드를 적으면 됩니다.

### 가사 내용

이 `block`에서 가사는 세 줄인데, `languages`에서 선언한 순서대로 적습니다.

### 예시

```
[02:24.30][amana]
伝えてもいいのかな
츠타에테모 이이노카나
전해도 괜찮은걸까

[02:29.00][tenka]
夢のような 時間に終わりはあるの？
유메노요-나 지칸니 오와리와아루노?
꿈만같은 시간에 끝은 있는거야?

[02:34.00][chiyuki]
わたし なみだ もろい
와타시 나미다 모로이
나는 눈물이 여려
```

**Love Addiction**의 2:24 부근의 가사입니다.

### 똑같은 내용, 이어붙이기, 중단하기

#### 똑같은 내용

일본어 노래 중, 영어 가사가 나올 때가 있습니다. 이럴 때는 굳이 한글 발음도, 한국어 번역도 적을 필요 없습니다. 이럴 때는 각 언어 자리에
`=`를 적어주세요.

```
[02:08.00][all]
So love me do
=
=
```

#### 이어붙이기

추임새(?) 같은 부분이 있는 가사가 있습니다. **Love Addiction**을 예시로 들면,

```
1:56 부근
치유키: 직감이 왔어
모두: (직감이 왔어)
치유키: 손을 잡아주었으니까
텐카: 틀림없어
모두: (틀림없어)
```

이처럼 모두가 합창하지만 괄호 안에 들어있는 가사가 있습니다. 이런 가사는 주요한 내용을 담고 있지 않습니다. 이런 경우는 가사를 **이어 붙이기** 할 수 있습니다.

`timing` 앞에 `+`를 붙여주세요.

```
[01:56.00][chiyuki]
ビビットきた
비빗토키타 
직감이 왔어 

+[01:56.50][all]
（ビビッときた）
(비빗토키타)
(직감이 왔어)

[01:58.00][chiyuki]
手を取ってくれたから
테오 톳테쿠레타카라
손을 잡아주었으니까

[02:02.00][tenka]
ちがいないわ
치가이나이와
틀림없어 

+[02:03.00][all]
（ちがいないわ）
(치가이나이와)
(틀림없어)
```

#### 중단하기

가사 중 반주에 들어가는 때가 있습니다. 그럼 반주 전의 가사가 끝나는 타이밍을 지정할 수 있습니다.

```
2:34 부근
치유키: 나는 눈물이 여려
아마나: 짓궂게 굴지는 말아줘
(반주)
```

이 때는 `code`, `가사 내용` 없이 `timing` 앞에 `-`를 적어주세요.

```
[02:34.00][chiyuki]
わたし なみだ もろい
와타시 나미다 모로이
나는 눈물이 여려

[02:37.90][amana]
いじわるはしないでね
이지와루와시나이데네
짓궂게 굴지는 말아줘

-[02:43.00]
```


## 체험해보기

이 파일 형식이 어떻게 돌아가는지 궁금하시면, 앞에서 계속 다뤘던 **Love Addiction** 파일이 있습니다. 다만 mp3 파일은 직접 구해서 
`musics/02. Love Addiction.mp3`로 넣어주셔야합니다.

그 다음, 아래의 멍령을 내려주세요.

```bash
go build tenka-go/
./tenka-go/tenka-go
```

## 목표

- [ ] 가사 제작 도우미 만들기
- [ ] JS로 데모 제작
- [ ] 3곡 이상 예시 추가
    - [ ] Arrive You ~그것이 운명이라도~
    - [ ] Brand New!
    - [ ] 홍백응원V