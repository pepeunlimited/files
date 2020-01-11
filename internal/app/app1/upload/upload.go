package upload

type MetaAPIArgs struct {
	Filename 		string `json:"filename"`
}

type UploadHeaders struct {
	ContentType 	string
	ContentLength	int64
	Authorization   string
}