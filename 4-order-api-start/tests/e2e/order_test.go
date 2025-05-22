package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"order-api/internal/auth"
	"order-api/internal/core/models"
	"order-api/internal/order"
	"os"
	"testing"
)

func initDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"))

	testDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to DB")
	}

	err = testDB.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{})
	if err != nil {
		log.Fatal(err)
	}
	return testDB
}

func initData(db *gorm.DB) {
	db.Create(&models.User{
		PhoneNumber: "+79112223344",
		SessionID:   "CCa5XNkv5al5",
	})

	db.Create(&models.Product{
		Article:     10000000001,
		Name:        "Яблоки",
		Price:       123.45,
		Description: "Сладкие",
		Image:       []string{"/home/product1.jpeg"},
	})
}

func removeData(db *gorm.DB, orderID uint) {
	db.Unscoped().Where("phone_number = ?", "+79112223344").Delete(&models.User{})
	db.Unscoped().Where("name = ?", "Яблоки").Delete(&models.Product{})
	db.Unscoped().Where("id = ?", orderID).Delete(&models.Order{})
}

func setup() (*httptest.Server, *gorm.DB) {
	db := initDB()
	initData(db)
	ts := httptest.NewServer(App())

	return ts, db
}

func getUserToken(t *testing.T, ts *httptest.Server) string {
	// Получение JWT
	// 	 получение sessionID
	loginData, _ := json.Marshal(auth.LoginRequest{Number: "+79112223344"})
	createUserRes, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewBuffer(loginData))
	if err != nil {
		t.Fatal(err)
	}
	body, _ := io.ReadAll(createUserRes.Body)

	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.SessionId == "" {
		t.Fatal("SessionId is empty")
	}

	//  получение токена
	verifyData, _ := json.Marshal(&auth.VerifyRequest{
		SessionId: resData.SessionId,
		Code:      345678,
	})

	verifyRes, err := http.Post(ts.URL+"/auth/verify", "application/json", bytes.NewBuffer(verifyData))
	if err != nil {
		t.Fatal(err)
	}
	verifyBody, _ := io.ReadAll(verifyRes.Body)

	var verifyResData auth.VerifyResponse
	err = json.Unmarshal(verifyBody, &verifyResData)
	if err != nil {
		t.Fatal(err)
	}

	if verifyResData.Token == "" {
		t.Fatal("Token is empty")
	}
	return verifyResData.Token
}

func makeRequest(t *testing.T, ts *httptest.Server, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, ts.URL+"/order", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+getUserToken(t, ts))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return res, err
}

func TestOrder_Create_Success(t *testing.T) {
	ts, db := setup()
	defer removeData(db, 1)
	defer ts.Close()

	// Подготавливаем body для запроса
	data, _ := json.Marshal(&order.CreateOrderRequest{Products: []int{10000000001}})

	// Делаем запрос
	res, err := makeRequest(t, ts, data)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("got %d, want %d", res.StatusCode, http.StatusCreated)
	}

	orderBody, _ := io.ReadAll(res.Body)
	var orderResData order.CreateOrderResponse
	err = json.Unmarshal(orderBody, &orderResData)

	if orderResData.ProductsCount != 1 {
		t.Errorf("got %d, want %d", orderResData.ProductsCount, 1)
	}

	if orderResData.TotalCost != 123.45 {
		t.Errorf("got %f, want %f", orderResData.TotalCost, 123.45)
	}
}

func TestOrder_Create_Fail(t *testing.T) {
	ts, db := setup()
	defer removeData(db, 1)
	defer ts.Close()

	// Подготавливаем body для запроса
	data, _ := json.Marshal(&order.CreateOrderRequest{Products: []int{}})

	// Делаем запрос
	res, err := makeRequest(t, ts, data)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("got %d, want %d", res.StatusCode, http.StatusCreated)
	}
}

func TestOrder_Create_UnknownArticle(t *testing.T) {
	ts, db := setup()
	defer removeData(db, 1)
	defer ts.Close()

	// Подготавливаем body для запроса
	data, _ := json.Marshal(&order.CreateOrderRequest{Products: []int{10000000099}})

	// Делаем запрос
	res, err := makeRequest(t, ts, data)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("got %d, want %d", res.StatusCode, http.StatusCreated)
	}
}
