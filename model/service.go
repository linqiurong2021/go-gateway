package model

import "github.com/linqiurong2021/go-gateway/mysql"

// Service 服务存在表
type Service struct {
	ID    uint   `gorm:"id;primaryKey"`
	Host  string `gorm:"host"`
	Port  string `gorm:"port"`
	URL   string `gorm:"url"`
	Name  string `gorm:"name"`
	Alive bool   `gorm:"alive"`
}

// GetAllService 获取所有服务
func (s *Service) GetAllService() ([]*Service, error) {
	var services []*Service
	if err := mysql.DB.Where("alive = ?", 1).Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}
