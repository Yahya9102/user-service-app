package repository

import (
	"testing"
	"user-service-app/internals/models"
)

/*Test add user function*/


func TestAddUser(t *testing.T){
	repo := NewUserRepository()
	user := models.User {
		Name: "Yahya",
		Age: 25,
	}
	result := repo.Add(user)

	if result.ID != 1 {
		t.Errorf("expected ID to be 1, got %d",result.ID)
	}

	if result.Name != "Yahya" {
		t.Errorf("expected name to be Yahya, got %s", result.Name)
	}

	if result.Age != 25 {
		t.Errorf("expected age to be 25, got %d", result.Age)
	}

	if len(repo.users) != 1 {
		t.Errorf("expected repo length to be 1, got %d", len(repo.users))
	}
}


/**/

func TestGetAllUser(t *testing.T) {
	repo := NewUserRepository()

	repo.Add(models.User{Name: "A", Age: 20})
	repo.Add(models.User{Name: "C", Age: 30})
	

	
	users := repo.GetAll()

	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}

	if users[0].Name != "A" {
		t.Errorf("expected first user to be A, got %s", users[0].Name)
	}

	if users[1].Name != "B" {
		t.Errorf("expected first user to be B, got %s", users[1].Name)
	}
}