package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type File struct {
	Path string
	Name string
}

// list all files in a slice
func ListFiles(path string) ([]File, error) {
	files := []File{}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return files, err
	}
	if fileInfo.IsDir() {
		fileInfos, err := ioutil.ReadDir(path)
		if err != nil {
			return files, err
		}
		for _, fileInfo := range fileInfos {
			files = append(files, File{path, fileInfo.Name()})
		}
	} else {
		files = append(files, File{path, fileInfo.Name()})
	}
	return files, nil
}

// rename a file
func rename(from, to string) error {
	return os.Rename(from, to)
}

func RenameFolderObjects(from string, to string) error {
	files, err := ListFiles(from)
	if err != nil {
		return err
	}
	for _, file := range files {
		from := file.Path + "\\" + file.Name
		println("Now rename: " + from + "? Y/n")
		var (
			flag  bool
			input string
		)
		fmt.Scanln(&input)
		flag = input == "Y" || input == "y"
		if flag {
			// cmd open picture
			err := exec.Command("cmd", "/c", from).Run()
			if err != nil {
				return err
			}
			fmt.Scanln(&input)
			target := to + "\\" + input + ".jpg"
			// check if file exists
			_, err = os.Stat(target)
			if err == nil {
				println("File exists. Overwrite? Y/n")
				fmt.Scanln(&input)
				flag = input == "Y" || input == "y"
				if flag {
					// do nothing
				} else {
					println("Input new name: ")
					fmt.Scanln(&input)
					target = to + "\\" + input + ".jpg"
				}
			}
			err = rename(from, target)
			if err != nil {
				return err
			}
			println(from + " -> " + target)
		} else {
			// move the file without rename
			err := os.Rename(from, to+"\\"+file.Name)
			if err != nil {
				return err
			}
			println(from + " -> " + to + "\\" + file.Name)
		}
	}
	return nil
}
