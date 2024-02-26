.PHONY: dev build proto build-osx build-client

dev:
	@echo "Starting development environment..."
	@{ \
	trap 'trap - SIGTERM && kill -- -$$' SIGINT SIGTERM EXIT; \
	cd client && yarn dev > /dev/null & \
	cd server && go run . --devmode & \
	wait; \
	}

build-client:
	@echo "Building client..."
	@cd client && yarn && yarn build

build-server:
	@echo "Building Go server for production..."
	@cd server && go build -o ../myapp -tags !dev,!app .
	@echo "Build complete. Execute ./myapp to run the server."

build:
	@echo "Building client..."
	@cd client && yarn && yarn build
	@echo "Building Go server for production..."
	@cd server && go build -o ../myapp -tags !dev,!app .
	@echo "Build complete. Execute ./myapp to run the server."

build-osx:
	@echo "Building client..."
	@cd client && yarn && yarn build
	@echo "Building OSX Shell..."
	@echo "Building Go server for production..."
	@cd server && go build -o ../myapp -tags !dev,!app .
	@cd ./containers/osx/SwiftFrame && xcodebuild -project SwiftFrame.xcodeproj -configuration Debug
	@echo "Build complete. Execute ./myapp to run the server."

proto-gen:
	@echo "Generating Go code from .proto files..."
	@protoc --go_out=paths=source_relative:./server --go_opt=paths=source_relative ./proto/*.proto
	@protoc --plugin=client/node_modules/ts-proto/protoc-gen-ts_proto --ts_proto_out=./client/src ./proto/*.proto --ts_proto_opt=esModuleInterop=true
	@echo "Proto compilation complete."