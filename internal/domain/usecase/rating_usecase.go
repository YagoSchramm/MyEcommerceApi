package usecase

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/rules"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/usecase/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
)

type RatingUsecase struct {
	repo *repository.RatingRepository
}

func NewRatingUsecase(repo *repository.RatingRepository) *RatingUsecase {
	return &RatingUsecase{repo: repo}
}
func (usc *RatingUsecase) CreateRating(ctx context.Context, rating *dto.CreateRatingDTO) error {
	err := rules.CreateRating(*rating)
	if err != nil {
		return err
	}
	ratingEntity := dto.ToRatingEntity(*rating)
	return usc.repo.CreateRating(ctx, *ratingEntity)
}
func (usc *RatingUsecase) UpdateRating(ctx context.Context, updateIt *dto.UpdateRatingDTO) error {
	err := rules.UpdateRating(*updateIt)
	if err != nil {
		return err
	}
	return usc.repo.UpdateRating(ctx, *updateIt)
}
func (usc *RatingUsecase) DeletRating(ctx context.Context, deletIt *dto.DeleteRatingDTO) error {
	return usc.repo.DeleteRating(ctx, deletIt)
}
func (usc *RatingUsecase) GetRatingById(ctx context.Context, input *dto.GetRatingByIdDTO) (*dto.RatingResponseDTO, error) {
	ratingEntity, err := usc.repo.GetRatingById(ctx, input.ID.String())
	if err != nil {
		return nil, err
	}
	ratingResponse := dto.ToRatingResponse(ratingEntity)
	return &ratingResponse, nil
}
func (usc *RatingUsecase) GetRatingByUserId(ctx context.Context, input *dto.GetRatingByUserIdDTO) ([]*dto.RatingResponseDTO, error) {
	ratingEntityList, err := usc.repo.GetRatingByUserId(ctx, input.UserID.String())
	if err != nil {
		return nil, err
	}
	var ratingResponseList []*dto.RatingResponseDTO
	for _, rating := range ratingEntityList {
		ratingResponse := dto.ToRatingResponse(rating)
		ratingResponseList = append(ratingResponseList, &ratingResponse)
	}
	return ratingResponseList, nil
}
func (usc *RatingUsecase) GetAllByProductId(ctx context.Context, input *dto.GetAllRatingByProductIdDTO) ([]*dto.RatingResponseDTO, error) {
	ratingEntityList, err := usc.repo.GetAllByProductId(ctx, input.ProductID.String())
	if err != nil {
		return nil, err
	}
	var ratingResponseList []*dto.RatingResponseDTO
	for _, rating := range ratingEntityList {
		ratingResponse := dto.ToRatingResponse(rating)
		ratingResponseList = append(ratingResponseList, &ratingResponse)
	}
	return ratingResponseList, nil
}
