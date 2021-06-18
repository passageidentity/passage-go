package passage_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/passageidentity/passage-go"
)

func TestValidAuthenticateRequest(t *testing.T) {
	// Create a mock request
	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add a VALID authentication token to the request
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwYXNzYWdlIiwic3ViIjoiQ2NGa0djWXdLTkxJcEFtWkpIdnQifQ.EdUKRJwgK7ICRm4uIfGcL3v8zZRyrPWm1zUg5bnQ0m2ovLPGk2yLARZfpE6xhYhd22aemBx8EtNvg6hBg3qOVxKgrAsL-V78Z7S6DKJRRxbKSfJlz55SZOfklo5pz_VwVtOE34oWlq9JRbFaFHxhJuST3Z9ykJEm-1Ihr9ACX453pX9iTcLJmVMWbT4AlPAPtAL4rNsAU9b-KDQC1FUlTNE5e7HfOAaSnwz9qQZpviDAHYomJcWxFCWMri4Zj6PVQb8l57nMUqhf2BU4_q5iVJpV7DlODwApkQIPNWCGZj3oqeJv4T0ldUivtgq1YzEdK4Cw6__13f2IRuDTBHFKsQ")

	// Set a VALID public key as a properly configured server would
	os.Setenv("PASSAGE_PUBLIC_KEY", "LS0tLS1CRUdJTiBSU0EgUFVCTElDIEtFWS0tLS0tCk1JSUJDZ0tDQVFFQTVrZGlybnk1WnFjZ2NxcTg3YklMamtycFI3Ly9FRXRHR3lVL0xESkwxSnU2THRyaVNJemkKdzVPY2h4RkFKb25OUWhwd3dwKzNoT05OWWpKeFY1WWpibWlsc2ZMNWxsajJzZlJxY3lTMjBlZFZ2RDJqZ3hYWQpvT1g2SzU5dWtNQllzU0ZCcWtzM0UvYmNxd2YxN0U2bEIwVFhNY3NYUEtBNUE2QlJoMUVQSGtvQjlKWHJiS3hXCnhGT3JsRkJSRnJMVncvVmNUYnM3SXZRSGU4c3kxdlZoZGlkSng5R095MXVaQVlPL2dTcVNGNFRYd2RaaXJBQ1MKMjJFQlJFSSsrbWh1WWR3aWpaL3lZU1JwRFM4V0h1aVVMYjJrYnhVNUhoVG1jcnNrSjNjeEdnNHM1ZFA0V0YvZwpoMUNnYmdUczQ0dzFPVnUxdkNpNmMxQTQxMng3Wmh3UnR3SURBUUFCCi0tLS0tRU5EIFJTQSBQVUJMSUMgS0VZLS0tLS0K")

	psg := passage.New("fakeAppHandle")
	user, err := psg.AuthenticateRequest(req)
	if err != nil {
		t.Error("unexpected error authenticating request:", err)
	}
	if user.Handle != "CcFkGcYwKNLIpAmZJHvt" {
		t.Error("expected user handle uLljHFDIvwSJEVeuYRvs but found", user.Handle)
	}
}

func TestInvalidAuthenticateRequest(t *testing.T) {
	// Create a mock request
	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add a VALID authentication token to the request
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwYXNzYWdlIiwic3ViIjoiQ2NGa0djWXdLTkxJcEFtWkpIdnQifQ.T0gDJTB0bd-DF6EgEO-PjKk2tbEm9ZN3qVjDeFLB1aGWnm9TYlkvxmAHtYiYmhLkRTvaO_eQIFHToQgIuRwQJlMuZyuzsISmuJCtY9dds-9q3WSoWiaiaupO6ljOSI-k31pAAx_SIaDgdl2IgH3pKLUtDZ29eZxT1HZpwFI_NJ0i_2rnrfWyDXaovgOlxTPlXW4b-GBxsPNzxLTF6-wHSMkP0hjzaLgxyWpuapG5SFIGVKBU2PtdFhChrAv-7Lk_5WxBy-SAsccrJUcYTXgVJcCIC3_rSeKlhnkbYcfQwJO9PkEq2fSCO9pt_v2LUgtC1TeMp2JcMC-6qLe2QQ44uw")

	// Set a VALID public key as a properly configured server would
	os.Setenv("PASSAGE_PUBLIC_KEY", "LS0tLS1CRUdJTiBSU0EgUFVCTElDIEtFWS0tLS0tCk1JSUJDZ0tDQVFFQTVrZGlybnk1WnFjZ2NxcTg3YklMamtycFI3Ly9FRXRHR3lVL0xESkwxSnU2THRyaVNJemkKdzVPY2h4RkFKb25OUWhwd3dwKzNoT05OWWpKeFY1WWpibWlsc2ZMNWxsajJzZlJxY3lTMjBlZFZ2RDJqZ3hYWQpvT1g2SzU5dWtNQllzU0ZCcWtzM0UvYmNxd2YxN0U2bEIwVFhNY3NYUEtBNUE2QlJoMUVQSGtvQjlKWHJiS3hXCnhGT3JsRkJSRnJMVncvVmNUYnM3SXZRSGU4c3kxdlZoZGlkSng5R095MXVaQVlPL2dTcVNGNFRYd2RaaXJBQ1MKMjJFQlJFSSsrbWh1WWR3aWpaL3lZU1JwRFM4V0h1aVVMYjJrYnhVNUhoVG1jcnNrSjNjeEdnNHM1ZFA0V0YvZwpoMUNnYmdUczQ0dzFPVnUxdkNpNmMxQTQxMng3Wmh3UnR3SURBUUFCCi0tLS0tRU5EIFJTQSBQVUJMSUMgS0VZLS0tLS0K")

	psg := passage.New("fakeAppHandle")
	_, err = psg.AuthenticateRequest(req)
	if err == nil {
		t.Error("an invalid request was successfully authenticated")
	}
}
