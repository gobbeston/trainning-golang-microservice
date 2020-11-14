package customer

import (
	"github.com/gobbeston/trainning-golang-microservice/app/entities"
	"github.com/gobbeston/trainning-golang-microservice/app/layers/repositories/customer"
)

type useCase struct {
	CustomerRepo customer.Repo
}

func InitUseCase(customerRepository customer.Repo) UseCase {
	return &useCase{
		CustomerRepo: customerRepository,
	}
}

// InitUseCase init auth use case
type UseCase interface {
	CreateUser(input *entities.Users) (*entities.Users, error)
	CreateRole(input *entities.Roles) (*entities.Roles, error)
}
