//parse DemultiplexingStats.xml
package main

import(
	"io/ioutil"
	"fmt"
	"os"
	"encoding/xml"
)
/*
type Result struct {
    Person []Person `xml:"person"`
}
type Person struct {
    Name      string    `xml:"name,attr"`
    Age       int       `xml:"age,attr"`
    Career    string    `xml:"career"`
    Interests Interests `xml:"interests"`
}
type Interests struct {
    Interest []string `xml:"interest"`
}
**/
type Result struct{
	Flowcell []Flowcell `xml:"Flowcell"`
}

type Flowcell struct{
	FlowcellId string `xml:"flowcellid,attr"`//Person []Person `xml:"person"`
	Project []Project `xml:"Project"`
}

type Project struct{
	Name string `xml:"name,attr"`	
	Sample []Sample
}

type Sample struct{
	Name string `xml:"name,attr"`	
	Barcode []Barcode
}

type Barcode struct{
	Name string `xml:"name,attr"`	
	Lane []Lane
}

type Lane struct{
	Number string `xml:"number,attr"`
	BarcodeCount int `xml:"BarcodeCount"`
}

func main(){
	info,err := os.Open("DemultiplexingStats.xml")
	if err != nil{
		panic(err)	
	}
	data,err := ioutil.ReadAll(info)
	if err != nil{
		panic(err)	
	}
	
	var v Result
	err = xml.Unmarshal(data,&v)
	if err != nil{
		panic(err)	
	}	
	
	//print xml table
	for _,flowcell := range v.Flowcell{
		//fmt.Printf("%s\n",flowcell)
		fmt.Printf("Flowcell:%s\n",flowcell.FlowcellId)
		if (len(flowcell.FlowcellId))	== 0 {
			continue
		}
		
		for _,projects := range flowcell.Project{
			fmt.Printf("\tProject:%s\n",projects.Name)
			if (len(flowcell.Project)) == 0 {
					continue
			}
			
			for _, samples := range projects.Sample{
				fmt.Printf("\t\tSample:%s\n",samples.Name)
				if (len(samples.Name)) == 0{
					continue	
				}
				
				for _,barcodes := range samples.Barcode{
					fmt.Printf("\t\t\tBarcode%s\n",barcodes.Name)
					if (len(samples.Barcode)) == 0 {
						continue	
					}
					
					for _,lanes := range barcodes.Lane{
						fmt.Printf("\t\t\t\tLane:Count:%s\t%b\n",barcodes.Name,lanes.BarcodeCount)
						if (len(lanes.Number)) == 0  {
							continue
						}
						
					}
				}
			}
			
		}
	}
}

