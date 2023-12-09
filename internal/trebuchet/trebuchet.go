package trebuchet

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

func GetDocumentAndCalibrate(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var calibrationValue int = 0
	var lines int = 1

	for scanner.Scan() {
		line := scanner.Text()
		var key int8 = 0

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				key += int8(line[i] - '0')
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				key = key*10 + int8(line[i]-'0')
				break
			}
		}
		// fmt.Println("line: ", lines, ", key: ", key)
		calibrationValue += int(key)
		lines++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return calibrationValue
}
