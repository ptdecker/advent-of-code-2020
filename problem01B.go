/* Problem 01-B
 *
 * The Elves in accounting are thankful for your help; one of them even offers
 * you a starfish coin they had left over from a past vacation. They offer you
 * a second one if you can find three numbers in your expense report that meet
 * the same criteria.
 *
 * Using the above example again, the three entries that sum to 2020 are 979,
 * 366, and 675. Multiplying them together produces the answer, 241861950.
 *
 * In your expense report, what is the product of the three entries that sum to 2020?
 */

 package main

 import (
	 "os"
	 "log"
	 "bufio"
	 "strconv"
 )
 
 func problem01B(fileName string) int {
 
	 var values []int64
 
	 // Open data file
	 file, err := os.Open(fileName)
	 if err != nil {
		 log.Fatal(err)
	 }
	 defer file.Close()
 
	 // Read data
	 scanner := bufio.NewScanner(file)
	 scanner.Split(bufio.ScanLines)
	 for scanner.Scan() {
		 value, err := strconv.ParseInt(scanner.Text(), 10, 0)
		 if err != nil {
			 log.Fatal(err)
		 }
		 values = append(values, value)
	 }
 
	 for i := 0; i < (len(values) - 2); i++ {
		 for j := (i + 1); j < (len(values) - 1); j++ {
			 for k := (j + 1); k < len(values); k++ {
			 	if (values[i] + values[j] + values[k] == 2020) {
				 	return (int)(values[i] * values[j] * values[k])
			 	}
			}
		 }
	 }
 
	 return 0
 }
