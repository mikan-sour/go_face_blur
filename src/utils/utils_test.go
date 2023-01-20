package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

var fakeIMG = &image.RGBA{
	Rect:   image.Rect(0, 0, 3, 3),
	Stride: 3 * 4,
	Pix: []uint8{
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
		0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
	},
}

func TestReadImage(t *testing.T) {

	cases := []struct {
		fileName string
		isJPEG   bool
	}{
		{
			fileName: "test.jpg",
			isJPEG:   true,
		},
		{
			fileName: "test.png",
			isJPEG:   false,
		},
		{
			fileName: "test.jpeg",
			isJPEG:   true,
		},
	}

	for _, tc := range cases {

		f, err := os.Create(tc.fileName)
		if err != nil {
			panic(err)
		}
		defer os.Remove(f.Name())

		if tc.isJPEG {
			if err = jpeg.Encode(f, fakeIMG, nil); err != nil {
				panic(err)
			}
		} else {
			if err = png.Encode(f, fakeIMG); err != nil {
				panic(err)
			}
		}

		_, err = ReadImage(f.Name())

		if err != nil {
			t.Fatalf("expected no err but got %s", err.Error())
		}

	}
}

func TestReadImageOsOpenFail(t *testing.T) {
	_, err := ReadImage("test")
	if err == nil {
		t.Fatalf("expected nil but got %s", err.Error())
	}
}

func TestReadImageOsDecodeFail(t *testing.T) {
	expected := "image: unknown format"
	f, err := os.Create("test")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())

	_, err = ReadImage("test")
	if err.Error() != expected {
		fmt.Println(err.Error() == expected)
		t.Fatalf("expected %s but got something else", expected)
	}
}

func TestWriteImageOsCreateErr(t *testing.T) {
	expectedErr := "open : no such file or directory"

	fn := os.Getenv("FAKE_ENV_DOES_NOT_HAVE_AN_IMAGE")

	var img = new(image.Image)

	err := WriteImage(*img, fn)
	if err == nil {
		t.Fatalf("expected %s, got nil", expectedErr)
	}
	if err != nil && err.Error() != expectedErr {
		t.Fatalf("expected %s, got %s", expectedErr, err.Error())
	}
}

func TestWriteImageSuccess(t *testing.T) {
	f, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())

	err = WriteImage(fakeIMG, f.Name())
	if err != nil {
		t.Fatalf("expected no err, got %s", err.Error())
	}
}
