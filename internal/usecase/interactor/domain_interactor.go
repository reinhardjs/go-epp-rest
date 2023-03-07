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

type domainInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.DomainPresenter
	XMLMapper           mapper.XMLMapper
}

func NewDomainInteractor(domainRepository repository.RegistrarRepository, presenter presenter.DomainPresenter, xmlMapper mapper.XMLMapper) usecase.DomainInteractor {
	return &domainInteractor{
		RegistrarRepository: domainRepository,
		Presenter:           presenter,
		XMLMapper:           xmlMapper,
	}
}

func (interactor *domainInteractor) Check(ctx infrastructure.Context, data interface{}, ext string, langTag string) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Check: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CheckDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	interactor.Presenter.CheckSuccess(ctx, *responseDTO)
}

func (interactor *domainInteractor) Create(ctx infrastructure.Context, data interface{}, ext string, langTag string) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CreateDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	interactor.Presenter.CreateSuccess(ctx, *responseDTO)
}

func (interactor *domainInteractor) Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.DeleteDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	interactor.Presenter.DeleteSuccess(ctx, *responseDTO)
}

func (interactor *domainInteractor) Info(ctx infrastructure.Context, data interface{}, ext string, langTag string) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.InfoDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	interactor.Presenter.InfoSuccess(ctx, *responseDTO)
}

func (interactor *domainInteractor) SecDNSUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor SecDNSUpdate: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.SecDNSUpdateResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	interactor.Presenter.SecDNSUpdateSuccess(ctx, *responseDTO)
}

func (interactor *domainInteractor) ContactUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor ContactUpdate: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.DomainUpdateResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		return
	}

	interactor.Presenter.ContactUpdateSuccess(ctx, *responseDTO)
}
