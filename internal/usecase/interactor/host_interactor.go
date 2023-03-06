package interactor

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter/mapper"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/repository"
)

type hostInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.HostPresenter
	XMLMapper           mapper.XMLMapper
}

func NewHostInteractor(repository repository.RegistrarRepository, presenter presenter.HostPresenter, xmlMapper mapper.XMLMapper) usecase.HostInteractor {
	return &hostInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
		XMLMapper:           xmlMapper,
	}
}

func (interactor *hostInteractor) Check(ctx infrastructure.Context, data interface{}, ext string, langTag string) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Send: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CheckHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	interactor.Presenter.CheckSuccess(ctx, *responseDTO)
}

func (interactor *hostInteractor) Create(ctx infrastructure.Context, data interface{}, ext string, langTag string) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CreateHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	interactor.Presenter.CreateSuccess(ctx, *responseDTO)
}

func (interactor *hostInteractor) Update(ctx infrastructure.Context, data interface{}, ext string, langTag string) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Update: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.UpdateHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	interactor.Presenter.UpdateSuccess(ctx, *responseDTO)
}

func (interactor *hostInteractor) Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Delete: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.DeleteHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	interactor.Presenter.DeleteSuccess(ctx, *responseDTO)
}

func (interactor *hostInteractor) Info(ctx infrastructure.Context, data interface{}, ext string, langTag string) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Info: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.InfoHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	interactor.Presenter.InfoSuccess(ctx, *responseDTO)
}
