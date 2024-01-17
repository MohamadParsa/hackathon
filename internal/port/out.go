package port

import "github.com/MohamadParsa/hackathon/internal/model"

type Db interface {
	GetQuickAccessList(user string) []*model.QuickAccess
	GetSpecificQuickAccess(user string, id string) *model.QuickAccess
	InsertQuickAccess(quickAccess *model.QuickAccess) error
	UpdateQuickAccess(quickAccess *model.QuickAccess) error
	DeleteQuickAccess(user string, id string) error
}
