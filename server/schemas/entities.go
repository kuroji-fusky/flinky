package schemas

type Profile struct {
	Username string `json:"username"`
	Handle   string `json:"handle"`
	JoinDate string `json:"join_date"`
	IsAdmin  bool   `json:"is_admin"`
	EntityRef
}

type AddedByUser struct {
	ApprovalDate   string  `json:"approval_date"`
	Status         string  `json:"status"`
	RejectedReason *string `json:"rejected_reason"`
	Profile
}

type EntityRef struct {
	URL         string      `json:"url"`
	AvatarURL   EntityImage `json:"avatar_url"`
	EndpointURL string      `json:"endpoint_url"`
}

type EntityImage struct {
	URL string  `json:"url"`
	Alt *string `json:"alt"`
}
