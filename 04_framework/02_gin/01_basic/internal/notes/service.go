package notes

import (
	"context"
	"errors"
	"strings"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context,req CreateNoteRequest,) (*Note, error) {

	req.Title = strings.TrimSpace(req.Title)
	req.Content = strings.TrimSpace(req.Content)

	if req.Title == "" {
		return nil, errors.New("Title is missing , Please provide proper field")
	}

	return s.repo.Create(ctx, req)
}

func (s *Service) List(ctx context.Context)([]Note,error){
	noteList,err:=s.repo.List(ctx)
	if err!=nil{
		return nil,err;
	}

	return noteList,nil;

}

func (s *Service)GetNoteById(ctx context.Context,id int)(*Note,error){
	note,err:=s.repo.GetById(ctx,id);
	if err!=nil{
		return nil,err
	}
	return note,nil;
}
/*
func (s *Service) GetByID(
	ctx context.Context,
	id int,
) (*Note, error) {

	return s.repo.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]Note, error) {
	return s.repo.List(ctx)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

*/