package interactor

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter/mapper"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type pollInteractor struct {
	EppPollRepository   repository.EppPollRepository
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.PollPresenter
	XMLMapper           mapper.XMLMapper
	DtoToEntityMapper   mapper.DtoToEntity
}

func NewPollInteractor(
	eppPollRepository repository.EppPollRepository,
	registrarRepository repository.RegistrarRepository,
	presenter presenter.PollPresenter,
	xmlMapper mapper.XMLMapper,
	dtoToEntityMapper mapper.DtoToEntity,
) usecase.PollInteractor {
	return &pollInteractor{
		EppPollRepository:   eppPollRepository,
		RegistrarRepository: registrarRepository,
		Presenter:           presenter,
		XMLMapper:           xmlMapper,
		DtoToEntityMapper:   dtoToEntityMapper,
	}
}

func (interactor *pollInteractor) Poll(ctx infrastructure.Context) (err error) {
	pollRequestData := types.Poll{
		Poll: types.PollCommand{
			Operation: types.PollOperationRequest,
		},
	}

	var responseDTO *response.PollRequestResponse = &response.PollRequestResponse{}
	var code int = -1

	for code != 1300 {
		responseByte, err := interactor.RegistrarRepository.SendCommand(pollRequestData)
		if err != nil {
			err = errors.Wrap(err, "PollInteractor Poll: interactor.RegistrarRepository.SendCommand")
			break
		}

		err = interactor.XMLMapper.Decode(responseByte, responseDTO)
		if err != nil {
			err = errors.Wrap(err, "PollInteractor Poll: interactor.XMLMapper.Decode")
			break
		}

		code = responseDTO.Result.Code

		if responseDTO.MessageQueue != nil {
			if code == 1301 {
				eppPoll, err := interactor.DtoToEntityMapper.MapPollRequestResponseToEppPollEntity(*responseDTO)
				if err != nil {
					err = errors.Wrap(err, "PollInteractor Poll: DtoToEntityMapper.MapPollRequestResponseToEppPollEntity")
					break
				}

				err = interactor.EppPollRepository.Insert(eppPoll)
				if err != nil {
					err = errors.Wrap(err, "PollInteractor Poll: EppPollRepository.Insert")
					break
				}

				// Acknowledge
				pollAcknowledgeData := types.Poll{
					Poll: types.PollCommand{
						Operation: types.PollOperationAcknowledge,
						MessageID: &responseDTO.MessageQueue.Id,
					},
				}

				_, err = interactor.RegistrarRepository.SendCommand(pollAcknowledgeData)
				if err != nil {
					err = errors.Wrap(err, "PollInteractor Poll: interactor.RegistrarRepository.SendCommand(pollAcknowledgeData)")
					break
				}
			} else {
				// Acknowledge
				pollAcknowledgeData := types.Poll{
					Poll: types.PollCommand{
						Operation: types.PollOperationAcknowledge,
						MessageID: &responseDTO.MessageQueue.Id,
					},
				}

				_, err := interactor.RegistrarRepository.SendCommand(pollAcknowledgeData)
				if err != nil {
					err = errors.Wrap(err, "PollInteractor Poll: interactor.RegistrarRepository.SendCommand(pollAcknowledgeData)")
					break
				}
			}

			code = 1301
			break
		} else {
			code = 1300
			break
		}
	}

	if err != nil {
		return
	}

	err = interactor.Presenter.PollSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "PollInteractor Poll")
		return
	}
	return
}
