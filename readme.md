# 功能

	对给定的文件计算各种散列值（MD5、SHA1、SHA256），主要用于检验网络下载文件的完整性。

# 构建

	go build checksums.go

# 用法

	checksums [-p] filename

	-p       : 完成后等待用户输入回车键再退出程序。
	filename : 需要计算散列值的目标文件。
