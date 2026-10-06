package group

import (
	"homewormanager/internal/entity"
)

type group struct {
	GroupID     entity.GroupID `db:"group_id"`
	Name        string         `db:"group_name"`
	Course      int8           `db:"course"`
	IcalLink    string         `db:"ical_link"`
	RegisterKey string         `db:"register_key"`
}

func (g group) toGroup() entity.Group {
	return entity.Group{
		GroupID:     g.GroupID,
		Name:        g.Name,
		Course:      g.Course,
		IcalLink:    g.IcalLink,
		RegisterKey: g.RegisterKey,
	}
}
