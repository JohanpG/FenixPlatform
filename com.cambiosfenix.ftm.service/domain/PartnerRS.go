package domain

//--------- Create Project
type CreatePartnerRS struct {
	Header  *Response_Header
	Partner *Partner
	Status  string
	Message string
}

func (m *CreatePartnerRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}

//--------- Delete Project

type DeletePartnerRS struct {
	Header  *Response_Header
	ID      int
	Name    string
	Status  string
	Message string
}

func (m *DeletePartnerRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}

//--------- Update Project
type UpdatePartnerRS struct {
	Header  *Response_Header
	Partner *Partner
	Status  string
	Message string
}

func (m *UpdatePartnerRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}

//--------- Get Project
type GetPartnersRS struct {
	Header  *Response_Header
	Partner []*Partner
	Status  string
	Message string
}

func (m *GetPartnersRS) GetHeader() *Response_Header {
	if m != nil {
		return m.Header
	}
	return nil
}
