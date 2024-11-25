module github.com/wwmoraes/site

go 1.22.4

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/adrg/frontmatter v0.2.0
	github.com/bmatcuk/doublestar/v4 v4.6.1
	github.com/charmbracelet/huh v0.2.3
	github.com/dsoprea/go-exif/v3 v3.0.1
	github.com/dsoprea/go-jpeg-image-structure/v2 v2.0.0-20221012074422-4f3f7e934102
	github.com/dsoprea/go-png-image-structure/v2 v2.0.0-20210512210324-29b889a6093d
	github.com/dsoprea/go-utility/v2 v2.0.0-20221003172846-a3e1774ef349
	github.com/gabriel-vasile/mimetype v1.4.4
	github.com/goccy/go-json v0.10.2
	github.com/gocolly/colly v1.2.0
	github.com/shurcooL/githubv4 v0.0.0-20231126234147-1cffa1f02456
	github.com/spf13/cast v1.6.0
	github.com/spf13/cobra v1.8.0
	github.com/wwmoraes/go-rwfs v0.0.0-20231101192853-2f37ed32d908
	golang.org/x/oauth2 v0.15.0
	golang.org/x/text v0.17.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/PuerkitoBio/goquery v1.8.1 // indirect
	github.com/alecthomas/chroma v0.10.0 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/antchfx/htmlquery v1.3.0 // indirect
	github.com/antchfx/xmlquery v1.3.18 // indirect
	github.com/antchfx/xpath v1.2.5 // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/catppuccin/go v0.2.0 // indirect
	github.com/charmbracelet/bubbles v0.17.1 // indirect
	github.com/charmbracelet/bubbletea v0.25.0 // indirect
	github.com/charmbracelet/glamour v0.6.0 // indirect
	github.com/charmbracelet/lipgloss v0.9.1 // indirect
	github.com/containerd/console v1.0.4-0.20230313162750-1ae8d489ac81 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/dsoprea/go-iptc v0.0.0-20200609062250-162ae6b44feb // indirect
	github.com/dsoprea/go-logging v0.0.0-20200710184922-b02d349568dd // indirect
	github.com/dsoprea/go-photoshop-info-format v0.0.0-20200609050348-3db9b63b202c // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-xmlfmt/xmlfmt v0.0.0-20191208150333-d5b6f63a941b // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/golang/geo v0.0.0-20210211234256-740aa86cb551 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/microcosm-cc/bluemonday v1.0.25 // indirect
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/saintfish/chardet v0.0.0-20230101081208-5e3ef4b5456d // indirect
	github.com/shurcooL/graphql v0.0.0-20230722043721-ed46e5a46466 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/temoto/robotstxt v1.1.2 // indirect
	github.com/yuin/goldmark v1.6.0 // indirect
	github.com/yuin/goldmark-emoji v1.0.2 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/term v0.23.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//// go get -u updates ALL direct and indirect dependencies... yet bep/overlayfs
//// past v0.6.0 is incompatible with hugo@v0.121.x:
//// # github.com/gohugoio/hugo/hugolib/filesystems
//// ../../go/pkg/mod/github.com/gohugoio/hugo@v0.121.1/hugolib/filesystems/basefs.go:630:69: cannot use hugofs.LanguageDirsMerger (variable of type func(lofi []fs.FileInfo, bofi []fs.FileInfo) []fs.FileInfo) as overlayfs.DirsMerger value in struct literal
//// ../../go/pkg/mod/github.com/gohugoio/hugo@v0.121.1/hugolib/filesystems/basefs.go:631:69: cannot use hugofs.LanguageDirsMerger (variable of type func(lofi []fs.FileInfo, bofi []fs.FileInfo) []fs.FileInfo) as overlayfs.DirsMerger value in struct literal
// replace github.com/bep/overlayfs => github.com/bep/overlayfs v0.6.0
