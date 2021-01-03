package redis

import (
	"encoding/json"

	"github.com/mrceylan/go-url-shortener/pkg/url"
)

func (conn *Connection) AddUrl(url url.ShortUrl) error {
	_json, err := url.ToJson()
	if err != nil {
		return err
	}
	err = conn.Client.Set(ctx, url.Hash, _json, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (conn *Connection) GetUrl(hash string) (url.ShortUrl, error) {
	model := url.ShortUrl{}
	val, err := conn.Client.Get(ctx, hash).Result()
	if err != nil {
		return model, err
	}
	model = url.ShortUrl{}
	err = json.Unmarshal([]byte(val), &model)
	if err != nil {
		return model, err
	}
	return model, nil
}
