package mapper

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/entities"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter/mapper"
)

type DtoToEntityImpl struct{}

func NewDtoToEntityMapper() mapper.DtoToEntity {
	return &DtoToEntityImpl{}
}

func (m *DtoToEntityImpl) MapPollRequestResponseToEppPollEntity(input response.PollRequestResponse) (output entities.EPPPoll, err error) {
	layoutFormat := "2006-01-02T15:04:05.999999999Z"

	queueDate, errParse := time.Parse(layoutFormat, input.MessageQueue.QueueDate)
	if errParse != nil {
		err = errors.Wrap(errParse, "PollInteractor Poll: QueueDate time.Parse")
		return
	}

	message := input.MessageQueue.Message
	var domain, status, requestingID, actingID string
	var reDateTime, exDateTime, acDateTime time.Time

	if input.ResultData.TransferData != nil {
		reDateTime, errParse = time.Parse(layoutFormat, input.ResultData.TransferData.RequestingDate)
		if errParse != nil {
			err = errors.Wrap(errParse, "PollInteractor Poll: RequestingDate time.Parse")
			return
		}

		exDateTime, _ = time.Parse(layoutFormat, input.ResultData.TransferData.ExpireDate)

		acDateTime, _ = time.Parse(layoutFormat, input.ResultData.TransferData.ActingDate)

		domain = input.ResultData.TransferData.Name
		status = string(input.ResultData.TransferData.TransferStatus)
		requestingID = input.ResultData.TransferData.RequestingID
		actingID = input.ResultData.TransferData.ActingID

		if status == "pending" {
			message = "Transfer Requested."
		}

		if status == "clientApproved" {
			message = "Transfer Approved."
		}

		if status == "clientCancelled" {
			message = "Transfer Cancelled."
		}

		if status == "clientRejected" {
			message = "Transfer Rejected."
		}

		if status == "serverApproved" {
			message = "Transfer Auto Approved."
		}

		if status == "serverCancelled" {
			message = "Transfer Auto Rejected."
		}
	}

	output = entities.EPPPoll{
		Registry:       "Verisign",
		Datetime:       time.Now(),
		MessageId:      input.MessageQueue.Id,
		MessageCount:   input.MessageQueue.Count,
		Message:        message,
		QDate:          queueDate,
		Domain:         domain,
		Status:         status,
		RequestingDate: reDateTime,
		ExpireDate:     exDateTime,
		ActingDate:     acDateTime,
		RequestingId:   requestingID,
		ActingId:       actingID,
	}

	return
}
