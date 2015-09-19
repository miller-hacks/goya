package main

import (
	"image"
	"io"
	"log"
	"path"
	"runtime"

	"github.com/lazywei/go-opencv/opencv"
)

type Face struct {
	PointX int `json:"x"`
	PointY int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

func currentDir() string {
	_, currentfile, _, _ := runtime.Caller(0)
	return path.Dir(currentfile)
}

func detect(cascade *opencv.HaarCascade, r io.Reader) ([]*Face, error) {
	img, format, err := image.Decode(r)
	if err != nil {
		return []*Face{}, err
	}
	log.Printf("Image format: %s", format)
	rects := cascade.DetectObjects(opencv.FromImage(img))

	faces := []*Face{}

	for _, value := range rects {
		face := &Face{
			PointX: value.X(),
			PointY: value.Y(),
			Width:  value.Width(),
			Height: value.Height(),
		}
		faces = append(faces, face)
	}
	return faces, nil
}

func faces(r io.Reader) ([]*Face, error) {
	cascade := opencv.LoadHaarClassifierCascade(path.Join(currentDir(), "haarcascade_frontalface_alt.xml"))
	return detect(cascade, r)
}
