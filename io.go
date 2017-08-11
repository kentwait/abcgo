package abcgo

import (
	"bytes"
	"os"
	"strconv"
)

func ToDelimited(samples [][]float64, delimiter, path string) error {
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
	buff.WriteTo(f)

	return nil
}

func ToTab(samples [][]float64, path string) error {
	return ToDelimited(samples, "\t", path)
}

func ToCsv(samples [][]float64, path string) error {
	return ToDelimited(samples, ",", path)
}
