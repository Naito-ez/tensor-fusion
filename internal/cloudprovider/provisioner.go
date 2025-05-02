package cloudprovider

import (
	"fmt"

	tfv1 "github.com/preciouscust/tensor-fusion/api/v1"
	"github.com/preciouscust/tensor-fusion/internal/cloudprovider/types"

	alibaba "github.com/preciouscust/tensor-fusion/internal/cloudprovider/alibaba"
	aws "github.com/preciouscust/tensor-fusion/internal/cloudprovider/aws"
	mock "github.com/preciouscust/tensor-fusion/internal/cloudprovider/mock"
)

func GetProvider(config tfv1.ComputingVendorConfig) (*types.GPUNodeProvider, error) {
	var err error
	var provider types.GPUNodeProvider
	switch config.Type {
	case "aws":
		provider, err = aws.NewAWSGPUNodeProvider(config)
	case "alibaba":
		provider, err = alibaba.NewAlibabaGPUNodeProvider(config)
	case "mock":
		provider, err = mock.NewMockGPUNodeProvider(config)
	default:
		return nil, fmt.Errorf("unsupported cloud provider: %s", config.Type)
	}
	return &provider, err
}
