//echo3, the simple one
package main

import(
	"os"
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Join(os.Args[0:],";"))
	for i := 1;i<len(os.Args);i++{
		fmt.Println(i,"\t",os.Args[i])
	}
}

