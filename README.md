# Trumpeter Trump Tweet Fetcher

A simple microservice conforming to the [Open Microservice Specification](https://microservice.guide/) (OMS) that fetches the most recent tweet from President Trump

# Building and Running

There are two ways to build and run this, directly as a binary via:

```
go build
CONSUMER_KEY=<> CONSUMER_SECRET=<> ACCESS_TOKEN=<> ACCESS_TOKEN_SECRET=<> ./trumpeter
```

or via Docker:

```
docker build -t trumpeter:latest .
docker run -e "CONSUMER_KEY=<>" -e "CONSUMER_SECRET=<>" -e "ACCESS_TOKEN=<>" -e "ACCESS_TOKEN_SECRET=<> "trumpeter
```

## Running the tests

The integration test suite is written using [Ginkgo](https://onsi.github.io/ginkgo/). If you have ginkgo installed, they can be run via:

```
CONSUMER_KEY=<> CONSUMER_SECRET=<> ACCESS_TOKEN=<> ACCESS_TOKEN_SECRET=<> ginkgo -r
```

Alternatively, with fewer features, using `go test`:

```
CONSUMER_KEY=<> CONSUMER_SECRET=<> ACCESS_TOKEN=<> ACCESS_TOKEN_SECRET=<> go test ./...
```
