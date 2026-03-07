package pdf

import (
	_ "embed"

	"github.com/go-pdf/fpdf"
)

// DejaVu Sans font files (Bitstream Vera / DejaVu license - see fonts/LICENSE-DejaVu.txt)
// These provide comprehensive UTF-8 / Unicode coverage including Latin, Cyrillic,
// Greek, Arabic, Hebrew, and many other scripts.

//go:embed fonts/DejaVuSans.ttf
var dejaVuSansRegular []byte

//go:embed fonts/DejaVuSans-Bold.ttf
var dejaVuSansBold []byte

//go:embed fonts/DejaVuSans-Oblique.ttf
var dejaVuSansOblique []byte

//go:embed fonts/DejaVuSans-BoldOblique.ttf
var dejaVuSansBoldOblique []byte

//go:embed fonts/DejaVuSansMono.ttf
var dejaVuSansMonoRegular []byte

//go:embed fonts/DejaVuSansMono-Bold.ttf
var dejaVuSansMonoBold []byte

// Noto Sans SC font files (SIL Open Font License - see fonts/LICENSE-NotoSansSC.txt)
// These provide CJK (Chinese, Japanese, Korean) glyph coverage.

//go:embed fonts/NotoSansSC-Regular.ttf
var notoSansSCRegular []byte

//go:embed fonts/NotoSansSC-Bold.ttf
var notoSansSCBold []byte

// Font family names used throughout the PDF generator.
const (
	fontSans = "DejaVuSans"
	fontMono = "DejaVuSansMono"
)

// isCJKLanguage reports whether a language code requires CJK glyphs.
func isCJKLanguage(lang string) bool {
	switch lang {
	case "zh-TW", "zh-CN", "zh", "ja", "ko":
		return true
	}
	return false
}

// registerUTF8Fonts adds the embedded UTF-8 fonts to the PDF instance.
// When lang is a CJK language, Noto Sans SC is registered under fontSans
// so that Chinese/Japanese/Korean text renders correctly.
// After calling this, use fontSans and fontMono as the family name in SetFont().
func registerUTF8Fonts(pdf *fpdf.Fpdf, lang string) {
	if isCJKLanguage(lang) {
		pdf.AddUTF8FontFromBytes(fontSans, "", notoSansSCRegular)
		pdf.AddUTF8FontFromBytes(fontSans, "B", notoSansSCBold)
		pdf.AddUTF8FontFromBytes(fontSans, "I", notoSansSCRegular)
		pdf.AddUTF8FontFromBytes(fontSans, "BI", notoSansSCBold)
	} else {
		pdf.AddUTF8FontFromBytes(fontSans, "", dejaVuSansRegular)
		pdf.AddUTF8FontFromBytes(fontSans, "B", dejaVuSansBold)
		pdf.AddUTF8FontFromBytes(fontSans, "I", dejaVuSansOblique)
		pdf.AddUTF8FontFromBytes(fontSans, "BI", dejaVuSansBoldOblique)
	}

	pdf.AddUTF8FontFromBytes(fontMono, "", dejaVuSansMonoRegular)
	pdf.AddUTF8FontFromBytes(fontMono, "B", dejaVuSansMonoBold)
}
