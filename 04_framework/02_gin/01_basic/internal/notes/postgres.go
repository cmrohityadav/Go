package notes

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct {
 pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{pool: pool}
}


func (r *postgresRepository) Create(ctx context.Context,req CreateNoteRequest,
) (*Note, error) {

	query := `
	INSERT INTO notes(title, content, pinned)
	VALUES ($1, $2, $3)
	RETURNING id, title, content, pinned, created_at, updated_at
	`

	var note Note
	time:=time.Now().UTC()
	note.CreatedAt=time
	note.UpdatedAt=time
	err := r.pool.QueryRow(ctx,query,
		req.Title,
		req.Content,
		req.Pinned,
	).Scan(&note.ID,&note.Title,&note.Content,&note.Pinned,&note.CreatedAt,&note.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *postgresRepository)List(ctx context.Context)([]Note,error){
	query:="SELECT * FROM notes"

	rows,err:=r.pool.Query(ctx,query);
	if err!=nil{
	  return []Note{},err;
	}

	defer rows.Close();
    var notes []Note;
	for rows.Next(){
		var note Note;

		err:=rows.Scan(&note.ID,&note.Title,&note.Content,&note.Pinned,&note.CreatedAt,&note.UpdatedAt)
		if err!=nil{
			return nil,err;
		}
		notes = append(notes, note);

	}
	return notes,nil;
}

func (r *postgresRepository)GetById(ctx context.Context,id int)(*Note,error){
	query := `
		SELECT *  FROM notes
		WHERE id = $1
	`

	var note Note

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.Pinned,
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func(r *postgresRepository)DeleteById(ctx context.Context,id int) (bool,error){
	query:=`DELETE FROM notes WHERE id=$1`

	result,err:=r.pool.Exec(ctx,query,id);
	if err!=nil{
		return false,err
	}

	
	if result.RowsAffected()>0{
		return true,nil
	}
	
	return false,nil

}

func (r *postgresRepository)UpdateById(ctx context.Context,id int,req *UpdateNoteRequest)(*Note, error){
	updatedValues:=[]string{}
	args:=[]any{}
	argPosition:=1

	if req.Content!=nil{
		updatedValues=append(updatedValues, fmt.Sprintf("content=$%d",argPosition))
		args=append(args, req.Content);
		argPosition++;
	}

	if req.Title!=nil{
		updatedValues=append(updatedValues, fmt.Sprintf("title=$%d",argPosition))
		args=append(args, req.Title);
		argPosition++;
	}

	if req.Pinned!=nil{
		updatedValues=append(updatedValues, fmt.Sprintf("title=$%d",argPosition))
		args=append(args, req.Title);
		argPosition++;
	}

	if len(updatedValues)==0{
		return nil,fmt.Errorf("No Update Please Provide Valid data for update");
	}

	query:=fmt.Sprintf(`UPDATE notes SET %s WHERE id =$%d RETURNING id, title, content, pinned, created_at, updated_at`,strings.Join(updatedValues,", "),argPosition)
	args=append(args, id)
	var updatedNotes Note;
	err:=r.pool.QueryRow(ctx,query,args...).Scan(&updatedNotes.ID,&updatedNotes.Title,&updatedNotes.Content,&updatedNotes.Pinned,&updatedNotes.CreatedAt,&updatedNotes.UpdatedAt)
	if err!=nil{
		return nil,err;
	}
	return &updatedNotes,nil

}
/*
func (r *postgresRepository) GetByID(
	ctx context.Context,
	id int,
) (*Note, error) {

	query := `
	SELECT id, title, content, pinned, created_at, updated_at
	FROM notes
	WHERE id = $1
	`

	var note Note

	err := r.db.QueryRow(ctx, query, id).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.Pinned,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &note, nil
}

*/


/*

func (r *postgresRepository) Update(
	ctx context.Context,
	id int,
	req UpdateNoteRequest,
) (*Note, error) {

	query := `
	UPDATE notes
	SET
		title = $1,
		content = $2,
		pinned = $3,
		updated_at = NOW()
	WHERE id = $4
	RETURNING id,title,content,pinned,created_at,updated_at
	`

	var note Note

	err := r.db.QueryRowContext(
		ctx,
		query,
		req.Title,
		req.Content,
		req.Pinned,
		id,
	).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.Pinned,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &note, nil
}


func (r *postgresRepository) Delete(
	ctx context.Context,
	id int,
) error {

	query := `
	DELETE FROM notes
	WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

*/