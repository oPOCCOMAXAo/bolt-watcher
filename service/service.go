package service

import (
	"bolt-watcher/bolt"
)

type Service struct {
	API *bolt.Client
}

func New(api *bolt.Client) *Service {
	return &Service{
		API: api,
	}
}
