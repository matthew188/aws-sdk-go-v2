# v1.25.1 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.25.0 (2023-03-17)

* **Feature**: S3 setting to create AWS Glue Data Catalog. Oracle setting to control conversion of timestamp column. Support for Kafka SASL Plain authentication. Setting to map boolean from PostgreSQL to Redshift. SQL Server settings to force lob lookup on inline LOBs and to control access of database logs.

# v1.24.1 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.24.0 (2023-03-07)

* **Feature**: This release adds DMS Fleet Advisor Target Recommendation APIs and exposes functionality for DMS Fleet Advisor. It adds functionality to start Target Recommendation calculation.

# v1.23.5 (2023-02-22)

* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.23.4 (2023-02-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.23.3 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.23.2 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.23.1 (2023-01-23)

* No change notes available for this release.

# v1.23.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).

# v1.22.3 (2022-12-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.2 (2022-12-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.1 (2022-11-22)

* No change notes available for this release.

# v1.22.0 (2022-11-17)

* **Feature**: Adds support for Internet Protocol Version 6 (IPv6) on DMS Replication Instances

# v1.21.16 (2022-11-16)

* No change notes available for this release.

# v1.21.15 (2022-11-10)

* No change notes available for this release.

# v1.21.14 (2022-10-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.13 (2022-10-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.12 (2022-09-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.11 (2022-09-14)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.10 (2022-09-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.9 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.8 (2022-08-30)

* No change notes available for this release.

# v1.21.7 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.6 (2022-08-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.5 (2022-08-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.4 (2022-08-08)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.3 (2022-08-04)

* **Documentation**: Documentation updates for Database Migration Service (DMS).

# v1.21.2 (2022-08-01)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.1 (2022-07-21)

* **Documentation**: Documentation updates for Database Migration Service (DMS).

# v1.21.0 (2022-07-07)

* **Feature**: New api to migrate event subscriptions to event bridge rules

# v1.20.1 (2022-07-05)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.0 (2022-07-01)

* **Feature**: Added new features for AWS DMS version 3.4.7 that includes new endpoint settings for S3, OpenSearch, Postgres, SQLServer and Oracle.

# v1.19.1 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.0 (2022-06-08)

* **Feature**: This release adds DMS Fleet Advisor APIs and exposes functionality for DMS Fleet Advisor. It adds functionality to create and modify fleet advisor instances, and to collect and analyze information about the local data infrastructure.

# v1.18.6 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.5 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.4 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.3 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.2 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.1 (2022-03-23)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/github.com/matthew188/aws-sdk-go-v2/aws/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.0 (2022-01-14)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Documentation**: Updated API models
* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.
* **Feature**: Updated to latest service endpoints

# v1.13.1 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented aws.EndpointResolverWithOptions from being used by the service client. ([#1514](https://github.com/matthew188/aws-sdk-go-v2/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.13.0 (2021-11-30)

* **Feature**: API client updated

# v1.12.0 (2021-11-19)

* **Feature**: API client updated
* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.0 (2021-11-12)

* **Feature**: Service clients now support custom endpoints that have an initial URI path defined.
* **Feature**: Waiters now have a `WaitForOutput` method, which can be used to retrieve the output of the successful wait operation. Thank you to [Andrew Haines](https://github.com/haines) for contributing this feature.

# v1.10.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Feature**: Updated service to latest API model.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.0 (2021-10-21)

* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.1 (2021-10-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2021-09-24)

* **Feature**: API client updated

# v1.7.1 (2021-09-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2021-08-27)

* **Feature**: Updated API model to latest revision.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.2 (2021-08-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.1 (2021-08-04)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.0 (2021-07-15)

* **Feature**: Updated service model to latest version.
* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2021-06-25)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.1 (2021-05-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2021-05-14)

* **Feature**: Constant has been added to modules to enable runtime version inspection for reporting.
* **Dependency Update**: Updated to the latest SDK module versions

