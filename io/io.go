// read and write pns files
package top_io

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"top/top"
)

type reading int8

const (
	nothing reading = iota
	poles
	nulls
	seps
	view
)

func ReadPNS(r io.Reader) *top.Topology {
	t := top.NewTopology()
	scanner := bufio.NewScanner(r)
	currently_reading := nothing
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line[0] == '%' {
			continue
		}
		if line[:3] == "END" || line[:5] == "ALPHA" {
			currently_reading = nothing
			continue
		}

		switch currently_reading {
		case nothing:
			switch line[:10] {
			case "BEGIN POLE":
				currently_reading = poles
			case "BEGIN NULL":
				currently_reading = nulls
			case "BEGIN SEPA":
				currently_reading = seps
			case "BEGIN VIEW":
				currently_reading = view
			default:
				fmt.Printf("Unexpected: %s\n", line)
			}
		case poles:
			pole, err := top.PoleFromString(line)
			if err != nil {
				fmt.Printf("Error reading pole: %v\n", err)
			} else {
				t.Poles[pole.Name] = pole
			}
		case nulls:
			null, err := top.NullFromString(line)
			if err != nil {
				fmt.Printf("Error reading null: %v\n", err)
			} else {
				t.Nulls[null.Name] = null
			}
		case seps:
			// seps := top.SepFromString(line)
			// t.Separators = append(t.Separators, seps)
		}

	}
	return t
}
