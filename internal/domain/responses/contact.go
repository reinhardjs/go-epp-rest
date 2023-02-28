package responses

import "gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"

// Contact Check
type CheckContactResponse struct {
	Result        Result                 `xml:"response>result"`
	ResultData    CheckContactResultData `xml:"response>resData"`
	TransactionID TransactionID          `xml:"response>trID"`
}

type CheckContactResultData struct {
	CheckDatas []CheckContactData `xml:"chkData>cd"`
}

type CheckContactData struct {
	Id     Id     `xml:"id"`
	Reason string `xml:"reason,omitempty"`
}

type Id struct {
	AvailKey int    `xml:"avail,attr"`
	Value    string `xml:",chardata"`
}

// Contact Create
type CreateContactResponse struct {
	Result        Result                  `xml:"response>result"`
	ResultData    CreateContactResultData `xml:"response>resData"`
	TransactionID TransactionID           `xml:"response>trID"`
}

type CreateContactResultData struct {
	CreateData CreateContactData `xml:"creData"`
}

type CreateContactData struct {
	Id         string `xml:"id"`
	CreateDate string `xml:"crDate"`
}

// Contact Update
type UpdateContactResponse struct {
	Result        Result        `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}

// Contact Delete
type DeleteContactResponse struct {
	Result        Result        `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}

// Contact Info
type InfoContactResponse struct {
	Result        Result                  `xml:"response>result"`
	ResultData    DomainContactResultData `xml:"response>resData"`
	TransactionID TransactionID           `xml:"response>trID"`
}

type DomainContactResultData struct {
	InfoData types.ContactInfoData `xml:"urn:ietf:params:xml:ns:domain-1.0 infData"`
}
