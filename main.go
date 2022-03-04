/*
BSD 3-Clause License

Copyright (c) 2022, William Jones
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

/*
Henon Phase
*/
package main

import (
	"fmt"
	"math"

	gd "github.com/misterunix/cgo-gd"
)

func main() {
	fmt.Println("Starting Henon Phase")

	width := 2000
	height := 2000
	x1 := float64(width / 2)
	y1 := float64(height / 2)

	ibuf0 := gd.CreateTrueColor(width, height)

	/*
	   xn+1 = xn cos(a) - (yn - xn2) sin(a)
	   yn+1 = xn sin(a) + (yn - xn2) cos(a)
	*/

	// Initialize the Henon Phase
	var xn float64 = .01
	var yn float64 = .01
	var a float64 = -10.0

	bkground := ibuf0.ColorAllocateAlpha(0x00, 0x00, 0x00, 0)
	white := ibuf0.ColorAllocateAlpha(0xFF, 0xFF, 0xFF, 70)
	ibuf0.FilledRectangle(0, 0, width, height, bkground)

	//ibuf0.SetPixel(0, 0, white)

	//total := 1000 * 1000 * 1000
	//var running int

	for k := 0; k < 100; k++ {
		for j := 0; j < 1000; j++ {
			// Iterate through the Henon Phase
			for i := 0; i < 1000; i++ {
				xtmp := xn*math.Cos(a) - (yn-(xn*xn))*math.Sin(a)
				ytmp := xn*math.Sin(a) + (yn-(xn*xn))*math.Cos(a)

				xn = xtmp
				yn = ytmp

				// Map the Henon Phase to the screen

				x := int((xn*.5)*x1) + width/2
				y := int((yn*.5)*y1) + height/2

				//	x := int((xn*float64(width/2) + float64(width/2)) * 0.1)
				//		y := int((yn*float64(height/2) + float64(height/2)) * 0.1)

				ibuf0.SetPixel(x, y, white)
				//running++
				//if running%1000 == 0 {
				//	r := float64(running) / float64(total)
				//	fmt.Printf("\r%f", r)
				//}
			}
			xn += .05
			yn += .05
			//ibuf0.Png("test.png")
		}
		xn += float64(k) * 0.02
		yn += float64(k) * 0.02
		filename := fmt.Sprintf("images/%04d.png", k)
		ibuf0.Png(filename)

	}
	fmt.Println("\nDone")
	// Save the Henon Phase
	ibuf0.Png("test.png")

	fmt.Println("Ending Henon Phase")

}
