package logic

import (
	"github.com/linqiurong2021/go-gateway/model"
	"github.com/linqiurong2021/go-gateway/service"
)

// Service 服务
type Service struct {
}

// GetAllService 获取所有服务
func (s *Service) GetAllService() ([]*model.Service, error) {
	var service = new(service.Service)
	return service.GetAllService()
}
