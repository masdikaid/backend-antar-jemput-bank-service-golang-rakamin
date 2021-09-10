package models

import (
	//"backend-a-antar-jemput/internal/databases"

	"gorm.io/gorm"
)

//type TransactionInterface interface{
	// Create(*TransactionInterface) (TransactionInterface,error)  //1
	//GetAll() ([]TransactionInterface,error)					//2
	// //GetByCust()([]TransactionInterface,error)					//3
	// //GetByAgent()([]TransactionInterface,error)				//4
	// //GetAgentRating() float									//5
	// GetByID(int) (TransactionInterface,error)					//6
	// Confirm() (TransactionInterface,error)						//7
	// Cancel() (TransactionInterface,error)						//8
	// Done() (TransactionInterface,error)							//9
	// Delete() (TransactionInterface,error)						//10

//}

type Transaction struct{
	gorm.Model
	CustomersID uint
	Customers Customers
	AgentsID uint
	Agents Agents
	LocationID uint
	Location Location
	Tipe string
	Amount int
	Status int
	Rating float64
	//id int
	//customerId customer
	//agetId Agent
	//locationID Location
}

// func (t *Transaction) Create() (*Transaction, error)  {
// 	return t,nil
// }

// func (t *Transaction) GetAll() ([]TransactionInterface,error)  {
// 	//return TransactionInterface.GetAll(),nil
// 	var trx []Transaction
// 	databases.DBCon.Find(&trx)
// 	return trx

// }

// func (t *Transaction) GetByCust()([]TransactionInterface,error)	  {
// 	return t,nil
// }

// func (t *Transaction) GetByAgent()([]TransactionInterface,error)  {
// 	return t,nil
// }

// func (t *Transaction) GetAgentRating() float	  {
// 	return t,nil
// }

// func (t *Transaction) GetByID(int) (*Transaction,error)  {
// 	return t,nil
// }

// func (t *Transaction) Confirm() (*Transaction,error)	  {
// 	return t,nil
// }

// func (t *Transaction) Cancel() (*Transaction,error)	  {
// 	return t,nil
// }

// func (t *Transaction) Done() (*Transaction,error)	 {
// 	return t,nil
// }

// func (t *Transaction) Delete() (*Transaction,error)  {
// 	return t,nil
// }

