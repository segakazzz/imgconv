package imgconv

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Convert(dirname string) (e error){
	files, e := getSourceFiles(dirname)
	if e != nil {
		return
	}
	e = convertFiles(files, dirname)
	if e != nil {
		return
	}
	return nil
}

func getSourceFiles(dirName string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		return []os.FileInfo{}, err
	}
	return files, nil
}

func convertFiles(files []os.FileInfo, dirName string) (e error){
	re, e := regexp.Compile(".jpg$")
	if e != nil{
		return
	}
	for _, file := range files {
		if re.MatchString(file.Name()) {
			e = jpg2png(dirName, file.Name())
			if e != nil {
				return
			}
		}
	}
	return nil
}

func jpg2png(dirname string, filename string) (e error) {
	input := filepath.Join(dirname, filename)
	outDir := filepath.Join(dirname, "out")
	output := filepath.Join(outDir, strings.Replace(strings.ToLower(filename), ".jpg", ".png", -1 ))
	fmt.Println(output)
	if !dirExists(outDir) {
		os.Mkdir(outDir, 0755)
	}

	in, _ := os.Open(input)
	var out *os.File
	if fileExists(output){
		out, e = os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if e != nil {
			return
		}
	} else {
		out, e = os.Create(output)
		if e != nil {
			return
		}
	}
	defer in.Close()
	defer out.Close()
	img, e := jpeg.Decode(in)
	if e !=nil {
		return
	}
	e = png.Encode(out, img)
	if e != nil {
		return
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func dirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}