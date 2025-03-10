//go:build integration
// +build integration

package docdb

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/matthew188/aws-sdk-go-v2/aws"
	"github.com/matthew188/aws-sdk-go-v2/service/docdb"
	"github.com/aws/smithy-go"

	"github.com/matthew188/aws-sdk-go-v2/service/internal/integrationtest"
)

func TestInteg_00_DescribeDBEngineVersions(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := docdb.NewFromConfig(cfg)
	params := &docdb.DescribeDBEngineVersionsInput{}
	_, err = client.DescribeDBEngineVersions(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestInteg_01_DescribeDBInstances(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := docdb.NewFromConfig(cfg)
	params := &docdb.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String("fake-id"),
	}
	_, err = client.DescribeDBInstances(ctx, params)
	if err == nil {
		t.Fatalf("expect request to fail")
	}

	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expect error to be API error, was not, %v", err)
	}
	if len(apiErr.ErrorCode()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if len(apiErr.ErrorMessage()) == 0 {
		t.Errorf("expect non-empty error message")
	}
}
