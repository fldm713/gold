package gold

type Validator interface {
	Validate(i any) error
}