package main

// Based on https://github.com/dps/go-xml-parse/blob/master/go-xml-parse.go

import (
	"fmt"
	"os"
	"flag"
	"encoding/xml"
)

var inputFile = flag.String("infile", "thwiki-20140507-pages-meta-current.xml",
	"Input file path")

type Redirect struct {
	Title string `xml:"title,attr"`
}

type Page struct {
	Title string `xml:"title"`
	Redir Redirect `xml:"redirect"`
	Text string `xml:"revision>text"`
}



func ExtractFromPage(p *Page) {
	fmt.Println(p.Title)
	fmt.Println(p.Text)
}

func main() {
	flag.Parse()

	xmlFile, err := os.Open(*inputFile)
	
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer xmlFile.Close()
		
	decoder := xml.NewDecoder(xmlFile)

	var inElement string
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		
		switch se := t.(type) {
		case xml.StartElement:		
			inElement = se.Name.Local
			if inElement == "page" {
				var p Page
				decoder.DecodeElement(&p, &se)
				if p.Redir.Title == "" {
					ExtractFromPage(&p)
				}
			}
		default:
		}		
	}
}
