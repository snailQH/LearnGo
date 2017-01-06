//Package ConversionStats is design for parse xml result from bcl2fastq
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Result IS THE HEADER OF THE XML FILE,top level,Stats attr
type Result struct {
	Flowcell []Flowcell `xml:"Flowcell"`
}

//Flowcell is the second level of the xml file[ConversionStats.xml]
type Flowcell struct {
	//attribute of the Flowcell
	FlowcellID string `xml:"flowcell-id,attr"`
	//two Project sub node for the Flowcell,3rd level of the Result
	ProjectList []Project `xml:"Project"`
	//egiht Lane sub node for the Flowcell,3rd level of the Result
	FlowcellLaneList []FlowcellLane `xml:"Lane"`
}

//FlowcellLane definition,one of the two sub-node of Flowcell,3rd level of the Result
type FlowcellLane struct {
	LaneNumber         string `xml:"number,attr"`
	TopUnknownBarcodes tubs
}

//tubs is the sub node of the FlowcellLane,4th level of the Result
type tubs struct {
	Barcode []tubBarcodes `xml:"Barcode"`
}

//tub is the sub node of tubs,ie Barcode
type tubBarcodes struct {
	//attribute1 of the tubBarcodes
	Bcount string `xml:"count,attr"`
	//attribute2 of the tubBarcodes
	Bseq string `xml:"sequence,attr"`
}

//Project definition,another sub-node of the Flowcell,3rd level of the Result
type Project struct {
	Name   string   `xml:"name,attr"`
	Sample []Sample `xml:"Sample"`
}

//Sample is the sub node of the Project,4th level of the Result
type Sample struct {
	Name              string          `xml:"name,attr"`
	SampleBarcodeList []SampleBarcode `xml:"Barcode"`
}

//SampleBarcode is the sub node of Sample,5th level of the Result
type SampleBarcode struct {
	BarcodeName  string         `xml:"name,attr"`
	BarcodeLanes []BarcodeLanes `xml:"Lane"`
}

//BarcodeLanes is the sub node of the Barcode,6th of the Result
type BarcodeLanes struct {
	LaneNumber string `xml:"number,attr"`
	Tile       []Tile `xml:"Tile"`
}

//Tile is the sub node of Lane,7th level of the Result
type Tile struct {
	Number string `xml:"number,attr"`
	Raw    RawReads
	Pf     PfReads
}

//RawReads is one of the two sub-node of Tile,8th level of Result
type RawReads struct {
	RawClusterCount string    `xml:"ClusterCount"`
	RawRead         []RawRead `xml:"Read"`
}

//RawRead is the sub node of Raw,9th level of Result
type RawRead struct {
	//Number is the attribute if read, value can be 1 or 2,means pair-end
	Number          string `xml:"number,attr"`
	Yield           string
	YieldQ30        string
	QualityScoreSum string
}

//PfReads is one of the two sub-node of Tile,8th level of Result
type PfReads struct {
	PfClusterCount string   `xml:"ClusterCount"`
	PfRead         []PfRead `xml:"Read"`
}

//PfRead is the sub node of Pf,9th level of Result
type PfRead struct {
	//Number is the attribute if read, value can be 1 or 2,means pair-end
	Number          string `xml:"number,attr"`
	Yield           string
	YieldQ30        string
	QualityScoreSum string
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
		fmt.Printf("[Flowcell]:%s\n", flowcell.FlowcellID)
		if len(flowcell.FlowcellID) == 0 {
			continue
		}

		for _, projects := range flowcell.ProjectList {
			fmt.Printf("#1[Project:]%s\n", projects.Name)
			if len(flowcell.ProjectList) == 0 {
				continue
			}

			for _, samples := range projects.Sample {
				fmt.Printf("#12[Project-Sample:]%s\n", samples.Name)
				if len(samples.Name) == 0 {
					continue
				}

				for _, barcodes := range samples.SampleBarcodeList {
					fmt.Printf("#123[Barcode:]%s\n", barcodes.BarcodeName)
					if len(barcodes.BarcodeName) == 0 {
						continue
					}

					for _, lanes := range barcodes.BarcodeLanes {
						fmt.Printf("#1234[LaneID:]%s\n", lanes.LaneNumber)
						if len(lanes.LaneNumber) == 0 {
							continue
						}

						for _, tiles := range lanes.Tile {
							fmt.Printf("#12345[Tile:]%s\n", tiles.Number)
							if len(tiles.Number) == 0 {
								continue
							}

							raw := tiles.Raw
							fmt.Printf("#123456[RawClusterCount:]%s\n", raw.RawClusterCount)
							for _, rrr := range raw.RawRead {
								fmt.Println("[RawReads:number:]", rrr.Number)
								fmt.Println("[RawReads:yield,yieldq30,QualityScoreSum:]", rrr.Yield, rrr.YieldQ30, rrr.QualityScoreSum)

							}

							pf := tiles.Pf
							fmt.Printf("RawClusterCount:%s", pf.PfClusterCount)
							for _, prr := range pf.PfRead {
								fmt.Println("[PfReads:number:]", prr.Number)
								fmt.Println("[PfReads:yield,yieldq30,QualityScoreSum:]", prr.Yield, prr.YieldQ30, prr.QualityScoreSum)

							}

						}

					}
				}
			}

		}

		for _, laneList := range flowcell.FlowcellLaneList {
			fmt.Printf("\t[FlowcellLaneListNumber:]%s\n", laneList.LaneNumber)
			tu := laneList.TopUnknownBarcodes
			for _, tbs := range tu.Barcode {

				fmt.Println("[TopUnknownBarcodes:count,sequence", tbs.Bcount, "\t", tbs.Bseq)
			}

		}
	}
}
