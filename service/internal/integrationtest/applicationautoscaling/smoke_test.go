//go:build integration
// +build integration

package applicationautoscaling

import (
	"context"
	"testing"
	"time"

	"github.com/matthew188/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/matthew188/aws-sdk-go-v2/service/applicationautoscaling/types"

	"github.com/matthew188/aws-sdk-go-v2/service/internal/integrationtest"
)

func TestInteg_00_DescribeScalableTargets(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := applicationautoscaling.NewFromConfig(cfg)
	params := &applicationautoscaling.DescribeScalableTargetsInput{
		ServiceNamespace: types.ServiceNamespaceEc2,
	}
	_, err = client.DescribeScalableTargets(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
