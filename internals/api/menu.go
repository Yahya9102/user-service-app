package api

import (
	"fmt"
	"user-service-app/internals/service"
)

func RuntMenu(userService *service.UserService) {

	for {
		fmt.Println("1. Create user")
		fmt.Println("2. List user")
		fmt.Println("3. Delete user")
		fmt.Println("4. Update user")
		fmt.Println("5. Exit")
	

		var choice int 
		fmt.Print("Pick an option")
		fmt.Scanln(&choice)


		switch choice {
			
		case 1:
			var name string
			var age int
			fmt.Print("Name:...")
			fmt.Scanln(&name)
			fmt.Print("Age:...")
			fmt.Scanln(&age)

			user := userService.CreateUser(name, age)

			fmt.Println("User created", user)

		case 2:
			users := userService.ListUsers()
			
			if len(users) == 0 {
				fmt.Println("No users found")
			} else {
				for _, u := range users {
					fmt.Printf("ID: %d, Name: %s, Age: %d \n", u.ID, u.Name, u.Age)
				}
			}

		case 3:
			var id int
			fmt.Print("Type in ID of the user you want to delete")
			fmt.Scanln(&id)

			if userService.DeleteUser(id) {
				fmt.Println("User deleted")
			} else {
				fmt.Println("User not found")
			}

		case 4:
			var id int
			var name string
			var age int

			fmt.Print("Type in ID of the user you want to update: ")
			fmt.Scanln(&id)

			fmt.Print("New name: ")
			fmt.Scanln(&name)

			fmt.Print("New age: ")
			fmt.Scanln(&age)

			user, ok := userService.UpdateUser(id, name, age)
			if !ok {
				fmt.Println("User not found")
			} else {
				fmt.Println("User updated:", user)
			}

		case 5:	
			return
		

		default:
			fmt.Println("Invalid choice, pick again")	

		}
		
	}

}