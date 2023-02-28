package queries

type TransferCheckQuery struct {
	Domain    string `form:"domain"`
	Extension string `form:"ext"`
	AuthInfo  string `form:"authinfo"`
}

type TransferRequestQuery struct {
	Domain    string `form:"domain"`
	Extension string `form:"ext"`
	AuthInfo  string `form:"authinfo"`
}

type TransferCancelQuery struct {
	Domain    string `form:"domain"`
	Extension string `form:"ext"`
	AuthInfo  string `form:"authinfo"`
}

type TransferApproveQuery struct {
	Domain    string `form:"domain"`
	Extension string `form:"ext"`
	AuthInfo  string `form:"authinfo"`
}

type TransferRejectQuery struct {
	Domain    string `form:"domain"`
	Extension string `form:"ext"`
	AuthInfo  string `form:"authinfo"`
}
