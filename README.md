# ascii-creator
This Go module converts an image to ascii art.

This module was initially used as a personal learning space, to explore the concepts and capabilities of Golang.

It continues to evolve, and must be considered work-in-progress. 

## Example use
To use the module, create a simple application which loads an image from file.
Create a new __Generator__ using a builder, and provide the character set you would like for the conversion and the loaded image.
Finally call __Generate__ to generate the ascii output. Here is a small example:

``` Golang
package main

import (
	"fmt"
	creator "github.com/faakern/ascii-creator"
	"image"
	"image/png"
	"os"
)

func main() {
	// Register supported image types
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	fmt.Println("Converting image...")

	// Open the input file and decode the image
	file, err := os.Open("./image.png")
	img, _, err := image.Decode(file)

	if err != nil {
		fmt.Printf("Could not find or open file 'image.png': %s\n", err)
		os.Exit(1)
	}

	// Create a generator builder for ascii conversion/creation
	builder := creator.NewBuilder()

	// Provide a list of characters - these should be arranged from 'darker' to 'lighter' values,
	// an input image, and build a property object to be used as basis for the conversion.
	generator := builder.WithCharSet(ascii_creator.CharSet{
		Characters: []byte{' ', '.', ',', ':', ';', '+', '*', '?', '%', '&', '#', '@'},
	}).WithInput().Image(img).Build()

	// Do the actual conversion/ascii generation
	out, err := generator.Generate()
	if err != nil {
		fmt.Printf("Error converting image: %s\n", err)
		os.Exit(1)
	}

	// Write the result to an output file
	file, err = os.Create("image.txt")
	if err != nil {
		fmt.Println("Could not create file image.txt")
		os.Exit(1)
	}

	size, err := file.Write(out)
	if err != nil {
		fmt.Println("Could not write output to file")
		os.Exit(1)
	}

	fmt.Printf("Wrote %d bytes to file 'image.txt'\n", size)

	err = file.Close()
	if err != nil {
		os.Exit(1)
	}
}
```

The example application requires the presence of a local PNG named 'image.png', and will provide the converted ascii art file called 'image.txt'.
For convenience's sake, the input file should not be too large. This will create an output which may be difficult to portray, as text size plays a role in its presentation.

An alternative approach, providing more freedom to choose input, can be achieved by replacing the first lines of code with the following:

```Golang
	// Register supported image types
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	fmt.Println("Converting image...")

	// Read the name of the input file
	args := os.Args[1:]
	if args == nil || len(args) == 0 {
		fmt.Println("No input file specified.")
		os.Exit(1)
	}

	// Open the input file and decode the image
	file, err := os.Open(args[0])
	img, _, err := image.Decode(file)
```

The input file name is expected to be provided as the first program argument.

## Output

The following image displays the output of the conversion, compared to the input:

![Senjou No Oubashi](https://github.com/faakern/ascii-creator/blob/main/senju_no_oubashi.png?raw=true)

## Gamma Correction

The ascii generation may produce an output which doesn't satisfy your aesthetic preferences.
Maybe it contains too much noise or has too dark values.

To allow for a cleaner output, the ascii generator can be specified with a gamma correction value:

```Golang
	generator := builder.WithCharSet(creator.CharSet{
		Characters: []byte{' ', '.', ',', ':', ';', '+', '*', '?', '%', '&', '#', '@'},
	}).WithGammaCorrection(1.5).WithInput().Image(img).Build()
```

![Added Gamma Correction](https://github.com/faakern/ascii-creator/blob/main/gamma_corrected.png?raw=true)
