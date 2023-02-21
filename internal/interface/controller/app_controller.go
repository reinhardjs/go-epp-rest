package controller

type AppController struct {
	Domain interface{ DomainController }
}
