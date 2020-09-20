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

package tidbyt

import (
	"errors"
	"os"
)

const (
	// TidbytDeviceIDEnv is the environment variable that NewConfigFromEnv is looking for when setting the device ID.
	TidbytDeviceIDEnv = "TIDBYT_DEVICE_ID"

	// TidbytTokenEnv is the environment variable that NewConfigFromEnv is looking for when setting the token.
	TidbyTokenEnv = "TIDBYT_AUTH_TOKEN"
)

// Config is a helper to provide configuration values when creating a Tidbyt client.
type Config struct {
	// DeviceID is the DeviceID used to make API requests. It probably shouldn't live here in retrospect.
	DeviceID string

	// Token is the API token generated from the Tidbyt app used to communicate with the Tidbyt API.
	Token string
}

// NewConfigFromEnv is a helper to generate a config object from environment variables.
func NewConfigFromEnv() (*Config, error) {
	// TODO - if we leave this here, we should likely make it optional.
	deviceID := os.Getenv(TidbytDeviceIDEnv)
	if deviceID == "" {
		return nil, errors.New("could not find TIDBYT_DEVICE_ID in environment")
	}

	token := os.Getenv(TidbyTokenEnv)
	if token == "" {
		return nil, errors.New("could not find TIDBYT_AUTH_TOKEN in environment")
	}

	return &Config{
		DeviceID: deviceID,
		Token: token,
	}, nil
}
