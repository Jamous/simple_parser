package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//Evil global variabels
var parseFile, matchFile, outFile string

func main() {
	//Parse user input
	parseInput()

	//Import the file to parse through
	parseString := importFile()

	//Import values to parse
	matchSlice := importMatch()

	//Find matches for matchSlice in parString, return outString
	outString := findMatches(parseString, matchSlice)

	//Print results
	fmt.Println(outString)

	//Write outString to file
	writeToFile(outString)
}

func parseInput() {
	//Parese inputs
	pfPtr := flag.String("p", "input.txt", "File to parse from.")
	mfPtr := flag.String("m", "match.txt", "File to match values from.")
	ofPtr := flag.String("o", "output.txt", "File to save results to.")
	flag.Parse()

	parseFile = *pfPtr
	matchFile = *mfPtr
	outFile = *ofPtr
}

func importFile() string {
	//Read file into ifile
	iFile, err := ioutil.ReadFile(parseFile)
	if err != nil {
		log.Fatal("failed to open parseFile")
	}

	//Convert iFile to string
	parseString := string(iFile)

	return parseString
}

func importMatch() []string {
	var matchSlice []string

	//Open file
	mFile, err := os.Open(matchFile)
	if err != nil {
		log.Fatal("failed to open matchFile")
	}

	//Create buffio scanner and split lines
	scanner := bufio.NewScanner(mFile)
	scanner.Split(bufio.ScanLines)

	//Iterate over scanned text
	for scanner.Scan() {
		matchSlice = append(matchSlice, scanner.Text())
	}

	return matchSlice
}

func findMatches(parseString string, matchSlice []string) string {
	outString := ""

	//Split lines and read every line
	lines := strings.Split(parseString, "\n")
	for _, line := range lines {
		//Evaluate line against everyting in matchSlice, if its contained print it
		for _, match := range matchSlice {
			if strings.Contains(line, match) {
				outString += line + "\n"
			}
		}
	}

	return outString
}

func writeToFile(outString string) {
	//setup outfile
	out := makeFile(outFile)
	defer out.Close()

	//Write to file
	_, err := out.WriteString(outString)
	if err != nil {
		log.Fatal("failed to write string to file")
	}
}

func makeFile(fileName string) *os.File {
	//Check if file exists, if it does increment by 1 and try again
	//variabels
	var tempName, finalName string
	ncount := 0

	//Seperate file extension and name
	extension := fileName[strings.Index(fileName, "."):]
	name := fileName[:strings.Index(fileName, ".")]

	//Start a for loop to find an appropiate file name
	for {
		//Set tempName to filename on the first iteration
		if ncount == 0 {
			tempName = fileName
		} else {
			tempName = fmt.Sprintf("%s_%d%s", name, ncount, extension)
		}

		//Check if file exists, an error will indicate it does not, which is what we ant.
		_, err := os.Stat(tempName)
		if err != nil {
			//Set finalName and exit loop
			finalName = tempName
			break
		}

		//Increment ncount and try again
		ncount++
	}

	out, err := os.Create(finalName)
	if err != nil {
		errm := fmt.Sprintf("Failed to create file for %s, error: %s", fileName, err)
		log.Fatal(errm)
	}

	return out
}
