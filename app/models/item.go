package models

var Fillable = [] string{
	"site",
	"price",
	"start_time",
	"name",
	"description",
	"nickname",
}

type Item struct {
	ID			uint64
	Site        string
	Price       string
	StartTime   string
	Name        string
	Description string
	Nickname    string
}
