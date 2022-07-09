package model

import "fmt"

type Message struct {
	ID        uint   `gorm:"primary_key" json:"id" `
	FromUid   int    `gorm:"index" json:"from_uid" form:"from_uid"`
	ToUid     int    `gorm:"index" json:"to_uid" form:"to_uid"`
	Content   string `json:"content" form:"content"`
	ImageUrl  string `json:"image_url" form:"image_url"`
	CreatedOn int    `gorm:"index" binding:"-" json:"created_on,omitempty"`
}

func SaveMessage(msg *Message) error {
	return db.Create(msg).Error
}

func GetMessages(fromUid, toUid, pageNum, pageSize int) ([]*Message, error) {
	var messages []*Message
	where := fmt.Sprintf("(`from_uid` = %d and `to_uid` = %d) or (`from_uid` = %d and `to_uid` = %d)", fromUid, toUid, toUid, fromUid)
	if err := db.Offset(pageNum).Limit(pageSize).Where(where).Order("id desc").Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}