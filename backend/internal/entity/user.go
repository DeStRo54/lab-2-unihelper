package entity

import "time"

type User struct {
	UserID    UserID
	Name      string
	Surname   *string
	Email     string
	Password  string
	Role      Role
	GroupID   GroupID
	CreatedAt time.Time
}

type UserCreds struct {
	UserID  UserID
	Role    Role
	GroupID GroupID
}

type UserFullInfo struct {
	UserID    UserID
	Name      string
	Surname   *string
	Email     string
	Role      Role
	GroupName string
}
