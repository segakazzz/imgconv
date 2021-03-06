package imgconv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type converter struct {
	dirname string
	input string
	output string
}


// NewConverter generates new converter and return it's pointer
func NewConverter(dirname string, input string, output string) (*converter, error){
	switch input {
	case "jpg", "png":
		input = strings.ToLower(input)
	default:
		return &converter{}, fmt.Errorf("Input extension is not valid. Select one from jpg/png")
	}
	switch output {
	case "jpg", "png":
		output = strings.ToLower(output)
	default:
		return &converter{}, fmt.Errorf("Output extension is not valid. Select one from jpg/png")
	}

	if input == output {
		return &converter{}, fmt.Errorf("Input and Output extensiton is the same. No convertion is needed")
	}
	return &converter{dirname: dirname, input:input, output:output}, nil
}

// Convert method is convert all jpg files in dirname to png. "out" folder is generated if it doesn't exist.
func (c *converter)Convert()(e error){
	files, e := c.getSourceFiles()
	if e != nil {
		return
	}
	e = c.convertFiles(files)
	if e != nil {
		return
	}
	return nil
}

func (c *converter)getSourceFiles() ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(c.dirname)
	if err != nil {
		return []os.FileInfo{}, err
	}
	return files, nil
}

func (c *converter)convertFiles(files []os.FileInfo) (e error){
	re, e := regexp.Compile("." + c.input + "$")
	if e != nil{
		return
	}
	for _, file := range files {
		if re.MatchString(file.Name()) {
			e = c.convertSingle(file.Name())
			if e != nil {
				return
			}
		}
	}
	return nil
}

func (c *converter) convertSingle(filename string) (e error) {
	input := filepath.Join(c.dirname, filename)
	outDir := filepath.Join(c.dirname, "out")
	output := filepath.Join(outDir, strings.Replace(strings.ToLower(filename), "." + c.input, "." + c.output, -1 ))
	fmt.Println(output)
	if !c.dirExists(outDir) {
		os.Mkdir(outDir, 0755)
	}

	in, _ := os.Open(input)
	var out *os.File
	if c.fileExists(output){
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

	var (
		img image.Image
	)
	switch c.input{
	case "jpg":
		img, e = jpeg.Decode(in)
	case "png":
		img, e=png.Decode(in)
	}

	if e !=nil {
		return
	}
	switch c.output {
	case "png":
		e = png.Encode(out, img)
	case "jpg":
		e = jpeg.Encode(out, img, nil)
	}
	if e != nil {
		return
	}
	return nil
}

func (c *converter)fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (c *converter)dirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}