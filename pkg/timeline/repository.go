package timeline

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type TimelineRepository struct {
	redisKey string
	rdb      *redis.Client
}

func NewTimelineRepository() *TimelineRepository {
	return &TimelineRepository{
		redisKey: "timeline",
		rdb: redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (repo *TimelineRepository) AddTimeline(ctx context.Context, post Post) error {
	member, err := json.Marshal(&PostJson{
		Title:       post.Title,
		Description: post.Description,
	})
	if err != nil {
		return err
	}
	if err = repo.rdb.ZAdd(ctx, repo.redisKey, redis.Z{
		Score:  float64(post.CreatedAt),
		Member: member,
	}).Err(); err != nil {
		return err
	}
	return nil
}

func (repo *TimelineRepository) GetTimeline(ctx context.Context, limit int, offset int) ([]*Post, error) {
	start := offset
	stop := start + limit - 1
	serializedMembersWithScores, err := repo.rdb.ZRevRangeWithScores(ctx, repo.redisKey, int64(start), int64(stop)).Result()
	if err != nil {
		return nil, err
	}
	member := &PostJson{}
	res := make([]*Post, 0, limit)
	for _, serializedMembersWithScore := range serializedMembersWithScores {
		serializedMember := serializedMembersWithScore.Member
		score := serializedMembersWithScore.Score
		if err := json.Unmarshal([]byte(serializedMember.(string)), member); err != nil {
			return nil, err
		}
		post := &Post{
			Title:       member.Title,
			Description: member.Description,
			CreatedAt:   int64(score),
		}
		res = append(res, post)
	}
	return res, nil
}
