package tile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"page"
)

type Tile struct {
	Row      int
	Col      int
	PageInfo page.PageInfo
	Data     []byte
}

func (t *Tile) GetUrl() string {
	x := t.GetX()
	y := t.GetY()
	w := t.PageInfo.TileWidth
	h := t.PageInfo.TileHeight
	return fmt.Sprintf("http://digital.nls.uk/imageserver/iipsrv.fcgi?iiif=%s/%d,%d,%d,%d/pct:100/0/native.jpg", t.PageInfo.Identifier, x, y, w, h)
}

func (t *Tile) GetX() int {
	return t.Col * t.PageInfo.TileWidth
}

func (t *Tile) GetY() int {
	return t.Row * t.PageInfo.TileHeight
}

func (t *Tile) Download() {
	response, err := http.Get(t.GetUrl())
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		t.Data = contents
	}
}
