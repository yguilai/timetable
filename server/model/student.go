package model

type Student struct {
	Model
	StuNum    string    `json:"stuNum" gorm:"type:char(12);uniqueIndex"`
	Password  string    `json:"-"`
	Timetable Timetable `json:"timetable" gorm:"foreignKey:StudentId" `
}

// 课表
type Timetable struct {
	Model
	StudentId uint      `json:"-"`
	Yearly    string    `json:"yearly" `
	Weekly    int       `json:"weekly"`
	Courses   []Course `json:"courses" gorm:"foreignKey:TimetableId"`
	Remarks   string    `json:"remarks"`
}

// 单个课程
type Course struct {
	Model
	TimetableId uint   `json:"-"`
	Name        string `json:"name"`
	Class       string `json:"class"`
	Time        int    `json:"time"`
	Week        int    `json:"week"`
	Weeks       string `json:"weeks"`
	Teacher     string `json:"teacher"`
	Location    string `json:"location"`
}
