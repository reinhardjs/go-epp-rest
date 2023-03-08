package request

type HostCheckQuery struct {
	Host      string `form:"ns1"`
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

type HostChangeQuery struct {
	Host      string `form:"host"`
	NewHost   string `form:"newhost"`
	Extension string `form:"ext"`
}

type HostCheckAndCreateQuery struct {
	Host   string `form:"ns1"`
	IPList string `form:"iplist"`
}
