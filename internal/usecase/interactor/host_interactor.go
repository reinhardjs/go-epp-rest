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
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
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

func (interactor *hostInteractor) Check(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Check: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CheckHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Check: interactor.XMLMapper.Decode (CheckHostResponse)")
		return
	}

	err = interactor.Presenter.Check(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Check")
		return
	}
	return
}

func (interactor *hostInteractor) Create(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Create: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.CreateHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Create: interactor.XMLMapper.Decode (CreateHostResponse)")
		return
	}

	err = interactor.Presenter.Create(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Create")
		return
	}
	return
}

func (interactor *hostInteractor) Update(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Update: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.UpdateHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Update: interactor.XMLMapper.Decode (UpdateHostResponse)")
		return
	}

	err = interactor.Presenter.Update(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Update")
		return
	}
	return
}

func (interactor *hostInteractor) Delete(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Delete: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.DeleteHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Delete: interactor.XMLMapper.Decode (DeleteHostResponse)")
		return
	}

	err = interactor.Presenter.Delete(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Delete")
		return
	}
	return
}

func (interactor *hostInteractor) Info(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	responseByte, err := interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Info: interactor.RegistrarRepository.SendCommand")
		return
	}

	responseDTO := &response.InfoHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)

	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Info: interactor.XMLMapper.Decode (InfoHostResponse)")
		return
	}

	err = interactor.Presenter.Info(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Info")
		return
	}
	return
}

func (interactor *hostInteractor) Change(ctx infrastructure.Context, data types.HostUpdateType, ext string, langTag string) (err error) {
	infData := types.HostInfoType{
		Info: types.HostInfo{
			Name: data.Update.Name,
		},
	}
	responseByte, err := interactor.RegistrarRepository.SendCommand(infData)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Change: interactor.RegistrarRepository.SendCommand (host info)")
		return
	}
	hostInfoResponseDTO := &response.InfoHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, hostInfoResponseDTO)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Change: interactor.XMLMapper.Decode (InfoHostResponse)")
		return
	}

	data.Update.Remove.Address = hostInfoResponseDTO.ResultData.InfoData.Address

	responseByte, err = interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Change: interactor.RegistrarRepository.SendCommand (host update)")
		return
	}

	responseDTO := &response.UpdateHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor Change: interactor.XMLMapper.Decode (UpdateHostResponse)")
		return
	}

	err = interactor.Presenter.Update(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor Update")
		return
	}
	return
}

func (interactor *hostInteractor) CheckAndCreate(ctx infrastructure.Context, data interface{}, ext string, langTag string) (err error) {
	checkData := types.HostCheckType{
		Check: types.HostCheck{
			Names: []string{
				data.(types.HostCreateType).Create.Name,
			},
		},
	}

	responseByte, err := interactor.RegistrarRepository.SendCommand(checkData)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor CheckAndCreate: interactor.RegistrarRepository.SendCommand (host check)")
		return
	}

	checkHostResponseDTO := &response.CheckHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, checkHostResponseDTO)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor CheckAndCreate: interactor.XMLMapper.Decode (CheckHostResponse)")
		return
	}

	if len(checkHostResponseDTO.ResultData.CheckDatas) > 0 {
		if checkHostResponseDTO.ResultData.CheckDatas[0].HostName.Available == 0 {
			// not available to be created
			err = errors.Wrap(&error_types.InteractorError{Original: &error_types.HostNameNotAvailableError{}}, "HostInteractor CheckAndCreate")
			return
		}
	}

	responseByte, err = interactor.RegistrarRepository.SendCommand(data)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor CheckAndCreate: interactor.RegistrarRepository.SendCommand (host create)")
		return
	}

	responseDTO := &response.CreateHostResponse{}
	err = interactor.XMLMapper.Decode(responseByte, responseDTO)
	if err != nil {
		err = errors.Wrap(&error_types.InteractorError{Original: err}, "HostInteractor CheckAndCreate: interactor.XMLMapper.Decode (CreateHostResponse)")
		return
	}

	err = interactor.Presenter.CheckAndCreate(ctx, *responseDTO)
	if err != nil {
		err = errors.Wrap(err, "HostInteractor CheckAndCreate")
		return
	}
	return
}
