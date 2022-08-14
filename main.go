package main

import (
	"fmt"
	"log"
	"time"

	"honeywell/printer"

	"github.com/beevik/etree"
)

const (
	host     = "192.168.137.1:9200"
	filename = "test.XML"

	base_qr = "https://github.com/andriantriputra/"
)

var values = []string{"andrian_001", "andrian_001", "andrian_001", "andrian_001"}

func main() {
	if err := printer.Connect(host); err != nil {
		log.Fatalf("Error Connect :%s", err.Error())
	}

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(filename); err != nil {
		log.Fatalf("Error ReadFromFile :%s", err.Error())
	}
	doc.Indent(10)

	root := doc.SelectElement("labels")
	rootin := root.SelectElement("label")

	j := 0
	for i, value := range values {
		for _, label := range rootin.SelectElements("variable") {
			lang := label.SelectAttrValue("name", "")
			if lang == "Data1" {
				label.SetText(base_qr + value)
			} else {
				label.SetText(value)
			}
		}

		val, _ := doc.WriteToString()
		if err := printer.Print(val); err != nil {
			log.Fatalf("Error Print :%s", err.Error())
		}
		j = i + 1
		fmt.Println(" ===================================================== ")
		log.Printf("%v;%s", j, value)
		fmt.Println(val)
		fmt.Println(" ===================================================== ")
		fmt.Println()

		time.Sleep(1 * time.Second)
	}
}
