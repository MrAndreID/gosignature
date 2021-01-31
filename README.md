# MrAndreID / Go Signature

[![Go Reference](https://pkg.go.dev/badge/github.com/MrAndreID/gosignature.svg)](https://pkg.go.dev/github.com/MrAndreID/gosignature) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The `MrAndreID/GoSignature` package is a collection of functions in the go language for Signature with rsa-sha256.

---

## Table of Contents

* [Install](#install)
* [Usage](#usage)
* [Full Example](#full-example)
* [Versioning](#versioning)
* [Authors](#authors)
* [License](#license)
* [Official Documentation for Go Language](#official-documentation-for-go-language)
* [More](#more)

---

## Install

To use The `MrAndreID/GoSignature` package, you must follow the steps below:

```sh
go get -u github.com/MrAndreID/gosignature
```

## Usage

### Generate a Signature

```go
signature, err := gosignature.Generate([]byte(`-----BEGIN RSA PRIVATE KEY-----
MIIBOQIBAAJAfBNTEHscBt+jNL+Ry1aqPoICoM8fuSpCkFN3KZ37zVEWu6Poav56
yp1i5D4ngVQ7vW2w2Go7M8TCfNS4aLna3wIDAQABAkAklCnj7Pd5S0s5TNT1poow
PXH66LVIiJ3xILo7ybincbp2jMc0Ah7TmuoTm9C0Uz8esoKWK2YGkLOykHSPxw25
AiEA2H3uURF/iqz4V5cJscbDMRQHc/qOpr22UF8utCkEMm0CIQCSt+ICBt2Y0Pir
BXjFmM6AcLHqwvy93gtypGx78dvS+wIgYCM0JH3/xGZhdgwVewPIBFBfquo2VOdk
Qbay98BLI9UCIHPzEJELffs0QyFdTKnUbnZBGcpvWLBwl9l9KiL17AUbAiEApiKl
aJkS7fmAd5XGd43BXgWJevG87yUzXxKPbsMp3mk=
-----END RSA PRIVATE KEY-----`), []byte(`Andrea Adam`))
```

Output:

```sh
ZMVHtYqKTgwsBItALtoE71ApBVJSQ1vxtW9a9oiugGZkhpIUBHtKTpL5e29CAeZnwlHTurUxpk1aH2RHx9sx3Q==
```

### Signature Verification

```go
verify, err := gosignature.Verify(signature, []byte(`-----BEGIN PUBLIC KEY-----
MFswDQYJKoZIhvcNAQEBBQADSgAwRwJAfBNTEHscBt+jNL+Ry1aqPoICoM8fuSpC
kFN3KZ37zVEWu6Poav56yp1i5D4ngVQ7vW2w2Go7M8TCfNS4aLna3wIDAQAB
-----END PUBLIC KEY-----`), []byte(`Andrea Adam`))

if verify == false {
    fmt.Println("Signature verification failed because : ", err)
} else {
    fmt.Println("Signature verification was successful")
}
```

Output:

```sh
Signature verification was successful
```

## Full Example

Full Example can be found on the [Go Playground website](https://play.golang.com/p/LfI8dDO-Vzt).

## Versioning

I use [SemVer](https://semver.org/) for versioning. For the versions available, see the tags on this repository. 

## Authors

**Andrea Adam** - [MrAndreID](https://github.com/MrAndreID/)

## License

MIT licensed. See the LICENSE file for details.

## Official Documentation for Go Language

Documentation for Go Language can be found on the [Go Language website](https://golang.org/doc/).

## More

Documentation can be found [on https://go.dev/](https://pkg.go.dev/github.com/MrAndreID/gosignature).
