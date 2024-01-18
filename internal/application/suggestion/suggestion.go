package suggestion

import (
	"encoding/json"
	"net/http"

	"github.com/MohamadParsa/hackathon/internal/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Suggestion struct {
}

func New() *Suggestion {
	return &Suggestion{}
}
func (suggestion *Suggestion) GetSuggestionList(user string) ([]byte, int) {
	items := []*model.Suggestion{
		{
			Id:          uuid.New().String(),
			Type:        "cab",
			Title:       "خونه به شرکت",
			Picture:     "https://hackathon.storage.iran.liara.space/cab.png",
			DeepLink:    "https://app.snapp.taxi/pre-ride?rideFrom={%22options%22:{%22serviceType%22:1,%22recommender%22:%22cab%22}}",
			Description: "دیر به شرکت نرسی، همین الان اسنپ بگیر",
		},
		{
			Id:          uuid.New().String(),
			Type:        "food",
			Title:       "شفارش از باگت (جردن)",
			Picture:     "https://hackathon.storage.iran.liara.space/baget.jpg",
			DeepLink:    "https://superapp.snappfood.ir/",
			Description: "تخفیف ۲۵٪برای انتخاب همیشگی",
		},
		{
			Id:          uuid.New().String(),
			Type:        "food",
			Title:       "شفارش از سالانی",
			Picture:     "https://hackathon.storage.iran.liara.space/salani.jpeg",
			DeepLink:    "https://superapp.snappfood.ir/",
			Description: "همون انتخاب همیشگی برای نهار",
		},
		{
			Id:          uuid.New().String(),
			Type:        "food",
			Title:       "شفارش از شیلا (پارک ملت)",
			Picture:     "https://hackathon.storage.iran.liara.space/shila.jpg",
			DeepLink:    "https://superapp.snappfood.ir/",
			Description: "یک پیشنهاد جدید برای نهار",
		},
		{
			Id:          uuid.New().String(),
			Type:        "cab",
			Title:       "شرکت به خونه",
			Picture:     "https://hackathon.storage.iran.liara.space/cab.png",
			DeepLink:    "https://app.snapp.taxi/pre-ride?rideFrom={%22options%22:{%22serviceType%22:1,%22recommender%22:%22cab%22}}",
			Description: "میخوای برای ساعت 17:00 اسنپ جلوی شرکت باشه",
		},
	}
	byteResult, err := json.Marshal(items)
	if err != nil {
		log.Errorf("converting GetSuggestionList result to byte returns error: %v", err)
		return nil, http.StatusInternalServerError
	}
	return byteResult, http.StatusOK
}
