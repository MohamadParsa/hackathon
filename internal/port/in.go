package port

import "github.com/MohamadParsa/hackathon/internal/model"

type QuickAccessApi interface {
	GetQuickAccessList(user string) ([]byte, int)
	GetSpecificQuickAccess(user string, id string) ([]byte, int)
	AddQuickAccess(quickAccess model.QuickAccess) int
	UpdateQuickAccess(quickAccess model.QuickAccess) int
	DeleteQuickAccess(user string, id string) int
}
type SuggestionApi interface {
	GetSuggestionList(user string) ([]byte, int)
}
