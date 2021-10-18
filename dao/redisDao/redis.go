package redisDao

import (
	"errors"
	goRedis "github.com/go-redis/redis/v8"
	"sync"
)

var (
	ErrNoSuchRedisConn       = errors.New("no such redis connection")
	ErrAliasNameAlreadyExist = errors.New("alias name already exist")
)

type Client struct {
	Client *goRedis.Client
	Config *goRedis.Options
	lock   *sync.Mutex
}

var redisMap = make(map[string]*Client)

func AddClient(aliasName string, config *goRedis.Options) error {
	if _, has := redisMap[aliasName]; has {
		return ErrAliasNameAlreadyExist
	}
	redisMap[aliasName] = &Client{
		Client: goRedis.NewClient(config),
		Config: config,
		lock:   new(sync.Mutex),
	}
	return nil
}

func Get(aliasName string) (*goRedis.Client, error) {
	if client, has := redisMap[aliasName]; has {
		return client.Client, nil
	}
	return nil, ErrNoSuchRedisConn
}
