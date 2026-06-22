package service

import (
	"context"
	"slices"
	"strings"

	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/repository"
)

type SearchService interface {
	Search(ctx context.Context, query string) ([]model.SearchResult, error)
}

type searchService struct {
	groupRepo    repository.GroupRepository
	positionRepo repository.PositionRepository
}

func NewSearchService(groupRepo repository.GroupRepository, positionRepo repository.PositionRepository) SearchService {
	return searchService{
		groupRepo,
		positionRepo,
	}
}

func (s searchService) Search(ctx context.Context, query string) ([]model.SearchResult, error) {
	searchResults := []model.SearchResult{}
	positions, err := s.positionRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	groups, err := s.groupRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Implement reasonable search algorithm
	for _, position := range positions {
		if strings.Contains(strings.ToLower(position.Name), strings.ToLower(query)) {
			searchResults = append(searchResults, model.SearchResult{
				Label: position.Name,
				Path:  "/positions/" + position.ID.String(),
			})
		}
	}

	for _, group := range groups {
		if strings.Contains(strings.ToLower(group.Name), strings.ToLower(query)) {
			searchResults = append(searchResults, model.SearchResult{
				Label: group.Name,
				Path:  "/groups/" + group.ID.String(),
			})
		}
	}

	slices.SortFunc(searchResults, func(a model.SearchResult, b model.SearchResult) int {
		return strings.Compare(a.Label, b.Label)
	})

	return searchResults, nil
}
