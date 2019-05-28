package models

type Note struct {
	Model
	NKey    string `gorm:"unique:not null"`
	UserID  int
	User    User
	Title   string `gorm:"type:varchar(200)"`
	Summary string `gorm:"type:varchar(800)"`
	Content string `gorm:"type:text"`
	Visit   int    `gorm:"default:0"`
	Praise  int    `gorm:"default:0"`
}

func UpdatedNoteByKeyAndUserId(nkey, title, content string, userid int) (note Note, err error) {
	return note, db.Model(&note).Where("n_key = ? and user_id=?", nkey, userid).Updates(map[string]interface{}{"title": title, "content": content}).Take(&note).Error
}

func SaveNote(note *Note) error {
	return db.Save(note).Error
}
