package md2html

import _ "embed"

//go:embed styles/github-dark_colorblind.css
var styleDarkColorblind string

//go:embed styles/github-dark_dimmed.css
var styleDarkDimmed string

//go:embed styles/github-dark_high_contrast.css
var styleDarkHighContrast string

//go:embed styles/github-dark_tritanopia.css
var styleDarkTritanopia string

//go:embed styles/github-dark.css
var styleDark string

//go:embed styles/github-light_colorblind.css
var styleLightColorblind string

//go:embed styles/github-light_high_contrast.css
var styleLightHighContrast string

//go:embed styles/github-light_tritanopia.css
var styleLightTritanopia string

//go:embed styles/github-light.css
var styleLight string

// supportedStyles Supported Markdown styles
var supportedStyles = map[string]string{
	"dark_colorblind":     styleDarkColorblind,
	"dark_dimmed":         styleDarkDimmed,
	"dark_high_contrast":  styleDarkHighContrast,
	"dark_tritanopia":     styleDarkTritanopia,
	"dark":                styleDark,
	"light_colorblind":    styleLightColorblind,
	"light_high_contrast": styleLightHighContrast,
	"light_tritanopia":    styleLightTritanopia,
	"light":               styleLight,
}
