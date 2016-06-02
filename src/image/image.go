package image

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"
	"page"
	"tile"
)

type Image struct {
	PageInfo page.PageInfo
	Image    *image.RGBA
}

func (i *Image) NewImage(p page.PageInfo) {
	i.PageInfo = p
	i.Image = image.NewRGBA(image.Rect(0, 0, i.PageInfo.Width, i.PageInfo.Height))
}

func (i *Image) AddTile(tile *tile.Tile) {
	reader := bytes.NewReader(tile.Data)
	tileImage, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Unable to parse image", err)
		return
	}
	dp := image.Point{tile.GetX(), tile.GetY()}
	draw.Draw(i.Image, image.Rectangle{dp, dp.Add(image.Pt(tile.PageInfo.Width, tile.PageInfo.Height))}, tileImage, image.ZP, draw.Src)
}

func (i *Image) SaveImage(filename string) {
	w, _ := os.Create(filename)
	defer w.Close()
	png.Encode(w, i.Image)
}
