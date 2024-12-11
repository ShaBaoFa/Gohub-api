package user

import (
	"gohub/pkg/hash"
	"gorm.io/gorm"
)

// BeforeSave GORM 模型的钩子
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}

	return
}
