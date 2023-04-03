package jsonDb

import (
	"app/models"
	"errors"
	"sort"
)

func (s *shopCartRepo) FilterProductsDate(req *models.User) ([]models.ShopCart, error) {
	shopCarts, err := s.Read()
	if err != nil {
		return []models.ShopCart{}, err
	}

	userShopCarts := []models.ShopCart{}

	for _, v := range shopCarts {
		if v.Date >= req.FromDate && v.Date <= req.ToDate {
			userShopCarts = append(userShopCarts, v)
		}
	}

	sortedDate := []string{}
	for _, v := range userShopCarts {
		sortedDate = append(sortedDate, v.Date)
	}
	sort.Strings(sortedDate)

	lastMap := []models.ShopCart{}
	i := 0
	for i != len(sortedDate)-1 {
		for _, v := range userShopCarts {
			if sortedDate[i] == v.Date {
				lastMap = append(lastMap, v)
				i += 1
			}
		}
	}

	return lastMap, nil
}

func (u *userRepo) GetByNam(req *models.User) (models.User, error) {
	users, err := u.Read()
	if err != nil {
		return models.User{}, err
	}

	for _, v := range users {
		if v.Name == req.Name {
			return v, nil
		}
	}

	return models.User{}, errors.New("There is no user with this data!")
}