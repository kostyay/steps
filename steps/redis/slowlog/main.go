package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/stackpulse/steps-sdk-go/env"
	"github.com/stackpulse/steps-sdk-go/step"
	"github.com/stackpulse/steps/redis/base"
)

type Args struct {
	base.Args
	LastEntries int64 `env:"LAST_ENTRIES" envDefault:"10"`
}

type RedisSlowLog struct {
	args        Args
	redisClient *redis.Client
}

func (l *RedisSlowLog) Init() error {
	err := env.Parse(&l.args)
	if err != nil {
		return err
	}

	l.redisClient, err = base.InitRedisClient(l.args.Args)

	return err
}

func (l *RedisSlowLog) Run() (exitCode int, output []byte, err error) {
	cmd := l.redisClient.SlowLogGet(context.Background(), l.args.LastEntries)
	if cmd == nil {
		return step.ExitCodeFailure, nil, fmt.Errorf("invalid result returned from slowlog operation")
	}
	result := cmd.Val()
	if result == nil {
		return step.ExitCodeFailure, nil, fmt.Errorf("run command")
	}

	jsonOutput, err := json.Marshal(result)
	if err != nil {
		return step.ExitCodeFailure, nil, fmt.Errorf("marshal output: %w", err)
	}

	return step.ExitCodeOK, jsonOutput, nil
}

func main() {
	step.Run(&RedisSlowLog{})
}
