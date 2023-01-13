package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/kanikim/src/entity"
)

type UserRepository interface {
	Get(id uuid.UUID) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Save(*entity.User) error
	Remove(id uuid.UUID) error
	Update(*entity.User) error
}

type UserRepositoryImpl struct {
	Conn *gorm.DB
}

func RealUserRepository(conn *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Conn: conn}
}

func (repo *UserRepositoryImpl) Get(id uuid.UUID) (*entity.User, error) {
	
}
