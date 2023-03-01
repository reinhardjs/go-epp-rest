package request

type ContactCheckQuery struct {
	ContactList string `form:"contactlist"`
}

type ContactCreateQuery struct {
	Contact   string `form:"contact"`
	Email     string `form:"email"`
	AuthInfo  string `form:"authinfo"`
	Phone     string `form:"phone"`
	Fax       string `form:"fax"`
	Company   string `form:"company"`
	Address1  string `form:"addr1"`
	Address2  string `form:"addr2"`
	City      string `form:"city"`
	State     string `form:"state"`
	Zip       string `form:"zip"`
	Country   string `form:"country"`
	Firstname string `form:"fname"`
	Lastname  string `form:"lname"`
}

type ContactUpdateQuery struct {
	Contact   string `form:"contact"`
	Email     string `form:"email"`
	AuthInfo  string `form:"authinfo"`
	Phone     string `form:"phone"`
	Fax       string `form:"fax"`
	Company   string `form:"company"`
	Address1  string `form:"addr1"`
	Address2  string `form:"addr2"`
	City      string `form:"city"`
	State     string `form:"state"`
	Zip       string `form:"zip"`
	Country   string `form:"country"`
	Firstname string `form:"fname"`
	Lastname  string `form:"lname"`
}

type ContactDeleteQuery struct {
	Contact string `form:"contact"`
}

type ContactInfoQuery struct {
	Contact string `form:"contact"`
}
