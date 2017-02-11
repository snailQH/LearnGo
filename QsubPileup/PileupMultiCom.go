//pileupMultiCom is designed for qsub multiple commandlines directly in shell to sub nodes
//test in centos system,you should build the .go before run the app, instead of run by " ... | go run pileupMultiCom.go""
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	hostname, err := os.Hostname()
	check(err)

	user, err := user.Current()
	username := user.Username
	check(err)

	bio := bufio.NewReader(os.Stdin)

	curDir, err := filepath.Abs(filepath.Dir(os.Args[0])) //this line get the dir of this program,not the Current working dir
	check(err)

	x := 0
	for {
		x++
		line, _ := bio.ReadString('\n') //read in the stdin,seperator "\n"
		if line == "" {
			break
		}
		qsub(line, hostname, username, curDir, x)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func qsub(s string, hostname string, user string, curDir string, id int) {
	strid := fmt.Sprintf("%d", id)

	pbsDir := curDir + "/Pileup_pbs"
	err := os.MkdirAll(pbsDir, 0755)
	check(err)

	filename := curDir + "/Pileup_pbs/Pileup_" + strid + ".pbs"
	fmt.Println("qsub ", filename)
	outpbs, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	check(err)
	defer outpbs.Close()
	outWriter := bufio.NewWriter(outpbs)

	//those are the header of pbs file
	strN := "#PBS -N pileupMultiCom" + strid + "\n"
	strO := "#PBS -o " + curDir + "/Pileup_pbs/Pileup_" + strid + ".out" + "\n"
	strE := "#PBS -e " + curDir + "/Pileup_pbs/Pileup_" + strid + ".err" + "\n"
	strNode := "#PBS -l nodes=1:ppn=4" + "\n"
	strR := "#PBS -r y" + "\n"
	strU := "#PBS -u " + user + "\n"
	strQ := "#PBS -q batch" + "\n\n"

	//this is the commandlines
	strC := s + "\n"

	outWriter.WriteString(strN)
	outWriter.WriteString(strO)
	outWriter.WriteString(strE)
	outWriter.WriteString(strNode)
	outWriter.WriteString(strR)
	outWriter.WriteString(strU)
	outWriter.WriteString(strQ)
	outWriter.WriteString(strC)

	outWriter.Flush()
}
