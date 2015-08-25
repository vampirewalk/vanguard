package gocvtext

// #cgo CXXFLAGS: -std=c++11
// #cgo darwin pkg-config: opencv
import "C"

import (
	"bytes"
	"image"
	"image/png"
)

func Recognize(img image.Image) (texts []string) {
	recognizer := NewTextRecognizer()

	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		return nil
	}

	imgBytes := NewByteVector()
	imgByteArray := buf.Bytes()
	for index := 0; index < len(imgByteArray); index++ {
		imgBytes.Add(imgByteArray[index])
	}
	words := recognizer.Recognize(imgBytes)
	for index := 0; index < int(words.Size()); index++ {
		texts = append(texts, string(words.Get(index)))
	}

	DeleteByteVector(imgBytes)
	DeleteTextRecognizer(recognizer)

	return texts
}
