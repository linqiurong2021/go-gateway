package service

import "github.com/linqiurong2021/go-gateway/model"

// Service 服务
type Service struct{}

var service *model.Service

func init() {
	service = new(model.Service)
}

// GetAllService 获取所有服务
func (s *Service) GetAllService() ([]*model.Service, error) {
	return service.GetAllService()
}
