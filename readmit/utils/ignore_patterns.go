package utils

// IgnorePatterns is a comprehensive list of files and directories
// that should be skipped when reading project files.
var IgnorePatterns = []string{

	"ReadMe.md", "README.MD", "readme.md", "ReadMe.MD", "Readme.md", "readme.MD", // common readme variations
	// Package manager dirs
	"node_modules", "bower_components", "jspm_packages", "vendor", "Pods", "target",

	// Build artifacts
	"dist", "build", "bin", "out", "obj", ".next", ".nuxt", ".parcel-cache", ".vercel", ".expo", ".angular", ".svelte-kit", ".astro",

	// System junk
	".DS_Store", "Thumbs.db", "ehthumbs.db", "Icon\r", "Desktop.ini",

	// VCS dirs
	".git", ".svn", ".hg", ".bzr", ".idea", ".vscode", ".fleet",

	// Env & config
	".env", ".env.local", ".env.development", ".env.production", ".env.test", ".secrets", ".credentials",

	// Logs / tmp
	"*.log", "npm-debug.log*", "yarn-debug.log*", "yarn-error.log*", "pnpm-debug.log*", "lerna-debug.log*",
	"tmp", "temp", ".temp", ".cache", "cache", ".pytest_cache", ".mypy_cache", ".parcel-cache", ".next/cache",

	// Dependencies lock files
	"go.mod", "package-lock.json", "yarn.lock", "pnpm-lock.yaml", "composer.lock", "Cargo.lock", "Pipfile.lock",

	// Archives & compressed
	"*.zip", "*.tar", "*.tar.gz", "*.tgz", "*.rar", "*.7z", "*.gz", "*.bz2",

	// Binary executables
	"*.exe", "*.dll", "*.so", "*.dylib", "*.bin", "*.out", "*.o", "*.class", "*.pyc", "*.pyo",

	// Disk images / installers
	"*.iso", "*.dmg", "*.msi", "*.deb", "*.rpm", "*.apk", "*.app", "*.jar",

	// Images
	"*.png", "*.jpg", "*.jpeg", "*.gif", "*.bmp", "*.tiff", "*.tif", "*.ico", "*.icns", "*.webp", "*.heic", "*.svg",

	// Audio
	"*.mp3", "*.wav", "*.ogg", "*.flac", "*.aac", "*.m4a", "*.wma", "*.aiff", "*.aif",

	// Video
	"*.mp4", "*.mkv", "*.avi", "*.mov", "*.wmv", "*.flv", "*.webm", "*.mpeg", "*.mpg", "*.3gp", "*.m4v",

	// Fonts
	"*.woff", "*.woff2", "*.ttf", "*.otf", "*.eot",

	// Office docs
	"*.doc", "*.docx", "*.xls", "*.xlsx", "*.ppt", "*.pptx", "*.pdf", "*.odt", "*.ods", "*.odp", "*.rtf",

	// Backups
	"*.bak", "*.old", "*.orig", "*.tmp", "*.swp", "*.swo", "*~",

	// Coverage / reports
	"coverage", "coverage-final.json", "lcov.info", ".nyc_output", "reports", "test-results",

	// Cloud & IDE specific
	".vercel", ".firebase", ".aws", ".terraform", ".circleci", ".github", ".gitlab", ".azure-pipelines", ".bitrise",

	// Misc dev junk
	".eslintcache", ".stylelintcache", ".tsbuildinfo", ".npm", ".yarn", ".gradle", ".pytest_cache", "__pycache__",

	// Databases / dumps
	"*.sqlite", "*.sqlite3", "*.db", "*.sql", "*.dump",
}
