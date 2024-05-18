# Porto

A utility to help with api mocks and testing.
Quickly establish a port response to make testing backend services 
easier.

## Configuration

Create a file named config.yaml with the following contents:

```yaml
services:
  - endpoint: test
    comment: "test endpoint" 
    default: "Hello Test endpoint" 
```

A quick guide to the fields

| Field | Description |
|-------|-------------|
| endpoint | Represents the route to be added e.g. /test |
| comment  | Description of the route to be applied |
| default  | HTML content to be returned from the route |

## Running the application

The application will default to TCP:8080.
Example how to change the port allocation:

```bash
PORT=9001
```

Then do the following:

1. Update the config.yaml with the required routes
2. Run porto
```bash
./porto
```
