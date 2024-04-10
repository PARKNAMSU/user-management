package user_type

type UserPermissionSchema struct {
	UserId       int `json:"userId" db:"userId"`
	PermissionId int `json:"permissionId" db:"permissionId"`
	Permission   int `json:"permission" db:"permission"`
}
