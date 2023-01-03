package ascii

import (
	"image"
	"io"
)

// Holds a list of characters used for conversion
// - these should be arranged from 'darker' to 'lighter' values
var characters = []byte{' ', '.', ',', ':', ';', '+', '*', '?', '%', '&', '#', '@'}

const alphaThreshold = 0

func rgbaToAscii(r uint32, g uint32, b uint32, a uint32) byte {
	// Find the luminescence of this pixel value
	// Handle transparency by checking the alpha threshold
	var y = float32(255)
	if float32(a/257) > float32(alphaThreshold) {
		y = (0.2126 * float32(r/257)) + (0.7152 * float32(g/257)) + (0.0722 * float32(b/257))
	}

	// Place the value within the range of the represented ascii characters,
	// and return its value.
	pos := int(y) * len(characters) / 257

	return characters[pos]
}

// Convert Returns an ascii (byte array) representation of an image (file)
func Convert(file io.Reader) ([]byte, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	width, height := img.Bounds().Max.X, img.Bounds().Max.Y

	var ascii []byte
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ascii = append(ascii, rgbaToAscii(img.At(x, y).RGBA()))
		}
		ascii = append(ascii, '\n') // Append a newline at the end of each row
	}

	return ascii, nil
}
