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

	CharAliasAlternate   CharacterTypeAlias = "ALTERNATE"
	CharAliasNickname    CharacterTypeAlias = "NICKNAME"
	CharAliasTranslation CharacterTypeAlias = "TRANSLATION"
	CharAliasMononym     CharacterTypeAlias = "MONONYM"

	GenderMale   GenderAlias = "male"
	GenderFemale GenderAlias = "female"
	GenderOther  GenderAlias = "other"
)
