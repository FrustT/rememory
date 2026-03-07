## What changed

Added Noto Sans SC as a CJK font for PDF generation. When the bundle language is Chinese (zh-TW, zh-CN), Japanese, or Korean, the PDF now uses Noto Sans SC instead of DejaVu Sans for the text font. DejaVu Sans Mono is still used for monospace sections (PEM blocks, metadata) since those are always ASCII.

The font registration in `registerUTF8Fonts` is now language-aware: it picks the right font family based on whether the language needs CJK glyphs. This means `readme.go` itself needed only a one-line change (passing the language to the font registration function).

Font files added:
- `NotoSansSC-Regular.ttf` (~10 MB) — static Regular weight extracted from the Google Fonts variable font
- `NotoSansSC-Bold.ttf` (~10 MB) — static Bold weight
- `LICENSE-NotoSansSC.txt` — SIL Open Font License

## Why

The existing DejaVu Sans font lacks CJK glyphs, so Chinese text in README.pdf renders as blank boxes. Since the project already supports zh-TW translations for bundle READMEs, the PDF needs a font that can actually render those characters.

Fixes #97

## Testing

- Added `TestGenerateReadmeCJK` — generates a PDF with zh-TW language, Chinese holder name, Chinese project name, and Chinese friend names. Verifies the PDF is produced without errors.
- Added `TestIsCJKLanguage` — verifies CJK language detection for zh-TW, zh-CN, zh, ja, ko, and confirms non-CJK languages return false.
- All existing PDF tests continue to pass (DejaVu Sans still used for non-CJK languages).
- Full test suite passes (`go test ./...`).
- Linting passes (`make lint`).
