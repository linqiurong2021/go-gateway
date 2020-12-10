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

// Create 创建
func (s *Service) Create(in *Service) {

}

// Update 创建
func (s *Service) Update(in *Service) {

}

// Delete 创建
func (s *Service) Delete(in *Service) {

}

// GetByID 通过ID获取
func (s *Service) GetByID(in *Service) {

}

// GetByURL 通过URL获取
func (s *Service) GetByURL(in *Service) {

}

// GetAllService 获取所有服务
func (s *Service) GetAllService() ([]*Service, error) {
	var services []*Service
	if err := mysql.DB.Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}
