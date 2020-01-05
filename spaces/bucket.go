package spaces

type Bucket interface {
	Delete(bucketName string) error
	Files(bucketName string)  Files
	Create(bucketName string) Files
	Exist(bucketName string)  (bool, error)
}

type bucket struct {
	name 		string
	isCreate 	bool
	isFiles    	bool
	spaces 		spaces
}

func (b *bucket) Exist(bucketName string) (bool, error) {
	return b.spaces.client.BucketExists(bucketName)
}

func (b *bucket) Delete(bucketName string) error {
	return b.spaces.client.RemoveBucket(bucketName)
}

func (b *bucket) Files(bucketName string) Files {
	executor := make(map[int]interface{})
	b.isFiles = true
	b.isCreate = false
	b.name = bucketName
	return &files{bucket:*b, executor:executor, order: 0}
}

func (b *bucket) Create(bucketName string) Files {
	executor := make(map[int]interface{})
	b.isCreate = true
	b.isFiles = false
	b.name = bucketName
	return &files{bucket:*b, executor:executor, order: 0}
}

