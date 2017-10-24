package tool

import (
	"fmt"
	"time"

	"FenixPlatform/com.cambiosfenix.ftm.service/dao"
	DOMAIN "FenixPlatform/com.cambiosfenix.ftm.service/domain"
	"FenixPlatform/com.cambiosfenix.ftm.service/log"
	"FenixPlatform/com.cambiosfenix.ftm.service/util"
)

func CreateTransaction(pRequest *DOMAIN.CreateTransactionRQ) *DOMAIN.CreateTransactionRS {

	timeResponse := time.Now()
	// Create response
	response := DOMAIN.CreateTransactionRS{}

	isValid, message := util.ValidateDates(&pRequest.Date, true)
	if !isValid {
		response.Message = message
		response.Transaction = nil
		response.Status = "Error"
		return &response
	}

	transaction := util.MappingCreateTransaction(pRequest)

	if transaction != nil {
		// Save in DB
		id, err := dao.AddTransaction(transaction)

		if err != nil {
			message := "Project wasn't insert"
			log.Error(message)
			//response.Message = message
			response.Message = err.Error()
			response.Transaction = nil
			response.Status = "Error"
			return &response
		}

		// Get Project inserted
		transaction = dao.GetTransactionById(id)
		response.Transaction = transaction
		response.Status = "OK"

		response.Header = util.BuildHeaderResponse(timeResponse)
		return &response
	}
	response.Transaction = nil
	response.Status = "Error"

	message = "Error in creation of project"
	log.Error(message)
	response.Message = message

	response.Header = util.BuildHeaderResponse(timeResponse)
	return &response
}

func UpdateTransaction(pRequest *DOMAIN.UpdateTransactionRQ) *DOMAIN.UpdateTransactionRS {
	timeResponse := time.Now()
	response := DOMAIN.UpdateTransactionRS{}

	isValid, message := util.ValidateDates(&pRequest.Date, false)
	if !isValid {
		response.Message = message
		response.Transaction = nil
		response.Status = "Error"
		return &response
	}

	oldTransaction := dao.GetTransactionById(pRequest.ID)
	if oldTransaction != nil {
		fmt.Sprintf("Found transaction with that id")
		if pRequest.Code != "" {
			oldTransaction.Code = pRequest.Code
		}
		if pRequest.Type != 0 {
			oldTransaction.Type = pRequest.Type
		}
		if pRequest.TRMID != 0 {
			oldTransaction.TRM = util.BuildTRM(pRequest.TRMID)
		}

		if pRequest.TRMValue != 0 {
			oldTransaction.TRM.TRMValue = pRequest.TRMValue
		}

		if pRequest.PartnerID != 0 {
			oldTransaction.Partner = util.BuildPartner(pRequest.PartnerID)
		}
		if pRequest.CurrencyFromQuantity != 0 {
			oldTransaction.CurrencyFromQuantity = pRequest.CurrencyFromQuantity
		}

		if pRequest.CurrencyToQuantity != 0 {
			oldTransaction.CurrencyToQuantity = pRequest.CurrencyToQuantity
		}
		if pRequest.Date != "" {
			startDate := new(string)
			startDate = &pRequest.Date
			if startDate == nil || *startDate == "" {
				log.Error("Dates undefined")
				return nil
			}
			startDateInt, err := util.ConvertirFechasPeticion(startDate)
			if err != nil {
				log.Error(err)
				return nil
			}
			if pRequest.Date != "" {
				oldTransaction.Date = time.Unix(startDateInt, 0)
			}

		}

		if pRequest.PayType != 0 {
			oldTransaction.PayType = pRequest.PayType
		}
		if pRequest.IncoTerm != 0 {
			oldTransaction.IncoTerm = pRequest.IncoTerm
		}
		if pRequest.Casher != "" {
			oldTransaction.Casher = pRequest.Casher
		}

		if pRequest.Notes != "" {
			oldTransaction.Notes = pRequest.Notes
		}
		// Save in DB
		rowsUpdated, err := dao.UpdateTransaction(oldTransaction)
		if err != nil || rowsUpdated <= 0 {
			message := "Transaction wasn't update"
			log.Error(message)
			//response.Message = message
			response.Message = err.Error()
			response.Transaction = nil
			response.Status = "Error"
			return &response
		}

		// Get Transaction updated
		transaction := dao.GetTransactionById(pRequest.ID)
		response.Transaction = transaction
		response.Status = "OK"

		response.Header = util.BuildHeaderResponse(timeResponse)

		return &response
	}

	message = "Resource wasn't found in DB"
	log.Error(message)
	response.Message = message
	response.Transaction = nil
	response.Status = "Error"

	response.Header = util.BuildHeaderResponse(timeResponse)

	return &response
}

func DeleteTransaction(pRequest *DOMAIN.DeleteTransactionRQ) *DOMAIN.DeleteTransactionRS {
	timeResponse := time.Now()
	response := DOMAIN.DeleteTransactionRS{}
	transactionToDelete := dao.GetTransactionById(pRequest.ID)
	if transactionToDelete != nil {
		// Delete in DB
		rowsDeleted, err := dao.DeleteTransaction(pRequest.ID)
		if err != nil || rowsDeleted <= 0 {
			message := "Transaction wasn't delete"
			log.Error(message)
			response.Message = message
			response.Status = "Error"
			return &response
		}

		response.ID = transactionToDelete.ID
		response.Code = transactionToDelete.Code
		response.Status = "OK"

		response.Header = util.BuildHeaderResponse(timeResponse)

		return &response
	}

	message := "Transaction wasn't found in DB"
	log.Error(message)
	response.ID = pRequest.ID
	response.Message = message
	response.Status = "Error"

	response.Header = util.BuildHeaderResponse(timeResponse)

	return &response
}

func GetTransactions(pRequest *DOMAIN.GetTransactionsRQ) *DOMAIN.GetTransactionsRS {
	timeResponse := time.Now()
	response := DOMAIN.GetTransactionsRS{}
	var message string
	isValid, message := util.ValidateDates(&pRequest.StartDate, false)
	if !isValid {
		response.Message = message
		response.Status = "Error"
		return &response
	}

	isValidEnd, messageEnd := util.ValidateDates(&pRequest.EndDate, false)
	if !isValidEnd {
		response.Message = messageEnd
		response.Status = "Error"
		return &response
	}

	filters := util.MappingFiltersTransactions(pRequest)
	if filters != nil {
		transactions, filterString := dao.GetTransactionsByFilters(filters, pRequest.StartDate, pRequest.EndDate)
		_ = filterString
		_ = transactions

		//		if len(transactions) == 0 && filterString == "" || transactions == nil {
		//			transactions = dao.GetAllTransactions()
		//		}
		//		if transactions != nil && len(transactions) > 0 {

		//			response.Transactions = transactions
		//			// Create response
		//			response.Status = "OK"

		//			response.Header = util.BuildHeaderResponse(timeResponse)

		//			return &response
		//		}
	}

	transactions := dao.GetAllTransactions()
	if transactions != nil && len(transactions) > 0 {
		//response.Transactions = transactions
		response.Transactions = append(response.Transactions, filters)
		// Create response
		response.Status = "OK"

		response.Header = util.BuildHeaderResponse(timeResponse)

		return &response
	}

	message = "Transactions were not found in DB"
	log.Error(message)
	response.Message = message
	response.Status = "Error"

	response.Header = util.BuildHeaderResponse(timeResponse)

	return &response
}

func newTrue() *bool {
	b := true
	return &b
}

func getFilterTransactions(pStartDate, pEndDate string) []*DOMAIN.Transaction {
	requestTransactions := DOMAIN.GetTransactionsRQ{}
	requestTransactions.StartDate = pStartDate
	requestTransactions.EndDate = pEndDate

	//TODO filter in query enabled projects.
	responseTransactions := GetTransactions(&requestTransactions)
	return responseTransactions.Transactions
}
