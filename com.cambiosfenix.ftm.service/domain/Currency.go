package domain

//-------- Enumeration Object
type CurrencyRQ struct {
	ID   int
	Code string
}

type CurrencyRS struct {
	Header     *Response_Header
	Currencies []*Currency
	Status     string
	Message    string
}

func (m *CurrencyRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}
