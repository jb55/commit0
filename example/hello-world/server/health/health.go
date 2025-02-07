package health

import (
	"context"
	api "github.com/yourrepo/hello-world-idl/gen/go/health"
)

type HealthServer struct {

}

func NewHealthServer() *HealthServer {
	return &HealthServer{}
}

func (s *HealthServer) Check(ctx context.Context, req *api.HealthCheckRequest) (*api.HealthCheckResponse, error) {
	resp := &api.HealthCheckResponse{
		Status: api.HealthCheckResponse_SERVING,
	}
	return resp,nil
}

func (s *HealthServer) Watch(req *api.HealthCheckRequest, server api.Health_WatchServer) error {
	return nil
}
