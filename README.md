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
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Parse arguments and place them in a map for later look-up
	var argMap map[string]string
	argMap = make(map[string]string)
	argMap["-g"] = "1" // Place defaults in the map, if you like

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		key := args[i][:2]
		value := args[i][2:]
		argMap[key] = value
	}

	if argMap["-i"] == "" {
		fmt.Println("No input file specified. Please provide the following argument '-i<filename>'")
		os.Exit(1)
	}

	// Register supported image types
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	fmt.Printf("Converting image '%s'...\n", argMap["-i"])

	// Open the input file and decode the image
	file, err := os.Open(argMap["-i"])
	img, _, err := image.Decode(file)

	if err != nil {
		fmt.Printf("Could not find or open file: %s\n", err)
		os.Exit(1)
	}

	// Create a builder for ascii conversion/creation
	builder := creator.NewBuilder()

	// Provide a list of characters - these should be arranged from 'darker' to 'lighter' values,
	// an input image, and build a generator to be used as basis for the conversion.
	gamma, err := strconv.ParseFloat(argMap["-g"], 32)
	if err != nil {
		fmt.Println("Gamma correction value is invalid")
		os.Exit(1)
	}

	generator := builder.WithCharSet(creator.CharSet{
		Characters: []byte{' ', '.', ',', ':', ';', '+', '*', '?', '%', '&', '#', '@'},
	}).WithGammaCorrection(float32(gamma)).WithInput().Image(img).Build()

	// Do the actual conversion/ascii generation
	var out creator.Result
	err = generator.Generate(&out)
	if err != nil {
		fmt.Printf("Error converting image: %s\n", err)
		os.Exit(1)
	}

	// Write the result to an output file
	file, err = os.Create(fmt.Sprintf("%s.txt", strings.Split(argMap["-i"], ".")[0]))
	if err != nil {
		fmt.Println("Could not create output file")
		os.Exit(1)
	}

	size, err := file.Write(out.Ascii)
	if err != nil {
		fmt.Println("Could not write output to file")
		os.Exit(1)
	}

	fmt.Printf("Wrote %d bytes to %s\n", size, file.Name())

	err = file.Close()
	if err != nil {
		os.Exit(1)
	}
}
```

The example application parses a command line argument for the file name, and will provide the converted ascii art file with the same name, only with '.txt' ending.
The command line parsing is rudimentary, but serves its purpose for this use. You may find modules handling command line parsing better - such as the [flag](https://pkg.go.dev/flag) module.

For convenience's sake, the input file should not be too large. This will create an output which may be difficult to portray, as font size plays a role in its presentation.

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
