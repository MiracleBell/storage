package customers

import (
	. "../../models/detail"
	"log"
)

func SaveBankDetails(bankDetails *BankDetails) error {
	query := "INSERT INTO bank_details(recipient, recipient_bank_name, recipient_bank_address) VALUES(?,?,?)"

	result, err := db.Exec(query,
		bankDetails.Recipient,
		bankDetails.RecipientsBankName,
		bankDetails.RecipientsBankAddress)

	if err != nil {
		log.Fatal("Can't save bank details in DB")
		return err
	}
	id, _ := result.LastInsertId()
	bankDetails.ID = uint64(id)
	return nil
}
