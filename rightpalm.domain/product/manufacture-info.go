package product

import "time"

type ManufactureInfo struct {
	ID                int
	ManufactureID     int
	LanguageID        int
	Url               string
	UrlClicked        string
	LastedDateClicked time.Time
	CreatedDate       time.Time
}
