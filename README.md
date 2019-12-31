# sago

Implementation of SAGA and Event Sourcing pattern in Go.

![sago design](svg/sago.svg?sanitize=true)

## Run

Run broker - redis

```
docker run \
    --name redis \
    -p 6379:6379 \
    redis --requirepass topsecret
```

Run order service

```
go run main.go
```

Run clients against service

```
go run client/order/client.go
```

Simulate customer SAGA

```
go run client/customer/client.go
```
