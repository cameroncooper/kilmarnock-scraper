package tile

import (
	"page"
	"testing"
)

func TestGetX(t *testing.T) {
	tile := Tile{Col: 0, Row: 0, PageInfo: page.PageInfo{Width: 257, Height: 257, TileWidth: 256, TileHeight: 256}}
	if tile.GetX() != 0 {
		t.Fatalf("X is not correct %d", tile.GetX())
	}

	tile.Col = 1
	if tile.GetX() != 256 {
		t.Fatalf("X is not correct %d", tile.GetX())
	}
}

func TestGetY(t *testing.T) {
	tile := Tile{Col: 0, Row: 0, PageInfo: page.PageInfo{Width: 257, Height: 257, TileWidth: 256, TileHeight: 256}}
	if tile.GetY() != 0 {
		t.Fatalf("Y is not correct %d", tile.GetY())
	}

	tile.Row = 1
	if tile.GetY() != 256 {
		t.Fatalf("Y is not correct %d", tile.GetY())
	}
}

func TestGetUrl(t *testing.T) {
	tile := Tile{Col: 0, Row: 0, PageInfo: page.PageInfo{Identifier: "/7445/74453208.5.jp2", Width: 257, Height: 257, TileWidth: 256, TileHeight: 256}}
	url := tile.GetUrl()
	if url != "http://digital.nls.uk/imageserver/iipsrv.fcgi?iiif=/7445/74453208.5.jp2/0,0,256,256/pct:100/0/native.jpg" {
		t.Fatalf("URL is not correct %s", url)
	}

	tile.Col = 1
	url = tile.GetUrl()
	if url != "http://digital.nls.uk/imageserver/iipsrv.fcgi?iiif=/7445/74453208.5.jp2/256,0,256,256/pct:100/0/native.jpg" {
		t.Fatalf("URL is not correct %s", url)
	}

	tile.Row = 1
	url = tile.GetUrl()
	if url != "http://digital.nls.uk/imageserver/iipsrv.fcgi?iiif=/7445/74453208.5.jp2/256,256,256,256/pct:100/0/native.jpg" {
		t.Fatalf("URL is not correct %s", url)
	}

}
