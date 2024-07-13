package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("error needs 3 arguments to process")
		return
	}

	maxBarCount, err := strconv.Atoi(os.Args[1])
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

	// map input file into data
	var sectionMap = make(map[string][][]string)
	var maxBarSize int = 0
	currentTitle := ""
	var header string = ""
	for scanner.Scan() {
		line := scanner.Text()

		if header == "" && line != "" {
			header = strings.TrimSpace(line)
			continue
		}

		if line == "" {
			currentTitle = ""
			continue
		}

		if currentTitle == "" {
			currentTitle = strings.TrimSpace(line)
			sectionMap[currentTitle] = [][]string{}
		} else {
			barText := strings.Fields(line)
			var bars [][]string
			for _, b := range barText {
				chords := strings.Split(strings.TrimSpace(b), ",")
				if len(chords) > maxBarSize {
					maxBarSize = len(chords)
				}
				bars = append(bars, chords)
			}
			sectionMap[currentTitle] = append(sectionMap[currentTitle], bars...)
		}
	}
	// fmt.Printf("%v %d", sectionMap, maxBarSize)

	// map data into output file
	fmt.Printf("%s\n\n", header)
	writer.WriteString(fmt.Sprintf("%s\n\n", header))
	for t, c := range sectionMap {
		fmt.Printf("%s\n", t)
		writer.WriteString(fmt.Sprintf("%s\n", t))

		maxRow := int(math.Ceil((float64(len(c)) / float64(maxBarCount))))

		for i := 0; i < maxRow; i++ {
			var printBar = make([]string, 0)

			start := i * maxBarCount
			end := start + maxBarCount
			if end >= len(c) {
				end = len(c)
			}
			row := c[start:end]
			for _, col := range row {
				var newCol []string

				switch maxBarSize {
				case 2:
					if len(col) < 2 {
						newCol = []string{col[0], ""}
					} else {
						newCol = append(newCol, col...)
					}
					printBar = append(printBar, fmt.Sprintf(" %-6s %-6s ", newCol[0], newCol[1]))
				case 4:
					if len(col) < 4 {
						newCol = []string{col[0], ""}
						switch len(col) {
						case 1:
							newCol = []string{col[0], "", "", ""}
						case 2:
							newCol = []string{col[0], "", col[1], ""}
						case 3:
							newCol = []string{col[0], col[1], col[2], ""}
						}
					} else {
						newCol = append(newCol, col...)
					}
					printBar = append(printBar, fmt.Sprintf(" %-6s %-6s %-6s %-6s ", newCol[0], newCol[1], newCol[2], newCol[3]))
				}
			}

			printLine := fmt.Sprintf("|%s|\n", strings.Join(printBar, "|"))
			fmt.Printf("%s", printLine)
			writer.WriteString(printLine)
		}
		fmt.Printf("\n")
		writer.WriteString("\n")
	}

	// flush the writer to ensure all buffered data is written to the new file
	if err := writer.Flush(); err != nil {
		log.Fatalln("error flushing writer:", err)
		return
	}
}
