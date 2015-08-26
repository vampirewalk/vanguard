package gocvtext

// #cgo CXXFLAGS: -std=c++11
// #cgo darwin pkg-config: opencv
import "C"

import (
	"bytes"
	"image"
	"image/png"
)

func Recognize(img image.Image) (texts []string, err error) {
	recognizer := NewTextRecognizer()

	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return nil, err
	}

	imgByteVector := NewByteVector()
	imgByteArray := buf.Bytes()
	for index := 0; index < len(imgByteArray); index++ {
		imgByteVector.Add(imgByteArray[index])
	}
	words := recognizer.Recognize(imgByteVector)
	for index := 0; index < int(words.Size()); index++ {
		texts = append(texts, string(words.Get(index)))
	}

	DeleteByteVector(imgByteVector)
	DeleteTextRecognizer(recognizer)

	return texts, nil
}
