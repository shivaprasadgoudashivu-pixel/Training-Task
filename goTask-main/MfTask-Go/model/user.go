package model

type USER struct {
	User_Id int     `json:"user_Id"`
	Name    string  `json:"name"`
	Address string  `json:"adddress"`
	Order   []ORDER `json:"orders" gorm:"foreignKey:UserId;references:User_Id`
}
