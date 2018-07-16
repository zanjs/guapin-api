package model

//SearchQ  type contains page info
type SearchQ struct {
	IDAutoModel
	TitleModel
	NameModel
	CategoryIDModel
	DescriptionModel
	PictureModel
	ContentModel
	TimeAllModel
	Category Category `json:"categories"`
}
