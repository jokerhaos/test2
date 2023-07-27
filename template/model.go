package {{.Package}}

import (
	"gorm.io/gorm"
)

type {{.Model}}Entity struct{
	DB *gorm.DB
}

func New{{.Model}}Model(db *gorm.DB) *{{.Model}}Entity {
	return &{{.Model}}Entity{
		DB:db,
	}
}

// 根据id查询单条
func (entity *{{.Model}}Entity) GetTacticsGroupById(id int32) (info *{{.Model}},err error) {
	err = global.DB.Model({{.Model}}{}).
		Limit(1).
		Find(&info, id).Error
	return info, err
}

