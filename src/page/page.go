package page

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type PageInfo struct {
	Identifier   string   `json:"identifier"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	ScaleFactors []int    `json:"scale_factors"`
	TileWidth    int      `json:"tile_width"`
	TileHeight   int      `json:"tile_height"`
	Formats      []string `json:"formats"`
	Qualities    []string `json:"quilities"`
	Profile      string   `json:"profile"`
}

type Page struct {
	URL          string
	content      string
	info_content string
	Info         PageInfo
}

func (p *Page) DownloadContent() {
	response, err := http.Get(p.URL)
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
		p.content = string(contents)
	}
}

func (p *Page) DownloadInfo() {
	url := "http://digital.nls.uk/imageserver/iipsrv.fcgi?iiif=" + p.GetIdentifier() + "/info.json"
	response, err := http.Get(url)
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
		p.info_content = string(contents)
	}
}

func (p *Page) GetIdentifier() string {
	r, _ := regexp.Compile("iiif=(/[0-9]+/[0-9]+.[0-9].jp2)'")
	return r.FindStringSubmatch(p.content)[1]
}

func (p *Page) GetTitle() string {
	r, _ := regexp.Compile("\"grouping_title\">(.*?)</h1>")
	return r.FindStringSubmatch(p.content)[1]
}

func (p *Page) ParseInfo() {
	err := json.Unmarshal([]byte(p.info_content), &p.Info)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON PageInfo %e", err)
	}
}
