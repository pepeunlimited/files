package storage


type Actions interface {
	Delete(bucket Buckets, filename string) error
	Upload(file File, meta FileMetaData, buckets Buckets) error
	CreateBucket(bucket Buckets) error
	DeleteBucket(bucket Buckets) error
}