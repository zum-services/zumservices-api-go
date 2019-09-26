<p align="center"><img src="https://raw.githubusercontent.com/zumcoin/zum-assets/master/ZumCoin/zumcoin_logo_design/3d_green_lite_bg/ZumLogo_800x200px_lite_bg.png" width="400"></p>

[![Go Report Card](https://goreportcard.com/badge/github.com/zum-services/zumservices-api-go)](https://goreportcard.com/report/github.com/zum-services/zumservices-api-go)
[![GoDoc](https://godoc.org/github.com/zum-services/zumservices-api-go?status.svg)](https://godoc.org/github.com/zum-services/zumservices-api-go)

# ZUM Service Golang API Interface

This wrapper allows you to easily interact with the [ZUM Services](https://zum.services) 1.0.1 API to quickly develop applications that interact with the [ZumCoin](https://zumcoin.org) Network.


# Table of Contents

1. [Installation](#installation)
2. [Intialization](#intialization)
3. [Documentation](#documentation)
   1. [Methods](#methods)


# Installation

```bash
go get github.com/zum-services/zumservices-api-go
```


# Intialization

```go
import (
    "fmt"

    "github.com/zum-services/zumservices-api-go"
)

ZS := ZUMServices.ZSwrapper {
    token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCIsImFwcElkIjo0LCJ1c2VySWQiOjYsInBlcm1pc3Npb25zIjpbImFkZHJlc3M6bmV3Il0sImlhdCI6MTUzNjU4NTM2NywiZXhwIjoxNTM5MTc3MzY3LCJhdWQiOiJ0dXJ0bGV3YWxsZXQuaW8iLCJpc3MiOiJUUlRMIFNlcnZpY2VzIiwianRpIjoiMzMifQ.AEHXmvTo8RfNuZ15Y3IGPRhZPaJxFSmOZvVv2YGN9L4We7bXslIPxhMv_n_5cNW8sIgE2Fr-46OTb5H5AFgpjA",
    timeout: 2000
}
```



# Documentation

API documentation is available at https://zum.services/documentation


## Methods

### createAddress()
Create a new ZUM addresses

```go
response, err := ZS.createAddress()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```

### getAddress(address)
Get address details by address
```go
response, err := ZS.getAddress("Zum1yfSrdpfiSNG5CtYmckgpGe1FiAc9gLCEZxKq29puNCX92DUkFYFfEGKugPS6EhWaJXmhAzhePGs3jXvNgK4NbWXG4yaGBHC")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```

### deleteAddress(address)
Delete a selected ZUM address

```go
response, err := ZS.deleteAddress("Zum1yfSrdpfiSNG5CtYmckgpGe1FiAc9gLCEZxKq29puNCX92DUkFYFfEGKugPS6EhWaJXmhAzhePGs3jXvNgK4NbWXG4yaGBHC")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getAddresses()
View all addresses.

```go
response, err := ZS.getAddresses()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### scanAddress(address, blockIndex)
Scan an address for transactions between a 100 block range starting from the specified blockIndex.

```go
response, err := ZS.scanAddress("Zum1yfSrdpfiSNG5CtYmckgpGe1FiAc9gLCEZxKq29puNCX92DUkFYFfEGKugPS6EhWaJXmhAzhePGs3jXvNgK4NbWXG4yaGBHC", 899093)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getAddressKeys(address)
Get the public and secret spend key of an address.


```go
response, err := ZS.getAddressKeys("Zum1yfSrdpfiSNG5CtYmckgpGe1FiAc9gLCEZxKq29puNCX92DUkFYFfEGKugPS6EhWaJXmhAzhePGs3jXvNgK4NbWXG4yaGBHC")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### integrateAddress(address, paymentId)
Create an integrated address with an address and payment ID.

```go
response, err := ZS.getAddressKeys("Zum1yfSrdpfiSNG5CtYmckgpGe1FiAc9gLCEZxKq29puNCX92DUkFYFfEGKugPS6EhWaJXmhAzhePGs3jXvNgK4NbWXG4yaGBHC", "7d89a2d16365a1198c46db5bbe1af03d2b503a06404f39496d1d94a0a46f8804")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getIntegratedAddresses(address)
Get all integrated addresses by address.

```go
response, err := ZS.getIntegratedAddresses("Zum1yfSrdpfiSNG5CtYmckgpGe1FiAc9gLCEZxKq29puNCX92DUkFYFfEGKugPS6EhWaJXmhAzhePGs3jXvNgK4NbWXG4yaGBHC", "7d89a2d16365a1198c46db5bbe1af03d2b503a06404f39496d1d94a0a46f8804")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getFee(amount)
Calculate the ZUM Services fee for an amount specified in ZUM with two decimal points.

```go
response, err := ZS.getFee(1092.19)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### createTransfer()
Send a ZUM transaction with an address with the amount specified two decimal points.

```go
response, err := ZS.createTransfer(
    "Zum1yfSrdpfiSNG5CtYmckgpGe1FiAc9gLCEZxKq29puNCX92DUkFYFfEGKugPS6EhWaJXmhAzhePGs3jXvNgK4NbWXG4yaGBHC", "Zum1yhbRwHsXj19c1hZgFzgxVcWDywsJcDKURDud83MqMNKoDTvKEDf6k7BoHnfCiPbj4kY2arEmQTwiVmhoELPv3UKhjYjCMcm",
    1000.00,
    1.00,
    "",
    "",
)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getTransfer()
Get a transaction details specified by transaction hash.

```go
response, err := ZS.getTransfer("EohMUzR1DELyeQM9RVVwpmn5Y1DP0lh1b1ZpLQrfXQsgtvGHnDdJSG31nX2yESYZ")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getWallet()
Get wallet container info and health check.

```go
response, err := ZS.getWallet()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


### getStatus()
Get the current status of the ZUM Services infrastructure.

```go
response, err := ZS.getStatus()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response)
}
```


# License

```
Copyright (c) 2019 ZumCoin Development Team Developers

Please see the included LICENSE file for more information.
```
