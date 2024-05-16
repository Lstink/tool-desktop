## 协议解析工具


### 使用方式


```bash
# 1. 拉取项目
git clone https://github.com/Lstink/tool-desktop.git
# 2. 进入项目目录
cd tool-desktop
# 3. 同步依赖
go mod tidy

```

### 启动本地项目

```bash
go run main.go
```

### 打包项目客户端
```bash
# 1. 安装 `fyne` 打包工具
go install fyne.io/fyne/v2/cmd/fyne@latest
# 2. 运行打包命令，这里示例打包 macOS 版本
fyne package -os darwin -icon icon.png
```

* windows打包命令：`fyne package -os windows -icon icon.png`
* macOS打包命令：`fyne package -os darwin -icon icon.png`

> 注意windows打包需要在windows环境下执行，且需要windows安装gcc环境

