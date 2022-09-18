package services

import (
	"fmt"
	"strings"
)

type CsvDecoder struct {
}

func (csvDecoder *CsvDecoder) Decode(headers []string, line string, delimiter string) (map[string]string, error) {
	if delimiter == "" {
		delimiter = ","
	}

	if len(headers) == 0 {
		return nil, fmt.Errorf("empty headers")
	}

	split := strings.Split(line, delimiter)

	if len(split) != len(headers) {
		return nil, fmt.Errorf("missing value from headers: %s", strings.Join(headers, ","))
	}

	res := make(map[string]string)
	for i, h := range headers {
		res[h] = split[i]
	}
	return res, nil
}

type JsonlnDecoder struct {
}

func (jsonlnDecoder *JsonlnDecoder) Decode(headers []string, line string, delimiter string) (map[string]string, error) {

	if len(headers) == 0 {
		return nil, fmt.Errorf("empty headers")
	}

	toReplace := []string{"[", "]", `"`, "'", " "}
	s := line

	for _, c := range toReplace {
		s = strings.ReplaceAll(s, c, "")
	}

	split := strings.Split(s, ",")

	if len(split) != len(headers) {
		return nil, fmt.Errorf("missing value from headers: %s", strings.Join(headers, ","))
	}

	res := make(map[string]string)
	for i, h := range headers {
		res[h] = split[i]
	}
	return res, nil
}
