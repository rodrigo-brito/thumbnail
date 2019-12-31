package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

type Option func(image.Image) image.Image

func Resize(w, h int) Option {
	return func(img image.Image) image.Image {
		return imaging.Fill(img, w, h, imaging.Center, imaging.Lanczos)
	}
}

func Blur(amount float64) Option {
	return func(img image.Image) image.Image {
		return imaging.Blur(img, amount)
	}
}

func Brightness(amount float64) Option {
	return func(img image.Image) image.Image {
		return imaging.AdjustBrightness(img, amount)
	}
}

func Text(title, subtitle string, width, height int, font string, fontSize float64) Option {
	return func(img image.Image) image.Image {
		dc := gg.NewContext(width, height)
		dc.Clear()
		dc.SetRGB(1, 1, 1)
		if err := dc.LoadFontFace(font, fontSize); err != nil {
			log.Fatal(fmt.Errorf("error on load font: %w", err))
		}
		dc.DrawImage(img, 0, 0)

		lines := strings.Split(title, "||")
		var yCoordinate float64
		for i, line := range lines {
			yCoordinate = float64(height)/2 + float64(len(lines)-1)*fontSize/2 - float64(len(lines)-i-1)*fontSize
			dc.DrawStringAnchored(line, float64(width)/2, yCoordinate, 0.5, 0.5)
		}

		if len(subtitle) > 0 {
			if err := dc.LoadFontFace(font, fontSize/2); err != nil {
				log.Fatal(fmt.Errorf("error on load font: %w", err))
			}
			yCoordinate = float64(height)/2 + float64(len(lines)-1)*fontSize/2 - float64(len(lines))*fontSize
			dc.DrawStringAnchored(subtitle, float64(width)/2, yCoordinate-fontSize/2, 0.5, 0.5)
		}

		dc.Clip()
		return dc.Image()
	}
}

func Process(path string, options ...Option) error {
	img, err := imaging.Open(path)
	if err != nil {
		return err
	}

	for _, option := range options {
		img = option(img)
	}

	out, err := os.Create("thumbnail.png")
	if err != nil {
		return err
	}
	return png.Encode(out, img)
}
