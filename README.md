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

````bash
go run main.go```  

### 打包项目客户端  
```bash  
# 1. 安装 `fyne` 打包工具  
go install fyne.io/fyne/v2/cmd/fyne@latest  
# 2. 运行打包命令，这里示例打包 macOS 版本  
fyne package -os darwin -icon icon.png  
````

- windows打包命令：`fyne package -os windows -icon icon.png`
- macOS打包命令：`fyne package -os darwin -icon icon.png`

## windows下打包方法

1. 安装 MSYS2  
   [ MSYS2下载地址](https://github.com/msys2/msys2-installer/releases/download/2023-07-18/msys2-x86_64-20230718.exe)
2. 启动UCRT64环境（在windows上安装程序中找）
3. 在弹出的命令行中运行以下命令(所有的项目都选y)

  ```bash
  pacman -Syu
  pacman -S git mingw-w64-x86_64-toolchain
  pacman -S mingw-w64-ucrt-x86_64-gcc
  gcc --version
  ```

4. 设置ucrt64下bin目录到系统环境变量中
5. 在项目根目录下运行打包命令：`fyne package -os windows -icon icon.png`