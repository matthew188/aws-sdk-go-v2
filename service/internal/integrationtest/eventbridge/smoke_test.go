//go:build integration
// +build integration

package eventbridge

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/matthew188/aws-sdk-go-v2/aws"
	"github.com/matthew188/aws-sdk-go-v2/service/eventbridge"

	"github.com/matthew188/aws-sdk-go-v2/service/internal/integrationtest"
	"github.com/aws/smithy-go"
)

func TestInteg_00_ListRules(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := eventbridge.NewFromConfig(cfg)
	params := &eventbridge.ListRulesInput{}
	_, err = client.ListRules(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestInteg_01_DescribeRule(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := eventbridge.NewFromConfig(cfg)
	params := &eventbridge.DescribeRuleInput{
		Name: aws.String("fake-rule"),
	}
	_, err = client.DescribeRule(ctx, params)
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
