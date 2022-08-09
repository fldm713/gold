package gold

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

type (
	Binder interface {
		Bind(i any, c *Context) error
	}

	DefaultBinder struct{}
)

func (b *DefaultBinder) Bind(i any, c *Context) error {
	return b.BindBody(i, c)
}

func (b *DefaultBinder) BindBody(i any, c *Context) error {
	r := c.R
	if r == nil {
		return nil
	}
	contentType := r.Header.Get("Content-Type")
	switch {
	case strings.HasPrefix(contentType, "application/json"):
		if err := json.NewDecoder(r.Body).Decode(i); err != nil {
			return err
		}
	case strings.HasPrefix(contentType, "application/xml"):
		if err := xml.NewDecoder(r.Body).Decode(i); err != nil {
			return err
		}
	default:
		return errors.New("unsupported content type")
	}
	return nil
}