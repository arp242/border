package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
)

func main() {
	colour := flag.String("color", "#999", "color to use")
	width := flag.Int("width", 2, "border width")
	write := flag.String("write", "_border", "insert before the extension in the output filename")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "must give at least one file")
		os.Exit(1)
	}

	r, g, b, err := hexcolor(*colour)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	border := color.RGBA{R: r, G: g, B: b, A: 255}

	for _, f := range flag.Args() {
		fp, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(fp)
		fp.Close()
		if err != nil {
			log.Fatal(err)
		}

		bounds := img.Bounds()
		nbounds := bounds
		nbounds.Max.X += *width * 2
		nbounds.Max.Y += *width * 2
		nimg := image.NewRGBA(nbounds)
		for y := nbounds.Min.Y; y < nbounds.Max.Y; y++ {
			for x := nbounds.Min.X; x < nbounds.Max.X; x++ {
				if x < *width || y < *width || x >= bounds.Max.X+*width || y >= bounds.Max.Y+*width {
					nimg.Set(x, y, border)
				} else {
					nimg.Set(x, y, img.At(x-*width, y-*width))
				}
			}
		}

		dot := strings.LastIndex(f, ".")
		destfp, err := os.Create(f[:dot] + *write + f[dot:])
		if err != nil {
			log.Fatal(err)
		}
		err = png.Encode(destfp, nimg)
		if err != nil {
			log.Fatal(err)
		}
		err = destfp.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func hexcolor(c string) (uint8, uint8, uint8, error) {
	if c[0] != '#' {
		return 0, 0, 0, fmt.Errorf("invalid -color: %q", c)
	}

	var rgb []byte
	if len(c) == 4 {
		c = "#" +
			strings.Repeat(string(c[1]), 2) +
			strings.Repeat(string(c[2]), 2) +
			strings.Repeat(string(c[3]), 2)
	}

	n, err := fmt.Sscanf(strings.ToLower(c), "#%x", &rgb)
	if n != 1 || len(rgb) != 3 || err != nil {
		return 0, 0, 0, fmt.Errorf("invalid -color: %q", c)
	}

	return rgb[0], rgb[1], rgb[2], nil
}
