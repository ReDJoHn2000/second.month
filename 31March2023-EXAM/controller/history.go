package controller

import (
	"app/models"
	"fmt"
)

func (c *Controller) GetHistory(req *models.HistoryProd) {
	id := ""
	user, err := c.store.User().GetByName(req)
	if err != nil {
		fmt.Println(err)
	}
	
	*&id = user.Id

	productPurchase, err := c.store.ShopCart().GetUserShopCart(&models.UserPrimaryKey{id})

	fmt.Println(productPurchase)
}

func (c *Controller) FilterByDate(req *models.HistoryProd)([]models.ShopCart) {
	data, err := c.store.ShopCart().FilterProductsByDate(&models.HistoryProd{FromDate: req.FromDate, ToDate: req.ToDate})
	if err != nil {
		fmt.Println("Unable to show by date: ",err)
		return data
	}
	return data
}
