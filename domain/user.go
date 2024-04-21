package domain

import "github.com/PARKNAMSU/user-management/model/user_type"

type UserInterface interface {
	GetUserPermissionByID() user_type.UserPermissionSchema
}

type UserServiceInterface interface {
	SetUser()
	GetUserByToken()
}
