package parser

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	scannerBuffer = 128 * 1024
)

// Parse read content from a io.Reader and returns the YAML tree.
func Parse(reader io.Reader, blockStart, blockEnd string) (interface{}, error) {
	front, err := fetchYAML(reader, blockStart, blockEnd)
	if err != nil {
		return nil, err
	}

	var res interface{}
	err = yaml.Unmarshal(front, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func fetchYAML(reader io.Reader, blockStart, blockEnd string) ([]byte, error) {
	buffer := make([]byte, scannerBuffer)

	scanner := bufio.NewScanner(reader)
	scanner.Buffer(buffer, scannerBuffer)

	res := &bytes.Buffer{}

	delimited := (len(blockStart) > 0) && (len(blockEnd) > 0)

	ln := 0
	for scanner.Scan() {
		line := scanner.Text()

		if delimited {
			if ln == 2 {
				break
			}
		}

		tmp := strings.TrimSpace(line)
		if delimited {
			if tmp == blockStart || tmp == blockEnd {
				ln++
				continue
			}
		}

		res.WriteString(line)
		res.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}
