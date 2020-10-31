package service

import (
	"ccsu/global"
	"ccsu/model"
	"errors"
)

func Login(s *model.Student) (err error, stuInter *model.Student) {
	var stu model.Student
	err = global.DB.
		Where(&model.Student{StuNum: s.StuNum, Password: s.Password}).First(&stu).Error
	return err, &stu
}

func IsExist(stuNum string) (error, bool) {
	var stu model.Student
	var count int64
	err := global.DB.Where(&model.Student{StuNum: stuNum}).First(&stu).Count(&count).Error
	return err, count != 0
}

func SetPassword(s *model.Student, newPwd string) (err error, stuInter *model.Student)  {
	var stu model.Student
	err = global.DB.
		Where(&model.Student{StuNum: s.StuNum, Password: s.Password}).First(&stu).
		Update("password", newPwd).Error
	return err, &stu
}

func Register(s model.Student) (err error, stuInter model.Student) {
	_, exist := IsExist(s.StuNum)
	if exist {
		return errors.New("学号已存在"), stuInter
	} else {
		err = global.DB.Create(&s).Error
	}
	return err, s
}

func GetStudentById(sid uint) (error, *model.Student) {
	var stu model.Student
	err := global.DB.Where("id = ?", sid).First(&stu).Error
	return err, &stu
}