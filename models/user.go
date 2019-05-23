package models

type User struct {
	Model
	Name string `gorm:"unique_index" json:"name"`
	Email string `gorm:"unique_index" json:"email"`
	Phone string `gorm:"unique_index" json:"phone"`
	Password string `json:"password"`
	Avatar string `json:"avatar"`
	Role int `json:"role"`//0代表管理员，1代表用户
}

func QueryUserByNameAndPassword(name, password string) (user User,err error) {
	return user,db.Where("name = ? and password = ?",name,password).Take(&user).Error
}

func QueryUserByName(name string) (user User, err error) {
	return user,db.Where("name = ?",name).Take(&user).Error
}

func QueryUserByEmail(email string) (user User, err error) {
	return user,db.Where("email = ?",email).Take(&user).Error
}

func QueryUserByPhone(phone string) (user User,err error) {
	return user,db.Where("phone = ?",phone).Take(&user).Error
}

func SaveUser(user *User) error {
	return db.Save(user).Error
}

func UpdatedInfoByName(name, phone, email string) (user User,err error) {
	return user,db.Model(&user).Where("name = ?",name).Updates(map[string]interface{}{"email":email,"phone":phone}).Take(&user).Error
}

func UpdatedPwdByName(name, password string) (user User,err error) {
	return user,db.Model(&user).Where("name = ?",name).Update("password",password).Take(&user).Error
}

func DeleteUserByName(name string) (error) {
	return db.Delete(&User{},"name = ?",name).Error
}

func QueryAllUser() (user []*User,err error){
	return user,db.Find(&user).Error
}