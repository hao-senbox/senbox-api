package usecase

import (
	"context"
	"fmt"
	"sen-global-api/internal/data/repository"
	"sen-global-api/internal/domain/entity"
	"sen-global-api/pkg/uploader"
)

type GetImageUseCase struct {
	uploader.UploadProvider
	*repository.ImageRepository
}

func (receiver *GetImageUseCase) GetAllByIds(ids []int) ([]entity.SImage, error) {
	return receiver.ImageRepository.GetAllByIds(ids)
}

func (receiver *GetImageUseCase) GetAllByName(imageName string) ([]entity.SImage, error) {
	return receiver.ImageRepository.GetAllByName(imageName)
}

func (receiver *GetImageUseCase) GetImageById(id uint64) (*entity.SImage, error) {
	return receiver.ImageRepository.GetByID(id)
}

func (receiver *GetImageUseCase) GetUrlByKey(key string) (*string, error) {
	img, err := receiver.ImageRepository.GetByKey(key)
	if err != nil {
		return nil, err
	}

	signedURL, err := receiver.GetFileUploaded(context.Background(), img.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to get signed URL: %w", err)
	}

	return signedURL, nil
}
