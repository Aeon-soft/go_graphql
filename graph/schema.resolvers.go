package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"github.com/pramishkarkee/gql-test-struct/graph/generated"
	"github.com/pramishkarkee/gql-test-struct/graph/model"
)

func (r *mutationResolver) CreateMovie(ctx context.Context, input model.NewMovie) (*model.Movie, error) {
	movie := model.Movie{
		Title: input.Title,
		URL:   input.URL,
		ID: fmt.Sprintf("T%d", rand.Int()),
	}
	fmt.Println("******************",movie)
	// _, err := r.DB.Model(&movie).Insert()
	// if err != nil {
	// 	return nil, fmt.Errorf("error inserting new movie: %v", err)
	// }
	if result := r.DB.Create(&movie); result.Error != nil {
        return nil, fmt.Errorf("error inserting new movie: %v", result.Error)
    }
 
	return &movie, nil
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	var movies []*model.Movie
	fmt.Println("******************",movies)
   err := r.DB.First(&movies)
   if err != nil {
       return nil, fmt.Errorf("no %v",err)
   }

   return movies, nil
// panic(fmt.Errorf("not implemented"))

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
