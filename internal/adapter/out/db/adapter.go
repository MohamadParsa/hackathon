package db

import (
	"github.com/MohamadParsa/hackathon/internal/model"
	"github.com/google/uuid"
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
	return []*model.QuickAccess{
		{
			Id:      uuid.New().String(),
			UserId:  user,
			Type:    "cab",
			Picture: "",
			Action: &model.Action{
				Id:            uuid.New().String(),
				OrderId:       uuid.New().String(),
				ActionCommand: "",
			},
			Title: "خونه عباس",
		},
		{
			Id:      uuid.New().String(),
			UserId:  user,
			Type:    "food",
			Picture: "",
			Action: &model.Action{
				Id:            uuid.New().String(),
				OrderId:       uuid.New().String(),
				ActionCommand: "",
			},
			Title: "نان - خونه",
		},
		{
			Id:      uuid.New().String(),
			UserId:  user,
			Type:    "cab",
			Picture: "",
			Action: &model.Action{
				Id:            uuid.New().String(),
				OrderId:       uuid.New().String(),
				ActionCommand: "",
			},
			Title: "خونه علی",
		},
		{
			Id:      uuid.New().String(),
			UserId:  user,
			Type:    "cab",
			Picture: "",
			Action: &model.Action{
				Id:            uuid.New().String(),
				OrderId:       uuid.New().String(),
				ActionCommand: "",
			},
			Title: "سارا",
		},
		{
			Id:      uuid.New().String(),
			UserId:  user,
			Type:    "food",
			Picture: "",
			Action: &model.Action{
				Id:            uuid.New().String(),
				OrderId:       uuid.New().String(),
				ActionCommand: "",
			},
			Title: "گوشت خونه",
		},
		{
			Id:      uuid.New().String(),
			UserId:  user,
			Type:    "food",
			Picture: "",
			Action: &model.Action{
				Id:            uuid.New().String(),
				OrderId:       uuid.New().String(),
				ActionCommand: "",
			},
			Title: "چلوبرگ فرشته",
		},
	}
}
func (adapter *Adapter) GetSpecificQuickAccess(user string, id string) *model.QuickAccess {
	return &model.QuickAccess{
		Id:      uuid.New().String(),
		UserId:  user,
		Type:    "cab",
		Picture: "",
		Action: &model.Action{
			Id:            uuid.New().String(),
			OrderId:       uuid.New().String(),
			ActionCommand: "",
		},
		Title: "خونه عباس",
	}

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
