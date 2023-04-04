//go:build integration
// +build integration

package sms

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/matthew188/aws-sdk-go-v2/aws"
	"github.com/matthew188/aws-sdk-go-v2/service/sms"
	"github.com/aws/smithy-go"

	"github.com/matthew188/aws-sdk-go-v2/service/internal/integrationtest"
)

func TestInteg_00_GetConnectors(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := sms.NewFromConfig(cfg)
	params := &sms.GetConnectorsInput{}
	_, err = client.GetConnectors(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestInteg_01_DeleteReplicationJob(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := sms.NewFromConfig(cfg)
	params := &sms.DeleteReplicationJobInput{
		ReplicationJobId: aws.String("invalidId"),
	}
	_, err = client.DeleteReplicationJob(ctx, params)
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
