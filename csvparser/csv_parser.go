package csvparser

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ParseCSV(fileName string, fileFolder string) (*[][]string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	csvFile, err := os.Open(rootDir + "/" + fileFolder + "/" + fileName)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	var data [][]string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
		}

		data = append(data, line)
	}

	return &data, nil
}
