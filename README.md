# Golang client for VBR REST API

[![Go Report Card](https://goreportcard.com/badge/github.com/veeamhub/veeam-vbr-sdk-go)](https://goreportcard.com/report/github.com/veeamhub/veeam-vbr-sdk-go)
[![GoDoc](https://godoc.org/github.com/veeamhub/veeam-vbr-sdk-go?status.svg)](https://godoc.org/github.com/veeamhub/veeam-vbr-sdk-go)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/license/mit/)

veeam-vbr-sdk-go is the unofficial Veeam Backup & Replication SDK for the Go programming language.
Client generation is based on https://github.com/deepmap/oapi-codegen generator.
Due to specific of Golang, we had to make some changes in the original specification to make it work with the generator.
You can find changes description in the [Specification](#specification) section.


## Specification
`openapi_spec.yaml` does not contain the whole original `VBR REST API` specification. We made several changes described below. 

### Changes in spec
For each schema which has both `OneOf` and `Properties` following changes were made:
* Properties were moved to separate schema with the same name as original schema but with `Base` prefix, e.g. `RepositoryModel` -> `BaseRepositoryModel`.
* `AllOf` was added to the original schema. It contains reference to the `Base` schema and `OneOf` from the original schema.
* All references to the original schema were replaced by reference to the `Base` schema.

### How to regenerate spec
Additional tool named `oapifixer` included in the project. You can find it in the `tools/oapifixer` directory. 
It applies all the changes described in the [Changes in spec](#changes-in-spec) section.
Tool integrated into the `Makefile` and can be used by the following command:
```bash
make convert
```
It will generate specification file which can be used to generate client. 
By default it expects that the original specification is placed in the `spec` directory and has the name `swagger.json`.
'Fixed' specification will be placed in the `spec` directory and will have the name `openapi_spec.yaml`.
To change the default behavior you can use the following command:
```bash
make convert vbr_spec=<path_to_original_vbr_specification> golang_spec=<path_to_result>
```

> It is possible to convert specification from JSON to YAML and vice versa. Just change the extension of the output files.


## How to generate code

To generate code just run the following command:
```bash
make generate
```
It will remove the previous version and generate the new one. The result of generation will be placed into the `vbrclient` directory.
The default value for specification is `./spec/openapi_spec.yaml`. To change it use the following command:
```bash
make generate golang_spec=<path_to_specification>
```

## How to use

### Create client and authenticate
Create a new client via vbrapi.NewClient()

```golang
serverAddress := "127.0.0.1"
serverPort := 1234 // You can set 0 to use DefaultPort value(see in const.go)
skipSSLVerify := false

client, err := vbrapi.NewClient(serverAddress, serverPort, skipSSLVerify)
```

As a result you will get client which can make requests which allow to be made without authentication.
To be able to make authenticated requests you have to authenticate your client:

```golang
var vbrSecret = &xgo.SecretVbr{
    Password: secret.Password,
    User:     secret.User,
}
err := vbrClient.Authenticate(ctx, vbrSecret); 
```

If the authentication went without errors, your client is now authenticated. No need to create another one or provide any context.

### Making requests

There are 2 types of Client interfaces provided by the library
* Client which returns unparsed `*http.Response`
* Client which returns parsed model

We use one with parsed models, so, for each endpoint you will have 2 functions. For example:
```golang
GetServerInfo(ctx context.Context, params *GetServerInfoParams, reqEditors ...RequestEditorFn) (*http.Response, error)
GetServerInfoWithResponse(ctx context.Context, params *GetServerInfoParams, reqEditors ...RequestEditorFn) (*GetServerInfoResponse, error)
```

GetServerInfoResponse contains parsed model:
```golang
type GetServerInfoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ServerInfoModel
	JSON401      *Error
	JSON403      *Error
	JSON500      *Error
}
```

As you can see there are several objects represents results for different HTTP codes.

Additionally, each response contains following functions to get HTTP Status and HTTP Code
```golang
// Status returns HTTPResponse.Status
func (r GetServerInfoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetServerInfoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
```

Please note that you should check status separately, the error will be returned only in case of failed request. If server returned an answer, no error will be returned.

Let's put it all together in the single example:
```golang
func (v *Client) GetServerInfo(ctx context.Context) (ServerInfo, error) {
	rs, err := v.GetServerInfoWithResponse(ctx, &vbrclient.GetServerInfoParams{
		XApiVersion: XapiVerV12,
	})
	if err != nil {
		return ServerInfo{}, kerrors.Wrap(err, "Failed to get server info")
	}

	if rs.StatusCode() != 200 {
		return ServerInfo{}, kerrors.New("Failed to get server info").WithField("status", rs.Status())
	}

	si := ServerInfo{
		VbrID:        *rs.JSON200.VbrId,
		Name:         rs.JSON200.Name,
		BuildVersion: rs.JSON200.BuildVersion,
	}

	return si, nil
}

```

## Known Issues
* If you convert specification from `JSON` to `YAML` usiang `oapifixer` tool the order of sections in the specification will be changed(to alphabetical). It is not a problem for the generator but it makes it difficult to review changes in the specification.


## Contributing
Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.