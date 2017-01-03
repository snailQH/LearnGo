package main

import (
    "os"
    "encoding/xml"
    // "encoding/json"
    "io/ioutil"
    "fmt"
)

type Location struct {
    CountryRegion []CountryRegion
}

type CountryRegion struct {
    Name string `xml:",attr"`
    Code string `xml:",attr"`
    State []State
}

type State struct {
    Name string `xml:",attr"`
    Code string `xml:",attr"`
    City []City
}

type City struct {
    Name string `xml:",attr"`
    Code string `xml:",attr"`
    Region []Region
}

type Region struct {
    Name string `xml:",attr"`
    Code string `xml:",attr"`
}

func main() {
    f, err := os.Open("LocList.xml")
    if err != nil {
        panic(err)
    }

    data, err := ioutil.ReadAll(f)
    if err != nil {
        panic(err)
    }

    // v := make(map[string]interface{})
    var v Location
    err = xml.Unmarshal(data, &v)
    if err != nil {
        panic(err)
    }

    // fmt.Printf("%#v\n", v)

    // table
    for _, countryRegion := range v.CountryRegion {
        // fmt.Printf("%s,%s\n", countryRegion.Code, countryRegion.Name)
        if len(countryRegion.State) == 0 {
            continue
        }
        for _, state := range countryRegion.State {
            // fmt.Printf("%s,%s,%s\n", countryRegion.Code, state.Code, state.Name)
            if len(state.City) == 0 {
                continue
            }
            for _, city := range state.City {
                // fmt.Printf("%s,%s,%s,%s\n", countryRegion.Code, state.Code, city.Code, city.Name)
                if len(city.Region) == 0 {
                    continue
                }
                for _, region := range city.Region {
                    fmt.Printf("%s,%s,%s,%s,%s\n", countryRegion.Code, state.Code, city.Code, region.Code, region.Name)
                }
            }
        }
    }
    
    // // json
    // js, err := json.Marshal(&v.CountryRegion[0])
    // if err != nil {
    //  panic(err)
    // }
    // fmt.Printf("%s\n", js)
}