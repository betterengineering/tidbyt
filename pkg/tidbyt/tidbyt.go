// Copyright 2020-Present Mark Spicer
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
// WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package tidbyt provides a helpful client to the Tidbyt API. See more at https://tidbyt.dev/
package tidbyt

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	publicapi "github.com/lodge93/tidbyt/api/public-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	// APIAddress is the publicly provided API address for Tidbyt.
	APIAddress = "api.tidbyt.com"

	// APIPort is the publicly provided API port for Tidbyt.
	APIPort = "443"
)

// TidbytClient is a wrapper around the generated PublicAPIClient so one does not have to write the boiler plate to
// setup a connection.
type TidbytClient struct {
	publicapi.PublicAPIClient
	conn *grpc.ClientConn
}

// NewTidbytClient setups a gRPC connection with a client to communicate with the Tidbyt API. Be sure to close the
// connection with the provided Close() method when you're done with it.
func NewTidbytClient(token string) (*TidbytClient, error) {
	conn, err := NewTidbytAPIConn(token)
	if err != nil {
		return nil, err
	}

	client := publicapi.NewPublicAPIClient(conn)

	return &TidbytClient{
		client,
		conn,
	}, nil
}

// NewTidbytAPIConn creates a new connection to the Tidbyt gRPC API. This is used both to setup a new client and to
// use reflection in generating the client.
func NewTidbytAPIConn(token string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return grpc.DialContext(ctx, net.JoinHostPort(APIAddress, APIPort),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(tokenAuth{
			token: token,
		}),
	)
}

// Close closes the underlying connection. Please close your connections!
func (t *TidbytClient) Close() error {
	return t.conn.Close()
}

