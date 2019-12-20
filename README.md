# grader-api-v2

> This is a proof-of-concept prototype of version 2 of the grader API.

![General Flow](https://static.swimlanes.io/ff07aa1e89a7f1032ebfa9b5ba88a108.png)

## Endpoints

### SubmitForGrading
```
Params:
- assignment URN
- zip key
- student name
- timeout
```

### CreateAssignment
```
Params:
- image name
- image tar
```

## Build
```
make build
```

## Run `graderd` locally
```
./bin/graderd [--addr address] [--port port] [--cert public cert] [--key private key]
```

## Clients for `graderd`

### Using gRPC client
```
./bin/grader-cli [-a address:port] [-c public cert] submit [-u assignment-urn] [-z zip-key] [-n student's name]
```

## Using `curl`
```
curl -X POST \
-k https://localhost:9090/api/submit \
-H "Content-Type: text/plain" -d \
'{
    "urn_key": "Hello",
    "zip_key": "Hello",
    "student_name": "Hello",
    "timeout": 600
}'
```