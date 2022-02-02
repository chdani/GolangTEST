package app

import (
	"GetGoApi/model"
	"fmt"
	"log"
	"net/http"
)

func (a *App) IndexHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Welcome to Post Api")

	}

}

func mapOrderToJson(p *model.Order) model.JsonOrder {

	return model.JsonOrder{
		Order_name:       p.Order_name,
		Product:          p.Product,
		Company_name:     p.Company_name,
		Customer_Name:    p.Customer_Name,
		Delivered_Amount: p.Delivered_Amount,
		Total_Amount:     p.Total_Amount,
	}

}

func (a *App) GetOrderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		Order, err := a.DB.GetOrders()

		if err != nil {

			log.Printf("Cannot parse post body, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]model.JsonOrder, len(Order))
		for idx, Order := range Order {

			resp[idx] = mapOrderToJson(Order)

		}
		sendResponse(w, r, resp, http.StatusOK)
	}

}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := model.OrderRequest{}
		err := parse(w, r, &req)
		if err != nil {

			log.Printf("Cannot parse post body, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
		}

	}

}
