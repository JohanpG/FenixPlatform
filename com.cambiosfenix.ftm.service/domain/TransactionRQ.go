package domain

type CreateTransactionRQ struct {
	Code                 string
	Type                 int
	Date                 string
	TRMID                int
	TRMValue             float64
	PartnerID            int
	CurrencyFromQuantity float64
	CurrencyToQuantity   float64
	PayType              int
	IncoTerm             int
	Casher               string
	Notes                string
}

type DeleteTransactionRQ struct {
	ID int
}

type GetTransactionsRQ struct {
	ID        int
	Code      string
	StartDate string
	EndDate   string
	Type      int
	TRMID     int
}

type UpdateTransactionRQ struct {
	ID                   int
	Code                 string
	Type                 int
	Date                 string
	TRMID                int
	TRMValue             float64
	PartnerID            int
	CurrencyFromQuantity float64
	CurrencyToQuantity   float64
	PayType              int
	IncoTerm             int
	Casher               string
	Notes                string
}
