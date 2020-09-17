# go-rest-api
A REST API implementation in golang.

## Packages used
1. HTTP Router - [gorilla/mux](https://github.com/gorilla/mux)
2. [Mysql driver for golang](https://github.com/Go-SQL-Driver/MySQL/)
3. [JWT for authentication](https://github.com/dgrijalva/jwt-go)

## Implementation
I have implemented a basic REST API to understand and delve deep into the language. MySQL is used as a database backend.

Most of the public APIs implement Rate Limiter upon various factors like IP Address, Authentication, subscription plans and others.
There are various implementations of Rate Limiter - Token Bucket, Leaky Bucket, Fixed Window counter and others.

I have used golang's [rate](https://godoc.org/golang.org/x/time/rate) package to implement a simple rate limiter, differential rate limiting basis user authentication. API will provide/limit basis user's authentication, an unsigned user will have different threshold (lower) compared to signed user (higher).

---
## Further Reading
1. [Go Packages](https://golangbot.com/go-packages/)
2. [Rate Limier](https://www.alexedwards.net/blog/how-to-rate-limit-http-requests)
3. [Rate Limiter](https://gobyexample.com/rate-limiting)
4. [Middleware](https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81)
5. [JWT](https://jwt.io/introduction/)
6. [Python implementation of REST API](https://github.com/karanrn/rest-api-python)
