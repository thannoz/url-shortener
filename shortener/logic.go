package shortener

import (
	"errors"
	"time"

	"github.com/teris-io/shortid"
	"go.step.sm/cli-utils/errs"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type redirectService struct {
	redirectRepo RedirectRepository
}

// Find calls the repository and returns its find method
func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

// Store validates the redirect struct.
// Generates a short code and time and call the store method of RedirectRepository
func (r *redirectService) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}

	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}

func NewRedirectService(redirectRepo RedirectRepository) RedirectService {
	return &redirectService{redirectRepo}
}
