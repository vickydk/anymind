# anymind
Project for Back-end Engineer (AnyLogi/Golang) at AnyMind Group

### Install Static check and linter

1. make staticcheck_install
2. make linter_install

### Validate code

1. make staticcheck
2. make linter

### Running App

1. make compose
2. Access localhost:8811

### Stop App

1. make compose_down

### Generate Data

1. make run_benchmark

### List API

Add Transaction
```azure
curl --location --request POST 'http://127.0.0.1:8811/api/v1/transactions' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_uuid": "232821ec-0a53-4c7e-9988-129050ec2bd2",
    "amount": 3.1,
    "date_time": "2023-01-24T16:05:05+07:00"
}
```

Check History

```azure
curl --location --request POST 'http://127.0.0.1:8811/api/v1/history/search' \
--header 'Content-Type: application/json' \
--data-raw '{
    "endDateTime": "2023-01-24T09:09:01+00:00"
}'
```

## Author

[VickyDk](https://github.com/vickydk)