# Trafilea Challenge

<img align="right" width="159px" src="https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.crunchbase.com%2Forganization%2Ftrafilea&psig=AOvVaw2jgmozHRwuMSUKcpcTplfi&ust=1705847308306000&source=images&cd=vfe&opi=89978449&ved=0CBIQjRxqFwoTCPDcjcmW7IMDFQAAAAAdAAAAABAH" alt="">

[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)

This is a project developed by [Gadiel Malagrino](https://github.com/gadielMa). The system prints all the numbers from 1 to 100. However, for multiples of 3,
instead of the number, print `Type 1`. For multiples of 5 print `Type 2`. For numbers
which are multiples of both 3 and 5, print `Type 3`.

### Description

70% of `unit tests` coverage.

In this project I decided on a 3-layer scaffold.

The first layer `cmd`, contains the initial configuration of the project and a `api` folder where the POST, PUT, etc. are.

The second layer is `service`, where the business logic is, there are also the `interfaces` to the repositories. This way we always put the interfaces on a layer before using them.

The last layer is `repositories`, where the uses of external services and interactions with the databases are.

I didn't need to create my own Docker images, since my `Dockerfile` was very simple.

**Language:** [go1.21.x](https://tip.golang.org/doc/go1.21)

<img src="https://i.imgur.com/3elNhQu.png" alt="">

**Chosen Architecture:** Clean Architecture
<img src="https://miro.medium.com/v2/resize:fit:800/1*0R0r00uF1RyRFxkxo3HVDg.png">

**Framework:** [go-chi](https://github.com/go-chi/chi)

<img src="https://camo.githubusercontent.com/f72d07b7d898f8935d557867df17416a1b430a2572f8ea1bae57d1700f5c754b/68747470733a2f2f63646e2e7261776769742e636f6d2f676f2d6368692f6368692f6d61737465722f5f6578616d706c65732f6368692e737667" alt="">

## Getting started

### Prerequisites
- **[Git](https://github.com/gadielMa/trafilea)**

### Steps

```sh
$ git clone https://github.com/gadielMa/trafilea.git
```

```sh
$ cd trafilea/
```

```sh
$  go run cmd/main.go
```

Wait a minute and...

Use a REST Client like [Postman](https://www.postman.com) to run the following curl:

```sh
curl --location --request GET 'http://localhost:3000/trafilea/divisible'
```
