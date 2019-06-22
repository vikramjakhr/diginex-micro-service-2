# diginex-micro-service-2

REST API for reversing a string

## How to use it:

* See usage with:

```
dignex-micro-service-2 help
```
```
Diginex Micro Service 2.

Usage:

  dignex-micro-service-2 [commands|flags]

The commands & flags are:

  help              prints help

  --port              port on which application will run (default: 8081)

Examples:

  # prints help:
  dignex-micro-service-2 help

  # sample usage
  dignex-micro-service-2 --port 8081

```

* Run test cases and code coverage (go > 1.12)
```
go test --cover
```

* Build the binary or [Download from here](./diginex-micro-service-2)
```
go build
```

* Run the binary as a server
```
dignex-micro-service-2                  // default port is 8081
dignex-micro-service-2 --port=9080      // with custom port
```

#### Accessing the API:

```
curl -X POST \
  http://localhost:8081/reverse \
  -d '{
 "message":"abcdef"
}'
```
Output:
```
{
    "message": "fedcba"
}
```