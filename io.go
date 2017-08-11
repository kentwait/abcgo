package abcgo

import (
	"bytes"
	"os"
	"strconv"
)

func FromDelimited(path, delimiter string) [][]float64 {

}

// ToDelimited writes a float64 2D slice to disk as a text file delimited
// by the given delimiter string.
func ToDelimited(samples [][]float64, path, delimiter string) error {
	var buff bytes.Buffer
	for i := 0; i < len(samples); i++ {
		for j := 0; j < len(samples[i])-1; j++ {
			buff.WriteString(strconv.FormatFloat(samples[i][j], 'f', 8, 64) + delimiter)
		}
		buff.WriteString(strconv.FormatFloat(samples[i][len(samples[i])], 'f', 8, 64) + "\n")
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	buff.WriteTo(f)

	return nil
}

// ToTab writes a float64 2D slice to disk as a tab-delimited text file.
func ToTab(samples [][]float64, path string) error {
	return ToDelimited(samples, "\t", path)
}

// ToCsv writes a float64 2D slice to disk as a CSV file.
func ToCsv(samples [][]float64, path string) error {
	return ToDelimited(samples, ",", path)
}
