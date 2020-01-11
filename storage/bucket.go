package storage


type Bucket interface {
	Delete() 			error
	Files()  			Object
	Create() 			Object
	Exist()  			(bool, error)
}

type Buckets struct {
	BucketName string
	Endpoint   string
}