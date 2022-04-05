# Implementation of Blockhain in Golang

This is simple implementation of blockhain in golang using _proof of work_ as a concensus protocol, There's no data inside this blockhain but you can add if you want on file `model/data.go` inside that you can modify `struc data` to whatever you want

## Installation

```bash
git clone https://github.com/rizki4106/blockhain.git
```

```bash
cd blockchain && go mod download
```

```bash
go run main.go // it run on port 4000
```

## Run using docker

```bash
docker-compose up
```

## Endpoint

#### 1. Create block

```text
http://localhost:4000/create-block
```

```json
{
  "total_item": 2,
  "message": "New block has been created",
  "block": [
    {
      "index": 4,
      "proof": 243727,
      "timestamp": "2022-04-05 02:35:05.575242056 +0000 UTC m=+86.392622159",
      "previous_hash": "d8552a98514d965b19afb0bcc7f4647184e5b89f7c4716584a29a8f1dd99747e",
      "data": null
    }
  ]
}
```

#### 2. Get all block

```text
http://localhost:4000/get-block
```

```json
{
  "total_item": 2,
  "message": "",
  "block": [
    {
      "index": 1,
      "proof": 1,
      "timestamp": "2022-04-05 02:33:39.183438159 +0000 UTC m=+0.000818312",
      "previous_hash": "0",
      "data": null
    },
    {
      "index": 2,
      "proof": 81243,
      "timestamp": "2022-04-05 02:35:01.907033339 +0000 UTC m=+82.724413462",
      "previous_hash": "d9f4e6d4f4d8a360b296f1da3554b223528a6612d84b7d8ea92d56bf0a754d02",
      "data": null
    }
  ]
}
```

#### 4. Validate blockchain

```text
http://localhost:4000/validate-block
```

```json
{
  "message": "All block looks good"
}
```

## What's Next ?

I still develop feature for handle a lot of nodes
