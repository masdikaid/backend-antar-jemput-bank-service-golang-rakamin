package repository

import (
	"backend-a-antar-jemput/internal/databases"
	"backend-a-antar-jemput/internal/entities"
	"errors"
	//"backend-a-antar-jemput/internal/models"
	//"backend-a-antar-jemput/internal/models"
)

type TransactionRepositoryInterface interface {
	Create(ent *entities.Transaction) (*entities.Transaction, error)  //1
	GetAll() ([]*entities.Transaction, error)                         //1
	GetAllByID(tipe string, id uint) ([]*entities.Transaction, error) //1
	// Update() (*entities.Transaction, error)                      //1
	// Delete() error                                               //10
}

type TransactionRepositoryMysql struct {
}

func (T TransactionRepositoryMysql) Create(ent *entities.Transaction) (*entities.Transaction, error) {
	databases.Load()
	res := databases.DBCon.Create(&ent)
	if res.Error != nil {
		return nil, res.Error
	}
	return ent, nil
}

func (T TransactionRepositoryMysql) GetAll() ([]*entities.Transaction, error) {
	databases.Load()
	res := []*entities.Transaction{}
	err := databases.DBCon.Find(&res)
	if err.Error != nil {
		return nil, err.Error
	}
	return res, nil
}

func (T TransactionRepositoryMysql) GetAllByID(tipe string, id uint) ([]*entities.Transaction, error) {
	var query string
	switch tipe {
	case "cust":
		query = "customer_id = ?"
	case "agent":
		query = "agent_id = ?"
	default:
		err := errors.New("wrong tipe")
		return nil, err
	}

	databases.Load()
	res := []*entities.Transaction{}
	err := databases.DBCon.Where(query, id).Find(res)
	if err.Error != nil {
		return nil, err.Error
	}
	return res, nil
}

// func GetAllTransactions() ([]contract.Transaction,error) {
// 	// var trx []models.Transaction
// 	// databases.DBCon.Find(&trx)
// 	// return trx

// 	databases.Load()
// 	var trx = contract.Transaction{}
// 	var result []contract.Transaction

// 	rows, err := databases.DBCon.Table("transactions").Select("transactions.tipe,transactions.amount,transactions.status").Rows()
// 	if err!=nil {
// 		return nil,err
// 	}

// 	for rows.Next(){
// 		err := rows.Scan(&trx.Tipe,&trx.Amount,&trx.Status)
// 		if err != nil{
// 			return nil,err
// 		}
// 		result = append(result, trx)
// 	}
// 	return result,nil

// }

// func GetAllTransactionsCust(customer_id string) ([]contract.TransactionCust,error)  {
// 	databases.Load()
// 	var trxcust = contract.TransactionCust{}
// 	var result []contract.TransactionCust

// 	rows,err  := databases.DBCon.Table("transactions").
// 				Select("transactions.status,transactions.tipe,transactions.amount,locations.province,locations.city,locations.district,locations.address,transactions.id,customers.customer_id").
// 				Joins("join locations on locations.id=transactions.location_id").
// 				Joins("join  customers on customers.login_id = locations.login_id").
// 				Where("transactions.customers_id = ?", customer_id).Rows()
// 	if err != nil {
// 		return nil, err
// 	}

// 	for rows.Next(){
// 		err := rows.Scan(&trxcust.Status,)
// 		if err != nil {
// 			return nil, err
// 		}
// 		result = append(result, trxcust)
// 	}
// 	return result,nil
// }
