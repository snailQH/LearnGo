//listFile is aimming to list all files in the specific directory
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"io"
)

func main(){
	for _ , folder := range os.Args[1:] {//specify the current dir
		listFile(folder)
	}
}

func listFile(folder string){
	files, _ := ioutil.ReadDir(folder) 
	for _,file := range files{
  	if file.IsDir(){
  		listFile(folder + "/" + file.Name())
    }else{
    	fmt.Println(folder + "/" + file.Name())
    	readFile(folder + "/" + file.Name())
    	x := folder + "/" + file.Name()
    	y := folder + "/"+ "helloWorld.go"
    	if file.Name() == "test.txt" {
    		os.Rename(x,y)
    	}
    }
   }
}

func readFile(file string) {
    inputFile, inputError := os.Open(file)
    if inputError != nil {
        fmt.Printf("An error occurred on opening the inputfile\n" +
            "Does the file exist?\n" +
            "Have you got acces to it?\n")
        return // exit the function on error
    }
    defer inputFile.Close()

    inputReader := bufio.NewReader(inputFile)
    for {
        inputString, readerError := inputReader.ReadString('\n')
        if readerError == io.EOF {
            return
        }
        fmt.Printf("The input was: %s", inputString)
    }
}
