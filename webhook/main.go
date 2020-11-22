package webhook

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to open data.db")
	}
	_ = db.AutoMigrate(&Webhook{})
}

func New(w *Webhook) {
	db.Create(w)
}

func Read(id uint) (w Webhook, found bool) {
	result := db.First(&w, id)
	found = !errors.Is(result.Error, gorm.ErrRecordNotFound)
	return
}

func GetByURL(url string) (w Webhook, found bool) {
	result := db.Where("url = ?", url).First(&w)
	found = !errors.Is(result.Error, gorm.ErrRecordNotFound)
	return
}

func All() (webhooks []Webhook) {
	db.Find(&webhooks)
	return
}

func Update(newValue *Webhook) {
	var w Webhook
	db.First(&w, newValue.ID)
	db.Model(&w).Updates(newValue)
}

func Delete(id uint) {
	var w Webhook
	db.Delete(&w, id)
}

func Search(keyword string) (webhooks []Webhook) {
	db.Where("name like '%" + keyword + "%'").
		Or("description like '%" + keyword + "%'").
		Or("executor like '%" + keyword + "%'").
		Or("url like '%" + keyword + "%'").
		Find(&webhooks)
	return
}
