package quickAccess

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/MohamadParsa/hackathon/internal/model"
	"github.com/MohamadParsa/hackathon/internal/port"
)

type QuickAccess struct {
	db port.Db
}

func New(db port.Db) *QuickAccess {
	return &QuickAccess{db: db}
}
func (quickAccess *QuickAccess) GetQuickAccessList(user string) ([]byte, int) {
	items := quickAccess.db.GetQuickAccessList(user)
	byteResult, err := json.Marshal(items)
	if err != nil {
		log.Errorf("converting GetQuickAccessList result to byte returns error: %v", err)
		return nil, http.StatusInternalServerError
	}
	return byteResult, http.StatusOK
}
func (quickAccess *QuickAccess) GetSpecificQuickAccess(user string, id string) ([]byte, int) {
	items := quickAccess.db.GetSpecificQuickAccess(user, id)
	byteResult, err := json.Marshal(items)
	if err != nil {
		log.Errorf("converting GetSpecificQuickAccess result to byte returns error: %v", err)
		return nil, http.StatusInternalServerError
	}
	return byteResult, http.StatusOK
}
func (quickAccess *QuickAccess) AddQuickAccess(item model.QuickAccess) int {
	err := quickAccess.db.InsertQuickAccess(&item)
	if err != nil {
		log.Errorf("InsertQuickAccess returns error: %v", err)
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
func (quickAccess *QuickAccess) UpdateQuickAccess(item model.QuickAccess) int {
	err := quickAccess.db.UpdateQuickAccess(&item)
	if err != nil {
		log.Errorf("UpdateQuickAccess returns error: %v", err)
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
func (quickAccess *QuickAccess) DeleteQuickAccess(user string, id string) int {
	err := quickAccess.db.DeleteQuickAccess(user, id)
	if err != nil {
		log.Errorf("InsertQuickAccess returns error: %v", err)
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
