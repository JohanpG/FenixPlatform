package domain

//-------- Enumeration Object
type EnumerationRQ struct {
	SetID int
	Value string
}

type EnumerationRS struct {
	Header  *Response_Header
	Types   []*Enumeration
	Status  string
	Message string
}

func (m *EnumerationRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}

type Response_Header struct {
	ResponseTime string
	RequestDate  string
}
