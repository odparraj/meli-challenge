package services

import (
	"encoding/json"
	"fmt"
	ML "goravel/contracts/services"
	"net/http"
	"strings"
)

var baseUrl = "https://api.mercadolibre.com"
var currencies = make(map[string]string)

type MeliApiService struct {
}

func init() {
	loadCurrencies()
}

func loadCurrencies() {
	curreciesToLoad, err := getCurrencies()

	if err != nil {
		panic("Error loading currencies")
	}

	for _, currency := range curreciesToLoad {
		currencies[currency.ID] = currency.Description
	}
}

func (meli *MeliApiService) GetItems(itemIds []string) (ML.Items, error) {
	attributes := "id,price,start_time,category_id,currency_id,seller_id,site_id"
	url := fmt.Sprintf("%s/items?ids=%s&attributes=%s", baseUrl, strings.Join(itemIds, ","), attributes)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	items := ML.Items{}
	json.NewDecoder(res.Body).Decode(&items)
	return items, nil
}

func (meli *MeliApiService) GetCategory(categoryId string) (*ML.Category, error) {
	attributes := "id,name"
	url := fmt.Sprintf("%s/categories/%s?attributes=%s", baseUrl, categoryId, attributes)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	category := &ML.Category{}
	json.NewDecoder(res.Body).Decode(category)
	return category, nil
}

func getCurrencies() (ML.Currencies, error) {
	attributes := "id,description"
	url := fmt.Sprintf("%s/currencies?attributes=%s", baseUrl, attributes)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	currencies := ML.Currencies{}
	json.NewDecoder(res.Body).Decode(&currencies)
	return currencies, nil
}

func (meli *MeliApiService) GetCurrencies() (ML.Currencies, error) {
	return getCurrencies()
}

func (meli *MeliApiService) GetUser(userId string) (*ML.User, error) {
	attributes := "id,nickname"
	url := fmt.Sprintf("%s/users/%s?attributes=%s", baseUrl, userId, attributes)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	user := &ML.User{}
	json.NewDecoder(res.Body).Decode(user)
	return user, nil
}

func (meli *MeliApiService) GetDataItem(itemId string) (*ML.DataItem, error) {
	items, err := meli.GetItems([]string{itemId})

	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, nil
	}

	if items[0].Status != 200 {
		return nil, fmt.Errorf("item not found [%s]", itemId)
	}

	return meli.dataItemResolver(items[0])
}

func (meli *MeliApiService) dataItemResolver(item ML.ItemResponse) (*ML.DataItem, error) {

	category, err := meli.GetCategory(item.Data.CategoryId)
	if err != nil {
		return nil, err
	}

	user, err := meli.GetUser(fmt.Sprintf("%d", item.Data.SellerId))
	if err != nil {
		return nil, err
	}

	return &ML.DataItem{
		ID:          item.Data.ID,
		Name:        category.Name,
		Nickname:    user.Nickname,
		StartTime:   item.Data.StartTime,
		Site:        item.Data.SiteId,
		Description: currencies[item.Data.CurrencyId],
		Price:       fmt.Sprintf("%f", item.Data.Price),
	}, nil
}
