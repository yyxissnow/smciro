package ioc

import (
	"google.golang.org/grpc"
)

var grpcIOC GrpcIOC

type GrpcIOC struct {
	Services map[string]GrpcService
}

type GrpcService interface {
	Register(*grpc.Server)
	Config()
	Name() string
}

func init() {
	grpcIOC = GrpcIOC{Services: map[string]GrpcService{}}
}

func RegistryGrpcService(service GrpcService) {
	_, ok := grpcIOC.Services[service.Name()]
	if !ok {
		grpcIOC.Services[service.Name()] = service
	}
}

func GetGrpcService(name string) GrpcService {
	service, ok := grpcIOC.Services[name]
	if !ok {
		return nil
	}
	return service
}

func GetAllGrpcServices() []string {
	var services []string
	for name := range grpcIOC.Services {
		services = append(services, name)
	}
	return services
}

func LoadGrpcIOC() error {
	for _, service := range ginIOC.Services {
		service.Config()
	}
	return nil
}
