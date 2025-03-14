# v1.28.7 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.6 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.5 (2023-02-22)

* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.28.4 (2023-02-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.3 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.28.2 (2023-02-09)

* **Documentation**: Removed Windows Server 2016 BYOL and made changes based on IAM campaign.

# v1.28.1 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).

# v1.27.2 (2022-12-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.1 (2022-12-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.0 (2022-11-17)

* **Feature**: The release introduces CreateStandbyWorkspaces, an API that allows you to create standby WorkSpaces associated with a primary WorkSpace in another Region. DescribeWorkspaces now includes related WorkSpaces properties. DescribeWorkspaceBundles and CreateWorkspaceBundle now return more bundle details.

# v1.26.0 (2022-11-15)

* **Feature**: This release introduces ModifyCertificateBasedAuthProperties, a new API that allows control of certificate-based auth properties associated with a WorkSpaces directory. The DescribeWorkspaceDirectories API will now additionally return certificate-based auth properties in its responses.

# v1.25.0 (2022-11-07)

* **Feature**: This release adds protocols attribute to workspaces properties data type. This enables customers to migrate workspaces from PC over IP (PCoIP) to WorkSpaces Streaming Protocol (WSP) using create and modify workspaces public APIs.

# v1.24.0 (2022-10-25)

* **Feature**: This release adds new enums for supporting Workspaces Core features, including creating Manual running mode workspaces, importing regular Workspaces Core images and importing g4dn Workspaces Core images.

# v1.23.2 (2022-10-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.23.1 (2022-10-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.23.0 (2022-09-29)

* **Feature**: This release includes diagnostic log uploading feature. If it is enabled, the log files of WorkSpaces Windows client will be sent to Amazon WorkSpaces automatically for troubleshooting. You can use modifyClientProperty api to enable/disable this feature.

# v1.22.9 (2022-09-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.8 (2022-09-15)

* No change notes available for this release.

# v1.22.7 (2022-09-14)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.6 (2022-09-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.5 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.4 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.3 (2022-08-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.2 (2022-08-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.1 (2022-08-08)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.0 (2022-08-01)

* **Feature**: This release introduces ModifySamlProperties, a new API that allows control of SAML properties associated with a WorkSpaces directory. The DescribeWorkspaceDirectories API will now additionally return SAML properties in its responses.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.0 (2022-07-27)

* **Feature**: Added CreateWorkspaceImage API to create a new WorkSpace image from an existing WorkSpace.

# v1.20.0 (2022-07-19)

* **Feature**: Increased the character limit of the login message from 850 to 2000 characters.

# v1.19.2 (2022-07-05)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.1 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.0 (2022-06-15)

* **Feature**: Added new field "reason" to OperationNotSupportedException. Receiving this exception in the DeregisterWorkspaceDirectory API will now return a reason giving more context on the failure.

# v1.18.4 (2022-06-10)

* No change notes available for this release.

# v1.18.3 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.2 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.1 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.0 (2022-04-11)

* **Feature**: Added API support that allows customers to create GPU-enabled WorkSpaces using EC2 G4dn instances.

# v1.17.0 (2022-03-31)

* **Feature**: Added APIs that allow you to customize the logo, login message, and help links in the WorkSpaces client login page. To learn more, visit https://docs.aws.amazon.com/workspaces/latest/adminguide/customize-branding.html

# v1.16.3 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.2 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.1 (2022-03-23)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/github.com/matthew188/aws-sdk-go-v2/aws/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2022-01-14)

* **Feature**: Updated API models
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.13.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.12.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.

# v1.11.1 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented aws.EndpointResolverWithOptions from being used by the service client. ([#1514](https://github.com/matthew188/aws-sdk-go-v2/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.0 (2021-11-30)

* **Feature**: API client updated

# v1.10.1 (2021-11-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.0 (2021-11-12)

* **Feature**: Service clients now support custom endpoints that have an initial URI path defined.

# v1.9.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2021-10-21)

* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.1 (2021-10-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2021-09-30)

* **Feature**: API client updated

# v1.6.1 (2021-09-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.0 (2021-08-27)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.3 (2021-08-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.2 (2021-08-04)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.1 (2021-07-15)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2021-06-25)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2021-05-25)

* **Feature**: API client updated

# v1.3.1 (2021-05-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2021-05-14)

* **Feature**: Constant has been added to modules to enable runtime version inspection for reporting.
* **Dependency Update**: Updated to the latest SDK module versions

