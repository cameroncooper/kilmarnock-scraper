package main

import (
	img "./image"
	"fmt"
	"page"
	"tile"
)

func main() {
	fmt.Printf("Got %d pages\n", len(page.GetPages()))

	for _, p := range page.GetPages() {
		fmt.Printf("Getting page from %s\n", p.URL)
		p.DownloadContent()
		fmt.Printf("Got page %s, %s\n", p.GetIdentifier(), p.GetTitle())
		p.DownloadInfo()
		p.ParseInfo()
		fmt.Printf("Got page info %+v\n", p.Info)
		im := img.Image{}
		im.NewImage(p.Info)
		cols := p.Info.Width / p.Info.TileWidth
		rows := p.Info.Height / p.Info.TileHeight
		fmt.Printf("Downloading tiles: ")
		for i := 0; i <= cols; i++ {
			for j := 0; j <= rows; j++ {
				fmt.Printf("%d,%d ", i, j)
				t := tile.Tile{Col: i, Row: j, PageInfo: p.Info}
				t.Download()
				im.AddTile(&t)
			}
		}
		fmt.Printf("\n")
		im.SaveImage("output/" + p.GetTitle() + ".png")
	}
}
