package queries

type HostCheckQuery struct {
	HostList  string `form:"hostlist"`
	Extension string `form:"ext"`
}

type HostCreateQuery struct {
	DNSList   string `form:"dnslist"`
	Host      string `form:"host"`
	IPList    string `form:"iplist"`
	Extension string `form:"ext"`
}

type HostUpdateQuery struct {
	DNSList      string `form:"dnslist"`
	Host         string `form:"host"`
	AddIPList    string `form:"addIP"`
	RemoveIPList string `form:"removeIP"`
	Extension    string `form:"ext"`
}

type HostDeleteQuery struct {
	DNSList   string `form:"dnslist"`
	Host      string `form:"host"`
	Extension string `form:"ext"`
}

type HostInfoQuery struct {
	DNSList   string `form:"dnslist"`
	Host      string `form:"host"`
	Extension string `form:"ext"`
}
