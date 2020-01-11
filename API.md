## cURL

#### GetFileByID
```
$ curl -X POST \
-H "Content-Type: application/json" \
"api.dev.pepeunlimited.com/twirp/pepeunlimited.files.DOFileService/GetFile" \
-d '{"file_id": 58}'
```
#### GetFileByFilename
```
$ curl -X POST \
-H "Content-Type: application/json" \
"api.dev.pepeunlimited.com/twirp/pepeunlimited.files.DOFileService/GetFile" \
-d '{"file_id": 58}' \
-d '{"filename": { "name": "filename.txt", "bucket_name": "bucketName", "bucket_id": 1 }}'
```
#### Delete
```
$ curl -X POST \
-H "Content-Type: application/json" \
"api.dev.pepeunlimited.com/twirp/pepeunlimited.files.DOFileService/Delete" \
-d '{"file_id": 58, "is_permanent": false}' \
-d '{"filename": { "name": "filename.txt", "bucket_name": "bucketName", "bucket_id": 1 }, "is_permanent": false}'
```
#### CreateBucket
```
$ curl -X POST \
-H "Content-Type: application/json" \
"api.dev.pepeunlimited.com/twirp/pepeunlimited.files.DOFileService/CreateBucket" \
-d '{"bucket_name": "test0r-666", "endpoint": "fra1.digitaloceanspaces.com"}'
```
#### Content-Length
```
stat -f%z ${filename}
```

```
curl -X POST \
-H "Authorization: Bearer aa" \
-H "Content-Type: plain/text" \
-H "Content-Length: 50" \
-H "Meta-API-Args: {\"filename\": \"trolli.txt\"}" \
localhost:8080/upload/do/v1/files \
--data-binary @const.go
```
#### UploadFile
```
curl -X POST \
-H "Authorization: Bearer aa" \
-H "Content-Type: plain/text" \
-H "Content-Length: 50" \
-H "Meta-API-Args: {\"filename\": \"trolli.txt\"}" \
api.dev.pepeunlimited.com/upload/do/v1/files \
--data-binary @const.go
```