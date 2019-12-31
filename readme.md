# Thumbnail Generator

A simple CLI for thumbnail creation.


## Instalatiion

`go get -u github.com/rodrigo-brito/thumbnail`


## Usage
```
NAME:
   thumbnail - thumbnail builder

USAGE:
   thumbnail command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --blur value        Gaussian blur in pixels (default: 0)
   --brightness value  Image brightness adjustment (default: 0)
   --width value       Thumbnail width (default: 1920)
   --height value      Thumbnail height (default: 1080)
   --text value        Text content
   --subtext value     Small title content
   --font value        Font face (default: "font/pacifico.ttf")
   --size value        Font size in pixels (default: 200)
   --help, -h          show help (default: false)
```

## Example of usage:
```
thumbnail --blur 10 \
    --brightness -10 \
    --size 200 \
    --text "We are the||Champions" \
    --subtext "~ Queen ~"  \
    background.jpg
```

## Result

<img src="https://user-images.githubusercontent.com/7620947/71635182-11983a80-2c01-11ea-9d81-40b4933f365a.png" />