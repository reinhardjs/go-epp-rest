package response

import "gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"

// Domain Check
type CheckDomainResponse struct {
	Result        Result                `xml:"response>result"`
	ResultData    CheckDomainResultData `xml:"response>resData"`
	TransactionID TransactionID         `xml:"response>trID"`
}

type CheckDomainResultData struct {
	CheckDatas []CheckDomainData `xml:"chkData>cd"`
}

type CheckDomainData struct {
	Name   Name   `xml:"name"`
	Reason string `xml:"reason,omitempty"`
}

type Name struct {
	AvailKey int    `xml:"avail,attr"`
	Value    string `xml:",chardata"`
}

// Domain Create
type CreateDomainResponse struct {
	Result        Result                 `xml:"response>result"`
	ResultData    CreateDomainResultData `xml:"response>resData"`
	TransactionID TransactionID          `xml:"response>trID"`
}

type CreateDomainResultData struct {
	CreatedData CreateDomainData `xml:"creData"`
}

type CreateDomainData struct {
	Name        string `xml:"name"`
	CreatedDate string `xml:"crDate"`
	ExpiredDate string `xml:"exDate"`
}

// Domain Delete
type DeleteDomainResponse struct {
	Result        Result        `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}

// Domain Info
type InfoDomainResponse struct {
	Result        Result               `xml:"response>result"`
	ResultData    DomainInfoResultData `xml:"response>resData"`
	TransactionID TransactionID        `xml:"response>trID"`
}

type DomainInfoResultData struct {
	InfoData types.DomainInfoData `xml:"urn:ietf:params:xml:ns:domain-1.0 infData"`
}

// Domain DNSSec Update
type SecDNSUpdateResponse struct {
	Result        Result                 `xml:"response>result"`
	ResultData    SecDNSUpdateResultData `xml:"response>resData"`
	TransactionID TransactionID          `xml:"response>trID"`
}

type SecDNSUpdateResultData struct {
	types.DomainInfoData `xml:"urn:ietf:params:xml:ns:domain-1.0 infData"`
}

// Domain Update
type DomainUpdateResponse struct {
	Result        Result        `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}

// Domain Renew
type RenewDomainResponse struct {
	Result        Result                `xml:"response>result"`
	ResultData    RenewDomainResultData `xml:"response>resData"`
	TransactionID TransactionID         `xml:"response>trID"`
}

type RenewDomainResultData struct {
	RenewedData RenewDomainData `xml:"renData"`
}

type RenewDomainData struct {
	Name        string `xml:"name"`
	ExpiredDate string `xml:"exDate"`
}
