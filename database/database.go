package database

import (
    "log"

    "ChronoTrack/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("chrono_track.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // 自动迁移数据库模型
    DB.AutoMigrate(&models.User{}, &models.Task{})
}