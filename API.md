## cURL

#### GetFileByID
```
$ curl -X POST \
-H "Content-Type: application/json" \
"api.dev.pepeunlimited.com/twirp/pepeunlimited.files.FilesService/GetFile" \
-d '{"file_id": 58}'
```
#### GetFileByFilename
```
$ curl -X POST \
-H "Content-Type: application/json" \
"api.dev.pepeunlimited.com/twirp/pepeunlimited.files.FilesService/GetFile" \
-d '{"file_id": 58}' \
-d '{"filename": { "name": "filename.txt", "bucket_name": "bucketName", "bucket_id": 1 }}'
```
#### DeleteByFileID
```
$ curl -X POST \
-H "Content-Type: application/json" \
"api.dev.pepeunlimited.com/twirp/pepeunlimited.files.FilesService/Delete" \
-d '{"file_id": 58, "is_permanent": false}'
```
#### DeleteByFilename
```
$ curl -X POST \
-H "Content-Type: application/json" \
"api.dev.pepeunlimited.com/twirp/pepeunlimited.files.FilesService/Delete" \
-d '{"filename": { "name": "filename.txt", "bucket_name": "bucketName", "bucket_id": 1 }, "is_permanent": false}'
```
#### CreateBucket
```
$ curl -X POST \
-H "Content-Type: application/json" \
"api.dev.pepeunlimited.com/twirp/pepeunlimited.files.FilesService/CreateBucket" \
-d '{"name": "test0r-666"}'
```
#### Content-Length
```
stat -f%z ${filename}
```
#### UploadFile
```
curl -X POST \
-H "Authorization: Bearer ${TOKEN}" \
-H "Content-Type: plain/text" \
-H "Content-Length: 50" \
-H "Meta-API-Args: {\"filename\": \"trolli.txt\"}" \
api.dev.pepeunlimited.com/upload/v1/files \
--data-binary @const.go
```