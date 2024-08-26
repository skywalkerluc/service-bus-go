package inventory

import (
	"fmt"
	"service-bus/pkg/models"
	"service-bus/pkg/utils"
)

func HandleOrder(orderData []byte) {
	var order models.Order
	if err := utils.FromJSON(string(orderData), &order); err != nil {
		fmt.Printf("error when desserializing order %v\n", err)
		return
	}

	fmt.Printf("adjusting inventory to order %s, product %s, quantity %d\n",
		order.OrderID, order.ProductID, order.Quantity)
}
