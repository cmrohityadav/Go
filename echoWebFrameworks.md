# Echo Framework

Echo is a high-performance, minimalist Go web framework designed for building scalable and maintainable server-side applications.

---

## REST (Representational State Transfer)

REST is an architectural style used to design web services.  
It leverages HTTP methods to perform CRUD (Create, Read, Update, Delete) operations on resources.

### Key Principles of REST
- **Stateless** – Every request is independent.
- **Client-Server Architecture** – Separation of concerns.
- **Uniform Interface** – Standardized communication using HTTP.
- **Cacheable** – Response data can be cached.
- **Layered System** – Intermediaries like proxies, load balancers allowed.
- **Resource-based** – Everything treated as a resource accessed via URLs.

# Content

---

## Terms

- `Serialization` (`struct → JSON/string/binary`)
- `Marshalling` = converting a Go struct → JSON
- `Deserialization` (`JSON/string/binary → struct/object`)
- `Unmarshalling` = converting JSON → Go struct

## Methods

- **GET**  
  Ye request sirf data **fetch** ya **read** karne ke liye use hoti hai.  
  Server ki state change nahi hoti.

- **POST**  
  Ye request **naya data create** karne ke liye use hoti hai.  
  Har POST request se generally naya resource ban jata hai.

- **PUT**  
  Ye **pure resource ko update/replace** karne ke liye use hota hai.  
  Agar aapko **sirf ek field** update karni ho, tab bhi **pura payload (saare fields)** bhejne padte hain.

- **PATCH**  
  Ye **partial update** ke liye hota hai.  
  Yani **kisi single field** ko update karna ho to **sirf wahi field** bhejna hota hai,  
  saare fields dene ki jarurat nahi hoti.

- **DELETE**  
  Ye request kisi resource ko **delete/remove** karne ke liye use hoti hai.

## Installation
```bash
go get github.com/labstack/echo/v4
```
## Basic Compare with Express JS

| Concept       | Express.js (Node) | Echo (Go)                     |
| ------------- | ----------------- | ----------------------------- |
| Functions     | Callback/async    | Normal functions + goroutines |
| Server        | app.listen()      | e.Logger.Fatal(e.Start())     |
| Routing       | app.get(), post() | e.GET(), POST()               |
| Middlewares   | app.use()         | e.Use()                       |
| Body parse    | req.body          | c.Bind()                      |
| Response JSON | res.json()        | return c.JSON()               |

## Basic Server
```go
package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Echo!")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
```
### echo.New()
- `echo` ek package (library) hai jo humne import kiya tha
- `New()` echo package ke andar ek function hai
- Ye function Echo framework ka main instance create karta hai
- Ye basically ek server ka object banata hai, jisme:
1. Router hota hai

2. Middleware stack hota hai

3. Logger hota hai

4. HTTP server settings hoti hain

4. Error handler hota hai

5. Context pool (memory optimization) hota hai
- Means, complete web application ki backbone.
- Ye return karta hai ek pointer to Echo struct, jiska type: `*echo.Echo`

### func(c echo.Context) error
- Context = Request + Response + Extras (One Object)
- Its contains:Request,Response,Path params,Query params,Body bind, etc

| Feature             | Express Equivalent               |
| ------------------- | -------------------------------- |
| Request             | `req`                            |
| Response            | `res`                            |
| Path params         | `req.params`                     |
| Query params        | `req.query`                      |
| Body bind           | `req.body`                       |
| File upload         | `req.files`                      |
| Cookie              | `req.cookies`                    |
| Store values        | `req.locals`                     |
| Middleware data     | `req.next()` + `locals`          |
| Request IDs, logger | Not default available in express |

#### Send JSON Response
```go
c.JSON(http.StatusOK, map[string]string{"msg": "Hello"})
```
#### Get Query Param
- URL: /user?name=rohit
```go
name := c.QueryParam("name")
```
#### Get URL Path Param
```go
id := c.Param("id")
```
#### Bind Body Data to Struct
- Like req.body
```go
u := new(User)
c.Bind(u)
```
#### Get Header
```go
token := c.Request().Header.Get("Authorization")
```
#### Set Cookie
```go
c.SetCookie(&http.Cookie{Name: "session", Value: "xyz"})

```
#### Store Custom Data (locals alternative)
```go
c.Set("userId", 20)
id := c.Get("userId")
```
#### Why Handler Returns error?
- Echo automatically error handle karta hai (logging, responses).
- Express me error middleware manually banana padta:
```js
app.use((err, req, res, next) => ...)
```
- Echo me ye built-in hai. Bas return err karo, kaam khatam.

### e.Start(":3000")
- Start() function HTTP server ko start karta hai.
- `:` ka matlab → machine ka default IP use karna.
```js
// similiar in Express
app.listen(3000)
```

| Code                          | Meaning                              |
| ----------------------------- | ------------------------------------ |
| `e.Start(":3000")`            | Bind to **all interfaces (0.0.0.0)** |
| `e.Start("127.0.0.1:3000")`   | Local only                           |
| `e.Start("192.168.x.x:3000")` | LAN network                          |
| `e.Start("<public-ip>:3000")` | Public facing server                 |



### e.Logger.Fatal(...)

- error ko log (print) karta hai, aur
- program ko turant bandh (exit) kar deta hai.

### e.StartTLS(":443", "cert.pem", "key.pem")

### e.Shutdown(ctx context.Context)
- Shutdown() Echo server ko gracefully stop (properly band) karta hai
- Graceful shutdown = background requests complete ho jaye, phir server band ho
- Zor se band nahi karta, lehje se politely band karta hai
- Server pending requests complete karke phir close hota hai

### e.Any()
- Echo me ek route banata hai jo saare HTTP methods accept karta hai
- GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD etc
```go
e.Any("/api/user", func(c echo.Context) error {
	return c.String(200, "This route accepts ANY HTTP method")
})
```
- Is route ko call karoge:
  -  GET /api/user
  -  POST /api/user
  -  DELETE /api/user
  -  PUT /api/user
  -  ➡️ Sabko yeh handle karega!
```js
app.all("/api/user", (req, res) => {
  res.send("This accepts any method");
});
```
### e.Group()
- e.Group() Echo me route grouping ke liye use hota hai
- Same prefix / same middleware / same version / same permission wale routes ko ek group me put kar sakte ho
```go
api := e.Group("/api")

api.GET("/users", getUsers)
api.POST("/users", createUser)
api.GET("/products", getProducts)

```
- Actual routes kya banenge?
| Route               |
| ------------------- |
| `/api/users`        |
| `/api/users` (POST) |
| `/api/products`     |

```js
//router.js
const router = express.Router();

router.get("/users", ...)
router.post("/users", ...)
//app.js
app.use("/api", router)
```

### e.File(routePath, fileLocation)
- e.File() Echo ka method hai jo client ko ek single file serve karta hai (PDF, Image, HTML, JS, etc.).
- `e.File("/", "public/index.html")`
