# TinyUrl-go

Implementing tiny-url with go, this is my first golang project, thanks for [go-gin-example](https://github.com/eddycjy/go-gin-example).

# Basement

MurmurHash32 + Redis Increment Key  -> base 62 

# How to run

## Required

Redis

## Conf

You should modify .env

```txt
RUN_MODE = test
PORT=8080
BASE_URL=http://localhost:8080

REDIS_HOST="localhost:6379"
```

# Load Tests

![](/assets/wrt_test.png)
