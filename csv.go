package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

func sum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

type statsFunc func([]float64) float64

func csv2float(r io.Reader, column int) ([]float64, error) {
	cR := csv.NewReader(r)
	cR.ReuseRecord = true
	column--
	var data []float64
	for i := 0; ; i++ {
		row, err := cR.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Cannot read data from file: %w", err)
		}
		if i == 0 {
			continue
		}
		if len(row) <= column {
			return nil,
				fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))
		}
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}
		data = append(data, v)
	}
	return data, nil
}
