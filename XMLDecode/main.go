package main

import (
	"encoding/xml"
	"fmt"
	// "io/ioutil"
	// "golang.org/x/net/html/charset"
	"os"
)

func main() {
	fmt.Println("Version: 0.1.0")

	type xmlProduct struct {
		Code                 string `xml:"codigo,attr"`
		Brand                string `xml:"marca,attr"`
		Category             string `xml:"categoria,attr"`
		Description          string `xml:"descricao,attr"`
		Unit                 string `xml:"unidade,attr"`
		Multiplo             string `xml:"multiplo,attr"`
		DealerPrice          string `xml:"preco,attr"`
		SuggestionPrice      string `xml:"precoeup,attr"`
		Weight               string `xml:"peso,attr"`
		TechnicalDescription string `xml:"descricao_tecnica,attr"`
		Availability         string `xml:"disponivel,attr"`
		Measurements         string `xml:"dimensoes,attr"`
		PictureLink          string `xml:"foto,attr"`
		WarrantyTime         string `xml:"tempo_garantia,attr"`
		RMAProcedure         string `xml:"procedimentos_rma,attr"`
	}

	type xmlDoc struct {
		Products []xmlProduct `xml:"produto"`
	}

	// file := os.Getenv("ZUNKAPATH") + "/xml/aldo/aldo-products-substitution.xml"
	// data, err := ioutil.ReadFile(file)
	// if err != nil {
	// fmt.Printf("error reading file. %v\n", err)
	// return
	// }

	aldoXML := xmlDoc{}

	decoder := xml.NewDecoder(os.Stdin)
	// decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&aldoXML)
	if err != nil {
		fmt.Printf("error decoding XML file. %v\n", err)
	}

	// err = xml.Unmarshal([]byte(data), &aldoXML)
	// if err != nil {
	// fmt.Printf("error unmarsahling. %v\n", err)
	// return
	// }

	fmt.Printf("codigo: %q\n", aldoXML.Products[0].Code)
	fmt.Printf("descricao_tecnica: %v\n", aldoXML.Products[0].TechnicalDescription)

	// fmt.Println(string(in))
}
