package frontmatter

const (
	// Hugo front matter.

	Build       Key = "_build"
	Date        Key = "date"
	Description Key = "description"
	Draft       Key = "draft"
	List        Key = "list"
	Render      Key = "render"
	Title       Key = "title"

	// Site front matter.

	TableOfContents Key = "table-of-contents"

	// Radar front matter.

	RadarIndex   Key = "radarIndex"
	RadarSection Key = "radarSection"
	RadarTier    Key = "radarTier"
	RadarX       Key = "radarX"
	RadarY       Key = "radarY"
)

type Key = string
