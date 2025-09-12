package models

type Category struct {
	CategoryID   int64   `json:"category_id" db:"CategoryID"`     // INTEGER, Auto Increment
	CategoryName *string `json:"category_name" db:"CategoryName"` // TEXT, nullable
	Description  *string `json:"description" db:"Description"`    // TEXT, nullable
	Picture      []byte  `json:"picture" db:"Picture"`            // BLOB, nullable
}
