package product

import "time"

type Part struct {
	i         int64
	versionNo *string
	partNo    string
	productNo string
	name      *string
	length    float64
	width     float64
	thickness float64
	quality   string
	density   float64
	site      string
	metafile  string
	// powderLen       float64 // always zero
	profCnt int64
	cutLen  float64
	area    float64
	drawn   int64
	// pointMark       int64 // always zero
	minRAng         float64
	minRLen         float64
	minRWidth       float64
	infoTxt         *string
	filename        string
	attributes      []byte
	innerArea       float64
	toolInfos       []byte
	insertDate      time.Time
	created         time.Time
	creator         string
	changed         *time.Time
	changer         *string
	sectionCopyData []byte
	partChecksum    string
}
