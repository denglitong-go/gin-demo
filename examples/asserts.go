package examples

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets950a25363eee220f7d8ce234bcc0b349e4ea9072 = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Index Template</title>\n</head>\n<body>\n    <p>Hello, {{.Foo}}</p>\n</body>\n</html>"
var _Assets8e1edddb536a70fc4f36d5f948de1bd51ea8571b = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <title>Bar Template</title>\n</head>\n<body>\n  <p>Can you see this? - {{.Bar}}</p>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"index.html", "bar.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1682732005, 1682732005187933414),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1682731752, 1682731752180272352),
		Data:     nil,
	}, "/html/index.html": &assets.File{
		Path:     "/html/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1682730843, 1682730843977984183),
		Data:     []byte(_Assets950a25363eee220f7d8ce234bcc0b349e4ea9072),
	}, "/html/bar.html": &assets.File{
		Path:     "/html/bar.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1682731752, 1682731752179962126),
		Data:     []byte(_Assets8e1edddb536a70fc4f36d5f948de1bd51ea8571b),
	}}, "")
