package validator


type ImageUploadServerValidator struct {}


func NewImageUploadServerValidator() ImageUploadServerValidator {
	return ImageUploadServerValidator{}
}

func (ImageUploadServerValidator) UploadImage() error {
	return nil
}