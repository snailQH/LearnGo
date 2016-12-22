//renameFileForJN is written for genenergy to rename the fastq file from Cloudhealth Genomics
//from:S398_02A_CHG016568-1211S1-ZY254-20161211-ATAGCG-ATAGCG_L007_R2.fastq.gz
//to:ZY254-20161211-ATAGCG_L007_R2.fastq.gz

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
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
			if strings.Contains(file.Name(), "Undetermined") {
				continue
			}
			x := folder + "/" + file.Name()
			y := ""
			for i,v := range strings.Split(file.Name(),"-") {
				if i == 2 {
					y = v
				}else if (i > 2 ){
					y += "-" + v 
				}
			}
			y = folder + "/" + y
			fmt.Println(">Renaming: " + x + "\t" + y)
			os.Rename(x,y)
		}
  }

}

