package theme

import (
	"encoding/xml"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"k8s.io/klog"
)

// JcinkXML is an XML file used by jcink
type JcinkXML struct {
	XMLName    xml.Name `xml:"skin"`
	Text       string   `xml:",chardata"`
	Skinname   string   `xml:"skinname"`
	Date       string   `xml:"date"`
	Stylesheet string   `xml:"stylesheet"`
	Wrappers   string   `xml:"wrappers"`

	Macros struct {
		Text string `xml:",chardata"`
		Item []Item `xml:"item"`
	} `xml:"macros"`

	Templates struct {
		Item []Item `xml:"item"`
	} `xml:"templates"`
}

// Item is a single macro or template.
type Item struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

// Construct returns a []byte of a skin XML file
func Construct(name string, stylesheet string, wrapper string, macroFolder string, templateFolder string) (*string, error) {
	timestamp := time.Now().Format("Monday 2 of Jan 2006 15:04:05 PM")

	output := &JcinkXML{
		Skinname: name,
		Date:     timestamp,
	}

	// Stylesheet
	stylesheetString, err := getFileString(stylesheet)
	if err != nil {
		return nil, err
	}
	// .post1 {} is required to pass xml upload validation in JCINK
	output.Stylesheet = fmt.Sprintf("\n%s\n%s\n", stylesheetString, ".post1 {} /* this is required to pass JCINK upload validation */")

	// Wrapper
	wrapperString, err := getFileString(wrapper)
	if err != nil {
		return nil, err
	}
	output.Wrappers = wrapperString

	// Macros
	macroFiles, err := listDirectory(macroFolder)
	if err != nil {
		return nil, err
	}
	//TODO: Macros requires all of them to be present
	// We are going to need to start with a default list and then do overrides
	macros, err := buildItems(macroFiles)
	if err != nil {
		return nil, err
	}
	output.Macros.Item = macros

	// Templates
	templateFiles, err := listDirectory(templateFolder)
	if err != nil {
		return nil, err
	}
	templates, err := buildItems(templateFiles)
	if err != nil {
		return nil, err
	}
	output.Templates.Item = templates

	data, err := xml.MarshalIndent(output, "	", "	")
	out := fmt.Sprintf("%s%s", xml.Header, html.UnescapeString(string(data)))
	return &out, nil
}

// buildItems returns a list of items from a list of files
func buildItems(files []string) ([]Item, error) {
	items := []Item{}
	for _, file := range files {
		fileString, err := getFileString(file)
		if err != nil {
			klog.Error(err)
			continue
		}

		item := Item{
			Name: getPlainFilename(file),
			Text: fileString,
		}
		items = append(items, item)
	}
	return items, nil
}

func getPlainFilename(fullname string) string {
	// strip the extension
	_, filename := filepath.Split(fullname)
	extension := filepath.Ext(filename)
	filename = filename[0 : len(filename)-len(extension)]
	return filename
}

// GetFileString reads a file into a string
func getFileString(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// listDirectory lists all of the files in a directory
func listDirectory(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

// Read reads a skin file in
func Read(filename string) error {
	skin := &JcinkXML{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, skin)
	if err != nil {
		return err
	}

	for _, template := range skin.Templates.Item {
		fmt.Println(template.Name)
	}

	return nil
}
