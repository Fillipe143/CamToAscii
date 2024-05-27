package main

import (
	"bytes"
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

const chars = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'.                                                                      "

func main() {
	webcam, _ := gocv.VideoCaptureDevice(0)

	frame := gocv.NewMat()
	for {
		webcam.Read(&frame)

		gocv.Resize(frame, &frame, image.Point{X: 200, Y: 100}, 0, 0, gocv.InterpolationDefault)
		gocv.CvtColor(frame, &frame, gocv.ColorBGRToRGB)
		gocv.Flip(frame, &frame, 1)

		var b bytes.Buffer
		for y := 0; y < frame.Rows(); y++ {
			for x := 0; x < frame.Cols(); x++ {
				pixel := frame.GetVecbAt(y, x)
				v := pixel[0] + pixel[1] + pixel[2]
				b.WriteRune(getCharForBrightness(int(v)))

				// Draw colors instead chars
				// b.WriteString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", pixel[0], pixel[1], pixel[2]))
				// b.WriteRune(' ')
				// b.WriteString("\x1b[0m")
			}

            b.WriteRune('\n')
		}
		fmt.Println(b.String())
		fmt.Printf("\033[%dA\033[%dD", frame.Rows(), frame.Cols())
	}
}

func getCharForBrightness(brigthness int) rune {
	index := (255 - brigthness) * (len(chars) - 1) / 255
	return []rune(chars)[index%len(chars)]
}
