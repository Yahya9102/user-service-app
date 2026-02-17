package service

import (
	"user-service-app/internals/models"
	"user-service-app/internals/repository"
)

type UserService struct {
	repo *repository.UserRepository
}


func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}



func  (s *UserService) CreateUser(name string, age int) models.User {
	user := models.User{Name: name, Age: age}
	return s.repo.Add(user)
}

func (s *UserService) ListUsers() []models.User {
	return s.repo.GetAll()
}

func (s *UserService) DeleteUser(id int) bool {
	return s.repo.Delete(id)
}


func (s *UserService) UpdateUser(id int, name string, age int) (models.User, bool) {
	updated := models.User{Name: name, Age: age}
	return s.repo.Update(id, updated)
}











// * DereferenceOperator  Hämtar värdet från en address
/*
x := 10

p := &x

fmt.Println(*p)
*/



// & ReferenceOperator Hämtar addressen till en variable
/*
 x := 10

 p := &x

 */
