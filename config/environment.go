package config

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Set default environments
var Env = struct {
	Env     string `mapstructure:"ENV"`                              // ระบุ environment ที่ใช้งาน เช่น development, staging, production
	Port    string `mapstructure:"PORT"`                             // พอร์ตที่แอปจะรันอยู่
	Cors    string `mapstructure:"CORS"`                             // รายการ origin ที่อนุญาต (CORS)
	AppHost string `mapstructure:"APP_HOST" validate:"required,uri"` // Host ของแอปพลิเคชัน

	// Signature (used for secure signing and token expiration)
	SignatureKey string        `mapstructure:"SIGNATURE_KEY"`     // คีย์สำหรับการสร้าง signature
	SignatureExp time.Duration `mapstructure:"SIGNATURE_EXPIRED"` // ระยะเวลาหมดอายุของ signature

	// Database configuration
	DBHost   string `mapstructure:"DB_HOST" validate:"required"`     // Host สำหรับเชื่อมต่อฐานข้อมูล
	DBPort   string `mapstructure:"DB_PORT" validate:"required"`     // Port สำหรับเชื่อมต่อฐานข้อมูล
	DBUser   string `mapstructure:"DB_USER" validate:"required"`     // User สำหรับเชื่อมต่อฐานข้อมูล
	DBPasswd string `mapstructure:"DB_PASSWORD" validate:"required"` // Password สำหรับเชื่อมต่อฐานข้อมูล
	DBName   string `mapstructure:"DB_NAME" validate:"required"`     // ชื่อฐานข้อมูล
}{
	Env:          "production",
	Port:         "8000",
	Cors:         "*",
	AppHost:      "http://localhost:8000",
	SignatureKey: "blog-api",
	SignatureExp: 24 * time.Hour,
}

func NewAppInitEnvironment() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Environment variables not used from .env")
	}

	// try load settings from env vars
	r := reflect.TypeOf(Env)
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i).Tag.Get("mapstructure")
		viper.BindEnv(f)
	}

	if err := viper.Unmarshal(&Env); err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(Env)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("🚫 Not found env :", err.Namespace(), err.Tag())
		}
		log.Fatalf("\n\n❌ Config failed server not start !!!\n\n")
	}

}
