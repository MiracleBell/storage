package orders

import (
	. "../"
	. "../../models/orders"
	"log"
)

var db, _ = OpenConnection()

func CreateNewOrder(order Order) error {
	query := "INSERT INTO orders(status, customer_id) VALUES(?,?)"
	secondQuery := "INSERT INTO orders_goods(order_id, goods_id, quantity) VALUES "

	result, err := db.Exec(query,
		order.OrderStatus.ID,
		order.Customer.ID)

	if err != nil {
		log.Fatal("Can't add goods in DB")
	}
	id, _ := result.LastInsertId()

	// Формируем запрос на вставку нескольких строк
	vals := []interface{}{}
	for _, orderGoods := range order.Card {
		secondQuery += "(?, ?, ?),"
		vals = append(vals, id, orderGoods.Goods.ID, orderGoods.Count)
	}
	secondQuery = secondQuery[:len(secondQuery)-1]

	// Выполняем запрос
	_, err = db.Exec(secondQuery, vals...)
	order.ID = uint64(id)

	return nil
}

func GetOrderById(id uint64) (Order, error) {
	query := "SELECT id, status, customer_id,  FROM orders WHERE id=?" +
		"JOIN orders_goods ON orders_goods.oreder_id=oreders.id" +
		"JOIN goods ON orders_goods.goods_id=goods.id"

	var order Order
	err := db.QueryRow(query, id).Scan(
		&order.ID,
		&order.OrderStatus.Name,
		&order.Customer.ID)

	if err != nil {
		log.Fatal("Can't get goods by id", id)
		return Order{}, err
	}
	return order, err
}
