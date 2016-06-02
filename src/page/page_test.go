package page

import (
	"io/ioutil"
	"testing"
)

func TestGetInfo(t *testing.T) {
	p := Page{}
	// p.GetContent()

	// Read in page.html fixture
	buf, err := ioutil.ReadFile("test_fixtures/page.html")
	if err != nil {
		t.Fatal("Unable to open page.html")
	}
	p.content = string(buf)

	id := p.GetIdentifier()

	if id != "/7445/74453044.5.jp2" {
		t.Fatalf("identifier is wrong: ", id)
	}

}

func TestGetTitle(t *testing.T) {
	p := Page{}

	// Read in page.html fixture
	buf, err := ioutil.ReadFile("test_fixtures/page.html")
	if err != nil {
		t.Fatal("Unable to open page.html")
	}
	p.content = string(buf)

	title := p.GetTitle()

	if title != "(1) Spine - Burns' poems" {
		t.Fatalf("identifier is wrong: ", title)
	}

}

func TestParseInfo(t *testing.T) {
	p := Page{}
	// p.GetContent()

	// Read in page.html fixture
	buf, err := ioutil.ReadFile("test_fixtures/info.json")
	if err != nil {
		t.Fatal("Unable to open info.json")
	}
	p.info_content = string(buf)

	p.ParseInfo()

	if p.Info.Identifier != "/7445/74453044.5.jp2" {
		t.Fatalf("identifier is wrong: ", p.Info.Identifier)
	}

	if p.Info.Width != 2500 {
		t.Fatalf("width is wrong: ", p.Info.Width)
	}

	if p.Info.Height != 14124 {
		t.Fatalf("height is wrong: ", p.Info.Height)
	}

	if p.Info.TileWidth != 256 {
		t.Fatalf("tile_width is wrong: ", p.Info.TileWidth)
	}

	if p.Info.TileHeight != 256 {
		t.Fatalf("tile_height is wrong: ", p.Info.TileHeight)
	}
}
