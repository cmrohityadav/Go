1. HTTP Fundamentals
HTTP Request & Response
Methods (GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS)
Status codes
Headers
Cookies
Query parameters
URL encoding
Content-Type
Accept
Authorization header
User-Agent
Host
HTTP versions (1.1 vs HTTP/2)
2. Starting a Server
http.ListenAndServe()
http.Server{}

Topics:

Custom server
ReadTimeout
WriteTimeout
IdleTimeout
Shutdown
Graceful shutdown
TLS server
HTTP/2 support
3. Routing

Using

http.NewServeMux()

Learn:

Handle
HandleFunc
Pattern matching
Route precedence
Wildcards (Go 1.22+)
Path parameters
Sub-routing
Method-based routing

Example:

GET /users/{id}
POST /users
DELETE /users/{id}
4. Request Object

Understand everything inside

http.Request

Important fields:

Method
URL
Header
Body
Form
MultipartForm
Context
Cookies
RemoteAddr
Host
ContentLength
TLS
5. ResponseWriter

Learn

http.ResponseWriter

Topics:

Write()
WriteHeader()
Header()
Content-Type
JSON response
XML response
File response
Redirect
Streaming response
6. Forms

Topics:

ParseForm()
FormValue()
PostForm
URL Query
URL Encoded forms
7. JSON APIs

Learn

json.Decoder
json.Encoder
Decode request
Encode response
Validation
Unknown fields
Pretty JSON
8. File Upload

Very important.

Topics:

Multipart
ParseMultipartForm
FormFile
MultipartReader
Multiple files
Large files
Streaming
MIME validation
File extension validation
Max upload size
Progress
S3 upload
Presigned URLs
9. File Download

Topics:

http.ServeFile()
http.ServeContent()

Headers

Content-Disposition

Download attachment

Resume download

Range requests

10. Static Files

Learn

http.FileServer()

Topics

Serve images
CSS
JS
SPA support
Prefix stripping
11. Middleware

How middleware works.

Implement

Logging
Authentication
Authorization
Panic recovery
Rate limiting
Request ID
CORS
Compression
Metrics
12. Context

Huge topic.

Learn

r.Context()

Topics

Deadlines
Cancellation
Values
Request timeout
DB cancellation
13. Authentication

Learn

Basic Auth
Bearer Token
JWT
API Keys
Sessions
Cookies
Refresh Tokens
14. Cookies

Topics

http.SetCookie()
r.Cookie()
Secure
HttpOnly
SameSite
Expiration
15. Sessions

Implement manually.

Understand

Session ID
Cookie storage
Redis storage
Memory storage
16. CORS

Understand

OPTIONS

Headers

Access-Control-Allow-Origin

etc.

17. Streaming

Topics

io.Copy()
Chunked encoding
Flush()
SSE
WebSockets (with Gorilla or nhooyr)
Live logs
Video streaming
18. HTTP Client

Learn

http.Client

Topics

GET
POST
PUT
DELETE
Timeouts
Custom Transport
Keep Alive
Redirects
Cookies
Proxy
19. Custom Transport

Very important.

Learn

http.Transport

Topics

Pooling
Idle connections
Max connections
TLS config
20. Reverse Proxy

Learn

httputil.ReverseProxy

Used in

API Gateway
Load Balancer
21. Graceful Shutdown

Topics

Server.Shutdown()

Handle

SIGINT
SIGTERM
22. Error Handling

Return

JSON errors
HTTP errors
Custom error responses
23. Logging

Learn

Request logging
Response logging
Structured logging
Correlation IDs
24. Security

Topics

HTTPS
TLS
CSRF
XSS
Path traversal
Directory traversal
Secure cookies
Header validation
Input sanitization
File upload security
25. Performance

Topics

Connection pooling
Streaming
Buffer reuse
Compression
Keep-Alive
Gzip
HTTP/2
26. Testing

Learn

httptest

Topics

NewRequest
ResponseRecorder
Test server
Integration testing
27. Advanced Routing (Go 1.22+)

Learn

GET /users/{id}
POST /users

Methods in patterns

Wildcards

PathValue()

SetPathValue()

28. Advanced Multipart
Nested multipart
Multipart mixed
Multipart alternative
Large uploads
Streaming uploads
29. HTTP Compression

Topics

gzip
deflate
Accept-Encoding
Content-Encoding
30. HTTP Caching

Headers

ETag
Last-Modified
Cache-Control
Expires
31. Rate Limiting

Implement

Token Bucket
Leaky Bucket
Sliding Window
32. Timeouts

Learn

Server timeout
Client timeout
Read timeout
Write timeout
Idle timeout
Context timeout
33. Production Deployment

Topics

Reverse proxy (Nginx/Caddy)
HTTPS certificates
Health checks (/healthz, /readyz)
Environment variables
Configuration management
Docker
Logging and monitoring
34. Observability
Metrics (Prometheus)
Distributed tracing (OpenTelemetry)
Structured logs
Request IDs
35. Common Packages Around net/http

Although not part of net/http, these are commonly used alongside it:

context
encoding/json
mime/multipart
net
net/url
net/http/httptest
net/http/httputil
io
os
path/filepath
time
crypto/tls
compress/gzip
Suggested Learning Order
HTTP basics
http.Server and ServeMux
http.Request and ResponseWriter
Routing
Query parameters and forms
JSON APIs
Middleware
File upload/download
Static file serving
Cookies and sessions
Authentication and authorization
Context and timeouts
HTTP client
Streaming (SSE/WebSockets)
Reverse proxy
Graceful shutdown
Testing with httptest
Security and performance
Production deployment and observability