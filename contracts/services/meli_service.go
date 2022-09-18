package services

type DataItem struct {
	ID          string
	Site        string
	Price       string
	StartTime   string
	Name        string
	Description string
	Nickname    string
}

type Currency struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Currencies []Currency

type ItemResponse struct {
	Status int  `json:"code"`
	Data   Item `json:"body"`
}

type Item struct {
	ID         string  `json:"id"`
	SiteId     string  `json:"site_id"`
	CurrencyId string  `json:"currency_id"`
	CategoryId string  `json:"category_id"`
	SellerId   uint64  `json:"seller_id"`
	Price      float64 `json:"price"`
	StartTime  string  `json:"start_time"`
}

type Items []ItemResponse

type User struct {
	ID       uint64 `json:"id"`
	Nickname string `json:"nickname"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MeliApiServiceInterface interface {
	GetItems(itemIds []string) (Items, error)
	GetCategory(categoryId string) (*Category, error)
	GetUser(userId string) (*User, error)
	GetDataItem(itemId string) (*DataItem, error)
	GetCurrencies() (Currencies, error)
}
