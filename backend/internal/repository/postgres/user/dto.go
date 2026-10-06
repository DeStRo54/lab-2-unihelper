package user

import (
	"time"

	"homewormanager/internal/entity"
)

type user struct {
	UserID    entity.UserID  `db:"user_id"`
	Name      string         `db:"name"`
	Surname   *string        `db:"surname"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	Role      entity.Role    `db:"role"`
	GroupID   entity.GroupID `db:"group_id"`
	CreatedAt time.Time      `db:"created_at"`
}

func (u user) toUser() entity.User {
	return entity.User{
		UserID:    u.UserID,
		Name:      u.Name,
		Surname:   u.Surname,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		GroupID:   u.GroupID,
		CreatedAt: u.CreatedAt,
	}
}

type userFull struct {
	UserID    entity.UserID `db:"user_id"`
	Name      string        `db:"name"`
	Surname   *string       `db:"surname"`
	Email     string        `db:"email"`
	Password  string        `db:"password"`
	Role      entity.Role   `db:"role"`
	GroupName string        `db:"group_name"`
}

func (u userFull) toUserFull() entity.UserFullInfo {
	return entity.UserFullInfo{
		UserID:    u.UserID,
		Name:      u.Name,
		Surname:   u.Surname,
		Email:     u.Email,
		Role:      u.Role,
		GroupName: u.GroupName,
	}
}

func toUsersFull(users []userFull) []entity.UserFullInfo {
	result := make([]entity.UserFullInfo, len(users))
	for i, u := range users {
		result[i] = u.toUserFull()
	}
	return result
}
