package ascii

import (
	"errors"
	"image"
)

const alphaThreshold = 0

func (c *CharSet) rgbaToAscii(r uint32, g uint32, b uint32, a uint32) byte {
	// Find the luminescence of this pixel value
	// Handle transparency by checking the alpha threshold
	var y = float32(255)
	if float32(a/257) > float32(alphaThreshold) {
		y = (0.2126 * float32(r/257)) + (0.7152 * float32(g/257)) + (0.0722 * float32(b/257))
	}

	// Place the value within the range of the represented properties characters,
	// and return its value.
	pos := int(y) * len(c.Characters) / 257

	return c.Characters[pos]
}

// Generate returns an ascii (byte array) representation,
// based on a parameterized generator consisting of an image and a charset.
func (g *Generator) Generate() ([]byte, error) {
	if g.img == nil || g.charset.Characters == nil {
		return nil, errors.New("no required image or charset provided")
	}

	img := g.img
	width, height := img.Bounds().Max.X, img.Bounds().Max.Y

	var ascii []byte
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ascii = append(ascii, g.charset.rgbaToAscii(img.At(x, y).RGBA()))
		}
		ascii = append(ascii, '\n') // Append a newline at the end of each row
	}

	return ascii, nil
}

type Generator struct {
	charset CharSet
	img     image.Image
}

type CharSet struct {
	Characters []byte
}
