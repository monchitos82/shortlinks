package main

import (
    "log"
    "fmt"
    "net/http"
    "net/url"
    "go-shortlinks/config"
    "go-shortlinks/database"
    "go-shortlinks/models"
    "go-shortlinks/handlers"
)

func main() {

    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    // Build DSN for PostgreSQL (prod)
    host := url.QueryEscape(cfg.Database.Prod.Host)
    user := url.QueryEscape(cfg.Database.Prod.User)
    password := url.QueryEscape(cfg.Database.Prod.Password)
    dbName := url.QueryEscape(cfg.Database.Prod.DBName)

    dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&timezone=%s",
        user,
        password,
        host,
        cfg.Database.Prod.Port,
        dbName,
        cfg.Database.Prod.SSLMode,
        cfg.Database.Prod.TimeZone)

    if err := database.InitDB(dsn); err != nil {
        log.Fatalf("failed to connect to PostgreSQL: %v", err)
    }

    database.DB.AutoMigrate(&models.URL{})

    http.HandleFunc("/", handlers.RedirectURL)
    http.HandleFunc("/healthcheck", handlers.Healthcheck)
    http.HandleFunc("/shorts", handlers.ShortenURL)
    http.HandleFunc("/updates", handlers.GetURLUpdates(cfg.Templates.Path))
    http.HandleFunc("/urls", handlers.GetAllURLs)
    
    address := fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port)
    log.Printf("Server is running on: %v", address)
    log.Fatal(http.ListenAndServe(address, nil))
}

