//test
package main

import (
	"fmt"
	"os"
	"encoding/xml"
	"io/ioutil"
)

type Location struct{
	CountryRegion []CountryRegion
}

type CountryRegion struct{
	Name string `xml:",attr"`
	Code string `xml:",attr"`
	State []State
}

type State struct{
	Name string `xml:",attr"`
	Code string `xml:",attr"`
	City []City
}

type City struct{
	Name string `xml:",attr"`
	Code string `xml:",attr"`
	Region []Region
}

type Region struct{
	Name string `xml:",attr"`
	Code string `xml:",attr"`
}

func main(){
	f,err := os.Open("LocList.xml")
	if err != nil{
		panic(err)	
	}
	
	data,err := ioutil.ReadAll(f)
	if err != nil{
		panic(err)	
	}
	
	var v Location// v:= make(map[string]interface{})
	err = xml.Unmarshal(data, &v)
	if err != nil{
		panic(err)	
	}
	
	for _,countryRegion := range v.CountryRegion{
		if len(countryRegion.State) == 0 {
			continue	
		}
		if countryRegion.Code != "1"{
			continue	
		}
		fmt.Printf("%s,%s\n",countryRegion.Name,countryRegion.Code)
		
		for _,State := range countryRegion.State{
			if len(countryRegion.State)	 == 0 {
				continue	
			}
			fmt.Printf("%s,%s,\n",countryRegion,State)
			
			for _,City
			
		}
		
	}
	
}
