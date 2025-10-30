package service

import (
	"context"
	"reflect"
	"sync"
	"testing"
	mock_model "todo-list/internal/mocks"
	"todo-list/internal/model"

	"github.com/golang/mock/gomock"
)

var task = &model.Task{
	Name:        "Task # 1",
	Description: "Task # 1 Description",
	Status:      model.NotStarted,
}

var userId int32 = 1

var taskId = "identity"

func TestTodoService_AddTask(t *testing.T) {
	type fields struct {
		todoRepository model.Repository
		ctx            context.Context
		wg             *sync.WaitGroup
	}
	type args struct {
		task   *model.Task
		userId int32
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_model.NewMockRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "Positive addition",
			args: args{
				task:   task,
				userId: userId,
			},
			fields: fields{
				todoRepository: repository,
				ctx:            context.TODO(),
				wg:             &sync.WaitGroup{},
			},
			wantErr: false,
			mock: func() {
				repository.EXPECT().AddTask(task, userId).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
				ctx:            tt.fields.ctx,
				wg:             tt.fields.wg,
			}
			if err := s.AddTask(tt.args.task, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("TodoService.AddTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoService_EditTask(t *testing.T) {
	type fields struct {
		todoRepository model.Repository
		ctx            context.Context
		wg             *sync.WaitGroup
	}
	type args struct {
		task   *model.Task
		userId int32
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_model.NewMockRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "Positive editing",
			args: args{
				task:   task,
				userId: userId,
			},
			fields: fields{
				todoRepository: repository,
				ctx:            context.TODO(),
				wg:             &sync.WaitGroup{},
			},
			wantErr: false,
			mock: func() {
				repository.EXPECT().EditTask(task, userId).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
				ctx:            tt.fields.ctx,
				wg:             tt.fields.wg,
			}
			if err := s.EditTask(tt.args.task, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("TodoService.EditTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoService_DeleteTask(t *testing.T) {
	type fields struct {
		todoRepository model.Repository
		ctx            context.Context
		wg             *sync.WaitGroup
	}
	type args struct {
		id     string
		userId int32
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_model.NewMockRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "Positive deletion",
			args: args{
				id:     taskId,
				userId: userId,
			},
			fields: fields{
				todoRepository: repository,
				ctx:            context.TODO(),
				wg:             &sync.WaitGroup{},
			},
			wantErr: false,
			mock: func() {
				repository.EXPECT().DeleteTask(taskId, userId).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
				ctx:            tt.fields.ctx,
				wg:             tt.fields.wg,
			}
			if err := s.DeleteTask(tt.args.id, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("TodoService.DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoService_ChangeStatus(t *testing.T) {
	type fields struct {
		todoRepository model.Repository
		ctx            context.Context
		wg             *sync.WaitGroup
	}
	type args struct {
		id     string
		status model.Status
		userId int32
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_model.NewMockRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "Positive status editing",
			args: args{
				id:     taskId,
				status: model.NotStarted,
				userId: userId,
			},
			fields: fields{
				todoRepository: repository,
				ctx:            context.TODO(),
				wg:             &sync.WaitGroup{},
			},
			wantErr: false,
			mock: func() {
				repository.EXPECT().ChangeStatus(taskId, model.NotStarted, userId).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
				ctx:            tt.fields.ctx,
				wg:             tt.fields.wg,
			}
			if err := s.ChangeStatus(tt.args.id, tt.args.status, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("TodoService.ChangeStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoService_GetTaskById(t *testing.T) {
	type fields struct {
		todoRepository model.Repository
		ctx            context.Context
		wg             *sync.WaitGroup
	}
	type args struct {
		id     string
		userId int32
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_model.NewMockRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Task
		wantErr bool
		mock    func()
	}{
		{
			name: "Positive getting of a task",
			args: args{
				id:     taskId,
				userId: userId,
			},
			fields: fields{
				todoRepository: repository,
				ctx:            context.TODO(),
				wg:             &sync.WaitGroup{},
			},
			want:    task,
			wantErr: false,
			mock: func() {
				repository.EXPECT().GetTaskById(taskId, userId).Return(task, nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
				ctx:            tt.fields.ctx,
				wg:             tt.fields.wg,
			}
			got, err := s.GetTaskById(tt.args.id, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoService.GetTaskById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoService.GetTaskById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoService_GetTasks(t *testing.T) {
	type fields struct {
		todoRepository model.Repository
		ctx            context.Context
		wg             *sync.WaitGroup
	}
	type args struct {
		filter *model.Filter
		userId int32
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_model.NewMockRepository(ctrl)

	filter := &model.Filter{}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Task
		wantErr bool
		mock    func()
	}{
		{
			name: "Positive getting of filtered tasks",
			args: args{
				filter: filter,
				userId: userId,
			},
			fields: fields{
				todoRepository: repository,
				ctx:            context.TODO(),
				wg:             &sync.WaitGroup{},
			},
			want:    []*model.Task{task},
			wantErr: false,
			mock: func() {
				repository.EXPECT().GetTasks(filter, userId).Return([]*model.Task{task}, nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
				ctx:            tt.fields.ctx,
				wg:             tt.fields.wg,
			}
			got, err := s.GetTasks(tt.args.filter, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoService.GetTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoService.GetTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoService_GetAllTasks(t *testing.T) {
	type fields struct {
		todoRepository model.Repository
		ctx            context.Context
		wg             *sync.WaitGroup
	}
	type args struct {
		userId int32
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_model.NewMockRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Task
		wantErr bool
		mock    func()
	}{
		{
			name: "Positive getting of filtered tasks",
			args: args{
				userId: userId,
			},
			fields: fields{
				todoRepository: repository,
				ctx:            context.TODO(),
				wg:             &sync.WaitGroup{},
			},
			want:    []*model.Task{task},
			wantErr: false,
			mock: func() {
				repository.EXPECT().GetAllTasks(userId).Return([]*model.Task{task}, nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
				ctx:            tt.fields.ctx,
				wg:             tt.fields.wg,
			}
			got, err := s.GetAllTasks(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoService.GetAllTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoService.GetAllTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoService_Login(t *testing.T) {
	type fields struct {
		todoRepository model.Repository
		ctx            context.Context
		wg             *sync.WaitGroup
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantUserId int32
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TodoService{
				todoRepository: tt.fields.todoRepository,
				ctx:            tt.fields.ctx,
				wg:             tt.fields.wg,
			}
			gotUserId, err := s.Login(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserId != tt.wantUserId {
				t.Errorf("TodoService.Login() = %v, want %v", gotUserId, tt.wantUserId)
			}
		})
	}
}

func Test_checkPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPasswordHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("checkPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		todoRepository model.Repository
		ctx            context.Context
		wg             *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
		want *TodoService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.todoRepository, tt.args.ctx, tt.args.wg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
