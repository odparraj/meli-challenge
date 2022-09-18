package services

type Decoder interface {
	Decode(headers []string, line string, delimiter string) (map[string]string, error)
}