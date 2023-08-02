package models

type Category struct {
	CategoryId string `json:"CategoryId" validate:"required"`
	Color      int64  `json:"Color"      validate:"required"`
	Label      string `json:"Label"      validate:"required"`
}

type DeletingCategory struct {
	CategoryId string `json:"CategoryId" validate:"required"`
}
