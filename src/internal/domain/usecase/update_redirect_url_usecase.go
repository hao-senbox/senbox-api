package usecase

import (
	log "github.com/sirupsen/logrus"
	"sen-global-api/internal/data/repository"
	"sen-global-api/internal/domain/entity"
	"sen-global-api/internal/domain/request"
)

type UpdateRedirectUrlUseCase struct {
	*repository.RedirectUrlRepository
}

func (receiver *UpdateRedirectUrlUseCase) Update(id int, req request.UpdateRedirectUrlRequest) (*entity.SRedirectUrl, error) {
	form, err := receiver.RedirectUrlRepository.GetById(uint64(id))
	if err != nil {
		return nil, err
	}
	if req.Password != nil {
		form.Password = req.Password
	}

	err = receiver.RedirectUrlRepository.Update(form)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	return form, nil
}
