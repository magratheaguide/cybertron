package theme

import (
	"encoding/xml"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"k8s.io/klog"
)

// JcinkXML is an XML file used by jcink
type JcinkXML struct {
	XMLName    xml.Name `xml:"skin"`
	Text       string   `xml:",chardata"`
	ThemeName  string   `xml:"skinname"`
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

// Construct returns a []byte of a theme XML file
func Construct(name string, stylesheet string, wrapper string, macroFolder string, templateFolder string) (*string, error) {
	timestamp := time.Now().Format("Monday 2 of Jan 2006 15:04:05 PM")

	output := &JcinkXML{
		ThemeName: name,
		Date:      timestamp,
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
	macroOverrides, err := buildItems(macroFiles)
	if err != nil {
		return nil, err
	}
	macros := combineMacros(macroOverrides)
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
	if err != nil {
		return nil, err
	}

	fixQuotes := regexp.MustCompile(`<item name=("(.+)")>`)
	data = fixQuotes.ReplaceAll(data, []byte(`<item name='$2'>`))

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

// combineMacros returns the default macros with any overrides
// works by taking the list of overrides, then adding all defaults that do not exist in overrides
func combineMacros(overrides []Item) []Item {
	ret := overrides

	for _, defaultMacro := range defaultMacros {
		if !itemListContains(overrides, defaultMacro) {
			ret = append(ret, defaultMacro)
			klog.V(2).Infof("adding default macro that has no override: %s", defaultMacro.Name)
		}
	}

	return ret
}

func itemListContains(list []Item, item Item) bool {
	for _, i := range list {
		if item.Name == i.Name {
			return true
		}
	}
	return false
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
		if info == nil {
			klog.Warningf("%s dir is empty", dir)
			return nil
		}
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

// Read reads a theme file in
func Read(filename string) error {
	theme := &JcinkXML{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, theme)
	if err != nil {
		return err
	}

	for _, template := range theme.Templates.Item {
		fmt.Println(template.Name)
	}

	for _, macro := range theme.Macros.Item {
		fmt.Printf("{Name: %s, Text: %s},\n", macro.Name, macro.Text)
	}

	return nil
}
