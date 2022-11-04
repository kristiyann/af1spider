package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/kristiyann/af1-spider/models"
)

type LogicImpl interface {
	GetProductFromNike(params GetProductFromNikeParams) error
}

type logicImpl struct{}

func NewLogicImpl() LogicImpl {
	return &logicImpl{}
}

type GetProductFromNikeParams struct {
	Url  string
	Size string
}

func (l *logicImpl) GetProductFromNike(params GetProductFromNikeParams) error {
	params.Url = strings.Replace(params.Url, "https://", "", -1)
	strArr := strings.Split(params.Url, "/")
	identifier := strArr[len(strArr)-1]
	marketplace := strArr[1]

	actualUrl := "https://api.nike.com/product_feed/threads/v2?filter=language(en-GB)&filter=marketplace(" + strings.ToUpper(marketplace) + ")&filter=channelId(d9a5bc42-4b9c-4976-858a-f159cf99c647)&filter=productInfo.merchProduct.styleColor(" + identifier + ")"

	resp, err := http.Get(actualUrl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var responseModel models.NikeProductResponse
	err = json.NewDecoder(resp.Body).Decode(&responseModel)
	if err != nil {
		return err
	}

	product := responseModel.Objects[0]
	skuList := product.ProductInfo[0].Skus

	var desiredSku models.NikeProductSku

	for i := 0; i < len(skuList); i++ {
		if skuList[i].NikeSize == params.Size {
			desiredSku = skuList[i]
			break
		}
	}

	availableSkuList := product.ProductInfo[0].AvailableSkus

	for i := 0; i < len(availableSkuList); i++ {
		if availableSkuList[i].ID == desiredSku.ID {
			fmt.Println("request-id: " + uuid.New().String())
			fmt.Printf("product: %s \n", product.ProductInfo[0].ProductContent.FullTitle)
			fmt.Printf("size: %s \n", desiredSku.NikeSize)
			fmt.Printf("available: %t \n", availableSkuList[i].Available)
			fmt.Println("request-time: " + time.Now().Local().String() + "\n")

			if availableSkuList[i].Available {
				sendEmail(product.ProductInfo[0].ProductContent.FullTitle, params.Size, params.Url)
			} else {
				time.Sleep(15 * time.Second)
				l.GetProductFromNike(params)
			}
		}
	}

	return nil
}

func sendEmail(productName string, size string, link string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fromEmail := os.Getenv("FROM_EMAIL")
	fromPassword := os.Getenv("FROM_PASS")
	toEmail := os.Getenv("TO_EMAIL")

	auth := smtp.PlainAuth("", fromEmail, fromPassword, "smtp.gmail.com")

	body := "Subject: AF1SPIDER - Product Available\n" + productName + " is available in size " + size + " - " + link

	smtp.SendMail("smtp.gmail.com:587", auth, fromEmail, []string{toEmail}, []byte(body))
}
