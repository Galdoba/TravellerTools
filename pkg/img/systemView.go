package main

import "github.com/fogleman/gg"

func main() {
	dc := gg.NewContext(4000, 4000)
	dc.SetRGB(0, 0, 0)
	// dc.DrawCircle(2000, 2000, 4000)
	// dc.Fill()

	// dc.SetRGB(0.5, 0.5, 0.5)
	// dc.DrawCircle(2000, 2000, 2000)
	// dc.Fill()
	// dc.SetRGB(0, 0, 0)
	// dc.DrawCircle(2000, 2000, 1410)
	// dc.Fill()

	// dc.SetRGB(0.5, 0.5, 0.5)
	// dc.DrawCircle(2000, 2000, 1010)
	// dc.Fill()
	// dc.SetRGB(0, 0, 0)
	// dc.DrawCircle(2000, 2000, 710)
	// dc.Fill()

	// dc.SetRGB(0.5, 0.5, 0.5)
	// dc.DrawCircle(2000, 2000, 510)
	// dc.Fill()
	// dc.SetRGB(0, 0, 0)
	// dc.DrawCircle(2000, 2000, 61)
	// dc.Fill()

	// dc.SetRGB(1, 0, 0)
	// dc.DrawCircle(2000, 2000, 1461)
	// dc.SetRGB(1, 0, 0)
	// dc.DrawCircle(2000, 2000, 1462)
	// dc.SetRGB(1, 0, 0)
	// dc.DrawCircle(2000, 2000, 1463)
	dc.SetRGB(1, 0, 0)
	dc.DrawEllipse(200, 200, 101, 81)
	dc.Fill()
	dc.SetRGB(1, 1, 1)
	dc.DrawEllipse(200, 200, 99, 80)
	dc.Fill()

	dc.SavePNG("out.png")
}
