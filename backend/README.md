# Testing

`$ go clean -testcache`

unit testing command

`$ go test -v -cover ./...`

integrate testing command

`$ go test -v -cover -tags=integrate ./...`

# gRPC file gen

`$ cd /backend`

`$ protoc -I ./proto/ --go_out=plugins=grpc:./proto-gen/todo --go_opt=module=github.com/t-ash0410/tdd-sample/backend/proto-gen/todo ./proto/todo.proto`

# Build image

`$ gcloud builds submit --config todo.cloudbuild.yml`

# Deploy image to Cloud Run

`$ gcloud run deploy tdd-sample-rest-todo --image us-central1-docker.pkg.dev/tdd-sample-327207/tdd-sample-repository/tdd-sample-rest-todo-image:tag1`
