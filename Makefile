# 定义变量
APP_NAME=vsp
SRC=./cmd/main.go
BUILD_DIR=./build
DIST_DIR=./dist

# 编译应用程序
build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC)

# 运行应用程序
run:
	$(BUILD_DIR)/$(APP_NAME)

dev:
	go run $(SRC)

# 清理构建目录
clean:
	rm -rf $(BUILD_DIR)/*

# 打包应用程序
dist:
	mkdir -p $(DIST_DIR)
	tar -czvf $(DIST_DIR)/$(APP_NAME).tar.gz $(BUILD_DIR)/$(APP_NAME)
