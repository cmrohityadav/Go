package notes

import "context"

type Repository interface {
	Create(ctx context.Context, req CreateNoteRequest) (*Note, error)
	GetById(ctx context.Context, id int) (*Note, error)
	List(ctx context.Context) ([]Note, error)

	UpdateById(ctx context.Context, id int, req *UpdateNoteRequest) (*Note, error)
	DeleteById(ctx context.Context, id int) (bool, error)
}
