protoc -I=. --go_out=./generated model.proto
protoc -I=. --go_out=./generated template.proto
protoc -I=. --go_out=./generated event-incoming.proto
protoc -I=. --go_out=./generated event-outgoing.proto
