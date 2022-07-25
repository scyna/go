protoc -I=. --go_out=./generated organization.proto
protoc -I=. --go_out=./generated module.proto
protoc -I=. --go_out=./generated service.proto
protoc -I=. --go_out=./generated client.proto
protoc -I=. --go_out=./generated sync.proto
protoc -I=. --go_out=./generated event.proto
protoc -I=. --go_out=./generated application.proto
protoc -I=. --go_out=./generated task.proto
