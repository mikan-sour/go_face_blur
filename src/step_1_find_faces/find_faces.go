package findfaces

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
	"math/rand"
	"os"
	"path/filepath"

	"gaussian-blur/src/utils"

	"github.com/Kagami/go-face"
	"github.com/anthonynsimon/bild/blur"
)

const (
	dataDir    = "IMAGES"
	models     = "MODELS"
	STEP_0     = "IMAGE_IN"
	STEP_1     = "IMAGE_OUT"
	RADIUS_MIN = 10.0
	RADIUS_MAX = 35.0
)

var (
	modelsDir   = filepath.Join(dataDir, models)
	stepZeroDir = filepath.Join(dataDir, STEP_0)
	stepOneDir  = filepath.Join(dataDir, STEP_1)
)

type FaceFinder interface {
	FindFaces() (*os.File, string, error)
	cropImage(img image.Image, crop image.Rectangle) (image.Image, error)
	Tilize(img image.Image, rects []image.Rectangle) (*os.File, string, error)
}

type FaceFinderImpl struct {
	imgPath string
}

func New(imgPath string) FaceFinder {
	return &FaceFinderImpl{imgPath: imgPath}
}

func (f *FaceFinderImpl) FindFaces() (*os.File, string, error) {
	rec, err := face.NewRecognizer(modelsDir)
	if err != nil {
		return nil, "", err
	}

	defer rec.Close()

	imageInput := filepath.Join(stepZeroDir, f.imgPath)

	faces, err := rec.RecognizeFile(imageInput)

	if err != nil {
		return nil, "", err
	}

	fmt.Println("Number of Faces in Image: ", len(faces))

	img, err := utils.ReadImage(imageInput)
	if err != nil {
		return nil, "", err
	}

	var rects []image.Rectangle
	for _, face := range faces {
		rects = append(rects, face.Rectangle)
	}

	return f.Tilize(img, rects)

}

// cropImage takes an image and crops it to the specified rectangle.
func (f *FaceFinderImpl) cropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(crop), nil
}

func (f *FaceFinderImpl) Tilize(img image.Image, rects []image.Rectangle) (*os.File, string, error) {

	b := img.Bounds()
	outputImage := image.NewRGBA(b)
	draw.Draw(outputImage, b, img, image.Point{}, draw.Src)

	for h := 0; h < len(rects); h++ {
		rect := rects[h]
		lengthOf := rect.Size().X
		tiles := 144
		rowLen := int(math.Sqrt(float64(tiles)))
		tileDimension := int(math.Floor(float64(lengthOf) / float64(rowLen)))

		var start_x, start_y int

		for i := 0; i < rowLen; i++ {
			start_x = rect.Min.X + (i * tileDimension)
			for j := 0; j < rowLen; j++ {
				start_y = rect.Min.Y + (j * tileDimension)

				subT := subTile(start_x, start_y, tileDimension)

				subTImage, err := f.cropImage(img, subT)
				if err != nil {
					return nil, "", fmt.Errorf("failed to crop: %s", err)
				}

				blurredSubTImage := blur.Box(subTImage, randomFloat())

				draw.Draw(outputImage, subT, blurredSubTImage, subT.Min, draw.Over)

			}
		}
	}

	outputPath := fmt.Sprintf("%s/output.jpg", stepOneDir)

	op, err := os.Create(outputPath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create: %s", err)
	}

	err = png.Encode(op, outputImage)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer op.Close()

	return op, outputPath, nil

}

func randomFloat() float64 {
	return RADIUS_MIN + rand.Float64()*(RADIUS_MAX-RADIUS_MIN)
}

func subTile(minx, miny, spacing int) image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: minx,
			Y: miny,
		},
		Max: image.Point{
			X: minx + spacing,
			Y: miny + spacing,
		},
	}
}
