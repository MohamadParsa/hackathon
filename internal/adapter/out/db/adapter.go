package db

import (
	"github.com/MohamadParsa/hackathon/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Adapter struct {
	db *gorm.DB
}

func New(src string) (*Adapter, error) {
	gdb, err := gorm.Open(postgres.Open(src), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return &Adapter{
		db: gdb,
	}, nil
}

func (adapter *Adapter) GetQuickAccessList(user string) []*model.QuickAccess {
	return []*model.QuickAccess{{}}
}
func (adapter *Adapter) GetSpecificQuickAccess(user string, id string) *model.QuickAccess {
	return &model.QuickAccess{}

}
func (adapter *Adapter) InsertQuickAccess(quickAccess *model.QuickAccess) error {
	return nil
}
func (adapter *Adapter) UpdateQuickAccess(quickAccess *model.QuickAccess) error {
	return nil

}
func (adapter *Adapter) DeleteQuickAccess(user string, id string) error {
	return nil

}
