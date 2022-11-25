package models

type Users struct {
	Gorm
	FullName string `gorm:"not null" json:"fullname" form:"fullname" valid:"required~Your fullname is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(8)~Password has to have a minimum length of 8 characters"`
}

type Profile struct {
	Gorm
	Age         uint   `json:"age"`
	Photo       string `json:"photo"`
	Hoby        string `json:"hoby"`
	Description string `json:"description"`
	UserId      uint64 `json:"user_id"`
}
