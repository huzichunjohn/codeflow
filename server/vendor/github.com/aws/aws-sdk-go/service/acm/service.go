// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package acm

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// Welcome to the AWS Certificate Manager (ACM) CLI Command Reference. This
// guide provides descriptions, syntax, and usage examples for each ACM CLI
// command. You can use AWS Certificate Manager to request ACM Certificates
// for your AWS-based websites and applications. For general information about
// using ACM and for more information about using the console, see the AWS Certificate
// Manager User Guide (url-acm-ug;acm-overview.html). For more information about
// using the ACM API, see the  AWS Certificate Manager API Reference (http://docs.aws.amazon.com/acm/latest/APIReference/Welcome.html).
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type ACM struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "acm"

// New creates a new instance of the ACM client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ACM client from just a session.
//     svc := acm.New(mySession)
//
//     // Create a ACM client with additional configuration
//     svc := acm.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ACM {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *ACM {
	svc := &ACM{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-12-08",
				JSONVersion:   "1.1",
				TargetPrefix:  "CertificateManager",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ACM operation and runs any
// custom request initialization.
func (c *ACM) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
