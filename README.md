# jwt-dump

This is a simple CLI utility for extracting JWT header and paylod.

## Installation

```
go install github.com/evlbit/jwt-dump@latest
```

## Usage

Just simply run the command, provide a token, and its contents will be outputted.

```
$ jwt-dump
Enter token:
eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxMjM0LCJpYXQiOjE3Mzc4ODA3NjUsImV4cCI6MTczNzkwOTU2NX0.vz0TO00wX8F_OhsMSsuDaRuKITJosIT2qmMoX8HTXf0

Header:
{
  "typ": "JWT",
  "alg": "HS256"
}

Payload:
{
  "user_id": 1234,
  "iat": 1737880765,
  "exp": 1737909565
}
```
