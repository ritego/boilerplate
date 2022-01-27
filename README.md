# Aellacredit - Jara

## Introduction
Jara is wallet interest calculator. It's a port of the (wallet interest service)[https://github.com/aellacredit/wallet-interest/] from node to go, with the expectation of drastically reduce the running time for the process.

## Goals
- Compute interest for a 100k users in less than 30mins
- A drop-in replacement for existing wallet interest service

## Approach
- Use goroutines extensivelly, they are cheap
- All storage activity should be append only, completely avoid updates and read only from cache (in-memory or disk)

## Usage
```bash
$ git clone https://github.com/aellacredit/jara
$ cd jara
$ go run main.go
```

```
$	go run main.go compute settlement --class=all --period=daily --from="2021-12-07 00:00:00" --to="2021-12-07 23:59:59" --chunk=12
```


## TODO
- [NO INTEREST WALLETS] batch insert interest = 0 to db, ignore if exist
- implement offset and limit for settlement n payout repos