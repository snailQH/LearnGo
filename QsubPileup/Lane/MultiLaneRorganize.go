//pileupMultiCom is designed for qsub multiple commandlines directly in shell to sub nodes
//test in centos system,you should build the .go before run the app, instead of run by " ... | go run pileupMultiCom.go""
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
)

//define the prefix of the dir or run-folder
var prefix = "/online/rawfastq/bcl2fastq/basecallings_scripts/"
var comName = "/bin/sh"

var runID = flag.String("runid", "", "The runid which you want to run the reorganize,like S414")

func main() {
	flag.Parse()

	hostname, err := os.Hostname()
	check(err)

	user, err := user.Current()
	username := user.Username
	check(err)

	//curDir, err := filepath.Abs(filepath.Dir(os.Args[0])) //this line get the dir of this program,not the Current working dir
	//check(err)

	runDir := prefix + *runID
	pbsDir := runDir + " /Pileup_pbs"
	err = os.MkdirAll(pbsDir, 0755)
	check(err)

	args := "grep python /online/rawfastq/bcl2fastq/basecallings_scripts/" + *runID + "/pbs_scripts/*pbs | cut -d: -f 2"
	cmd1 := exec.Command(comName, "-c", args)
	cmdReader, err := cmd1.StdoutPipe()
	check(err)

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println("Ouuuuuuut:\n", scanner.Text())
			lane := scanner.Text()[144:148]
			fw := scanner.Text()[154:157]
			text := *runID + "_" + fw + "_" + lane
			qsubLane(scanner.Text(), hostname, username, runDir, text)
			//qsub(scanner.Text(), hostname, username, runDir, text)
		}
	}()

	err = cmd1.Start()
	check(err)

	err = cmd1.Wait()
	check(err)

	/*
		bio := bufio.NewReader(os.Stdin)
		x := 0
			for {
				x++
				line, _ := bio.ReadString('\n') //read in the stdin,seperator "\n"
				if line == "" {
					break
				}

				qsub(line, hostname, username, runDir, x)
			}
			**/
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func qsubLane(s string, hostname string, user string, runDir string, strid string) {
	filename := runDir + "/Pileup_pbs/Pileup_" + strid + ".pbs"
	outpbs, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	check(err)
	defer outpbs.Close()
	outWriter := bufio.NewWriter(outpbs)

	//those are the header of pbs file
	strN := "#PBS -N pileupMultiCom" + strid + "\n"
	strO := "#PBS -o " + runDir + "/Pileup_pbs/Pileup_" + strid + ".out" + "\n"
	strE := "#PBS -e " + runDir + "/Pileup_pbs/Pileup_" + strid + ".err" + "\n"
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
