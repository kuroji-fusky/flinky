package schemas

type (
	GenderAlias          string
	FranchiseMediumAlias string
	CharacterTypeAlias   string
)

const (
	MediumMangaComic FranchiseMediumAlias = "manga-comic"
	MediumFilm       FranchiseMediumAlias = "film"
	MediumGame       FranchiseMediumAlias = "game"
	MediumMixed      FranchiseMediumAlias = "mixed"
	MediumOthers     FranchiseMediumAlias = "others"
	MediumSeries     FranchiseMediumAlias = "series"

	CharAliasAlternate   CharacterTypeAlias = "alternate"
	CharAliasNickname    CharacterTypeAlias = "nickname"
	CharAliasTranslation CharacterTypeAlias = "translation"

	GenderMale   GenderAlias = "male"
	GenderFemale GenderAlias = "female"
	GenderOther  GenderAlias = "other"
)
