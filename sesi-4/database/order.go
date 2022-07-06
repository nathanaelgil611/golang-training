package database

import (
	"sesi-4/model"
	"time"
)

func CreateOrder(order model.Order) error {
	var o = model.Order{}

	sqlStatementOrder := `INSERT INTO orders (customer_name) VALUES($1) RETURNING order_id`
	err := db.QueryRow(sqlStatementOrder, order.CustomerName).Scan(&o.OrderID)

	sqlStatementItem := `INSERT INTO items (item_code, description, quantity, order_id) VALUES($1, $2, $3, $4)`
	for _, item := range order.Items {
		_, err = db.Exec(sqlStatementItem, item.ItemCode, item.Description, item.Quantity, o.OrderID)
	}

	if err != nil {
		panic(err)
	}
	return err
}

func GetAllOrder() ([]model.Order, error) {
	var orderList = make(map[int]model.Order)

	sqlStatement := `SELECT o.*, i.* FROM "orders" as o INNER JOIN "items" as i ON o.order_id=i.order_id`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var temp struct {
		OrderID      int       `json:"order_id"`
		CustomerName string    `json:"customer_name"`
		OrderedAt    time.Time `json:"ordered_at"`
		ItemID       int       `json:"item_id"`
		ItemCode     string    `json:"item_code"`
		Description  string    `json:"description"`
		Quantity     int       `json:"quantity"`
	}
	for rows.Next() {

		err := rows.Scan(&temp.OrderID, &temp.CustomerName, &temp.OrderedAt, &temp.ItemID, &temp.ItemCode, &temp.Description, &temp.Quantity, &temp.OrderID)
		if err != nil {
			panic(err)
		}

		if order, ok := orderList[temp.OrderID]; !ok {
			orderList[temp.OrderID] = model.Order{
				OrderID:      temp.OrderID,
				CustomerName: temp.CustomerName,
				OrderedAt:    temp.OrderedAt,
				Items:        []model.Item{model.Item{temp.ItemID, temp.ItemCode, temp.Description, temp.Quantity, temp.OrderID}},
			}
			// orderList[temp.OrderID].Items = append(orderList[temp.OrderID].Items, model.Item)
		} else {
			order.Items = append(order.Items, model.Item{temp.ItemID, temp.ItemCode, temp.Description, temp.Quantity, temp.OrderID})
			orderList[temp.OrderID] = order
		}

	}

	data := make([]model.Order, 0, len(orderList))

	for _, value := range orderList {
		data = append(data, value)
	}

	return data, err
}

func GetOrder(id int) (model.Order, error) {
	var orderList = make(map[int]model.Order)

	var temp struct {
		OrderID      int       `json:"order_id"`
		CustomerName string    `json:"customer_name"`
		OrderedAt    time.Time `json:"ordered_at"`
		ItemID       int       `json:"item_id"`
		ItemCode     string    `json:"item_code"`
		Description  string    `json:"description"`
		Quantity     int       `json:"quantity"`
	}

	sqlStatement := `SELECT o.*, i.* FROM "orders" as o INNER JOIN "items" as i ON o.order_id=i.order_id WHERE o.order_id = $1`
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&temp.OrderID, &temp.CustomerName, &temp.OrderedAt, &temp.ItemID, &temp.ItemCode, &temp.Description, &temp.Quantity, &temp.OrderID)
		if err != nil {
			panic(err)
		}

		if order, ok := orderList[temp.OrderID]; !ok {
			orderList[temp.OrderID] = model.Order{
				OrderID:      temp.OrderID,
				CustomerName: temp.CustomerName,
				OrderedAt:    temp.OrderedAt,
				Items:        []model.Item{model.Item{temp.ItemID, temp.ItemCode, temp.Description, temp.Quantity, temp.OrderID}},
			}
			// orderList[temp.OrderID].Items = append(orderList[temp.OrderID].Items, model.Item)
		} else {
			order.Items = append(order.Items, model.Item{temp.ItemID, temp.ItemCode, temp.Description, temp.Quantity, temp.OrderID})
			orderList[temp.OrderID] = order
		}

	}

	return orderList[id], err
}

func UpdateOrder(order *model.Order, id int) error {
	sqlStatement := `UPDATE orders SET customer_name = $1 WHERE order_id = $2`
	_, err := db.Exec(sqlStatement, order.CustomerName, id)

	for _, item := range order.Items {
		sqlStatementItem := `UPDATE items SET item_code = $1, description = $2, quantity = $3 WHERE order_id = $4 AND item_id = $5`
		_, err = db.Exec(sqlStatementItem, item.ItemCode, item.Description, item.Quantity, id, item.ItemID)
	}

	return err
}

func DeleteOrder(id int) error {

	sqlStatementItem := `DELETE FROM items WHERE order_id = $1`
	_, err := db.Exec(sqlStatementItem, id)
	if err != nil {
		panic(err)

	}

	sqlStatement := `DELETE FROM orders WHERE order_id = $1`
	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)

	}

	return err
}
