package main
import (
    "encoding/xml"
    "io/ioutil"
    "log"
   // "fmt"
)
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
func (person *Person) Chkis18() (flag bool) {
    if person.Age > 18 {
        flag = true
    }
    return flag
}
type Checker interface {
    Chkis18() (flag bool)
}
func main() {
    content, err := ioutil.ReadFile("test.xml")
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Println(content)
    var result Result
    err = xml.Unmarshal(content, &result)
    if err != nil {
        log.Fatal(err)
    }
    //A := "AAAA"
    log.Println(result.Person)
}
