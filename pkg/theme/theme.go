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

var defaultMacros = []Item{
	{Name: "A_LOCKED_B", Text: "<img src='style_images/2/t_closed.gif' border='0' alt='Closed Topic' />"},
	{Name: "A_MOVED_B", Text: "<img src='style_images/2/t_moved.gif' border='0' alt='Moved Topic'/>"},
	{Name: "A_POLLONLY_B", Text: "<img src='style_images/2/t_closed.gif' border='0' alt='Poll Only' />"},
	{Name: "A_POST", Text: "<img src='style_images/2/t_new.gif' border='0' alt='Start new topic' />"},
	{Name: "A_REPLY", Text: "<img src='style_images/2/t_reply.gif' border='0' alt='Reply to this topic' />"},
	{Name: "A_POLL", Text: "<img src='style_images/2/t_poll.gif' border='0' alt='Start Poll' />"},
	{Name: "A_STAR", Text: "<img src='style_images/2/pip.gif' border='0' alt='*' />"},
	{Name: "B_HOT", Text: "<img src='style_images/2/f_hot.gif' border='0' alt='Hot topic' />"},
	{Name: "B_HOT_NN", Text: "<img src='style_images/2/f_hot_no.gif' border='0' alt='No new' />"},
	{Name: "B_LOCKED", Text: "<img src='style_images/2/f_closed.gif' border='0' alt='Closed' />"},
	{Name: "B_MOVED", Text: "<img src='style_images/2/f_moved.gif' border='0' alt='Moved' />"},
	{Name: "B_NEW", Text: "<img src='style_images/2/f_norm.gif' border='0' alt='New Posts' />"},
	{Name: "B_NORM", Text: "<img src='style_images/2/f_norm_no.gif' border='0' alt='No New Posts' />"},
	{Name: "B_PIN", Text: "<img src='style_images/2/f_pinned.gif' border='0' alt='Pinned' />"},
	{Name: "B_POLL", Text: "<img src='style_images/2/f_poll.gif' border='0' alt='Poll' />"},
	{Name: "B_POLL_NN", Text: "<img src='style_images/2/f_poll_no.gif' border='0' alt='No new votes' />"},
	{Name: "B_HOT_DOT", Text: "<img src='style_images/2/f_hot_dot.gif' border='0' alt='New Posts' />"},
	{Name: "B_NEW_DOT", Text: "<img src='style_images/2/f_norm_dot.gif' border='0' alt='No New Posts' />"},
	{Name: "B_HOT_NN_DOT", Text: "<img src='style_images/2/f_hot_no_dot.gif' border='0' alt='No New Posts*' />"},
	{Name: "B_NORM_DOT", Text: "<img src='style_images/2/f_norm_no_dot.gif' border='0' alt='No New Posts*' />"},
	{Name: "B_POLL_DOT", Text: "<img src='style_images/2/f_poll_dot.gif' border='0' alt='Poll*' />"},
	{Name: "B_POLL_NN_DOT", Text: "<img src='style_images/2/f_poll_no_dot.gif' border='0' alt='No New Votes*' />"},
	{Name: "C_LOCKED", Text: "<img src='style_images/2/bf_readonly.gif' border='0' alt='Read Only Forum' />"},
	{Name: "C_OFF", Text: "<img src='style_images/2/bf_nonew.gif' border='0' alt='No New Posts' />"},
	{Name: "C_OFF_CAT", Text: "<img src='style_images/2/bc_nonew.gif' border='0' alt='No New Posts' />"},
	{Name: "C_OFF_RES", Text: "<img src='style_images/2/br_nonew.gif' border='0' alt='No New Posts' />"},
	{Name: "C_ON", Text: "<img src='style_images/2/bf_new.gif' border='0' alt='New Posts' />"},
	{Name: "C_ON_CAT", Text: "<img src='style_images/2/bc_new.gif' border='0' alt='New Posts' />"},
	{Name: "C_ON_RES", Text: "<img src='style_images/2/br_new.gif' border='0' alt='New Posts' />"},
	{Name: "F_ACTIVE", Text: "<img src='style_images/2/user.gif' border='0' alt='Active Users' />"},
	{Name: "F_NAV_SEP", Text: ""},
	{Name: "F_NAV", Text: "<img src='style_images/2/nav.gif' border='0' alt='&gt;' />"},
	{Name: "F_STATS", Text: "<img src='style_images/2/stats.gif' border='0' alt='Board Stats' />"},
	{Name: "NO_PHOTO", Text: "<img src='style_images/2/nophoto.gif' border='0' alt='No Photo' />"},
	{Name: "CAMERA", Text: "<img src='style_images/2/camera.gif' border='0' alt='Photo' />"},
	{Name: "M_READ", Text: "<img src='style_images/2/f_norm_no.gif' border='0' alt='Read Msg' />"},
	{Name: "M_UNREAD", Text: "<img src='style_images/2/f_norm.gif' border='0' alt='Unread Msg' />"},
	{Name: "P_AOL", Text: "<img src='style_images/2/p_aim.gif' border='0' alt='AOL' />"},
	{Name: "P_DELETE", Text: "<img src='style_images/2/p_delete.gif' border='0' alt='Delete Post' />"},
	{Name: "P_EDIT", Text: "<img src='style_images/2/p_edit.gif' border='0' alt='Edit Post' />"},
	{Name: "P_EMAIL", Text: "<img src='style_images/2/p_email.gif' border='0' alt='Email Poster' />"},
	{Name: "P_ICQ", Text: "<img src='style_images/2/p_icq.gif' border='0' alt='ICQ' />"},
	{Name: "P_MSG", Text: "<img src='style_images/2/p_pm.gif' border='0' alt='PM' />"},
	{Name: "P_QUOTE", Text: "<img src='style_images/2/p_quote.gif' border='0' alt='Quote Post' />"},
	{Name: "P_WEBSITE", Text: "<img src='style_images/2/p_www.gif' border='0' alt='Users Website' />"},
	{Name: "P_YIM", Text: "<img src='style_images/2/p_yim.gif' border='0' alt='Yahoo' />"},
	{Name: "P_REPORT", Text: "<img src='style_images/2/p_report.gif' border='0' alt='Report Post' />"},
	{Name: "P_MSN", Text: "<img src='style_images/2/p_msn.gif' border='0' alt='MSN' />"},
	{Name: "CAT_IMG", Text: ""},
	{Name: "NEW_POST", Text: "<img src='style_images/2/newpost.gif' border='0' alt='Goto last unread' title='Goto last unread' hspace='2' />"},
	{Name: "LAST_POST", Text: "<img src='style_images/2/lastpost.gif' border='0' alt='Last Post' />"},
	{Name: "BR_REDIRECT", Text: "<img src='style_images/2/br_redirect.gif' border='0' alt='Redirect' />"},
	{Name: "INTEGRITY_MSGR", Text: "<img src='style_images/2/p_im.gif' border='0' alt='Integrity Messenger IM' />"},
	{Name: "ADDRESS_CARD", Text: "<img src='style_images/2/addresscard.gif' border='0' alt='Mini Profile' />"},
	{Name: "T_QREPLY", Text: "<img src='style_images/2/t_qr.gif' border='0' alt='Fast Reply' />"},
	{Name: "T_OPTS", Text: "<img src='style_images/2/t_options.gif' border='0' alt='Topic Options' />"},
	{Name: "CAL_NEWEVENT", Text: "<img src='style_images/2/cal_newevent.gif' border='0' alt='Add New Event' />"},
	{Name: "F_RULES", Text: "<img src='style_images/2/forum_rules.gif' border='0' alt='Forum Rules' />"},
	{Name: "WARN_0", Text: "<img src='style_images/2/warn0.gif' border='0' alt='-----' />"},
	{Name: "WARN_1", Text: "<img src='style_images/2/warn1.gif' border='0' alt='X----' />"},
	{Name: "WARN_2", Text: "<img src='style_images/2/warn2.gif' border='0' alt='XX---' />"},
	{Name: "WARN_3", Text: "<img src='style_images/2/warn3.gif' border='0' alt='XXX--' />"},
	{Name: "WARN_4", Text: "<img src='style_images/2/warn4.gif' border='0' alt='XXXX-' />"},
	{Name: "WARN_5", Text: "<img src='style_images/2/warn5.gif' border='0' alt='XXXXX' />"},
	{Name: "WARN_ADD", Text: "<img src='style_images/2/warn_add.gif' border='0' alt='+' />"},
	{Name: "WARN_MINUS", Text: "<img src='style_images/2/warn_minus.gif' border='0' alt='-' />"},
	{Name: "LOGO", Text: "<img src='style_images/2/ipbalt-1.gif' />"},
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
	//TODO: Macros requires all of them to be present
	// We are going to need to start with a default list and then do overrides
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
