package models

type UserSession struct {
	AccessUUID  string
	Authorized  bool
	Email       string
	Exp         float64
	UserId      int64
	Permissions []string
	RoleId      int64
	TeamId      string
}
