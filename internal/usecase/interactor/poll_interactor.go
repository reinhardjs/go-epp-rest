package interactor

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/common/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/entities"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/repository"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type pollInteractor struct {
	EppPollRepository   repository.EppPollRepository
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.PollPresenter
	XMLMapper           adapter.XMLMapper
}

type PollInteractor interface {
	Poll() (res string, err error)
}

func NewPollInteractor(eppPollRepository repository.EppPollRepository, registrarRepository repository.RegistrarRepository, presenter presenter.PollPresenter, xmlMapper adapter.XMLMapper) PollInteractor {
	return &pollInteractor{
		EppPollRepository:   eppPollRepository,
		RegistrarRepository: registrarRepository,
		Presenter:           presenter,
		XMLMapper:           xmlMapper,
	}
}

func (interactor *pollInteractor) Poll() (res string, err error) {

	pollRequestData := types.Poll{
		Poll: types.PollCommand{
			Operation: types.PollOperationRequest,
		},
	}

	var responseDTO response.PollRequestResponse
	var code int = -1

	for code != 1300 {
		responseByte, err := interactor.RegistrarRepository.SendCommand(pollRequestData)
		if err != nil {
			err = errors.Wrap(err, "PollInteractor Poll: interactor.RegistrarRepository.SendCommand")
			break
		}

		err = interactor.XMLMapper.Decode(responseByte, &responseDTO)
		code = responseDTO.Result.Code

		if responseDTO.MessageQueue != nil {

			if code == 1301 {
				// https://gosamples.dev/date-time-format-cheatsheet/
				layoutFormat := "2006-01-02T15:04:05.999999999-0700"

				queueDate, errParse := time.Parse(layoutFormat, responseDTO.MessageQueue.QueueDate)
				if errParse != nil {
					err = errors.Wrap(errParse, "PollInteractor Poll: QueueDate time.Parse")
					break
				}

				reDateTime, errParse := time.Parse(layoutFormat, responseDTO.ResultData.TransferData.RequestingDate)
				if errParse != nil {
					err = errors.Wrap(errParse, "PollInteractor Poll: RequestingDate time.Parse")
					break
				}

				exDateTime, errParse := time.Parse(layoutFormat, responseDTO.ResultData.TransferData.ExpireDate)
				if errParse != nil {
					err = errors.Wrap(errParse, "PollInteractor Poll: ExpireDate time.Parse")
					break
				}

				acDateTime, errParse := time.Parse(layoutFormat, responseDTO.ResultData.TransferData.ActingDate)
				if errParse != nil {
					err = errors.Wrap(errParse, "PollInteractor Poll: ActingDate time.Parse")
					break
				}

				errInsert := interactor.EppPollRepository.Insert(entities.EPPPoll{
					Registry:       "Verisign",
					Datetime:       time.Now(),
					MessageId:      responseDTO.MessageQueue.Id,
					MessageCount:   responseDTO.MessageQueue.Count,
					Message:        responseDTO.MessageQueue.Message,
					QDate:          queueDate,
					Domain:         responseDTO.ResultData.TransferData.Name,
					Status:         string(responseDTO.ResultData.TransferData.TransferStatus),
					RequestingDate: reDateTime,
					ExpireDate:     exDateTime,
					ActingDate:     acDateTime,
					RequestingId:   responseDTO.ResultData.TransferData.RequestingID,
					ActingId:       responseDTO.ResultData.TransferData.ActingID,
				})

				if errInsert != nil {
					err = errors.Wrap(errParse, "PollInteractor Poll: EppPollRepository.Insert")
					break
				}

				// Acknowledge
				pollAcknowledgeData := types.Poll{
					Poll: types.PollCommand{
						Operation: types.PollOperationAcknowledge,
						MessageID: &responseDTO.MessageQueue.Id,
					},
				}
				interactor.RegistrarRepository.SendCommand(pollAcknowledgeData)

			} else {

				// Acknowledge
				pollAcknowledgeData := types.Poll{
					Poll: types.PollCommand{
						Operation: types.PollOperationAcknowledge,
						MessageID: &responseDTO.MessageQueue.Id,
					},
				}
				interactor.RegistrarRepository.SendCommand(pollAcknowledgeData)

			}

			code = 1301
			responseDTO.Result.Code = 1000
			responseDTO.Result.Message = "Command Completed Successfully"
		} else {
			code = 1300
			break
		}
	}

	if code == 1300 {
		responseDTO.Result.Code = 1000
		responseDTO.Result.Message = "No Message"
	}

	res = interactor.Presenter.Request(&presenter.PollRequestResponseImpl{
		DTO: &responseDTO,
	})

	return
}
