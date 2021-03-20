package utils

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Gpx struct {
	XMLName xml.Name `xml:"gpx"`
	Trk     Trk      `xml:"trk"`
}

type Trk struct {
	XMLName xml.Name `xml:"trk"`
	Name    Name     `xml:"name"`
	Trkseg  Trkseg   `xml:"trkseg"`
}

type Name struct {
	Name string `xml:"name"`
}

type Trkseg struct {
	XMLName xml.Name `xml:"trkseg"`
	Trkpt   []Trkpt  `xml:"trkpt"`
}

type Trkpt struct {
	XMLName xml.Name `xml:"trkpt"`
	Lat     string   `xml:"lat,attr"`
	Lon     string   `xml:"lon,attr"`
	Ele     string   `xml:"ele"`
	Time    string   `xml:"time"`
}

func Run() {
	// 프로젝트 root 경로 기준
	p := filepath.Join("./files", filepath.Base("20210320143858.gpx"))
	fp, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// xml 파일 읽기
	data, err := ioutil.ReadAll(fp)

	// xml 디코딩
	var Gpx Gpx
	xmlerr := xml.Unmarshal(data, &Gpx)
	if xmlerr != nil {
		panic(xmlerr)
	}

	for index, coord := range Gpx.Trk.Trkseg.Trkpt {
		fmt.Printf("[index : %d] lat : %s, lon: %s", index, coord.Lat, coord.Lon)
	}
}
