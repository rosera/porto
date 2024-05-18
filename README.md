# Porto

A utility to help with API mocks and testing.
Quickly establish a port response to make testing backend services 
easier.

## Configuration

Create a file named `config.yaml` with the following contents:

```yaml
services:
  - endpoint: test
    comment: "test endpoint"
    default: "<h1>Test endpoint</h1>\n<p>This is a test</p>"
    format: "html"
  - endpoint: json-test
    comment: "JSON test endpoint"
    default: "{\"message\": \"Hello JSON Test endpoint\"}"
    format: "json"
  - endpoint: plain-test
    comment: "Plain text test endpoint"
    default: "Hello Plain Text Test endpoint"
    format: "plain"
  - endpoint: image-test
    comment: "Image test endpoint"
    default: "https://via.placeholder.com/150"
    format: "image"
```

A quick guide to the fields

| Field | Description |
|-------|-------------|
| endpoint | Represents the route to be added e.g. /test |
| comment  | Description of the route to be applied |
| default  | Inplace or URL content to be returned from the route |
| format   | How the default should handled. html:json:plain:image |

## Running the application

The application will default to TCP:8080.

To change the port allocation:

1. At the command line create an environment variable for the PORT
```bash
PORT=9001
```

Then do the following:

1. Update the `config.yaml` with the required routes
2. Run porto
```bash
./porto
```
