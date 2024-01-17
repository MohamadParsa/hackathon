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
			Id:       uuid.New().String(),
			Type:     "taxi",
			Title:    "خونه به شرکت",
			Picture:  "",
			DeepLink: "https://app.snapp.taxi/pre-ride?rideFrom={%22options%22:{%22serviceType%22:1,%22recommender%22:%22cab%22}}",
		},
		{
			Id:       uuid.New().String(),
			Type:     "food",
			Title:    "برگراتور",
			Picture:  "",
			DeepLink: "https://superapp.snappfood.ir/",
		},
		{
			Id:       uuid.New().String(),
			Type:     "food",
			Title:    "آشپزخانه مزه",
			Picture:  "",
			DeepLink: "https://superapp.snappfood.ir/",
		},
		{
			Id:       uuid.New().String(),
			Type:     "food",
			Title:    "آشپزخانه پاچین",
			Picture:  "",
			DeepLink: "https://superapp.snappfood.ir/",
		},
		{
			Id:       uuid.New().String(),
			Type:     "taxi",
			Title:    "شرکت به خونه",
			Picture:  "",
			DeepLink: "https://app.snapp.taxi/pre-ride?rideFrom={%22options%22:{%22serviceType%22:1,%22recommender%22:%22cab%22}}",
		},
	}
	byteResult, err := json.Marshal(items)
	if err != nil {
		log.Errorf("converting GetSuggestionList result to byte returns error: %v", err)
		return nil, http.StatusInternalServerError
	}
	return byteResult, http.StatusOK
}
