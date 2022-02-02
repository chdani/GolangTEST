package database

import "GetGoApi/model"

func (d *DB) GetOrders() ([]*model.Order, error) {

	var orders []*model.Order
	err := d.db.Select(&orders, "SELECT  o.order_name,oi.product,cc.company_name,cs.name as Customer_Name, o.created_at ,  NULLIF( d.delivered_quantity, '')::int * NULLIF(oi.price_per_unit, '')::float4   as Delivered_Amount ,   NULLIF(oi.price_per_unit, '')::float4 *  NULLIF(oi.quantity, '')::int  as Total_Amount FROM public.orders o left join public.order_items oi on oi.order_id = o.id left join public.deliveries d on d.order_item_id = oi.id left join public.customers cs on cs.user_id = o.customer_id left join public.customer_companies  cc on cc.company_id =  cs.company_id")

	if err != nil {
		return orders, err

	}

	return orders, nil

}
