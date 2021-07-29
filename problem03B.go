/* Problem 03-A
 *
 * Answer
 */

package main

import (
	"bufio"
	"log"
	"os"
)

func isTreeAt(col int, mapLine string) bool {
	return mapLine[col] == '#'
}

func problem03B(fileName string) int {

	// Open data file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Walk the map
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	col := 0
	treesHit := 0
	for scanner.Scan() {
		mapText := scanner.Text()
		//fmt.Printf("%3d-%s\n", col, mapText)
		if mapText[col] == '#' {
			treesHit++
		}
		col = (col + 3) % len(mapText)
	}

	return treesHit
}
