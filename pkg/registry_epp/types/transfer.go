package types

// TransferType represents a transfer command.
type TransferType struct {
	TransferParent Transfer `xml:"command>transfer"`
}

type Transfer struct {
	Operation string         `xml:"op,attr"`
	Detail    TransferDetail `xml:"urn:ietf:params:xml:ns:domain-1.0 transfer"`
}

// TransferDetail represents a transfer request detail to the EPP server.
type TransferDetail struct {
	Name     string    `xml:"name"`
	AuthInfo *AuthInfo `xml:"authInfo,omitempty"`
}

// TransferData represents the response for a transfer command.
type TransferData struct {
	Name           string `xml:"name"`
	TransferStatus string `xml:"trStatus"`
	ReID           string `xml:"reId"`
	ReDate         string `xml:"reDate"`
}
