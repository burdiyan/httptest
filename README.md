Thin wrapper around standard `httptest` library for creating mocked JSON HTTP responses easily.

# Example

Create a server by calling `MockServer` function providing the necessary endpoints defining the path, HTTP response code and the content of the response.  

All responses will include `Content-Type` header as `application/json`.

```
server := httptest.MockServer(
    &Endpoint{"/foo", 200, `{"success": true}`},
    &Endpoint{"/bar", 404, `{"success": false}`},
)
```

After that you can make real HTTP calls to `server.URL` + path you have defined, and the server will respond with the content you have set on the endpoint.
