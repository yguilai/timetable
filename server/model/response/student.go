package response

import "ccsu/model"

type LoginResponse struct {
	Stu       model.Student `json:"stu"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}

type StuInfoStruct struct {
	ID     uint   `json:"id"`
	StuNum string `json:"stuNum"`
}
