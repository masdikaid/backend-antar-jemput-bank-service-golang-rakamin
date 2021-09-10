package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"

	//"github.com/gofiber/fiber/v2"
	//"go/constant"
)

func GetAllTransactions() []contract.Transaction {
	// allTrx, err := repository.GetAllTransactions()
	// if err != nil {
	// 	resp := contract.Transaction{
	// 		Status: 400,
	// 		Message: err.Error(),
	// 	}
	// }
	// resp := contract.Transaction{
	// 	Tipe: tipe,
	// 	Amount: amount,
	// 	Status: status,
	// 	Rating: rating,
	// }
	trx, err := repository.GetAllTransactions()
	
	if err != nil {
		// resp := contract.TransactionResponse{
		// 	Status: 400,
		// 	Message: err.Error(),
		panic(err.Error())
	}
	var trxResponse []contract.Transaction
	// resp := contract.TransactionResponse{
	// 	Status: ,
	// }
	for _, v := range trx {
		t := contract.Transaction{
			Tipe:   v.Tipe,
			Amount: v.Amount,
			Status: v.Status,
		}
		trxResponse = append(trxResponse, t)
	}
	return trxResponse

	// if err != nil {
	// 	trxResponse := contract.TransactionResponse{
	// 		Status:  400,
	// 		Message: err.Error(),
	// 	}
	// } else {
	// 	var trxResponse []contract.Transaction
	// 	for _, v := range trx {
	// 		t := contract.Transaction{
	// 			Tipe:   v.Tipe,
	// 			Amount: v.Amount,
	// 			Status: v.Status,
	// 			Rating: v.Rating,
	// 		}
	// 		trxResponse = append(trxResponse, t)
	// 	}
	// 	return trxResponse
	// }

	//return trxResponse
}
