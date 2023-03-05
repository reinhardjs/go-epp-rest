package mapper

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/entities"
)

type DtoToEntityImpl struct{}

func (m *DtoToEntityImpl) MapPollRequestResponseToEppPoll(input response.PollRequestResponse) (output entities.EPPPoll, err error) {
	layoutFormat := "2006-01-02T15:04:05.999999999-0700"

	queueDate, errParse := time.Parse(layoutFormat, input.MessageQueue.QueueDate)
	if errParse != nil {
		err = errors.Wrap(errParse, "PollInteractor Poll: QueueDate time.Parse")
		return
	}

	reDateTime, errParse := time.Parse(layoutFormat, input.ResultData.TransferData.RequestingDate)
	if errParse != nil {
		err = errors.Wrap(errParse, "PollInteractor Poll: RequestingDate time.Parse")
		return
	}

	exDateTime, errParse := time.Parse(layoutFormat, input.ResultData.TransferData.ExpireDate)
	if errParse != nil {
		err = errors.Wrap(errParse, "PollInteractor Poll: ExpireDate time.Parse")
		return
	}

	acDateTime, errParse := time.Parse(layoutFormat, input.ResultData.TransferData.ActingDate)
	if errParse != nil {
		err = errors.Wrap(errParse, "PollInteractor Poll: ActingDate time.Parse")
		return
	}

	output = entities.EPPPoll{
		Registry:       "Verisign",
		Datetime:       time.Now(),
		MessageId:      input.MessageQueue.Id,
		MessageCount:   input.MessageQueue.Count,
		Message:        input.MessageQueue.Message,
		QDate:          queueDate,
		Domain:         input.ResultData.TransferData.Name,
		Status:         string(input.ResultData.TransferData.TransferStatus),
		RequestingDate: reDateTime,
		ExpireDate:     exDateTime,
		ActingDate:     acDateTime,
		RequestingId:   input.ResultData.TransferData.RequestingID,
		ActingId:       input.ResultData.TransferData.ActingID,
	}

	return
}
