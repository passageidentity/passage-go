___To use the Passage Go SDK, you'll need your Passage App ID. You can create a new Passage App in the [console](https://console.passage.id).___

### Authenticating a Request

Passage makes it easy to associate an HTTP request with an authenticated user. The following code can be used to validate that a request was made by an authenticated user.

```go
import (
	"net/http"

	"github.com/passageidentity/passage-go"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {

	// Authenticate this request using the Passage SDK:
	psg, _ := passage.New("<PASSAGE_APP_ID>", nil)
	_, err := psg.AuthenticateRequest(r)
	if err != nil {
		// 🚨 Authentication failed!
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// ✅ Authentication successful. Proceed...

}
```

### Authorizing a User

It is important to remember that `psg.AuthenticateRequest()`  validates that a request is properly authenticated and returns the authenticated user's Passage identifier, but an additional _authorization_ check is typically required.

```go
import (
	"net/http"

	"github.com/passageidentity/passage-go"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {

	// Authenticate this request using the Passage SDK:
	psg, _ := passage.New("<PASSAGE_APP_ID>", nil)
	passageID, err := psg.AuthenticateRequest(r)
	if err != nil {
		// 🚨 Authentication failed!
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// ✋ Authentication successful, but let's check authorization:
	allowed := myAuthorizationCheck(passageID, "role.manager")
	if !allowed {
		// 🚨 Authorization failed!
		w.WriteHeader(http.StatusForbidden)
		return
	}
	
	// ✅ Authentication & authorization successful. Proceed...

}
```

### User Management

In addition to authenticating requests, the Passage SDK also provides a way to securely manage your users. These functions require authentication using a Passage API key. API keys can be managed in the [Passage Console](https://console.passage.id). 

___Passage API Keys are sensitive! You should store them securely along with your other application secrets.___

```go
import (
	"net/http"

	"github.com/passageidentity/passage-go"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {

  psg, _ := passage.New("<PASSAGE_APP_ID>", &passage.Config{
    APIKey: "<PASSAGE_API_KEY>",
  })
  passageID, err := psg.AuthenticateRequest(r)
  if err != nil {
		// 🚨 Authentication failed!
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
  
  // The passageID returned above can be used to get user information:
  passageUser, err := psg.GetUser(passageID) 
  if err != nil {
		// 💀 Couldn't get the Passage User for some reason...
		w.WriteHeader(http.StatusInternalServerError)
		return
  }
  
  // The passageUser struct can now be inspected for detailed information
  // about the user. A full list of fields can be found here:
  //   https://github.com/passageidentity/passage-go/blob/main/user.go
  _ = passageUser
  
}
```
