package main

import "fmt"

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, node := range f.children {
		node.print(indentation + indentation)
	}
}
func (f *folder) clone() inode {
	cloneFolder := &folder{
		name: f.name + "_clone",
	}
	var copyChildren []inode
	for _, value := range f.children {
		copyValue := value.clone()
		copyChildren = append(copyChildren, copyValue)
	}
	cloneFolder.children = copyChildren
	return cloneFolder
}
