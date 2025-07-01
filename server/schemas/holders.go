package schemas

type HolderResponse struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	EntityRef
}
