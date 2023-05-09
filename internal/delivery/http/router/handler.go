package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/controllers"
)

type handler struct {
	appController controllers.AppController
}

func (h *handler) apiV1(c *gin.Context) {
	action := c.Query("action")

	switch action {
	case "queryns":
		h.hostCheck(c)
	case "query":
		h.domainCheck(c)
	case "addcontact":
		h.contactCreate(c)
	case "register":
		h.domainCreate(c)
	case "registerns":
		h.hostCreate(c)
	case "queryaddns":
		h.hostCheckAndCreate(c)
	case "updatecontact":
		h.contactUpdate(c)
	case "updatens":
		h.hostUpdate(c)
	case "delete":
		h.domainDelete(c)
	case "deletecontact":
		h.contactDelete(c)
	case "deletens":
		h.hostDelete(c)
	case "infocontact":
		h.contactInfo(c)
	case "infohost":
		h.hostInfo(c)
	case "info":
		h.domainInfo(c)
	case "renew":
		h.domainRenew(c)
	case "querytransfer":
		h.transferCheck(c)
	case "request":
		h.transferRequest(c)
	case "cancel":
		h.transferCancel(c)
	case "approve":
		h.transferApprove(c)
	case "reject":
		h.transferReject(c)
	case "updatedomdnssec":
		h.domainSecDNSUpdate(c)
	case "poll":
		h.poll(c)
	}
}

func (h *handler) domainCheck(c *gin.Context) {
	h.appController.Domain.Check(c)
}

func (h *handler) domainCreate(c *gin.Context) {
	h.appController.Domain.Create(c)
}

func (h *handler) domainDelete(c *gin.Context) {
	h.appController.Domain.Delete(c)
}

func (h *handler) domainInfo(c *gin.Context) {
	h.appController.Domain.Info(c)
}

func (h *handler) domainSecDNSUpdate(c *gin.Context) {
	h.appController.Domain.SecDNSUpdate(c)
}

func (h *handler) domainContactUpdate(c *gin.Context) {
	h.appController.Domain.ContactUpdate(c)
}

func (h *handler) domainStatusUpdate(c *gin.Context) {
	h.appController.Domain.StatusUpdate(c)
}

func (h *handler) domainAuthInfoUpdate(c *gin.Context) {
	h.appController.Domain.AuthInfoUpdate(c)
}

func (h *handler) domainNameserverUpdate(c *gin.Context) {
	h.appController.Domain.NameserverUpdate(c)
}

func (h *handler) domainRenew(c *gin.Context) {
	h.appController.Domain.Renew(c)
}

func (h *handler) hostCheck(c *gin.Context) {
	h.appController.Host.Check(c)
}

func (h *handler) hostCreate(c *gin.Context) {
	h.appController.Host.Create(c)
}

func (h *handler) hostUpdate(c *gin.Context) {
	h.appController.Host.Update(c)
}

func (h *handler) hostDelete(c *gin.Context) {
	h.appController.Host.Delete(c)
}

func (h *handler) hostInfo(c *gin.Context) {
	h.appController.Host.Info(c)
}

func (h *handler) hostChange(c *gin.Context) {
	h.appController.Host.Change(c)
}

func (h *handler) hostCheckAndCreate(c *gin.Context) {
	h.appController.Host.CheckAndCreate(c)
}

func (h *handler) contactCheck(c *gin.Context) {
	h.appController.Contact.Check(c)
}

func (h *handler) contactCreate(c *gin.Context) {
	h.appController.Contact.Create(c)
}

func (h *handler) contactUpdate(c *gin.Context) {
	h.appController.Contact.Update(c)
}

func (h *handler) contactDelete(c *gin.Context) {
	h.appController.Contact.Delete(c)
}

func (h *handler) contactInfo(c *gin.Context) {
	h.appController.Contact.Info(c)
}

func (h *handler) transferCheck(c *gin.Context) {
	h.appController.Transfer.Check(c)
}

func (h *handler) transferRequest(c *gin.Context) {
	h.appController.Transfer.Request(c)
}

func (h *handler) transferCancel(c *gin.Context) {
	h.appController.Transfer.Cancel(c)
}

func (h *handler) transferApprove(c *gin.Context) {
	h.appController.Transfer.Approve(c)
}

func (h *handler) transferReject(c *gin.Context) {
	h.appController.Transfer.Approve(c)
}

func (h *handler) poll(c *gin.Context) {
	h.appController.Poll.Poll(c)
}
