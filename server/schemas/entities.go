package schemas

type EntityRef struct {
	URL         string      `json:"url"`
	AvatarURL   EntityImage `json:"avatar_url"`
	EndpointURL string      `json:"endpoint_url"`
}

type EntityImage struct {
	URL string  `json:"url"`
	Alt *string `json:"alt"`
}
