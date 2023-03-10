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

type contactInteractor struct {
	RegistrarRepository repository.RegistrarRepository
	Presenter           presenter.ContactPresenter
	XMLMapper           mapper.XMLMapper
}

func NewContactInteractor(repository repository.RegistrarRepository, presenter presenter.ContactPresenter, xmlMapper mapper.XMLMapper) usecase.ContactInteractor {
	return &contactInteractor{
		RegistrarRepository: repository,
		Presenter:           presenter,
		XMLMapper:           xmlMapper,
	}
}

func (interactor *contactInteractor) Check(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Check: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CheckContactResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Check: interactor.XMLMapper.Decode (CheckContactResponse)")
		return
	}

	err = interactor.Presenter.CheckSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Check")
		return
	}
	return
}

func (interactor *contactInteractor) Create(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CreateContactResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Create: interactor.XMLMapper.Decode (CreateContactResponse)")
		return
	}

	err = interactor.Presenter.CreateSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Create")
		return
	}
	return
}

func (interactor *contactInteractor) Update(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Update: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.UpdateContactResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Update: interactor.XMLMapper.Decode (UpdateContactResponse)")
		return
	}

	err = interactor.Presenter.UpdateSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Update")
		return
	}
	return
}

func (interactor *contactInteractor) Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Delete: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.DeleteContactResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Delete: interactor.XMLMapper.Decode (DeleteContactResponse)")
		return
	}

	err = interactor.Presenter.DeleteSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Delete")
		return
	}
	return
}

func (interactor *contactInteractor) Info(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Info: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.InfoContactResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Info: interactor.XMLMapper.Decode (InfoContactResponse)")
		return
	}

	err = interactor.Presenter.InfoSuccess(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "ContactInteractor Info")
		return
	}
	return
}
