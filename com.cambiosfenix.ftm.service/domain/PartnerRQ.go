package domain

type CreatePartnerRQ struct {
	Code        string
	Name        string
	Address     string
	City        string
	ZipCode     string
	State       string
	CountryCode string
	Phone       string
	Email       string
}

type DeletePartnerRQ struct {
	ID int
}

type GetPartnersRQ struct {
	ID   int
	Code string
	Name string
}

type UpdatePartnerRQ struct {
	ID          int
	Code        string
	Name        string
	Address     string
	City        string
	ZipCode     string
	State       string
	CountryCode string
	Phone       string
	Email       string
}
