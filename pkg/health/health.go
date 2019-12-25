package health

// Healther ...
type Healther interface {
	Check() bool
}

type health struct {
}

// Check - check if health is good
func (h *health) Check() bool {
	return true
}

// NewHealther ...
func NewHealther() Healther {
	return &health{}
}
