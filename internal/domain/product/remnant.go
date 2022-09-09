package product

import "time"

type Remnant struct {
	ID           int64
	ProductNo    string
	Name         *string
	Length       float64
	Width        float64
	Thick        *float64
	Quality      *string
	Density      float64
	Site         string
	Metafile     string
	PowderLen    *float64
	ProfCnt      *int64
	CutLen       *float64
	Area         *float64
	Drawn        int64
	PointMark    *int64
	MinRAng      *float64
	MinRLen      *float64
	MinRWidth    *float64
	Filename     string
	InnerArea    *float64
	ToolInfos    []byte
	InsertDate   time.Time
	Created      time.Time
	Creator      string
	Changed      *time.Time
	Changer      *string
	PartChecksum *string
}
