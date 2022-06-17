protoc -I=. --go_out=../scyna scyna.proto
protoc -I=. --go_out=../scyna error.proto
protoc -I=. --go_out=../scyna engine.proto
protoc -I=. --go_out=../scyna scheduler.proto

