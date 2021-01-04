# tenka

**This is just a proposal.**

A project for translated anime songs or idol songs lyrics file format.

It is pretty hard to identify who is singing
if you are new at some anime/idol music.
Moreover, if you are not a Japanese, you don't 
understand what they are saying.

I am one of people who love anime music but 
don't know Japanese, so I'm trying to make the 
lyrics file format.

## meta
You need to tell `tenka` about the music.

```toml
title = "Love Addiction"

languages = [
    "ja",
    "ja-en",
    "en",
]

# titles translated
titles = { ja = "Love Addiction", ja-en = "Love Addiction", en = "Love Addiction" }

artist = "アルストロメリア"
album = "THE IDOLM@STER SHINY COLORS FR@GMENT WING 05"
composer = "Jam9,家原正樹"
genre = ""
year = "2019"
disk = 1
track = 2
duration = "3:55"

series = "THE iDOLM@STER SHINY COLORS"

[singers]
    [singers.all] # overrides default 'all' properties
    color = "#000000"
    ja = "みんな"
    ja-en= "all"
    en = "all"

    [singers.amana]
    color = "#f53c71"
    ja = "大崎甘奈"
    ja-en = "Osaki Amana "
    en = "Amana Osaki"
    actor = "黒木ほの香"

    [singers.chiyuki]
    color = "#fbfbfb"
    ja = "桑山千雪"
    ja-en = "Kuwayama Chiyuki"
    en = "Chiyuki Kuwayama"
    actor = "芝崎典子"

    [singers.tenka]
    color = "#e75bec"
    ja = "大崎甜花"
    ja-en = "Osaki Tenka"
    en = "Tenka Osaki"
    actor = "前川涼子"
```

`title`, `artist`, `album`, `composer`, `genre`, `year`, `disk`, `track` are the same with mp3 tags.
You don't have to write here again, but please write `title`. It is used to find `ttsl` file with the same name.

`languages` is an array which contains languages used here. The first language should be the original language.
Please use `ISO 639-1 Code`. `ja-en` is how `ja` sounds in `en`. 
This languages order is important. 


## lyrics

```ttsl
[00:10.00][chiyuki]
アイスティーもペストリーも
AISUTII mo PESUTORII mo
Both iced tea and pastries are 

[00:14.00][chiyuki]
マンネリだよ
MANNERI da yo
routine and boring

[00:18.00][chiyuki]
いつもより贅沢に
Itsumo yori zeitaku ni
in a more luxurious way than usual
```

This is an example of `ttsl`(Tenka Translated and Synced Lyrics) file
for the music [Love Addiction](https://youtu.be/moVb4o6xn-k).
Each passage is called `block`, and the first line of `block` contains `timing` and `singer code`.
Following 2nd, 3rd, 4th lines are the `lyrics`, followd by Japanese(Original), Romanized, and English as specified in the `meta` file.

So, the format looks like below:

```ttsl
[hh:mm:ss.ms][code]
ja
ja-en
en
```

- Hours don't have to be specified.
- `ms` is milliseconds, but multiplied 10 times. (00:05.91 == 5910ms)


## try

In `tenka-go`, you can test how it really works, but you need *Love Addiction.mp3* in `musics` directory.
Put the music in, and run golang files.

```bash
cd tenka-go
go build
cd ..
.tenka-go/tenka-go
```