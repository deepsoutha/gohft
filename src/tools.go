package src

import (
	"encoding/csv"
	"os"
)

func WriteToCSV(fileName string, sliceOfSlice [][]string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	var rows [][]string
	for _, s := range sliceOfSlice {
		row := make([]string, len(s))
		copy(row, s)
		rows = append(rows, row)
	}
	writer := csv.NewWriter(file)
	writer.WriteAll(rows)
	writer.Flush()
	return nil
}
