package interactor

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter/mapper"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/domain"
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
) usecase.PollInteractor {
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

	var responseDTO domain.PollRequestResponseDTO
	var code int = -1

	for code != 1300 {
		responseByte, err := interactor.RegistrarRepository.SendCommand(pollRequestData)
		if err != nil {
			err = errors.Wrap(err, "PollInteractor Poll: interactor.RegistrarRepository.SendCommand")
			break
		}

		responseDTO, err = interactor.XMLMapper.ToPollRequestResponseDTO(responseByte)
		if err != nil {
			err = errors.Wrap(err, "PollInteractor Poll: interactor.XMLMapper.ToPollRequestResponseDTO")
			break
		}

		code = responseDTO.GetResultCode()

		if responseDTO.GetMessageQueue() != nil {

			if code == 1301 {

				eppPoll := interactor.DtoToEntityMapper.MapPollRequestResponseToEppPollEntity(responseDTO)

				err = interactor.EppPollRepository.Insert(eppPoll)
				if err != nil {
					err = errors.Wrap(err, "PollInteractor Poll: EppPollRepository.Insert")
					break
				}

				// Acknowledge
				pollAcknowledgeData := types.Poll{
					Poll: types.PollCommand{
						Operation: types.PollOperationAcknowledge,
						MessageID: responseDTO.GetMessageQueueId(),
					},
				}
				interactor.RegistrarRepository.SendCommand(pollAcknowledgeData)

			} else {

				// Acknowledge
				pollAcknowledgeData := types.Poll{
					Poll: types.PollCommand{
						Operation: types.PollOperationAcknowledge,
						MessageID: responseDTO.GetMessageQueueId(),
					},
				}
				interactor.RegistrarRepository.SendCommand(pollAcknowledgeData)

			}

			code = 1301
			break
		} else {
			code = 1300
			break
		}
	}

	res = interactor.Presenter.Poll(responseDTO.(presenter.PollRequestResponseDTO))

	return
}
