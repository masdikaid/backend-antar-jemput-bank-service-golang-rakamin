package repository

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/databases"
	//"backend-a-antar-jemput/internal/models"
	//"backend-a-antar-jemput/internal/models"
)

func GetAllTransactions() ([]contract.Transaction,error) {
	// var trx []models.Transaction
	// databases.DBCon.Find(&trx)
	// return trx

	databases.Load()
	var trx = contract.Transaction{}
	var result []contract.Transaction

	rows, err := databases.DBCon.Table("transactions").Select("transactions.tipe,transactions.amount,transactions.status").Rows()
	if err!=nil {
		return nil,err
	}

	for rows.Next(){
		err := rows.Scan(&trx.Tipe,&trx.Amount,&trx.Status)
		if err != nil{
			return nil,err
		}
		result = append(result, trx) 
	}
	return result,nil

}