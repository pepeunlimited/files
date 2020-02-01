package upload


type UploadDOV1Files struct {
	Endpoint  		string `json:"endpoint"`
	CDNEndpoint 	string `json:"cdn_endpoint"`
	FileId	  		int64  `json:"file_id"`
	Filename  		string `json:"filename"`
	URI       		string `json:"uri"`
}

type MetaAPIArgs struct {
	Filename 		string `json:"filename"`
}

type UploadHeaders struct {
	ContentType 	string
	ContentLength	int64
	Authorization   string
}