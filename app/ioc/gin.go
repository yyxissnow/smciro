package ioc

import (
	"github.com/gin-gonic/gin"
)

var ginIOC GinIOC

type GinIOC struct {
	RouterGroup *gin.RouterGroup
	Services    map[string]GinService
}

type GinService interface {
	Register(group *gin.RouterGroup)
	Config()
	Name() string
}

func init() {
	ginIOC = GinIOC{
		RouterGroup: nil,
		Services:    map[string]GinService{},
	}
}

func RegistryGinService(service GinService) {
	_, ok := ginIOC.Services[service.Name()]
	if !ok {
		ginIOC.Services[service.Name()] = service
	}
}

func GetGinService(name string) GinService {
	service, ok := ginIOC.Services[name]
	if !ok {
		return nil
	}
	return service
}

func GetAllGinServices() []string {
	var services []string
	for name := range ginIOC.Services {
		services = append(services, name)
	}
	return services
}

func InitGinIOC(router *gin.RouterGroup) {
	ginIOC.RouterGroup = router
}

func LoadGinIOC() {
	for _, service := range ginIOC.Services {
		service.Config()
		service.Register(ginIOC.RouterGroup)
	}
}
