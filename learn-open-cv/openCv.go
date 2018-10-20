package main

import (
	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	windows := gocv.NewWindow("testing camera")
	picture := gocv.NewMat()

	for {
		webcam.Read(&picture)
		windows.IMShow(picture)
		windows.WaitKey(1)
	}
}
