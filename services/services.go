package services

import "github.com/ashish00304/test-test/model"

var Users []model.User

func AddUserDetails(user model.User) {
	Users = append(Users, user)
}

func GetAllUserDetails() []model.User {
	return Users
}

func DeleteUserDetails(id int) bool {
	for index, user := range Users {
		if user.ID == id {
			Users = append(Users[:index], Users[index+1:]...)
			return true
		}

	}
	return false
}

func UpdateUserDetails(id int, updatedUser model.User) (model.User, bool) {
	for index, user := range Users {
		if user.ID == id {
			Users[index].Name = updatedUser.Name
			Users[index].Role = updatedUser.Role
			Users[index].Phone_Number = updatedUser.Phone_Number
			return Users[index], true

		}
	}
	return model.User{}, false
}
