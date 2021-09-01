package ch3

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const (
	width, h = 600, 320            //画布大小 （单位像素）
	cells    = 100                 // 格子
	xyrange  = 30.0                // axis ranges (-xyrange..+xyrange) 轴范围
	xyscale  = width / 2 / xyrange //pixels per x or y unit 每x或y单位像素
	zscale   = h * 0.4             // pixels per z unit 每个z像素
	angel    = math.Pi / 3         // 90℃/3 30℃
)

var sin30, cos30 = math.Sin(angel), math.Sin(angel)

func Svg() {

	file, err := os.OpenFile("polygon.svg", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal("OpenFile error", err)
	}
	str := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, h)
	_, err = io.WriteString(file, str)
	if err != nil {
		log.Fatal("WriteString error", err)
	}
	for i := 0; i < cells; {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			str = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			_, err = io.WriteString(file, str)
			if err != nil {
				log.Fatal("WriteString error", err)
			}
		}
		i = i +8
	}
	_, err = io.WriteString(file, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := h/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0) 勾股定理
	return math.Sin(r) / r
}
