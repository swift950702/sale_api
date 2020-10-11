package users

import "github.com/jinzhu/gorm"

type User struct {
	ID       uint   `gorm:"id",json:"id,omitempty"`
	Pwd      string `gorm:"pwd",json:"pwd"`
	Username string `gorm:"username",json:"username"`
	Phone    string `gorm:"phone",json:"phone,omitempty"`
}

//gorm框架与数据库表的映射主要这个方法返回的值
func (u *User) TableName() string {
	return "users"
}

func (this *User) SearchByName(db *gorm.DB) (bool, int) {
	user := &User{}
	res := db.Where("username = ?", this.Username).First(user)
	if res.RecordNotFound() {
		return false, 3
	} else if user.Pwd == this.Pwd {
		return true, 1
	} else {
		return false, 2
	}
}
