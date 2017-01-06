//echo3, the simple one
package main

import(
	"os"
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Join(os.Args[0:],";"))
}

