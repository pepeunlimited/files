package validator


type ImageServerValidator struct {}


func NewImageServerValidator() ImageServerValidator {
	return ImageServerValidator{}
}

func (ImageServerValidator) CreateImage() error {
	return nil
}
