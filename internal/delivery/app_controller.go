package delivery

type AppController struct {
	Domain   interface{ DomainController }
	Contact  interface{ ContactController }
	Host     interface{ HostController }
	Transfer interface{ TransferController }
}
