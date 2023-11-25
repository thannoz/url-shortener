package json

import (
	"encoding/json"
	"hex/shortener"

	"github.com/pkg/errors"
)

type Redirect struct{}

// Decode transforms json into our custom (Redirect) object.
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}

	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode transforms our custom (Redirect) object into json.
func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
