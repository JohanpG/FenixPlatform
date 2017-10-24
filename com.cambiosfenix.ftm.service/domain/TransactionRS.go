package domain

//--------- Create Project
type CreateTransactionRS struct {
	Header      *Response_Header
	Transaction *Transaction
	Status      string
	Message     string
}

func (m *CreateTransactionRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}

//--------- Delete Project

type DeleteTransactionRS struct {
	Header  *Response_Header
	ID      int
	Code    string
	Status  string
	Message string
}

func (m *DeleteTransactionRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}

//--------- Update Project
type UpdateTransactionRS struct {
	Header      *Response_Header
	Transaction *Transaction
	Status      string
	Message     string
}

func (m *UpdateTransactionRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}

//--------- Get Project
type GetTransactionsRS struct {
	Header       *Response_Header
	Transactions []*Transaction
	Status       string
	Message      string
}

func (m *GetTransactionsRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}
