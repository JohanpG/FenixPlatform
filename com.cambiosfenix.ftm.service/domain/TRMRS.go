package domain

//--------- Delete Project

type DeleteTRMRS struct {
	Header  *Response_Header
	ID      int
	Name    string
	Status  string
	Message string
}

func (m *DeleteTRMRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}

//--------- Update Project
type UpdateTRMRS struct {
	Header  *Response_Header
	TRM     *TRM
	Status  string
	Message string
}

func (m *UpdateTRMRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}

//--------- Get Project
type GetTRMsRS struct {
	Header  *Response_Header
	TRM     []*TRM
	Status  string
	Message string
}

func (m *GetTRMsRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}
