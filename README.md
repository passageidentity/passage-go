# passage-go

This Go SDK allows for easy server-side authentication for applications using [Passage](https://passage.id).

## Authenticating an HTTP request

To authenticate an HTTP request to a specific user, you can pass an `http.Request` type to the Passage `AuthenticateRequest` function. For example:

```go
package main

import (
	"net/http"

	"github.com/passageidentity/passage-go"
)

func exampleRequestHandler(w http.ResponseWriter, r *http.Request) {

  psg := passage.New()
  _, err := psg.AuthenticateRequest(r)
  if err != nil {
    // Authentication check failed!
    w.WriteHeader(http.StatusUnauthorized)
    return
  }

  // Proceed with request handler

}
```
