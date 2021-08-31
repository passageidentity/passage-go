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

  psg := passage.New("<APP_ID>")
  _, err := psg.AuthenticateRequest(r)
  if err != nil {
    // Authentication check failed!
    w.WriteHeader(http.StatusUnauthorized)
    return
  }

  // Proceed with request handler

}
```
## Retrieve User Info
To retrieve information about a user, you should use the `getUser` method. You will need to use a Passage API key, which can be created in the Passage Console under your Application Settings. This API key grants your web server access to the Passage management APIs to get and update information about users.
This API key must be protected and stored in an appropriate secure storage location. It should never be hard-coded in the repository.

```go 

package main

import (
	"net/http"

	"github.com/passageidentity/passage-go"
)

func exampleRequestHandler(w http.ResponseWriter, r *http.Request) {

  psg := passage.New("<APP_ID>", "<PASSAGE_API_KEY>")
  user, err := psg.GetUser("<USER_ID>") 
  if err != nil {
      //Handle err cases
      // - user not found
      // - invalid PASSAGE_API_KEY
      // ...
  }
```