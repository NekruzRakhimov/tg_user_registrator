package models

type User struct {
	ID             int    `json:"id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	EmailConfirmed bool   `json:"email_confirmed"`
	IsAdmin        bool   `json:"is_admin"`
	IsSuperAdmin   bool   `json:"is_super_admin"`
	OldPassword    string `json:"old_password"`
}

type AdminActivity struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
