# v1.3.2 (2023-03-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.1 (2023-03-10)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2023-02-24)

* **Feature**: This release adds the ability to delete domains through the DeleteDomain API. For more information see https://docs.aws.amazon.com/cases/latest/APIReference/Welcome.html

# v1.2.6 (2023-02-22)

* **Bug Fix**: Prevent nil pointer dereference when retrieving error codes.

# v1.2.5 (2023-02-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.4 (2023-02-16)

* No change notes available for this release.

# v1.2.3 (2023-02-15)

* **Announcement**: When receiving an error response in restJson-based services, an incorrect error type may have been returned based on the content of the response. This has been fixed via PR #2012 tracked in issue #1910.
* **Bug Fix**: Correct error type parsing for restJson services.

# v1.2.2 (2023-02-03)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.1 (2023-01-26)

* No change notes available for this release.

# v1.2.0 (2023-01-05)

* **Feature**: Add `ErrorCodeOverride` field to all error structs (aws/smithy-go#401).

# v1.1.2 (2022-12-15)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.1 (2022-12-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.0 (2022-11-09)

* **Feature**: This release adds the ability to disable templates through the UpdateTemplate API. Disabling templates prevents customers from creating cases using the template. For more information see https://docs.aws.amazon.com/cases/latest/APIReference/Welcome.html

# v1.0.2 (2022-10-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.1 (2022-10-21)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.0 (2022-10-04)

* **Release**: New AWS service client module
* **Feature**: This release adds APIs for Amazon Connect Cases. Cases allows your agents to quickly track and manage customer issues that require multiple interactions, follow-up tasks, and teams in your contact center.  For more information, see https://docs.aws.amazon.com/cases/latest/APIReference/Welcome.html

