package schemas

type EntityURL struct {
	URL string `json:"url"`
}

type EntityRef struct {
	EntityURL
	AvatarURL   string `json:"avatar_url"`
	EndpointURL string `json:"endpoint_url"`
}

type EntityImage struct {
	EntityURL
	Alt *string `json:"alt"`
}
