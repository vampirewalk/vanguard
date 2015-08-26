## Requirements

- Mac OS X
- Go 1.4
- [OpenCV 3.0 beta](http://opencv.org/downloads.html)
- [opencv_contrib](https://github.com/itseez/opencv_contrib)
- pkg-config

## Installation
```
go get github.com/vampirewalk/gocvtext
```
## Usage

```
package main

import (
	"fmt"
	"github.com/vampirewalk/gocvtext"
	"image"
	"os"
)

func main() {
	imgFilePath := "path/to/image"
	infile, err := os.Open(imgFilePath)
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	if err != nil {
		panic(err)
	}

	texts, err := gocvtext.Recognize(img)
	if err != nil {
		panic(err)
	}
	fmt.Println(texts)
}

```