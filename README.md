## grout
Grout is a simple router written in Go which supports middlewares and route decoration.

## Example
It is pretty easy to setup and use - replace the handlers by your own ones.
```Go
func main() {
    router := grout.NewRouter()

    router.AddMiddleware("CheckAccess", accessCheckMiddleware)
    router.CreateRoute("UserCreate", "/user", view.CreateUser, "PUT")
    router.CreateRoute("GetUser", "/user/:id", view.GetUser, "GET")

    log.Printf("Listening on %s...", port)
    log.Fatal(router.Serve(port))
}
```

