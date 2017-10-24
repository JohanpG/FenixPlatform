package domain

type CreateTRMRQ struct {
	CurrencyFromID int
	CurrencyToID   int
	TRMValue       float64
	TRMSaleValue   float64
}

type DeleteTRMRQ struct {
	ID int
}

type GetTRMRQ struct {
	ID             int
	CurrencyFromID int
	CurrencyToID   int
}

type UpdateTRMRQ struct {
	ID             int
	CurrencyFromID int
	CurrencyToID   int
	TRMValue       float64
	TRMSaleValue   float64
}
