package passage

import "gopkg.in/resty.v1"

func newRequest() *resty.Request {
	return resty.New().SetHeader("Passage-Version", version).R()
}
