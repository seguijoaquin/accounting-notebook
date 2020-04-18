# Accounting Notebook API

## Launching the app

The app might be launched either with Docker, with go command-line or as a binary.

- Running with docker
```bash
$ docker build -t api-img -f Dockerfile.min . // To generate a minimal image
$ docker build -t api-img -f Dockerfile .     // To generate a normal image
$ docker run -p 8080:8080 -name=api-container api-img
```

- Running with command-line
```bash
$ go run cmd/api/main.go
```

- Compile and run
```$bash
$ go build -o app ./cmd/api
$ ./app
```


### Commiting a transaction

To commit a transaction

```bash
curl -X POST -d '{"ammount":10,"type":"credit"}' http://localhost:8080/api/v1/transactions
```

_Response_
```
http status 201 CREATED
```

```json
{
    "id": "5",
    "type": "credit",
    "ammount": 10,
    "effective_date": "2020-04-18T06:12:25Z"
}
```

> Trying to commit a debit transaction while not havving enough funds will result in a _http status UnprocessableEntity_ error

### Getting all transactions

To get all transactions

```bash
curl -X GET http://localhost:8080/api/v1/transactions
```

_Response_
```
http status 200 OK
```

```json
[
    {
        "id": "1",
        "type": "credit",
        "ammount": 10,
        "effective_date": "2020-04-18T06:12:25Z"
    },
    {
        "id": "2",
        "type": "credit",
        "ammount": 20,
        "effective_date": "2020-04-18T06:12:25Z"
    }
]
```

### Getting a transaction by id

To get a transaction by id

```bash
curl -X GET http://localhost:8080/api/v1/transactions/:id
```

_Response_
```
http status 200 OK
```

```json
{
    "id": "1",
    "type": "credit",
    "ammount": 10,
    "effective_date": "2020-04-18T06:12:25Z"
}
```

### Getting the account balance

To get the account balance

```bash
curl -X GET http://localhost:8080/api/v1/
```

_Response_
```
http status 200 OK
```

```json
{
    "balance": "99999",
}
```