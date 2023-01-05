package ascii

import (
	"errors"
	"image"
)

func (c *CharSet) rgbaToAscii(r uint32, g uint32, b uint32) byte {
	// Find the luminescence of this pixel value
	y := (0.2126 * float32(r/257)) + (0.7152 * float32(g/257)) + (0.0722 * float32(b/257))

	// Place the value within the range of the represented properties characters,
	// and return its value.
	pos := int(y) * len(c.Characters) / 257

	return c.Characters[pos]
}

// Generate returns an ascii (byte array) representation,
// based on a parameterized generator consisting of an image and a charset.
func (gen *Generator) Generate() ([]byte, error) {
	if gen.img == nil || gen.charset.Characters == nil {
		return nil, errors.New("no required image or charset provided")
	}

	img := gen.img
	width, height := img.Bounds().Max.X, img.Bounds().Max.Y

	var ascii []byte
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if float32(a/257) > float32(gen.alphaThreshold) {
				ascii = append(ascii, gen.charset.rgbaToAscii(r, g, b))
			} else {
				ascii = append(ascii, gen.alphaValue)
			}
		}
		ascii = append(ascii, '\n') // Append a newline at the end of each row
	}

	return ascii, nil
}

// Generator provides parameters for generating ascii output.
// The alphaThreshold specifies what threshold should be used to return a replacement alphaValue.
// Charset specifies what characters to be used for the ascii conversion.
// These should be ordered from 'darker' to 'lighter' values - ex. [' ', '.', '*', '@'], for a 'natural' look,
// but you can experiment with this range, for various artistic expressions!!
// The more characters used, the more nuances the ascii image will gain.
// img holds the image to be used an input for the ascii generation.
type Generator struct {
	alphaThreshold int
	alphaValue     byte
	charset        CharSet
	img            image.Image
}

type CharSet struct {
	Characters []byte
}
