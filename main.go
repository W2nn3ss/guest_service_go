package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"guest-service/config"
	"guest-service/controllers"
	"guest-service/middleware"
	"guest-service/models"
	"guest-service/repository"
	"guest-service/service"
	"log"
	"os"
	"runtime"
	"time"
)

func GenerateToken(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func AddTokenToEnv(token string) error {
	fileEnv, err := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fileEnv.Close()

	_, err = fileEnv.WriteString(fmt.Sprintf("API_TOKEN=%s\n", token))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	token := os.Getenv("API_TOKEN")
	if token == "" {

		token, err = GenerateToken(16)
		if err != nil {
			log.Fatalf("Ошибка генерации токена: %v", err)
		}

		err = AddTokenToEnv(token)
		if err != nil {
			log.Fatalf("Ошибка добавления токена в .env файл: %v", err)
		}

		log.Println("Токен успешно добавлен в .env файл.")
	} else {
		log.Println("Токен уже существует в .env файле.")
	}

	r := gin.Default()

	r.Use(middleware.DebugMiddleware(), middleware.CheckBearerToken())

	config.ConnectDatabase()

	err = config.DB.AutoMigrate(&models.Guest{})
	if err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	guestRepo := &repository.PostgresGuestRepository{DB: config.DB}
	guestService := &service.GuestService{Repo: guestRepo}
	guestController := &controllers.GuestController{Service: guestService}

	r.POST("/guests", func(c *gin.Context) {
		start := time.Now()

		guestController.CreateGuest(c)

		duration := time.Since(start).Milliseconds()
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		memoryUsage := memStats.Alloc / 1024

		c.Header("X-Debug-Time", fmt.Sprintf("%dms", duration))
		c.Header("X-Debug-Memory", fmt.Sprintf("%dkb", memoryUsage))

		c.Next()
	})

	r.GET("/guests/:id", func(c *gin.Context) {
		start := time.Now()

		guestController.GetGuestById(c)

		duration := time.Since(start).Milliseconds()
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		memoryUsage := memStats.Alloc / 1024

		c.Header("X-Debug-Time", fmt.Sprintf("%dms", duration))
		c.Header("X-Debug-Memory", fmt.Sprintf("%dkb", memoryUsage))

		c.Next()
	})

	r.GET("/guests", func(c *gin.Context) {
		start := time.Now()

		guestController.GetGuest(c)

		duration := time.Since(start).Milliseconds()
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		memoryUsage := memStats.Alloc / 1024

		c.Header("X-Debug-Time", fmt.Sprintf("%dms", duration))
		c.Header("X-Debug-Memory", fmt.Sprintf("%dkb", memoryUsage))

		c.Next()
	})

	r.PUT("/guests/:id", func(c *gin.Context) {
		start := time.Now()

		guestController.UpdateGuest(c)

		duration := time.Since(start).Milliseconds()
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		memoryUsage := memStats.Alloc / 1024

		c.Header("X-Debug-Time", fmt.Sprintf("%dms", duration))
		c.Header("X-Debug-Memory", fmt.Sprintf("%dkb", memoryUsage))

		c.Next()
	})

	r.DELETE("/guests/:id", func(c *gin.Context) {
		start := time.Now()

		guestController.DeleteGuest(c)

		duration := time.Since(start).Milliseconds()
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		memoryUsage := memStats.Alloc / 1024

		c.Header("X-Debug-Time", fmt.Sprintf("%dms", duration))
		c.Header("X-Debug-Memory", fmt.Sprintf("%dkb", memoryUsage))

		c.Next()
	})

	err = r.Run()
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
