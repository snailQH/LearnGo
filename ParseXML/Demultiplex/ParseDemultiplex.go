//parse DemultiplexingStats.xml
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Result is the top structure of the xml file[demultiplex.xml]
type Result struct {
	Flowcell []Flowcell `xml:"Flowcell"`
}

//Flowcell is the first level structure of the Result
type Flowcell struct {
	FlowcellID string    `xml:"flowcellid,attr"` //Person []Person `xml:"person"`
	Project    []Project `xml:"Project"`
}

//Project is the second level structure of the Result
type Project struct {
	Name   string `xml:"name,attr"`
	Sample []Sample
}

//Sample is the sub level of Project
type Sample struct {
	Name    string `xml:"name,attr"`
	Barcode []Barcode
}

//Barcode is the sub level of the Sample
type Barcode struct {
	Name string `xml:"name,attr"`
	Lane []Lane
}

//Lane is the only sub level of Barcode
type Lane struct {
	Number       string `xml:"number,attr"`
	BarcodeCount int    `xml:"BarcodeCount"`
}

func main() {
	info, err := os.Open("DemultiplexingStats.xml")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(info)
	if err != nil {
		panic(err)
	}

	var v Result
	err = xml.Unmarshal(data, &v)
	if err != nil {
		panic(err)
	}

	//print xml table
	for _, flowcell := range v.Flowcell {
		fmt.Printf("[Flowcell]:%s\n", flowcell.FlowcellID)
		if (len(flowcell.FlowcellID)) == 0 {
			continue
		}

		for _, projects := range flowcell.Project {
			fmt.Printf("\t[Project]:%s\n", projects.Name)
			if (len(flowcell.Project)) == 0 {
				continue
			}

			for _, samples := range projects.Sample {
				fmt.Printf("\t\t[Sample]:%s\n", samples.Name)
				if (len(samples.Name)) == 0 {
					continue
				}

				for _, barcodes := range samples.Barcode {
					fmt.Printf("\t\t\t[Barcode]%s\n", barcodes.Name)
					if (len(samples.Barcode)) == 0 {
						continue
					}

					for _, lanes := range barcodes.Lane {
						fmt.Printf("\t\t\t\t[Lane:Count]:%s\t%b\n", barcodes.Name, lanes.BarcodeCount)
						if (len(lanes.Number)) == 0 {
							continue
						}
					}
				}
			}

		}
	}
}
