package goods

import (
	. "../"
	. "../../models/goods"
	"log"
)

var db, _ = OpenConnection()

func AddNewGoods(goods *Goods) {
	query := "INSERT INTO goods(name, description) VALUES(?,?)"

	result, err := db.Exec(query,
		goods.Name,
		goods.Description)

	if err != nil {
		log.Fatal("Can't add goods in DB")
	}
	id, _ := result.LastInsertId()
	goods.ID = uint64(id)
}

func GetGoodsById(id uint64) (Goods, error) {
	query := "SELECT id, name, description FROM goods WHERE id=?"

	var goods Goods
	err := db.QueryRow(query, id).Scan(
		&goods.ID,
		&goods.Name,
		&goods.Description)

	if err != nil {
		log.Fatal("Can't get goods by id", id)
		return Goods{}, err
	}
	return goods, err
}

func GetGoodsFromConcreteStorage(storageId uint64) ([]Goods, error) {
	//TODO: прописать таблицу товар-склад
	query := "SELECT goods.id, goods.name, description FROM goods"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Can't get goods")
		return nil, err
	}
	defer rows.Close()

	var goods []Goods
	for rows.Next() {
		var product Goods
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description)
		if err != nil {
			return nil, err
		}
		goods = append(goods, product)
	}
	return goods, err
}

func GetGoods() ([]Goods, error) {
	query := "SELECT id, name, description FROM goods"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Can't get goods")
		return nil, err
	}
	defer rows.Close()

	var goods []Goods
	for rows.Next() {
		var product Goods
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description)
		if err != nil {
			return nil, err
		}
		goods = append(goods, product)
	}
	return goods, err

}

func RemoveGoods(id uint64) error {
	deleteQuery := "DELETE FROM goods WHERE id=?"
	_, err := db.Exec(deleteQuery, id)
	if err != nil {
		log.Fatal("Can't remove goods")
	}
	return nil
}
