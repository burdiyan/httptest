# Description

Thin wrapper around standard `httptest` library for creating mocked JSON HTTP responses easily.

## Example

Create a server by calling `MockServer` function providing the necessary endpoints defining the path, HTTP response code and the content of the response.  

You can specify custom content type or use handy constructors for JSON and HTML types.

```
import "github.com/burdiyan/httptest"

server := httptest.MockServer(
    httptest.JSONEndpoint("/foo", 200, `{"success": true}`),
    httptest.HTMLEndpoint("/bar", 404, `<p>Hello World!</p>`),
)
```

After that you can make real HTTP calls to `server.URL` + path you have defined, and the server will respond with the content you have set on the endpoint.
