# GO Face Blur
An app I wrote to blur faces

## About
- The app uses [Kagami](https://github.com/Kagami/go-face) to find faces in images.
- The app then tilizes the face using [this Library](https://github.com/anthonynsimon/bild)

## To run
1. Put an image in `IMAGES/IMAGE_IN`
2. Run the below command
> go run src/main.go (name of file without path to) 

↑ ↑ ↑ for example, if the file is called `image.jpg`, just put `image.jpg` not the `/path/to/image.jpg`↑ ↑ ↑ 

3. Check `IMAGES/IMAGE_OUT` for the result

## Special Thanks
- [this article](https://www.golangprograms.com/how-to-add-watermark-or-merge-two-image.html)
- [this stackoverflow comment](https://stackoverflow.com/questions/67081246/how-to-capture-each-face-in-golang-go-face)
- Any of the libraries I used