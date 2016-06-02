package image

import (
	"io/ioutil"
	"page"
	"testing"
	"tile"
)

func TestNewImage(t *testing.T) {
	image := Image{}
	image.NewImage(page.PageInfo{Width: 640, Height: 480})

	if image.Image.Bounds().Min.X != 0 {
		t.Fatalf("Min image x bounds is wrong: ", image.Image.Bounds().Min.X)
	}

	if image.Image.Bounds().Max.X != 640 {
		t.Fatalf("Max image x bounds is wrong: ", image.Image.Bounds().Max.X)
	}

	if image.Image.Bounds().Min.Y != 0 {
		t.Fatalf("Min image y bounds is wrong: ", image.Image.Bounds().Min.Y)
	}

	if image.Image.Bounds().Max.Y != 480 {
		t.Fatalf("Max image y bounds is wrong: ", image.Image.Bounds().Max.Y)
	}
}

func TestAddTile(t *testing.T) {
	image := Image{}
	image.NewImage(page.PageInfo{Width: 640, Height: 480, TileWidth: 256, TileHeight: 256})

	tile := tile.Tile{Row: 0, Col: 0, PageInfo: image.PageInfo}
	buf, err := ioutil.ReadFile("test_fixtures/00.jpg")
	if err != nil {
		t.Fatal("Unable to open 00.jpg")
	}
	tile.Data = buf
	image.AddTile(&tile)

	tile.Col = 1
	tile.Row = 0
	buf, err = ioutil.ReadFile("test_fixtures/10.jpg")
	if err != nil {
		t.Fatal("Unable to open 10.jpg")
	}
	tile.Data = buf
	image.AddTile(&tile)

	image.SaveImage("tiles.png")

	tile.Col = 0
	tile.Row = 1
	buf, err = ioutil.ReadFile("test_fixtures/01.jpg")
	if err != nil {
		t.Fatal("Unable to open 01.jpg")
	}
	tile.Data = buf
	image.AddTile(&tile)

	image.SaveImage("tiles.png")

	tile.Col = 1
	tile.Row = 1
	buf, err = ioutil.ReadFile("test_fixtures/11.jpg")
	if err != nil {
		t.Fatal("Unable to open 11.jpg")
	}
	tile.Data = buf
	image.AddTile(&tile)

	image.SaveImage("tiles.png")
}
