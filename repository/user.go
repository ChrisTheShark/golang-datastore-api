package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/datastore"
	"github.com/ChrisTheShark/golang-datastore-api/models"
)

// UserRepository inteface describes rrepository operations on Users
type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(string) (*models.User, error)
	Create(models.User) (string, error)
	Delete(models.User) error
}

// UserRepositoryImpl houses logic to retrieve users from a datastore repository
type UserRepositoryImpl struct {
	db *datastore.Client
}

// NewUserRepository convience function to create a UserRepository
func NewUserRepository(db *datastore.Client) UserRepository {
	return &UserRepositoryImpl{db}
}

// GetAll get all users from the repository
func (r UserRepositoryImpl) GetAll() ([]models.User, error) {
	users := []models.User{}

	keys, err := r.db.GetAll(context.Background(), datastore.NewQuery("users"), &users)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve users due to: %v", err)
	}

	for i, k := range keys {
		users[i].ID = strconv.FormatInt(k.ID, 10)
	}

	return users, nil
}

// GetByID get a user by string identifier
func (r UserRepositoryImpl) GetByID(id string) (*models.User, error) {
	identifier, err := strconv.Atoi(id)
	if err != nil {
		return nil, models.UserNotFoundError{
			Message: fmt.Sprintf("unable to locate user with id: %v", id),
		}
	}

	k := datastore.IDKey("users", int64(identifier), nil)
	user := new(models.User)

	if err := r.db.Get(context.Background(), k, user); err != nil {
		return nil, fmt.Errorf(
			"unable to get user id: %v due to: %v", id, err)
	}

	user.ID = id
	return user, nil
}

// Create a User to the repository
func (r UserRepositoryImpl) Create(user models.User) (string, error) {
	k := datastore.IncompleteKey("users", nil)
	key, err := r.db.Put(context.Background(), k, &user)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("unable to create user due to: %v", err)
	}

	return strconv.FormatInt(key.ID, 10), nil
}

// Delete a User from the repository
func (r UserRepositoryImpl) Delete(user models.User) error {
	identifier, err := strconv.Atoi(user.ID)
	if err != nil {
		return models.UserNotFoundError{
			Message: fmt.Sprintf("unable to delete user with id: %v", user.ID),
		}
	}

	k := datastore.IDKey("users", int64(identifier), nil)
	if err := r.db.Delete(context.Background(), k); err != nil {
		return fmt.Errorf("unable to delete user due to: %v", err)
	}

	return nil
}
