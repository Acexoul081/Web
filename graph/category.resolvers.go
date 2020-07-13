package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"BackEnd/graph/model"
	"context"
	"errors"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, input *model.NewCategory) (*model.Kategori, error) {
	kategori := model.Kategori{
		CategoryName: input.CategoryName,
	}
	_, err := r.DB.Model(&kategori).Insert()

	if err != nil {
		return nil, errors.New("insert category failed")
	}
	return &kategori, nil
}

func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, input *model.NewCategory) (*model.Kategori, error) {
	var kategori model.Kategori

	err := r.DB.Model(&kategori).Where("id=?", id).First()
	if err != nil {
		return nil, errors.New("category not found")
	}

	kategori.CategoryName = input.CategoryName

	_, updateErr := r.DB.Model(&kategori).Where("id=?", id).Update()

	if updateErr != nil {
		return nil, errors.New("update category failed")
	}
	return &kategori, nil
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (bool, error) {
	var kategori model.Kategori

	err := r.DB.Model(&kategori).Where("id=?", id).First()
	if err != nil {
		return false, errors.New("category not found")
	}
	_, deleteError := r.DB.Model(&kategori).Where("id=?", id).Delete()

	if deleteError != nil {
		return false, errors.New("delete error")
	}
	return true, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Kategori, error) {
	var categories []*model.Kategori

	err := r.DB.Model(&categories).Order("id").Select()

	if err != nil {
		return nil, errors.New("query failed")
	}

	return categories, nil
}
