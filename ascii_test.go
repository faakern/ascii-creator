package ascii

import (
	"image"
	"image/color"
	"testing"
)

// TestGenerator_Generate tests generating ascii output
func TestGenerator_Generate(t *testing.T) {

	// Create an image to test - put one black dot among whites
	img := image.NewRGBA(image.Rect(0, 0, 30, 30))
	img.Set(0, 0, color.Black)

	generator := Generator{
		alphaValue:     '@',
		alphaThreshold: 0,
		charset: CharSet{
			[]byte{' ', '@'},
		},
		img: img,
	}

	var output Result
	err := generator.Generate(&output)
	if output.Ascii == nil || err != nil {
		t.Fail()
	}

	// Check if length of the buffer is correct - expect an additional character (newline) for each row
	if len(output.Ascii) != 930 {
		t.Errorf("Output length is not correct: %d", len(output.Ascii))
	}

	// Pick some random places to check for values, in addition to the expected first value
	if output.Ascii[0] != ' ' || output.Ascii[230] != '@' || output.Ascii[390] != '@' || output.Ascii[740] != '@' {
		t.Errorf("Output does not contain correct value")
	}
}

func TestGenerator_GenerateShouldFailOnNoImage(t *testing.T) {
	generator := Generator{
		alphaValue:     '@',
		alphaThreshold: 0,
		charset: CharSet{
			[]byte{' ', '@'},
		},
	}

	var output Result
	err := generator.Generate(&output)
	if output.Ascii != nil || err == nil {
		t.Fail()
	}
}

func TestGenerator_GenerateShouldFailOnNoCharSet(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 30, 30))
	generator := Generator{
		alphaValue:     '@',
		alphaThreshold: 0,
		img:            img,
	}

	var output Result
	err := generator.Generate(&output)
	if output.Ascii != nil || err == nil {
		t.Fail()
	}
}
