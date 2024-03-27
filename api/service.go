package api

import "github.com/google/uuid"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SaveResponse(response Response) (Response, error) {
	return s.repo.SaveResponse(response)
}

func (s *Service) calculateScore(answers []Answer) int {
	return s.repo.calculateScore(answers)
}

func (s *Service) GenerateUserID(user User) (User, error) {
	return s.repo.GenerateUserID(user)
}

func (s *Service) GenerateSessionID(session Session) (Session, error) {
	return s.repo.GenerateSessionID(session)
}

func (s *Service) GetUserByID(id uuid.UUID) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *Service) UpdateUser(user *User) error {
	return s.repo.UpdateUser(user)
}
