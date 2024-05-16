package passport

type E interface {
	Err() error
}

func Validate(errs ...E) error {
	for _, item := range errs {
		if err := item.Err(); err != nil {
			return err
		}
	}
	return nil
}

func ValidateErrors(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
