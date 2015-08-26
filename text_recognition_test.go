package gocvtext

import (
	"image"
	_ "image/jpeg" // Register JPEG format
	_ "image/png"  // Register PNG  format
	"os"
	"reflect"
	"testing"
)

func TestRecognize(t *testing.T) {
	imgFilePath := "./scenetext01.jpg"
	infile, err := os.Open(imgFilePath)
	if err != nil {
		t.Error(err.Error())
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	if err != nil {
		t.Error(err.Error())
	}

	expectedTexts := []string{"AT", "All", "TIME", "PARKING", "PROHIBITED", "DOUBLE", "NOTICE"}

	texts, err := Recognize(img)
	if err != nil {
		t.Error(err.Error())
	}

	if reflect.DeepEqual(expectedTexts, texts) {
		t.Log("Test passed")
	} else {
		t.Error("Test didn't pass")
	}
}
