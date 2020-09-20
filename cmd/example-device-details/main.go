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

package main

import (
	"context"
	"log"
	"time"

	publicapi "github.com/lodge93/tidbyt/api/public-api/proto"
	"github.com/lodge93/tidbyt/pkg/tidbyt"
)

// example-device-details is an example program to get the device details of your Tidbyt and log them to the screen!
func main() {
	cfg, err := tidbyt.NewConfigFromEnv()
	if err != nil {
		log.Fatalf("could not get config: %s\n", err)
	}

	client, err := tidbyt.NewTidbytClient(cfg.Token)
	if err != nil {
		log.Fatalf("could not create tidbyt client: %s\n", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := client.GetDevice(ctx, &publicapi.GetDeviceRequest{DeviceID: cfg.DeviceID})
	if err != nil {
		log.Fatalf("could not get device details: %s\n", err)
	}

	log.Println(req)
}
