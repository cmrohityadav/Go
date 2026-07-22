package notes

import "context"

type Repository interface {
	Create(ctx context.Context,req CreateNoteRequest)(*Note,error)
	// GetByID(ctx context.Context,id int)(*Note,error)
	List(ctx context.Context) ([]Note, error)
	
	// Update(ctx context.Context, id int, req UpdateNoteRequest) (*Note, error)
	// Delete(ctx context.Context, id int) error
}