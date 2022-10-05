package passage

import (
	_ "embed"
	"fmt"

	"gopkg.in/resty.v1"
)

//go:embed version.txt
var version string

func newRequest() *resty.Request {
	return resty.New().SetHeader("Passage-Version", fmt.Sprintf("passage-go %s", version)).R()
}
