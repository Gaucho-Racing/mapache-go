package mapache

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name" gorm:"index"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index,unique"`
	Password  string    `json:"password"`
	Subteam   string    `json:"subteam"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;precision:6"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;precision:6"`
	Roles     []string  `json:"roles" gorm:"-"`
}

func (User) TableName() string {
	return "user"
}

func (u User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

type UserRole struct {
	UserID    string    `json:"user_id" gorm:"primaryKey"`
	Role      string    `json:"role" gorm:"primaryKey"`
	CreatedAt time.Time `json:"time" gorm:"autoCreateTime;precision:6"`
}

func (UserRole) TableName() string {
	return "user_role"
}
