package w3

import (
	"context"
	"fmt"
	"reflect"

	"github.com/celo-org/celo-blockchain/rpc"
	"github.com/lmittmann/w3/w3types"
)

// Client represents a connection to an RPC endpoint.
type Client struct {
	client *rpc.Client
}

// NewClient returns a new Client given an rpc.Client client.
func NewClient(client *rpc.Client) *Client {
	return &Client{
		client: client,
	}
}

// Dial returns a new Client connected to the URL rawurl. An error is returned
// if the connection establishment failes.
//
// The supported URL schemes are "http", "https", "ws" and "wss". If rawurl is a
// file name with no URL scheme, a local IPC socket connection is established.
func Dial(rawurl string) (*Client, error) {
	client, err := rpc.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return &Client{
		client: client,
	}, nil
}

// MustDial is like [Dial] but panics if the connection establishment failes.
func MustDial(rawurl string) *Client {
	client, err := Dial(rawurl)
	if err != nil {
		panic(fmt.Sprintf("w3: %s", err))
	}
	return client
}

// Close closes the RPC connection and cancels any in-flight requests.
//
// Close implements the [io.Closer] interface.
func (c *Client) Close() error {
	c.client.Close()
	return nil
}

// CallCtx creates the final RPC request, sends it, and handles the RPC
// response.
//
// An error is returned if RPC request creation, networking, or RPC response
// handeling fails.
func (c *Client) CallCtx(ctx context.Context, calls ...w3types.Caller) error {
	// no requests = nothing to do
	if len(calls) <= 0 {
		return nil
	}

	batchElems := make([]rpc.BatchElem, len(calls))
	var err error

	// create requests
	for i, req := range calls {
		batchElems[i], err = req.CreateRequest()
		if err != nil {
			return err
		}
	}

	// do requests
	if len(batchElems) > 1 {
		// batch requests if >1 request
		err = c.client.BatchCallContext(ctx, batchElems)
		if err != nil {
			return err
		}
	} else {
		// non-batch requests if 1 request
		batchElem := batchElems[0]
		err = c.client.CallContext(ctx, batchElem.Result, batchElem.Method, batchElem.Args...)
		if err != nil {
			switch reflect.TypeOf(err).String() {
			case "*rpc.jsonError":
				batchElems[0].Error = err
			default:
				return err
			}
		}
	}

	// handle responses
	var callErrs CallErrors
	for i, req := range calls {
		err = req.HandleResponse(batchElems[i])
		if err != nil {
			if callErrs == nil {
				callErrs = make(CallErrors, len(calls))
			}
			callErrs[i] = err
		}
	}
	if len(callErrs) > 0 {
		return callErrs
	}
	return nil
}

// Call is like [Client.CallCtx] with ctx equal to context.Background().
func (c *Client) Call(calls ...w3types.Caller) error {
	return c.CallCtx(context.Background(), calls...)
}

// CallErrors is an error type that contains the errors of multiple calls. The
// length of the error slice is equal to the number of calls. Each error at a
// given index corresponds to the call at the same index. An error is nil if the
// corresponding call was successful.
type CallErrors []error

func (e CallErrors) Error() string {
	if len(e) == 1 {
		return fmt.Sprintf("w3: call failed: %s", e[0])
	}
	return "w3: one ore more calls failed"
}

func (e CallErrors) Is(target error) bool {
	_, ok := target.(CallErrors)
	return ok
}
