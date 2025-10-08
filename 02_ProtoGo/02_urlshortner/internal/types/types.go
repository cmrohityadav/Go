package types

type UrlNode struct{
	ID string
	OriginalUrl string
	ShortUrl string
}

type CreateShortUrlRequest struct{
	URL string `json:"url"`
}