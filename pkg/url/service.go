package url

import (
	"crypto/rand"
	"fmt"
)

type urlService struct {
	storage URLStorage
}

func InitService(st URLStorage) URLService {
	return &urlService{st}
}

func (urlSrv *urlService) CreateNewURL(url string) (string, error) {
	hash, err := generateHash()
	if err != nil {
		return "", err
	}
	_model := ShortUrl{Url: url, Hash: hash}
	err = urlSrv.storage.AddUrl(_model)
	if err != nil {
		return "", err
	}
	return _model.Hash, nil
}

func (urlSrv *urlService) GetURL(hash string) (string, error) {
	result, err := urlSrv.storage.GetUrl(hash)
	if err != nil {
		return "", err
	}
	return result.Url, nil
}

func generateHash() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", nil
	}
	uuid := fmt.Sprintf("%x", b[0:6])

	return uuid, nil
}
