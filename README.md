# Description

Thin wrapper around standard `httptest` library for creating mocked JSON HTTP responses easily.

## Example

Create a server by calling `MockServer` function providing the necessary endpoints defining the path, HTTP response code and the content of the response.  

All responses will have `application/json` in `Content-Type` header. 

```
import "github.com/burdiyan/httptest"

server := httptest.MockServer(
    &httptest.Endpoint{"/foo", 200, `{"success": true}`},
    &httptest.Endpoint{"/bar", 404, `{"success": false}`},
)
```

After that you can make real HTTP calls to `server.URL` + path you have defined, and the server will respond with the content you have set on the endpoint.
