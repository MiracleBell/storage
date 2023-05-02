package storages

import (
	. "../"
	. "../../models/goods"
	. "../../models/storage"
	"log"
)

var db, _ = OpenConnection()

func AddProductToStorage(storage Storage, goods Goods) error {
	//TODO: см. todo goods
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
