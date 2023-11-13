package facades

import "gorm.io/gorm"

func Database() *gorm.DB {
	return App().GetDatabase()
}
