package simplerouter

import (
	"encoding/json"
	"net/http"
)

type Ctx struct {
	w http.ResponseWriter
	r *http.Request
}

func (c *Ctx) String(code int, res string) error {
	c.w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.w.WriteHeader(code)
	_, err := c.w.Write([]byte(res))
	return err
}

func (c *Ctx) JSON(code int, val any) error {
	c.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.w.WriteHeader(code)
	return json.NewEncoder(c.w).Encode(val)
}
