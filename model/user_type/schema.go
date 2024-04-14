package user_type

import (
	"database/sql"
	"time"
)

type UserSchema struct {
	UserId    uint         `json:"userId" db:"userId"`
	UserIp    string       `json:"userIp" db:"userIp"`
	Status    uint         `json:"status" db:"status"`
	CreatedAt time.Time    `json:"createdAt" db:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt" db:"updatedAt"`
}

type UserInformationSchema struct {
	LoginId     string         `json:"loginId" db:"loginId"`
	Password    string         `json:"password" db:"password"`
	VisibleName sql.NullString `json:"visibleName" db:"visibleName"`
}

type UserPermissionSchema struct {
	UserId       uint `json:"userId" db:"userId"`
	PermissionId uint `json:"permissionId" db:"permissionId"`
	Permission   uint `json:"permission" db:"permission"`
}
