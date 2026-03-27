package organizations

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateOrganization(ctx context.Context, req CreateOrganizationRequest) (*Organization, error) {

	org := &Organization{
		Name:   req.Name,
		Email:  req.Email,
		Status: req.Status,
	}

	err := s.repo.CreateOrganization(ctx, org)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (s *Service) GetOrganization(ctx context.Context, id string) (*Organization, error) {

	org, err := s.repo.GetOrganizationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (s *Service) GetAllOrganization(ctx context.Context) ([]Organization, error) {
	org, err := s.repo.ListOrganizations(ctx)
	if err != nil {
		return nil, err
	}
	return org, nil
}
