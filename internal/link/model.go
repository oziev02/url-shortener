package link

import (
	"math/rand"

	"github.com/oziev02/url-shortener/internal/stat"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraints:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash() // Задаем хэш
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RandStringRunes(6) // Чтобы перегенерить хэш
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn((len(letterRunes)))]
	}
	return string(b)
}
