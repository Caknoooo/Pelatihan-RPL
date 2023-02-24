package entities

import(
	"time"
)

type User struct{
	Id int64 `json:"id"`
	Nama string `json:"nama"`
	Aktifitas string `json:"aktifitas"`
	Mulai time.Time `json:"mulai"`
	Selesai time.Time `json:"selesai"`
	Done bool `json:"done"` // Perlu diperbaiki
}