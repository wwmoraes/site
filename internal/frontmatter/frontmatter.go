package frontmatter

const (
	// Hugo front matter predefined keys.

	Aliases       Key = "aliases"
	Audio         Key = "audio"
	Cascade       Key = "cascade"
	Date          Key = "date"
	Description   Key = "description"
	Draft         Key = "draft"
	ExpiryDate    Key = "expiryDate"
	Headless      Key = "headless"
	Images        Key = "images"
	IsCJKLanguage Key = "isCJKLanguage"
	Keywords      Key = "keywords"
	Lastmod       Key = "lastmod"
	Layout        Key = "layout"
	LinkTitle     Key = "linkTitle"
	Markup        Key = "markup"
	Outputs       Key = "outputs"
	PublishDate   Key = "publishDate"
	Resources     Key = "resources"
	Series        Key = "series"
	Slug          Key = "slug"
	Summary       Key = "summary"
	Taxonomies    Key = "taxonomies"
	Title         Key = "title"
	Type          Key = "type"
	URL           Key = "url"
	Videos        Key = "videos"
	Weight        Key = "weight"

	// hidden-ish predefined keys.

	Build  Key = "_build"
	Target Key = "_target"

	// target sub-keys.

	Background  Key = "background"
	Environment Key = "environment"
	Kind        Key = "kind"
	Lang        Key = "lang"
	Path        Key = "path"

	// build sub-keys.

	List             Key = "list"
	PublishResources Key = "publishResources"
	Render           Key = "render"

	// Site-/theme-specific keys.

	TableOfContents Key = "table-of-contents"

	// Radar section keys.

	RadarIndex   Key = "radarIndex"
	RadarSection Key = "radarSection"
	RadarTier    Key = "radarTier"
	RadarX       Key = "radarX"
	RadarY       Key = "radarY"
)

type Key = string
