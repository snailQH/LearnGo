//parse ConversionStats.xml
package ConversionStats

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//THIS IS THE HEADER OF THE XML FILE
type Result struct {
	Flowcell   []Flowcell `xml:"Flowcell"`
}

type Flowcell struct {
	flowcellID string     `xml:"flowcell-id,attr"`
	Project      []Project      `xml:"Project"`
	flowcellLane []flowcellLane `xml:"Lane"`
}

type FlowcellLane struct {
	Number             string               `xml:"number,attr"`
	topUnknownBarcodes []topUnknownBarcodes `xml:"TopUnknownBarcodes"`
}

type TopUnknownBarcodes struct {
	Count    int    `xml:"Barcode,attr"`
	Sequcnce string `xml:"sequence,attr"`
	/**Name string `xml:",attr"`
	  Code string `xml:",attr"`*/
}

//

type Project struct {
	Name   string `xml:"name,attr"`
	Sample []Sample
}

type Sample struct {
	Name    string `xml:"name,attr"`
	Barcode []Barcode
}

type Barcode struct {
	Name string `xml:"name,attr"`
	Lane []Lane
}

type Lane struct {
	Number string `xml:"number,attr"`
	//BarcodeCount int `xml:"BarcodeCount"`
	tile []tile
}

type Tile struct {
	Number string `xml:"number,attr"`
	raw    []raw
	pf     []pf
}

type Raw struct {
	RawClusterCount int `xml:"ClusterCount"`
	rawRead         []rawRead
}

type RawRead struct {
	Number          string `xml:"number"`
	Yield           int    `xml:"Yield"`
	YieldQ30        int    `xml:"YieldQ30"`
	QualityScoreSum int    `xml:"QualityScoreSum"`
}

type Pf struct {
	PfClusterCount int `xml:"ClusterCount"`
	pfRead         []pfRead
}

type PfRead struct {
	Number          string `xml:"number"`
	Yield           int    `xml:"Yield"`
	YieldQ30        int    `xml:"YieldQ30"`
	QualityScoreSum int    `xml:"QualityScoreSum"`
}

func main() {
	info, err := os.Open("ConversionStats.xml")
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
		//fmt.Printf("%s\n",flowcell)
		fmt.Printf("Flowcell:%s\n", flowcell.FlowcellId)
		if (len(flowcell.FlowcellId)) == 0 {
			continue
		}

		for _, projects := range flowcell.Project {
			fmt.Printf("\tProject:%s\n", projects.Name)
			if (len(flowcell.Project)) == 0 {
				continue
			}

			for _, samples := range projects.Sample {
				fmt.Printf("\t\tSample:%s\n", samples.Name)
				if (len(samples.Name)) == 0 {
					continue
				}

				for _, barcodes := range samples.Barcode {
					fmt.Printf("\t\t\tBarcode%s\n", barcodes.Name)
					if (len(samples.Barcode)) == 0 {
						continue
					}

					for _, lanes := range barcodes.Lane {
						fmt.Printf("\t\t\t\tLane:Count:%s\t%b\n", barcodes.Name, lanes.BarcodeCount)
						if (len(lanes.Number)) == 0 {
							continue
						}

					}
				}
			}

		}
	}
}
