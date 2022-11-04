GONAME=actions-build-time

darwin:
	@echo "Building $(GONAME)-darwin_x86_64 binary"
	GOOS=darwin GOARCH=amd64 go build -o $(GONAME)-darwin_x86_64 main.go
	@echo "Compressing $(GONAME)-darwin_x86_64 binary..."
	tar -czvf $(GONAME)-darwin_x86_64.tar.gz ./$(GONAME)-darwin_x86_64
	@echo " "

darwin_arm:
	@echo "Building $(GONAME)-darwin_arm64 binary..."
	GOOS=darwin GOARCH=arm64 go build -o $(GONAME)-darwin_arm64 main.go
	@echo "Compressing $(GONAME)-darwin_arm64 binary..."
	tar -czvf $(GONAME)-darwin_arm64.tar.gz ./$(GONAME)-darwin_arm64
	@echo " "

linux:
	@echo "Building $(GONAME)-linux_x86_64 binary..."
	GOOS=linux GOARCH=amd64 go build -o $(GONAME)-linux_x86_64 main.go
	@echo "Compressing $(GONAME)-linux_x86_64 binary"
	tar -czvf $(GONAME)-linux_x86_64.tar.gz ./$(GONAME)-linux_x86_64
	@echo " "

linux_arm:
	@echo "Building binary for linux..."
	GOOS=linux GOARCH=arm64 go build -o $(GONAME)-linux_arm64 main.go
	@echo "Compressing $(GONAME)-linux_arm64 binary..."
	tar -czvf $(GONAME)-linux_arm64.tar.gz ./$(GONAME)-linux_arm64
	@echo " "

all: linux linux_arm darwin darwin_arm

clean:
	@echo "Removing all ${GONAME} binaries and tar files from current directory..."
	rm ${GONAME}*
	@echo "Done!"
	@echo " "

.PHONY: linux linux_arm darwin darwin_arm clean all