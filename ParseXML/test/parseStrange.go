//parseStrange
package main

import(
    "fmt"
    "io/ioutil"
    "encoding/xml"
    "os"
)

type result struct{
    FlowcellList []Flowcell `xml:"Flowcell"`
    
}

type Flowcell struct{
    FlowcellID string `xml:"flowcell-id,attr"`
    LaneList []Lane `xml:"Lane"`
}

type Lane struct{
    LaneNumber string `xml:"number,attr"`
    TopUnknownBarcodes tubs//strange <TopUnknownBarcodes>,without attr
}

type tubs struct{
    BarcodeList []Barcode `xml:"Barcode"`
}

type Barcode struct{
    Bcount int `xml:"count,attr"`
    Bseq string `xml:"sequence,attr"`
}


func main(){
    info,err := os.Open("strange.xml")
    if err != nil {
        panic(err)
    }
    data,err := ioutil.ReadAll(info)
    var v result
    xml.Unmarshal(data,&v)
    for _, i := range v.FlowcellList{
        fmt.Println(i.FlowcellID)
        for _,j := range i.LaneList{
            fmt.Println("Lane number:",j.LaneNumber)
            tu :=  j.TopUnknownBarcodes
            
            for _, tubbbbb := range tu.BarcodeList{//fmt.Println(tu.BarcodeList)
                fmt.Println(tubbbbb.Bcount,"\t",tubbbbb.Bseq)
            }
           
        }
    }
    

}
