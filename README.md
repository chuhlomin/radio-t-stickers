# Неофициальный Телеграм-стикерпак Радио-Т

[![Drone Build Status](https://ci.chuhlomin.com/api/badges/chuhlomin/radio-t-stickers/status.svg)](https://ci.chuhlomin.com/chuhlomin/radio-t-stickers)

## Что это

[Набор Телеграм-стикеров](https://t.me/addstickers/radiot) вдохновленный [подкастом Радио-Т](http://radio-t.com/).

## Как это работает

Приложение на Go генерирует стикеры в формате SVG на основе `speakers.csv`, `stickers.csv` и аватаров в папке `img`.

```bash
make run
```

Затем, c помощью [svgexport](https://github.com/shakiba/svgexport) SVG-стикеры конвертируются в PNG.

```bash
find out -type f -name "out/*.svg" -exec bash -c 'svgexport $0 $0.png' {} \;
```

Наконец, стикеры вручную отправляются боту [@Stickers](https://t.me/Stickers).

## Источники

* github.com/elquatro/rt-quotes
* lurkmore.to/Umputun
* lurkmore.to/Бобук

## Контакты

* mail@chuhlomin.com
* [t.me/chuhlomin](https://t.me/chuhlomin)
