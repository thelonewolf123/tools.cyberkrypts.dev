package interfaces

type Format struct {
	QualityLabel string
	URL          string
	FileName     string
	FileType     string
}

type VideoResponse struct {
	Title   string
	Author  string
	Formats []Format
}
