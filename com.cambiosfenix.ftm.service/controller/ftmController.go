package controller

import (
	"FenixPlatform/com.cambiosfenix.ftm.service/domain"
	"FenixPlatform/com.cambiosfenix.ftm.service/tool"
)

//Partners
//func ProcessCreatePartner(pRequest *domain.CreatePartnerRQ) *domain.CreatePartnerRS {
//	response := tool.CreatePartner(pRequest)
//	// Return response
//	return response
//}

//func ProcessUpdatePartner(pRequest *domain.UpdatePartnerRQ) *domain.UpdatePartnerRS {
//	response := tool.UpdatePartner(pRequest)
//	// Return response
//	return response
//}

//func ProcessGetPartners(pRequest *domain.GetPartnersRQ) *domain.GetPartnersRS {
//	response := tool.GetPartners(pRequest)
//	// Return response
//	return response
//}

//func ProcessDeletePartner(pRequest *domain.DeletePartnerRQ) *domain.DeletePartnerRS {
//	response := tool.DeletePartner(pRequest)
//	// Return response
//	return response
//}

//Transactions

func ProcessCreateTransaction(pRequest *domain.CreateTransactionRQ) *domain.CreateTransactionRS {
	response := tool.CreateTransaction(pRequest)
	// Return response
	return response
}

func ProcessUpdateTransaction(pRequest *domain.UpdateTransactionRQ) *domain.UpdateTransactionRS {
	response := tool.UpdateTransaction(pRequest)
	// Return response
	return response
}

func ProcessDeleteTransaction(pRequest *domain.DeleteTransactionRQ) *domain.DeleteTransactionRS {
	response := tool.DeleteTransaction(pRequest)
	// Return response
	return response
}

func ProcessGetTransactions(pRequest *domain.GetTransactionsRQ) *domain.GetTransactionsRS {
	response := tool.GetTransactions(pRequest)
	// Return response
	return response
}
