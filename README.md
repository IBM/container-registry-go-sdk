[![Build Status](https://travis-ci.com/IBM/container-registry-go-sdk.svg?branch=main)](https://travis-ci.com/IBM/container-registry-go-sdk)
[![Release](https://img.shields.io/github/v/release/IBM/container-registry-go-sdk)](https://github.com/IBM/container-registry-go-sdk/releases/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IBM/container-registry-go-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud Container Registry Go SDK 1.1.6
Go client library to interact with the [IBM Cloud Container Registry API](https://cloud.ibm.com/apidocs/container-registry), and [IBM Cloud Container Registry Vulnerability Advisor API](https://cloud.ibm.com/apidocs/container-registry/va-v4)

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
    + [`go get` command](#go-get-command)
    + [Go modules](#go-modules)
    + [`dep` dependency manager](#dep-dependency-manager)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Container Registry Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[Container Registry image management API](https://cloud.ibm.com/apidocs/container-registry) | containerregistryv1
[Container Registry Vulnerability Advisor API](https://cloud.ibm.com/apidocs/container-registry/va-v4) | vulnerabilityadvisorv4

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.21 or above.

## Installation
The current version of this SDK: 1.1.6

There are a few different ways to download and install the Container Registry Go SDK project for use by your
Go application:

#### `go get` command  
Use this command to download and install the SDK to allow your Go application to
use it:

```
go get github.com/IBM/container-registry-go-sdk
```

#### Go modules  
If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
  "github.com/IBM/container-registry-go-sdk/containerregistryv1"
  "github.com/IBM/container-registry-go-sdk/vulnerabilityadvisorv4"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

#### `dep` dependency manager  
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```
[[constraint]]
  name = "github.com/IBM/container-registry-go-sdk"
  version = "1.1.6"

```

then run `dep ensure`.

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md)

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/container-registry-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
