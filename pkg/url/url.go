package url

import (
	"encoding/json"
)

type ShortUrl struct {
	Hash string
	Url  string
}

func (_s ShortUrl) ToJson() (string, error) {
	b, err := json.Marshal(_s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
