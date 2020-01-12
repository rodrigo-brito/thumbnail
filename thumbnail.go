package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli/v2"
)

func Download(URL string) (string, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	file, err := ioutil.TempFile("", "thumbnail-temp"+path.Ext(URL))
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return file.Name(), err
}

func main() {
	app := &cli.App{
		Name:  "thumbnail",
		Usage: "thumbnail builder",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "blur",
				Usage: "Gaussian blur in pixels",
			},
			&cli.Float64Flag{
				Name:  "brightness",
				Usage: "Image brightness adjustment",
			},
			&cli.IntFlag{
				Name:  "width,w",
				Usage: "Thumbnail width",
				Value: 1920,
			},
			&cli.IntFlag{
				Name:  "height,h",
				Usage: "Thumbnail height",
				Value: 1080,
			},
			&cli.StringFlag{
				Name:  "text,t",
				Usage: "Text content",
			},
			&cli.StringFlag{
				Name:  "subtext",
				Usage: "Small title content",
			},
			&cli.StringFlag{
				Name:  "font,f",
				Usage: "Font face",
				Value: "font/good_brush.ttf",
			},
			&cli.Float64Flag{
				Name:  "size",
				Usage: "Font size in pixels",
				Value: 96,
			},
		},
		Action: func(c *cli.Context) error {
			var err error
			args := c.Args()
			if args.Len() != 1 {
				log.Print("Invalid argument. Please inform an image path or url")
				cli.ShowCommandHelpAndExit(c, "", 2)
			}

			path := args.First()
			if strings.HasPrefix(path, "http") {
				path, err = Download(path)
				if err != nil {
					log.Fatal(err)
				}
			}

			options := []Option{Resize(c.Int("width"), c.Int("height"))}

			if blur := c.Float64("blur"); blur > 0 {
				options = append(options, Blur(blur))
			}

			if brightness := c.Float64("brightness"); brightness != 0 {
				options = append(options, Brightness(brightness))
			}

			if text := c.String("text"); len(text) > 0 {
				options = append(options, Text(
					text,
					c.String("subtext"),
					c.Int("width"),
					c.Int("height"),
					c.String("font"),
					c.Float64("size"),
				))
			}

			return Process(path, options...)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
