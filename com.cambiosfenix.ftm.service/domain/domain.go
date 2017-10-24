package domain

import "time"

type Transaction struct {
	ID                   int       `db:"TransactionID,omitempty"`
	Code                 string    `db:"TransactionCode"`
	Type                 int       `db:",inline"`
	Date                 time.Time `db:"Date"`
	CurrencyFrom         *Currency
	CurrencyTo           *Currency
	TRM                  *TRM     `db:",inline"`
	Partner              *Partner `db:",inline"`
	CurrencyFromQuantity float64  `db:"CurrencyFromQuantity"`
	CurrencyToQuantity   float64  `db:"CurrencyToQuantity"`
	PayType              int      `db:"PayTypeID"`
	IncoTerm             int      `db:"IncoTermID"`
	Casher               string   `db:"Casher"`
	Notes                string   `db:"Notes"`
}

type Enumeration struct {
	ID    int    `db:"EnumerationID,omitempty"`
	Value string `db:"EnumerationValue"`
	Text  string `db:"EnumerationText"`
}

type Currency struct {
	ID          int    `db:"CurrencyID"`
	Code        string `db:"CurrencyCode"`
	Description string `db:"Description"`
}

type TRM struct {
	ID             int     `db:"TRMID,omitempty"`
	CurrencyFromID int     `db:"CurrencyFromID"`
	CurrencyToID   int     `db:"CurrencyToID"`
	TRMValue       float64 `db:"TRMValue"`
	TRMSaleValue   float64 `db:"TRMSaleValue"`
}

type Partner struct {
	ID          int    `db:"PartnerID,omitempty"`
	Code        string `db:"PartnerCode"`
	Name        string `db:"Name"`
	Address     string `db:"Address"`
	City        string `db:"City"`
	ZipCode     string `db:"ZipCode"`
	State       string `db:"State"`
	CountryCode string `db:"CountryCode"`
	Phone       string `db:"Phone"`
	Email       string `db:"Email"`
}

type RangeTransactionDates struct {
	StartDate string
	EndDate   string
	Hours     float64
}
