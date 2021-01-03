package url

type URLStorage interface {
	AddUrl(url ShortUrl) error
	GetUrl(hash string) (ShortUrl, error)
}

type URLService interface {
	CreateNewURL(url string) (string, error)
	GetURL(hash string) (string, error)
}
