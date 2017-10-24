package dao

import (
	"bytes"
	//"strconv"
	"time"

	DOMAIN "FenixPlatform/com.cambiosfenix.ftm.service/domain"
	"FenixPlatform/com.cambiosfenix.ftm.service/log"
	"upper.io/db.v3"
)

/**
*	Name : getProjectCollection
*	Return: db.Collection
*	Description: Get table Project in a session
 */
func getTransactionCollection() db.Collection {
	// Get a session
	session = GetSession()
	// Return table project in the session
	return session.Collection("Transaction")
}

/**
*	Name : GetAllProjects
*	Return: []*DOMAIN.Project
*	Description: Get all projects in a project table
 */
func GetAllTransactions() []*DOMAIN.Transaction {
	// Slice to keep all Transactions
	var transactions []*DOMAIN.Transaction
	// Add all projects in Transactions variable
	err := getTransactionCollection().Find().All(&transactions)
	// Close session when ends the method
	defer session.Close()
	if err != nil {
		log.Error(err)
	}
	return transactions
}

/**
*	Name : GetProjectById
*	Params: pId
*	Return: *DOMAIN.Project
*	Description: Get a project by ID in a project table
 */
func GetTransactionById(pId int) *DOMAIN.Transaction {
	// Project structure
	transaction := DOMAIN.Transaction{}
	// Add in Transaction variable, the project where ID is the same that the param
	res := getTransactionCollection().Find(db.Cond{"TransactionID": pId})

	//project.ProjectType = GetTypesByProjectId(pId)

	// Close session when ends the method
	defer session.Close()
	err := res.One(&transaction)
	if err != nil {
		log.Error(err)
		return nil
	}

	return &transaction
}

/**
*	Name : GetProjectsByDateRange
*	Params: pStartDate, pEndDate
*	Return: []*DOMAIN.Project
*	Description: Get a project in a date range in a project table
 */
func GetTransactionsByDateRange(pStartDate, pEndDate int64) []*DOMAIN.Transaction {
	// Slice to keep all projects
	transactions := []*DOMAIN.Transaction{}
	startDate := time.Unix(pStartDate, 0).Format("20060102")
	endDate := time.Unix(pEndDate, 0).Format("20060102")
	// Filter Transaction by date range
	res := getTransactionCollection().Find().Where("Date >= ?", startDate).And("Date <= ?", endDate)
	// Close session when ends the method
	defer session.Close()
	// Add all Transaction in projects variable
	err := res.All(&transactions)
	if err != nil {
		log.Error(err)
	}
	return transactions
}

/**
*	Name : AddProject
*	Params: pProject
*	Return: int, error
*	Description: Add project in DB
 */
func AddTransaction(pTransaction *DOMAIN.Transaction) (int, error) {
	// Get a session
	session = GetSession()
	// Close session when ends the method
	defer session.Close()
	// Insert in DB
	res, err := session.InsertInto("Transaction").Columns(
		"TransactionCode",
		"Type",
		"Date",
		"TRMID",
		"TRMValue",
		"PartnerID",
		"CurrencyFromQuantity",
		"CurrencyToQuantity",
		"PayTypeID",
		"IncoTermID",
		"Casher",
		"Notes").Values(
		pTransaction.Code,
		pTransaction.Type,
		pTransaction.Date,
		pTransaction.TRM.ID,
		pTransaction.TRM.TRMValue,
		pTransaction.Partner.ID,
		pTransaction.CurrencyFromQuantity,
		pTransaction.CurrencyToQuantity,
		pTransaction.PayType,
		pTransaction.IncoTerm,
		pTransaction.Casher,
		pTransaction.Notes).Exec()
	if err != nil {
		log.Error(err)
		return 0, err
	}
	// Get rows inserted
	insertId, err := res.LastInsertId()

	return int(insertId), nil
}

/**
*	Name : UpdateTransaction
*	Params: pTransaction
*	Return: int, error
*	Description: Update Transaction in DB
 */
func UpdateTransaction(pTransaction *DOMAIN.Transaction) (int, error) {
	// Get a session
	session = GetSession()
	// Close session when ends the method
	defer session.Close()
	// Update project in DB
	q := session.Update("Transaction").Set("TransactionCode = ?, Type = ?, Date = ?, TRMID = ?, TRMValue = ?, PartnerID = ?, CurrencyFromQuantity = ?, CurrencyToQuantity = ?, PayTypeID = ?,IncoTermID = ?, Casher = ?, Notes = ?", pTransaction.Code, pTransaction.Type, pTransaction.Date, pTransaction.TRM.ID, pTransaction.TRM.TRMValue, pTransaction.Partner.ID, pTransaction.CurrencyFromQuantity, pTransaction.CurrencyToQuantity, pTransaction.PayType, pTransaction.IncoTerm, pTransaction.Casher, pTransaction.Notes).Where("TransactionID = ?", int(pTransaction.ID))

	res, err := q.Exec()
	if err != nil {
		log.Error(err)
		return 0, err
	}
	// Get rows updated
	updateCount, err := res.RowsAffected()
	return int(updateCount), nil
}

/**
*	Name : DeleteProject
*	Params: pProjectId
*	Return: int, error
*	Description: Delete project in DB
 */
func DeleteTransaction(pTransactionId int) (int, error) {
	// Get a session
	session = GetSession()
	// Close session when ends the method
	defer session.Close()
	// Delete project in DB
	q := session.DeleteFrom("Transaction").Where("TransactionID", pTransactionId)
	res, err := q.Exec()
	if err != nil {
		log.Error(err)
		return 0, err
	}
	// Get rows deleted
	deleteCount, err := res.RowsAffected()
	return int(deleteCount), nil
}

func GetTransactionsByFilters(pTransactionFilters *DOMAIN.Transaction, pStartDate, pEndDate string) ([]*DOMAIN.Transaction, string) {
	// Slice to keep all resources
	//transactions := []*DOMAIN.Transaction{}
	transactions := GetAllTransactions()
	//result := getTransactionCollection().Find(db.Cond{"TransactionID": 2})

	// Close session when ends the method
	//defer session.Close()

	var filters bytes.Buffer
	if pTransactionFilters != nil {

		//		if pTransactionFilters.ID != 0 {
		//			filters.WriteString("TransactionID= '")
		//			filters.WriteString(strconv.Itoa(pTransactionFilters.ID))
		//			filters.WriteString("'")
		//		}
		//		if pTransactionFilters.Code != "" {
		//			if filters.String() != "" {
		//				filters.WriteString(" and ")
		//			}
		//			filters.WriteString("TransactionCode = '")
		//			filters.WriteString(pTransactionFilters.Code)
		//			filters.WriteString("'")
		//		}

		//		if pStartDate != "" {
		//			if filters.String() != "" {
		//				filters.WriteString(" and ")
		//			}
		//			filters.WriteString("Date >= '")
		//			filters.WriteString(pStartDate)
		//			filters.WriteString("'")
		//		}
		//		if pEndDate != "" {
		//			if filters.String() != "" {
		//				filters.WriteString(" and ")
		//			}
		//			filters.WriteString("Date <= '")
		//			filters.WriteString(pEndDate)
		//			filters.WriteString("'")
		//		}

		//		if pTransactionFilters.Type != 0 {
		//			filters.WriteString("Type= '")
		//			filters.WriteString(strconv.Itoa(pTransactionFilters.Type))
		//			filters.WriteString("'")
		//		}

		//		if pTransactionFilters.TRM.ID != 0 {
		//			filters.WriteString("TRMID= '")
		//			filters.WriteString(strconv.Itoa(pTransactionFilters.TRM.ID))
		//			filters.WriteString("'")
		//}

		//	if filters.String() != "" {
		//		err := result.Where(filters.String()).All(&transactions)

		//		if err != nil {
		//			log.Error(err)
		//		}
		//	}
		//_ = result
	}

	// Add all Transaction in projects variable

	return transactions, filters.String()
}
