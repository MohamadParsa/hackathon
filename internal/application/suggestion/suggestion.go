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
			Title:       "سفارش از پیتزا هات (سهروردی)",
			Picture:     "https://cdn.snappfood.ir/media/cache/vendor_logo/uploads/images/vendors/logos/2938_635791366331279296_s.jpg",
			DeepLink:    "https://snappfood.ir/restaurant/menu/%D9%BE%DB%8C%D8%AA%D8%B2%D8%A7_%D9%87%D8%A7%D8%AA__%D8%B3%D9%87%D8%B1%D9%88%D8%B1%D8%AF%DB%8C_-r-pz26wp",
			Description: "تخفیف ۲۵٪برای انتخاب همیشگی",
		},
		{
			Id:          uuid.New().String(),
			Type:        "food",
			Title:       "سفارش از رستوران نایب (سهروردی)",
			Picture:     "https://cdn.snappfood.ir/media/cache/vendor_logo/uploads/images/vendors/logos/270_634658443922812500_s.jpg",
			DeepLink:    "https://snappfood.ir/restaurant/menu/%D8%B1%D8%B3%D8%AA%D9%88%D8%B1%D8%A7%D9%86_%D9%86%D8%A7%DB%8C%D8%A8__%D8%B3%D9%87%D8%B1%D9%88%D8%B1%D8%AF%DB%8C_-r-0lre90",
			Description: "همون انتخاب همیشگی برای نهار",
		},
		{
			Id:          uuid.New().String(),
			Type:        "food",
			Title:       "شفارش از کباب بناب کاج شریعتی",
			Picture:     "https://cdn.snappfood.ir/media/cache/vendor_logo/uploads/images/vendors/logos/5f786cf6d3a97.jpg",
			DeepLink:    "https://snappfood.ir/restaurant/menu/%DA%A9%D8%A8%D8%A7%D8%A8_%D8%A8%D9%86%D8%A7%D8%A8_%DA%A9%D8%A7%D8%AC_%D8%B4%D8%B1%DB%8C%D8%B9%D8%AA%DB%8C__%D8%B4%D8%B9%D8%A8%D9%87_%D8%AD%D8%B3%DB%8C%D9%86%DB%8C%D9%87_%D8%A7%D8%B1%D8%B4%D8%A7%D8%AF_-r-p67vyn",
			Description: "یک پیشنهاد جدید برای نهار",
		},
		{
			Id:          uuid.New().String(),
			Type:        "cab",
			Title:       "شرکت به خونه",
			Picture:     "https://hackathon.storage.iran.liara.space/cab.png",
			DeepLink:    "https://snappfood.ir/restaurant/menu/%DA%A9%D8%A8%D8%A7%D8%A8_%D8%A8%D9%86%D8%A7%D8%A8_%DA%A9%D8%A7%D8%AC_%D8%B4%D8%B1%DB%8C%D8%B9%D8%AA%DB%8C__%D8%B4%D8%B9%D8%A8%D9%87_%D8%AD%D8%B3%DB%8C%D9%86%DB%8C%D9%87_%D8%A7%D8%B1%D8%B4%D8%A7%D8%AF_-r-p67vyn",
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
