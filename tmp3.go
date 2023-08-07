package main

import (
	"image/png"
	"os"

	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
)

func main() {
	// Start mouse drag and identify the coordinates.
	println("Press and hold the mouse button to select an area...")
	x1, y1 := robotgo.GetMousePos()
	robotgo.MouseToggle("down")
	for robotgo.MouseToggleStats() == "down" {
		// Wait until the mouse button is released.
	}
	x2, y2 := robotgo.GetMousePos()
	robotgo.MouseToggle("up")

	// Determine the rectangle for the screenshot.
	left, top, right, bottom := min(x1, x2), min(y1, y2), max(x1, x2), max(y1, y2)

	// Capture the screenshot.
	img, err := screenshot.CaptureRect(screenshot.Rect(left, top, right-left, bottom-top))
	if err != nil {
		panic(err)
	}

	// Save the screenshot to a file.
	file, err := os.Create("screenshot.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)

	println("Screenshot saved to screenshot.png")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
