package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"main.go/graph/generated"
	"main.go/graph/model"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.DataTask) (*model.Task, error) {
	
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.DataCategory) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	var tasks []*model.Task
	dummyTask := model.Task{
		Title: "Tarea dummy",
		Completed: false,
		Category: &model.Category{Name: "testC", Description: "1212321"},
	}
	tasks = append(tasks, &dummyTask)
	return tasks, nil
	
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
