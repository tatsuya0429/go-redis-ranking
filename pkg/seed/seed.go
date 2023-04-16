package seed

import (
	"context"
	"fmt"
	"go-redis-sample/pkg/timeline"

	"github.com/go-faker/faker/v4"
)

var ctx = context.Background()

func Seed() {
	repo := timeline.NewTimelineRepository()
	for i := 0; i < 1000; i++ {
		post := timeline.Post{}
		err := faker.FakeData(&post)
		if err != nil {
			fmt.Println(err)
		}
		repo.AddTimeline(ctx, post)
	}
}

func main() {
	Seed()
}
