package service

import (
	"20241112/pre/1/repository"
)

type MyService struct {
	MyRepository *repository.MyRepository
}

func NewMyService(myRepository *repository.MyRepository) *MyService {
	return &MyService{MyRepository: myRepository}
}
