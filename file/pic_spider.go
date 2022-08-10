package file

import (
	"errors"
	"os"
	"regexp"
)

func filter_pic(path string) bool {
	// filter .jpg .jpeg .png .git .bmp ... using regexp
	regex := regexp.MustCompile(`\.(jpg|jpeg|png|gif|bmp)$`)
	// check if path meets the regex
	return regex.MatchString(path)
}

// go through every dir and sub-dir in the input dir
func Spider(path string, output string) error {
	// open the dir
	dir, err := os.Open(path)
	if err != nil {
		return errors.New("open dir error " + err.Error() + " at:" + path)
	}
	// get file info
	fileInfo, err := dir.Readdir(-1)
	if err != nil {
		return errors.New("read dir error " + err.Error() + " at:" + path)
	}
	//print pwd
	println(path)
	// go through every file in the dir
	for _, file := range fileInfo {
		// if it is a dir, go through it
		if file.IsDir() {
			Spider(path+"/"+file.Name(), output)
		} else {
			// println("Judging " + path + "/" + file.Name())
			// if it is a pic, scale it
			if filter_pic(path + "/" + file.Name()) {
				err := scale(path+"/"+file.Name(), output+"/"+file.Name())
				println(path + "/" + file.Name() + " -> " + output + "/" + file.Name())
				if err != nil {
					return errors.New("scale error " + err.Error() + " at:" + "/" + file.Name() + "to: " + output + "/" + file.Name())
				}
			}
		}
	}
	return nil
}

func scale(path string, target string) error {
	// mv the pic to the output dir
	return os.Rename(path, target)
}
