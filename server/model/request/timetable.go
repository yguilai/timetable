package request

type TimeTableStruct struct {
	Yearly   string `json:"yearly" form:"yearly" validate:"required"`
	Weekly   int    `json:"weekly" form:"weekly"`
	StuNum   string `json:"stuNum"`
	StuId    uint   `json:"stuId"`
	Password string `json:"password"`
}
