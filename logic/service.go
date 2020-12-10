package logic

import (
	"github.com/linqiurong2021/go-gateway/model"
	"github.com/linqiurong2021/go-gateway/service"
)

var serve *service.Service

// Service 服务
type Service struct {
}

func init() {
	serve = new(service.Service)
}

// GetAllService 获取所有服务
func (s *Service) GetAllService() ([]*model.Service, error) {
	return serve.GetAllService()
}
