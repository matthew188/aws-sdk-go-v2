# v1.30.1 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.30.0 (2023-03-17)

* **Feature**: This release adds resourceType enums for types released from October 2022 through February 2023.

# v1.29.6 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.29.5 (2023-02-22)

* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.29.4 (2023-02-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.29.3 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.29.2 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.29.1 (2023-01-23)

* No change notes available for this release.

# v1.29.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).

# v1.28.2 (2022-12-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.1 (2022-12-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.0 (2022-11-29)

* **Feature**: With this release, you can use AWS Config to evaluate your resources for compliance with Config rules before they are created or updated. Using Config rules in proactive mode enables you to test and build compliant resource templates or check resource configurations at the time they are provisioned.

# v1.27.5 (2022-11-22)

* No change notes available for this release.

# v1.27.4 (2022-11-16)

* No change notes available for this release.

# v1.27.3 (2022-11-10)

* No change notes available for this release.

# v1.27.2 (2022-10-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.1 (2022-10-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.0 (2022-10-19)

* **Feature**: This release adds resourceType enums for AppConfig, AppSync, DataSync, EC2, EKS, Glue, GuardDuty, SageMaker, ServiceDiscovery, SES, Route53 types.

# v1.26.1 (2022-09-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.26.0 (2022-09-14)

* **Feature**: Fixed a bug in the API client generation which caused some operation parameters to be incorrectly generated as value types instead of pointer types. The service API always required these affected parameters to be nilable. This fixes the SDK client to match the expectations of the the service API.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.4 (2022-09-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.3 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.2 (2022-08-30)

* No change notes available for this release.

# v1.25.1 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.0 (2022-08-24)

* **Feature**: AWS Config now supports ConformancePackTemplate documents in SSM Docs for the deployment and update of conformance packs.

# v1.24.3 (2022-08-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.24.2 (2022-08-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.24.1 (2022-08-08)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.24.0 (2022-08-04)

* **Feature**: Add resourceType enums for Athena, GlobalAccelerator, Detective and EC2 types

# v1.23.1 (2022-08-01)

* **Documentation**: Documentation update for PutConfigRule and PutOrganizationConfigRule
* **Dependency Update**: Updated to the latest SDK module versions

# v1.23.0 (2022-07-27)

* **Feature**: This release adds ListConformancePackComplianceScores API to support the new compliance score feature, which provides a percentage of the number of compliant rule-resource combinations in a conformance pack compared to the number of total possible rule-resource combinations in the conformance pack.

# v1.22.0 (2022-07-14)

* **Feature**: Update ResourceType enum with values for Route53Resolver, Batch, DMS, Workspaces, Stepfunctions, SageMaker, ElasticLoadBalancingV2, MSK types

# v1.21.5 (2022-07-05)

* **Documentation**: Updating documentation service limits
* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.4 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.3 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.2 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.1 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.0 (2022-04-06)

* **Feature**: Add resourceType enums for AWS::EMR::SecurityConfiguration and AWS::SageMaker::CodeRepository

# v1.20.1 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.0 (2022-03-24)

* **Feature**: Added new APIs GetCustomRulePolicy and GetOrganizationCustomRulePolicy, and updated existing APIs PutConfigRule, DescribeConfigRule, DescribeConfigRuleEvaluationStatus, PutOrganizationConfigRule, DescribeConfigRule to support a new feature for building AWS Config rules with AWS CloudFormation Guard
* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.1 (2022-03-23)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.0 (2022-03-14)

* **Feature**: Add resourceType enums for AWS::ECR::PublicRepository and AWS::EC2::LaunchTemplate

# v1.18.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/github.com/matthew188/aws-sdk-go-v2/aws/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.0 (2022-01-28)

* **Feature**: Updated to latest API model.

# v1.15.0 (2022-01-14)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.13.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.
* **Feature**: Updated to latest service endpoints

# v1.12.2 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented aws.EndpointResolverWithOptions from being used by the service client. ([#1514](https://github.com/matthew188/aws-sdk-go-v2/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.12.1 (2021-11-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.12.0 (2021-11-12)

* **Feature**: Service clients now support custom endpoints that have an initial URI path defined.

# v1.11.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.0 (2021-10-21)

* **Feature**: API client updated
* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.2 (2021-10-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.1 (2021-09-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.0 (2021-09-02)

* **Feature**: API client updated

# v1.8.0 (2021-08-27)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2021-08-19)

* **Feature**: API client updated
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.2 (2021-08-04)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.1 (2021-07-15)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.0 (2021-06-25)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.1 (2021-05-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2021-05-14)

* **Feature**: Constant has been added to modules to enable runtime version inspection for reporting.
* **Feature**: Updated to latest service API model.
* **Dependency Update**: Updated to the latest SDK module versions

