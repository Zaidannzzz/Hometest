package main

var photo = []*Photo{}

type Photo struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func Getid() []*Photo {
	return photo
}

func Selectid(id uint) *Photo {
	for _, each := range photo {
		if each.ID == id {
			return each
		}
	}

	return nil
}

func init() {
	photo = append(photo, &Photo{ID: 1, Title: "bourne", Message: "tes 1"})
	photo = append(photo, &Photo{ID: 2, Title: "ethan", Message: "tes 2"})
	photo = append(photo, &Photo{ID: 3, Title: "wick", Message: "tes 3"})
}
