package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"todo-list/internal/model"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/stdlib"
)

type TodoRepository struct {
	db  *sql.DB
	ctx context.Context
	wg  *sync.WaitGroup
}

func (r *TodoRepository) AddTask(task *model.Task, userId int32) error {
	query := `
	insert into tasks (name, description, status, user_Id)
	values ($1, $2, $3, $4)
	`
	res, err := r.db.ExecContext(r.ctx, query, task.Name, task.Description, task.Status, userId)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task not added")
	}

	return nil
}

func (r *TodoRepository) EditTask(task *model.Task, userId int32) error {
	query := `
	update tasks
	set
		status = $3,
		name = $4,
		description = $5
	WHERE id = $1 and user_Id = $2`

	return r.updateTask(query, task.Id, userId, task.Status, task.Name, task.Description)
}

func (r *TodoRepository) DeleteTask(id string, userId int32) error {
	query := `
	delete from tasks 
	where id = $1 and user_Id = $2`

	_, err := r.db.ExecContext(r.ctx, query, id, userId)

	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) ChangeStatus(id string, status model.Status, userId int32) error {
	query := `
	update tasks 
	set status = $3
	WHERE id = $1 and user_Id = $2`

	return r.updateTask(query, id, userId, status)
}

func (r *TodoRepository) updateTask(query string, args ...any) error {
	res, err := r.db.ExecContext(r.ctx, query, args...)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task with id %s for current user not found", args[0])
	}

	return nil
}

func (r *TodoRepository) GetTaskById(id string, userId int32) (*model.Task, error) {
	res := &model.Task{}

	query := `
	select id, name, description, status 
	FROM tasks 
	WHERE id = $1 and user_Id = $2`

	taskId, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID format: %w", err)
	}

	err = r.db.QueryRowContext(r.ctx, query, taskId, userId).Scan(&res.Id, &res.Name, &res.Description, &res.Status)

	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *TodoRepository) GetTasks(filter *model.Filter, userId int32) ([]*model.Task, error) {
	query := `
	select id, name, description, status
	from tasks
	where user_Id = $1
		and date(uuid_extract_timestamp(id)) = $2
		and status & $3 = status
	order by id`

	return r.getTasks(query, userId, filter.CreatedDate, filter.Status)
}

func (r *TodoRepository) GetAllTasks(userId int32) ([]*model.Task, error) {
	query := `
	select id, name, description, status
	from tasks
	where user_Id = $1
	order by id`

	return r.getTasks(query, userId)
}

func (r *TodoRepository) getTasks(query string, args ...any) ([]*model.Task, error) {
	res := []*model.Task{}

	rows, err := r.db.QueryContext(r.ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	for rows.Next() {
		item := &model.Task{}
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Status)

		if err != nil {
			return nil, err
		}

		res = append(res, item)
	}

	return res, nil
}

func (r *TodoRepository) GetUserByName(username string) (*model.User, error) {
	res := &model.User{}

	query := "SELECT id, password FROM users WHERE name = $1"

	err := r.db.QueryRowContext(r.ctx, query, username).Scan(&res.Id, &res.Password)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func New(ctx context.Context, wg *sync.WaitGroup) *TodoRepository {
	res := &TodoRepository{
		ctx: ctx,
		wg:  wg,
	}

	res.initDb()
	res.listenForClosingDb()

	return res
}

func (r *TodoRepository) initDb() {
	var err error

	r.db, err = sql.Open("pgx", os.Getenv("POSTGRES_CONN"))

	if err != nil {
		log.Fatalf("failed to load driver: %v", err)
	}

	s := os.Getenv("POSTGRES_CONN")
	fmt.Println(s)

	err = r.db.PingContext(r.ctx)
	if err != nil {
		fmt.Println(s)
		log.Fatalf("failed to connect to db: %v, %v", err, os.Getenv("POSTGRES_CONN"))
	}

	log.Println("connected to postgres.")
}

func (r *TodoRepository) listenForClosingDb() {
	r.wg.Add(1)

	go func() {
		defer r.wg.Done()

		<-r.ctx.Done()

		if r.db != nil {
			if err := r.db.Close(); err != nil {
				log.Fatal(err)
			}

			log.Println("disconnected from postgres.")
		}
	}()
}
