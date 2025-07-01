package schemas

// The juicy stuff

type CharacterResponse struct {
	Name             string         `json:"name"`
	Id               int            `json:"id"`
	Description      string         `json:"description"`
	Sources          *[]string      `json:"sources,omitempty"` // Sources are optional but recommended for verifying things, you bitches
	Tags             []string       `json:"tags"`
	IsLocked         bool           `json:"is_locked"`
	IsNSFW           bool           `json:"is_nsfw"`
	VerifiedByAuthor *[]interface{} `json:"verified_by_author,omitempty"`

	Image EntityImage `json:"img"`

	Gallery []struct {
		EntityImage
		Source  string       `json:"source"`
		AddedBy *AddedByUser `json:"added_by"`
	} `json:"gallery"`

	Citations []struct {
		Source  string      `json:"source"`
		Url     string      `json:"url"`
		AddedBy AddedByUser `json:"added_by"`
	} `json:"citations"`

	Attributes struct {
		Gender GenderAlias `json:"gender"`
		Traits []string    `json:"traits"`

		Franchise struct {
			Name   string               `json:"name"`
			Year   int                  `json:"year"`
			Medium FranchiseMediumAlias `json:"medium"`

			Holder []struct {
				Name string `json:"name"`
				EntityRef
			} `json:"holder"`
		} `json:"franchise"`

		Quotes []struct {
			Quote   string       `json:"quote"`
			AddedBy *AddedByUser `json:"added_by"`
		} `json:"quotes"`

		Aliases *[]struct {
			Type CharacterTypeAlias `json:"type"`
			Name string             `json:"name"`
			Lang *string            `json:"lang,omitempty"`
		} `json:"aliases,omitempty"`

		VoiceActor *[]struct {
			Name string
			As   string
			EntityRef
		} `json:"voice_actor,omitempty"`

		Others []any `json:"others"`

		// Normally an object type, but this could theortically mold into whatever type if a character
		// has a complex family tree
		FamilyTree *[]interface{} `json:"family_tree,omitempty"`
	} `json:"attributes"`

	AddedBy AddedByUser `json:"added_by"`
}
