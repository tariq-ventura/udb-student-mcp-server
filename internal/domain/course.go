package domain

type Course struct {
	ID        int    `json:"id"`
	FullName  string `json:"fullname"`
	ShortName string `json:"shortname"`
}
