package product

import "time"

type Frame struct {
	ID           int64
	PartNo       string
	ProductNo    string
	Length       float64
	Width        float64
	Density      float64
	Metafile     string
	PowderLen    float64
	ProfCnt      int64
	CutLen       float64
	Area         float64
	Drawn        int64
	PointMark    int64   // always zero
	MinRAng      float64 // always zero
	MinRLen      float64
	MinRWidth    float64
	Filename     string
	InnerArea    float64 // always zero
	ToolInfos    []byte
	InsertDate   time.Time
	Creator      string
	Created      time.Time
	Changer      *string
	Changed      *time.Time
	PartChecksum string
}
