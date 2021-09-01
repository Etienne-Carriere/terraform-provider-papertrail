# terraform-provider-papertrail

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x and upper
-	[Go](https://golang.org/doc/install) 1.14 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/oogway/terraform-provider-papertrail`

```sh
$ mkdir -p $GOPATH/src/github.com/oogway; cd $GOPATH/src/github.com/oogway
$ git clone git@github.com:oogway/terraform-provider-papertrail
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/oogway/terraform-provider-papertrail
$ go get
$ go build
```

For Usage, have a look at docs in `website` directory.

Running Tests
-------------
```sh
$ cd $GOPATH/src/github.com/oogway/terraform-provider-papertrail/papertrail
$ PAPERTRAIL_TOKEN=<token> DESTINATION_PORT=<log_destination_port> go tests -v
```
