package service

import (
	"context"

	"github.com/YagoSchramm/myecommerce-api/internal/domain/rules"
	"github.com/YagoSchramm/myecommerce-api/internal/domain/service/dto"
	"github.com/YagoSchramm/myecommerce-api/internal/infrastructure/datastore/repository"
)

type RatingService struct {
	repo *repository.RatingRepository
}

func NewRatingService(repo *repository.RatingRepository) *RatingService {
	return &RatingService{repo: repo}
}
func (srv *RatingService) CreateRating(ctx context.Context, rating *dto.CreateRatingDTO) error {
	err := rules.CreateRating(*rating)
	if err != nil {
		return err
	}
	ratingEntity := dto.ToRatingEntity(*rating)
	return srv.repo.CreateRating(ctx, *ratingEntity)
}
func (srv *RatingService) UpdateRating(ctx context.Context, updateIt *dto.UpdateRatingDTO) error {
	err := rules.UpdateRating(*updateIt)
	if err != nil {
		return err
	}
	return srv.repo.UpdateRating(ctx, *updateIt)
}
func (srv *RatingService) DeletRating(ctx context.Context, deletIt *dto.DeleteRatingDTO) error {
	return srv.repo.DeleteRating(ctx, deletIt)
}
func (srv *RatingService) GetRatingById(ctx context.Context, input *dto.GetRatingByIdDTO) (*dto.RatingResponseDTO, error) {
	ratingEntity, err := srv.repo.GetRatingById(ctx, input.ID.String())
	if err != nil {
		return nil, err
	}
	ratingResponse := dto.ToRatingResponse(ratingEntity)
	return &ratingResponse, nil
}
func (srv *RatingService) GetRatingByUserId(ctx context.Context, input *dto.GetRatingByUserIdDTO) ([]*dto.RatingResponseDTO, error) {
	ratingEntityList, err := srv.repo.GetRatingByUserId(ctx, input.UserID.String())
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
func (srv *RatingService) GetAllByProductId(ctx context.Context, input *dto.GetAllRatingByProductIdDTO) ([]*dto.RatingResponseDTO, error) {
	ratingEntityList, err := srv.repo.GetAllByProductId(ctx, input.ProductID.String())
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
