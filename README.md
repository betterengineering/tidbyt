# Tidbyt
> **WARNING** This is an early version of this client. It's functional, however, the structure and methods could change
> substantially.

This is an unofficial client for the [Tidbyt API](https://tidbyt.dev/). It uses gRPC reflection to get proto definitions
for the services the API exposes and generates a client using protoc.

If there are official implementations released, this will either be updated to utilize those clients or be deleted as
unnecessary.

## Example
Check out [cmd/example-device-details/main.go](cmd/example-device-details/main.go) for an example of using the client.
To run the program as is, set your device details as environment variables and run the following:
```bash
export TIDBYT_DEVICE_ID="<device-id>"
export TIDBYT_AUTH_TOKEN="<auth-token>"
go cmd/example-device-details/main.go
```

## Regenerate Client
To regenerate the client (you likely don't need to do this), run the following:
```bash
export TIDBYT_DEVICE_ID="<device-id>"
export TIDBYT_AUTH_TOKEN="<auth-token>"
./scripts/gen-proto.sh
```

See the script comments for more information about the assumptions it's making.