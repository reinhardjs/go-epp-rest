package request

type DomainCheckQuery struct {
	DomainList string `form:"domainlist"`
}

type DomainCreateQuery struct {
	Domain string `form:"domain"`
	// Nameserver        string `form:"ns"`
	RegistrantContact string `form:"regcon"`
	AdminContact      string `form:"admcon"`
	TechContact       string `form:"techcon"`
	BillingContact    string `form:"bilcon"`
	AuthInfo          string `form:"authinfo"`
	Period            string `form:"period"`
	Extension         string `form:"ext"`
}

type DomainDeleteQuery struct {
	Domain    string `form:"domain"`
	Extension string `form:"ext"`
}

type DomainInfoQuery struct {
	Domain    string `form:"domain"`
	Extension string `form:"ext"`
}

type SecDNSUpdateQuery struct {
	Domain      string `form:"domain"`
	Extension   string `form:"ext"`
	IsRemoveAll string `form:"isremoveall"`

	DdKeytag0     string `form:"dd_keytag0"`
	DdAlgorithm0  string `form:"dd_algorithm0"`
	DdDigestType0 string `form:"dd_digesttype0"`
	DdDigest0     string `form:"dd_digest0"`
	KdFlag0       string `form:"kd_flag0"`
	KdProtocol0   string `form:"kd_protocol0"`
	KdAlgorithm0  string `form:"kd_algorithm0"`
	KdPublicKey0  string `form:"kd_publickey0"`

	DdKeytag1     string `form:"dd_keytag1"`
	DdAlgorithm1  string `form:"dd_algorithm1"`
	DdDigestType1 string `form:"dd_digesttype1"`
	DdDigest1     string `form:"dd_digest1"`
	KdFlag1       string `form:"kd_flag1"`
	KdProtocol1   string `form:"kd_protocol1"`
	KdAlgorithm1  string `form:"kd_algorithm1"`
	KdPublicKey1  string `form:"kd_publickey1"`

	XddKeytag0     string `form:"xdd_keytag0"`
	XddAlgorithm0  string `form:"xdd_algorithm0"`
	XddDigest0     string `form:"xdd_digest0"`
	XddDigestType0 string `form:"xdd_digesttype0"`
	XkdFlag0       string `form:"xkd_flag0"`
	XkdProtocol0   string `form:"xkd_protocol0"`
	XkdAlgorithm0  string `form:"xkd_algorithm0"`
	XkdPublicKey0  string `form:"xkd_publickey0"`

	XddKeytag1     string `form:"xdd_keytag1"`
	XddAlgorithm1  string `form:"xdd_algorithm1"`
	XddDigest1     string `form:"xdd_digest1"`
	XddDigestType1 string `form:"xdd_digesttype1"`
	XkdFlag1       string `form:"xkd_flag1"`
	XkdProtocol1   string `form:"xkd_protocol1"`
	XkdAlgorithm1  string `form:"xkd_algorithm1"`
	XkdPublicKey1  string `form:"xkd_publickey1"`
}

type DomainContactUpdateQuery struct {
	Domain               string `form:"domain"`
	RegistrantContact    string `form:"regcon"`
	AdminContact         string `form:"admcon"`
	TechContact          string `form:"techcon"`
	BillingContact       string `form:"bilcon"`
	DeleteAdminContact   string `form:"xadmcon"`
	DeleteTechContact    string `form:"xtechcon"`
	DeleteBillingContact string `form:"xbilcon"`
	Extension            string `form:"ext"`
}

type DomainStatusUpdateQuery struct {
	Domain    string `form:"domain"`
	Status    string `form:"status"`
	Extension string `form:"ext"`
}

type DomainAuthInfoUpdateQuery struct {
	Domain    string `form:"domain"`
	AuthInfo  string `form:"authinfo"`
	Extension string `form:"ext"`
}

type DomainNameserverUpdateQuery struct {
	Domain    string `form:"domain"`
	AuthInfo  string `form:"authinfo"`
	Extension string `form:"ext"`
}
