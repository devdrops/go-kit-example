package main

type StringService interface {
	Reverse(string) (string, error)
}

type stringSvc struct{}

type (
	reverseReq struct {
		S string `json:"s"`
	}

	reverseRes struct {
		V   string `json:"v"`
		Err string `json:"err,omitempty"`
	}
)

func (svc stringSvc) Reverse(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}

	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r), nil
}
