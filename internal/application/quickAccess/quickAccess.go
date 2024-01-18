package quickAccess

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/MohamadParsa/hackathon/internal/model"
	"github.com/MohamadParsa/hackathon/internal/port"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
func (quickAccess *QuickAccess) PurcahseHistory(user string, serviceType string) ([]byte, int) {
	items := model.PurcahseHistoryList{
		{
			OrderId:     uuid.New().String(),
			Type:        "cab",
			Title:       "درخواست تاکسی",
			Description: "از جردن، نجف دریابندری به بلوار سعادت آباد، هجدهم غربی",
		},
		{
			OrderId:     uuid.New().String(),
			Type:        "cab",
			Title:       "درخواست تاکسی",
			Description: "از یادگار امام، بوستان سعادت به شهید محمد کچویی",
		},
		{
			OrderId:     uuid.New().String(),
			Type:        "cab",
			Title:       "درخواست تاکسی",
			Description: "از ولی عصر، استاد کردوانی به امامزاده قاسم، دکتر باهنر",
		},
		{
			OrderId:     uuid.New().String(),
			Type:        "food",
			Title:       "درخواست عذا از شیلا جردن",
			Description: "استرامبولی استیک، چیکن چیز... برای آدرس بوستان سعادت آباد",
		},
		{
			OrderId:     uuid.New().String(),
			Type:        "food",
			Title:       "درخواست غذا از باگت",
			Description: "پیتزا گوشت و قارچ، سیب زمینی سرخ کرده... برای آدرس بوستان سعادت آباد",
		},
		{
			OrderId:     uuid.New().String(),
			Type:        "express",
			Title:       "درخواست اسنپ اکسپرس از اسمارت میرداماد",
			Description: "سفارش هایپ، شیر... برای آدرس بلوار سعادت آباد",
		},
		{
			OrderId:     uuid.New().String(),
			Type:        "doctor",
			Title:       "درخواست دارو",
			Description: "قرص کلسیم برای آدرس بلوار سعادت آباد",
		},
	}
	if serviceType != "" {
		items = items.FilerByType(serviceType)
	}
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

func (quickAccess *QuickAccess) UploadFile(fileContent io.Reader, fileName string) ([]byte, int) {
	extension := filepath.Ext(fileName)
	if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".svg" {
		log.Errorf("file type is not valid: %v", extension)
		return nil, http.StatusBadRequest
	}
	newFileName := uuid.New().String() + extension
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(viper.GetString("AWS_REGION")))
	if err != nil {
		log.Errorf("LoadDefaultConfig error: %v", err)
		return nil, http.StatusInternalServerError
	}

	// Define AWS credentials and bucket information
	awsAccessKeyID := viper.GetString("AWS_ACCESS_KEY")
	awsSecretAccessKey := viper.GetString("AWS_SECRET_KEY")
	endpoint := viper.GetString("AWS_ENDPOINT")
	bucketName := viper.GetString("AWS_BUCKET_NAME")

	// Initialize S3 client with custom configuration
	cfg.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     awsAccessKeyID,
			SecretAccessKey: awsSecretAccessKey,
		}, nil
	})

	cfg.BaseEndpoint = aws.String(endpoint)

	client := s3.NewFromConfig(cfg)

	// Specify the destination key in the bucket
	destinationKey := "uploads/" + newFileName

	// Use the S3 client to upload the file
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(destinationKey),
		Body:   fileContent,
	})
	if err != nil {
		log.Errorf("PutObjectInputreturns error: %v", err)
		return nil, http.StatusInternalServerError
	}
	type uploadResult struct {
		FileName string `json:"fileName"`
	}
	result := uploadResult{FileName: newFileName}
	byteResult, err := json.Marshal(result)
	if err != nil {
		log.Errorf("converting uploadResult result to byte returns error: %v", err)
		return nil, http.StatusInternalServerError
	}
	return byteResult, http.StatusOK
}
