package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"regexp"
)

// readImage reads a image file from disk.
func ReadImage(name string) (image.Image, error) {

	fd, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	isJPEG, err := regexp.MatchString(".jpeg", name)
	if err != nil {
		return nil, err
	}

	// image.Decode requires that you import the right image package. We've
	// decode jpeg files then we would need to import "image/jpeg".

	var img image.Image
	if isJPEG {
		img, err = jpeg.Decode(fd)
	} else {
		img, _, err = image.Decode(fd)
	}

	if err != nil {
		return nil, err
	}

	return img, nil
}

// writeImage writes an Image back to the disk.
func WriteImage(img image.Image, name string) error {
	fd, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fd.Close()

	return png.Encode(fd, img)
}
