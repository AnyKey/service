package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"time"
)

type Repository struct {
	Redis *redis.Client
}
// New will create new an Repository object representation of user.Repository interface
func New(redis *redis.Client) *Repository {
	return &Repository{
		Redis: redis,
	}
}

func (ur *Repository) SetToken(ctx context.Context, bytes []byte) {
	ur.Redis.Set(ctx, "JWT:", bytes, 60*time.Minute)
	return

}

func (ur *Repository) GetToken(ctx context.Context) *string {
	var token string
	res := ur.Redis.Get(ctx, "JWT:")
	if res.Err() != nil {
		log.Errorln(res.Err())
		return nil
	}
	bytes, err := res.Bytes()
	if err != nil {
		return nil
	}
	err = json.Unmarshal(bytes, &token)
	if err != nil {
		return nil
	}
	return &token
}
