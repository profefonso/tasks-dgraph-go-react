package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/profefonso/tasks-dgraph-go-react/graph/generated"
	"github.com/profefonso/tasks-dgraph-go-react/graph/model"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.DataTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.DataCategory) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input model.DataTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, input model.DataCategory) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	var tasks []*model.Task

	dgraphClient := DgraphClient()
	txn := dgraphClient.NewTxn()
	defer txn.Discard(ctx)

	task := &model.Task{
		ID:        fmt.Sprintf("T%d", rand.Int()),
		Title:     "entreda_descrip",
		Category:  &model.Category{ID: "87897987", Name: "Nova"},
		Completed: false,
	}

	tasks = append(tasks, task)

	pb, err := json.Marshal(task)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		SetJson: pb,
	}
	res, err := txn.Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

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
