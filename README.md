# ascii-creator
This Go module converts an image to ascii art.

## Example use
To use the module, create a simple application to load an image from file, and call the __Convert__ function, with the file as input parameter. Here is a small example:

``` Golang
package main

import (
	"fmt"
	ascii_creator "github.com/faakern/ascii-creator"
	"image"
	"image/png"
	"os"
)

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	fmt.Println("Converting image...")

	file, err := os.Open("./image.png")

	if err != nil {
		fmt.Printf("Could not find or open file 'image.png': %s\n", err)
		os.Exit(1)
	}

	img, err := ascii_creator.Convert(file)
	if err != nil {
		fmt.Printf("Error converting image: %s\n", err)
	}

	file, err = os.Create("image.txt")
	if err != nil {
		fmt.Println("Could not create file image.txt")
	}

	size, err := file.Write(img)
	if err != nil {

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
