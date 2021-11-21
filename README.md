# Неофициальный Телеграм-стикерпак Радио-Т

[![main](https://github.com/chuhlomin/radio-t-stickers/actions/workflows/main.yml/badge.svg?branch=main&event=push)](https://github.com/chuhlomin/radio-t-stickers/actions/workflows/main.yml)

## Что это

[Набор Телеграм-стикеров](https://t.me/addstickers/radiot) вдохновленный [подкастом Радио-Т](http://radio-t.com/).

## Как это работает

Приложение на [Go](https://golang.org) генерирует стикеры в формате SVG на основе `speakers.csv`, `stickers.csv` и аватаров в папке `img`.

```bash
make run
```

Для конвертации SVG-стикеров в PNG используется [svgexport](https://github.com/shakiba/svgexport):

```bash
npm install svgexport -g
```

Для конвертации всех SVG файлов в `out` директории:

```bash
make convert
```

Наконец, стикеры вручную отправляются боту [@Stickers](https://t.me/Stickers).

## Источники

* [github.com/elquatro/rt-quotes](https://github.com/elquatro/rt-quotes)
* [lurkmore.to/Umputun](https://lurkmore.to/Umputun)
* [lurkmore.to/Бобук](https://lurkmore.to/Бобук)

## Контакты

* mail@chuhlomin.com
* [t.me/chuhlomin](https://t.me/chuhlomin)
