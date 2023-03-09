package interactor

import (
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/error_types"
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

func (interactor *domainInteractor) Check(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Check: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CheckDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "DomainInteractor Check: interactor.XMLMapper.Decode (CheckDomainResponse)")
		return
	}

	var resultCode = responseDTO.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseDTO.Result}, "DomainInteractor Check: epp command error")
		return
	}

	err = interactor.Presenter.CheckSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Check")
		return
	}
	return
}

func (interactor *domainInteractor) Create(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CreateDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "DomainInteractor Create: interactor.XMLMapper.Decode (CreateDomainResponse)")
		return
	}

	var resultCode = responseDTO.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseDTO.Result}, "DomainInteractor Create: epp command error")
		return
	}

	err = interactor.Presenter.CreateSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create")
		return
	}
	return
}

func (interactor *domainInteractor) Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.DeleteDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "DomainInteractor Delete: interactor.XMLMapper.Decode (DeleteDomainResponse)")
		return
	}

	var resultCode = responseDTO.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseDTO.Result}, "DomainInteractor Delete: epp command error")
		return
	}

	err = interactor.Presenter.DeleteSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Delete")
		return
	}
	return
}

func (interactor *domainInteractor) Info(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.InfoDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "DomainInteractor Info: interactor.XMLMapper.Decode (InfoDomainResponse)")
		return
	}

	var resultCode = responseDTO.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseDTO.Result}, "DomainInteractor Info: epp command error")
		return
	}

	err = interactor.Presenter.InfoSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Info")
		return
	}
	return
}

func (interactor *domainInteractor) SecDNSUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor SecDNSUpdate: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.SecDNSUpdateResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "DomainInteractor SecDNSUpdate: interactor.XMLMapper.Decode (SecDNSUpdateResponse)")
		return
	}

	var resultCode = responseDTO.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseDTO.Result}, "DomainInteractor SecDNSUpdate: epp command error")
		return
	}

	err = interactor.Presenter.SecDNSUpdateSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor SecDNSUpdate")
		return
	}
	return
}

func (interactor *domainInteractor) ContactUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor ContactUpdate: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.DomainUpdateResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "DomainInteractor ContactUpdate: interactor.XMLMapper.Decode (DomainUpdateResponse)")
		return
	}

	var resultCode = responseDTO.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseDTO.Result}, "DomainInteractor ContactUpdate: epp command error")
		return
	}

	err = interactor.Presenter.ContactUpdateSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor ContactUpdate")
		return
	}
	return
}

func (interactor *domainInteractor) StatusUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor StatusUpdate: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.DomainUpdateResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "DomainInteractor StatusUpdate: interactor.XMLMapper.Decode (DomainUpdateResponse)")
		return
	}

	var resultCode = responseDTO.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseDTO.Result}, "DomainInteractor StatusUpdate: epp command error")
		return
	}

	err = interactor.Presenter.StatusUpdateSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor StatusUpdate")
		return
	}
	return
}

func (interactor *domainInteractor) AuthInfoUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor AuthInfoUpdate: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.DomainUpdateResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "DomainInteractor AuthInfoUpdate: interactor.XMLMapper.Decode (DomainUpdateResponse)")
		return
	}

	var resultCode = responseDTO.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseDTO.Result}, "DomainInteractor AuthInfoUpdate: epp command error")
		return
	}

	err = interactor.Presenter.AuthInfoUpdateSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor AuthInfoUpdate")
		return
	}
	return
}

func (interactor *domainInteractor) NameserverUpdate(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor NameserverUpdate: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.DomainUpdateResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "DomainInteractor NameserverUpdate: interactor.XMLMapper.Decode (DomainUpdateResponse)")
		return
	}

	var resultCode = responseDTO.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseDTO.Result}, "DomainInteractor NameserverUpdate: epp command error")
		return
	}

	err = interactor.Presenter.NameserverUpdateSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor NameserverUpdate")
		return
	}
	return
}

func (interactor *domainInteractor) Renew(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Renew: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.RenewDomainResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "DomainInteractor Renew: interactor.XMLMapper.Decode (DomainUpdateResponse)")
		return
	}

	var resultCode = responseDTO.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseDTO.Result}, "DomainInteractor Renew: epp command error")
		return
	}

	err = interactor.Presenter.RenewSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "DomainInteractor Renew")
		return
	}
	return
}
