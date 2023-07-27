// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Password   string    `gorm:"column:password;not null" json:"password"`
	Age        int32     `gorm:"column:age;not null" json:"age"`
	Sex        int32     `gorm:"column:sex;not null" json:"sex"`
	CreateTime time.Time `gorm:"column:create_time;not null" json:"create_time"`
}

type User2 struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Password   string    `gorm:"column:password;not null" json:"password"`
	Age        int32     `gorm:"column:age;not null" json:"age"`
	Sex        int32     `gorm:"column:sex;not null" json:"sex"`
	CreateTime time.Time `gorm:"column:create_time;not null" json:"create_time"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
