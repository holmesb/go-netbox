go-netbox 
=========

[![GoDoc](http://godoc.org/github.com/fbreckle/go-netbox?status.svg)](http://godoc.org/github.com/fbreckle/go-netbox) [![Build Status](https://github.com/fbreckle/go-netbox/workflows/main/badge.svg?branch=master)](https://github.com/fbreckle/go-netbox/actions) [![Report Card](https://goreportcard.com/badge/github.com/fbreckle/go-netbox)](https://goreportcard.com/report/github.com/fbreckle/go-netbox)

Package `netbox` provides an API 2.0 client for [netbox-community's NetBox](https://github.com/netbox-community/netbox)
IPAM and DCIM service.

This package assumes you are using NetBox 2.0, as the NetBox 1.0 API no longer exists.

Why this fork exists
====================

This fork exists solely to support [e-breuninger/terraform-provider-netbox](https://github.com/e-breuninger/terraform-provider-netbox). As such, some changes in this fork do only make sense in that context.


Versioning
==========

tbd. In the meanwhile, look at branches and tags.


Changes in this fork
====================

Change `models.ip_address.AssignedObject` type to prevent json marshalling errors since [this change](https://github.com/netbox-community/netbox/pull/4781)

Add `x-omitempty: false` to some attributes, allowing them to be set to their empty value. [issue](https://github.com/netbox-community/go-netbox/issues/107)

Change ConfigContext type for VMs and Devices [#2](https://github.com/fbreckle/go-netbox/pull/2)

Fix LocalConfigContext to support arbitrary JSON object [#4](https://github.com/fbreckle/go-netbox/pull/4)


Using the client
================

The `github.com/fbreckle/go-netbox/netbox` package has some convenience functions for creating clients with the most common
configurations you are likely to need while connecting to NetBox. `NewNetboxAt` allows you to specify a hostname
(including port, if you need it), and `NewNetboxWithAPIKey` allows you to specify both a hostname:port and API token.
```golang
import (
    "github.com/fbreckle/go-netbox/netbox"
)
...
    c := netbox.NewNetboxAt("your.netbox.host:8000")
    // OR
    c := netbox.NewNetboxWithAPIKey("your.netbox.host:8000", "your_netbox_token")
```

If you specify the API key, you do not need to pass an additional `authInfo` to operations that need authentication, and
can pass `nil`:
```golang
    c.Dcim.DcimDeviceTypesCreate(createRequest, nil)
```

If you connect to netbox via HTTPS you have to create an HTTPS configured transport:
```
package main

import (
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"

	log "github.com/sirupsen/logrus"
)

func main() {
	token := os.Getenv("NETBOX_TOKEN")
	if token == "" {
		log.Fatalf("Please provide netbox API token via env var NETBOX_TOKEN")
	}

	netboxHost := os.Getenv("NETBOX_HOST")
	if netboxHost == "" {
		log.Fatalf("Please provide netbox host via env var NETBOX_HOST")
	}

	transport := httptransport.New(netboxHost, client.DefaultBasePath, []string{"https"})
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+token)

	c := client.New(transport, nil)

	req := dcim.NewDcimSitesListParams()
	res, err := c.Dcim.DcimSitesList(req, nil)
	if err != nil {
		log.Fatalf("Cannot get sites list: %v", err)
	}
	log.Infof("res: %v", res)
}
```

Go Module support
================

Go 1.13+

`go get github.com/fbreckle/go-netbox`


More complex client configuration
=================================

The client is generated using [go-swagger](https://github.com/go-swagger/go-swagger). This means the generated client
makes use of [github.com/go-openapi/runtime/client](https://godoc.org/github.com/go-openapi/runtime/client). If you need
a more complex configuration, it is probably possible with a combination of this generated client and the runtime
options.

The [godocs for the go-openapi/runtime/client module](https://godoc.org/github.com/go-openapi/runtime/client) explain
the client options in detail, including different authentication and debugging options. One thing I want to flag because
it is so useful: setting the `DEBUG` environment variable will dump all requests to standard out.

Regenerating the client
=======================

To regenerate the client with a new or different swagger schema, first clean the existing client, then replace
swagger.json, run the json preprocessor (requires python3) and finally re-generate:
```
make clean
make preprocess
make generate
```
