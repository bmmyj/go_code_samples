package main

import (
	"os"
	"log"
	"fmt"
)
// 背景知识：
// Linux链接分两种，一种被称为硬链接（Hard Link），另一种被称为符号链接（Symbolic Link）。默认情况下，ln命令产生硬链接。
//--硬连接 硬连接指通过索引节点来进行连接。在Linux的文件系统中，保存在磁盘分区中的文件不管是什么类型都给它分配一个编号，称为
// 索引节点号(Inode Index)。在Linux中，多个文件名指向同一索引节点是存在的。一般这种连接就是硬连接。硬连接的作用是允许一个文件
// 拥有多个有效路径名，这样用户 就可以建立硬连接到重要文件，以防止“误删”的功能。其原因如上所述，因为对应该目录的索引节点有一个以上
// 的连接。只删除一个连接并不影响索引节点本身和 其它的连接，只有当最后一个连接被删除后，文件的数据块及目录的连接才会被释放。
// 也就是说，文件真正删除的条件是与之相关的所有硬连接文件均被删除。

//【软连接】
// 另外一种连接称之为符号连接（Symbolic Link），也叫软连接。软链接文件有类似于Windows的快捷方式。它实际上是一个特殊的文件。
// 在符号连接中，文件实际上是一个文本文件，其中包含的有另一文件的位置信息。
func main() {
	// Create a hard link
	// You will have two file names that point to the same contents
	// Changing the contents of one will change the other
	// Deleting/renaming one will not affect the other
	err := os.Link("original.txt", "original_also.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("creating sym")
	// Create a symlink
	err = os.Symlink("original.txt", "original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Lstat will return file info, but if it is actually
	// a symlink, it will return info about the symlink.
	// It will not follow the link and give information
	// about the real file
	// Symlinks do not work in Windows
	fileInfo, err := os.Lstat("original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Link info: %+v", fileInfo)

	// Change ownership of a symlink only
	// and not the file it points to
	err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}