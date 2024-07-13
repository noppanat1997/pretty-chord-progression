package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("error needs 2 arguments to process")
		return
	}

	maxBarSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("error first argument must be an integer")
		return
	}
	inputPath := os.Args[2]
	outputPath := os.Args[3]

	// open input file
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalln("error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// create output file (overwrite old one)
	newFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatalln("error creating new file:", err)
		return
	}
	defer newFile.Close()
	writer := bufio.NewWriter(newFile)

	// read and reformat data in each line
	preserveText := ""
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			fmt.Printf("%s\n", line)
			writer.WriteString(fmt.Sprintf("%s\n", line))
			preserveText = ""
			continue
		}

		if preserveText == "" {
			preserveText = line
			fmt.Printf("%s\n", line)
			writer.WriteString(fmt.Sprintf("%s\n", line))
		} else {
			barLine := strings.Fields(line)
			var bars [][]string
			for _, b := range barLine {
				chords := strings.Split(b, ",")
				bars = append(bars, chords)
			}

			var printBar = make([]string, 0)
			for _, b := range bars {
				var newBar []string

				switch maxBarSize {
				case 2:
					if len(b) < 2 {
						newBar = []string{b[0], ""}
					} else {
						newBar = append(newBar, b...)
					}
					printBar = append(printBar, fmt.Sprintf(" %-6s %-6s ", newBar[0], newBar[1]))
				case 4:
					if len(b) < 4 {
						newBar = []string{b[0], ""}
						switch len(b) {
						case 1:
							newBar = []string{b[0], "", "", ""}
						case 2:
							newBar = []string{b[0], "", b[1], ""}
						case 3:
							newBar = []string{b[0], b[1], b[2], ""}
						}
					} else {
						newBar = append(newBar, b...)
					}
					printBar = append(printBar, fmt.Sprintf(" %-6s %-6s %-6s %-6s ", newBar[0], newBar[1], newBar[2], newBar[3]))
				}
			}

			printLine := fmt.Sprintf("|%s|\n", strings.Join(printBar, "|"))
			fmt.Printf("%s", printLine)
			writer.WriteString(printLine)
		}
	}

	// flush the writer to ensure all buffered data is written to the new file
	if err := writer.Flush(); err != nil {
		log.Fatalln("error flushing writer:", err)
		return
	}
}
