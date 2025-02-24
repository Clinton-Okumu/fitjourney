package models

import(
  "fmt"
  "log"
  "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(db *gorm.DB) {
  DB = db

  err := DB.AutoMigrate(&User{}, &Exercise{}, &Workout{}, &Role{})
  if err != nil {
    log.Fatal("Error migrating database:", err)
  }
  fmt.Println("Database migration complete")

  seedRoles()

 }

 func seedRoles() {
    roles := []Role{{Name: "Admin"}, {Name: "User"}}
    for _, r := range roles {
        if err ;= DB.Where("name = ?", r.Name).FirstOrCreate(&r).Error; err != nil {
        log.Fatal("Error seeding roles:", err)
      }
    }
  }
  fmt.Println("Roles seeded")




