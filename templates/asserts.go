package templates

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets14fe4559026d4c5b5eb530ee70300c52d99e70d7 = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Index Template</title>\n</head>\n<body>\n    <p>Hello, {{.Foo}}</p>\n</body>\n</html>"
var _Assets82000b211c7454e536f1b5fcf63d1a2f1941bff4 = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Index Template</title>\n</head>\n<body>\n    <p>Hello, {{.Foo}}</p>\n</body>\n</html>"
var _Assetsb0abd4aceef5b07bf5a6bef10e9528f3bb1fe528 = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <title>Bar Template</title>\n</head>\n<body>\n  <p>Can you see this? - {{.Bar}}</p>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"bar.html", "index.html", "users"}, "/users": []string{"index.html"}}, map[string]*assets.File{
	"/users/index.html": &assets.File{
		Path:     "/users/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1682730843, 1682730843977984183),
		Data:     []byte(_Assets82000b211c7454e536f1b5fcf63d1a2f1941bff4),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1683439308, 1683439308749599096),
		Data:     nil,
	}, "/bar.html": &assets.File{
		Path:     "/bar.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1682731752, 1682731752179962126),
		Data:     []byte(_Assetsb0abd4aceef5b07bf5a6bef10e9528f3bb1fe528),
	}, "/index.html": &assets.File{
		Path:     "/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1682730843, 1682730843977984183),
		Data:     []byte(_Assets14fe4559026d4c5b5eb530ee70300c52d99e70d7),
	}, "/users": &assets.File{
		Path:     "/users",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1683438993, 1683438993864685760),
		Data:     nil,
	}}, "")
