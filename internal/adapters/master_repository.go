package adapters

import (
	"bytes"
	"compress/gzip"
	"database/sql/driver"
	"encoding/xml"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type MasterRepository struct {
	nest string
}

func NewMasterRepository(path, site string) *MasterRepository {
	nestFolder := filepath.Join(path, site, "nest")
	return &MasterRepository{
		nest: nestFolder,
	}
}

func (r *MasterRepository) ReadNest(id string) (any, error) {
	filename := filepath.Join(r.nest, innerFolder(id), id+".nest.nxl")

	_, err := os.Stat(filename)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}

	// content, err := io.ReadAll(reader)
	// if err != nil {
	// 	return nil, err
	// }

	res := make(map[any]any)
	err = xml.NewDecoder(reader).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *MasterRepository) RemoveNest(id string) error {
	filename := filepath.Join(r.nest, innerFolder(id), id+".nest.nxl")

	_, err := os.Stat(filename)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		return err
	}
	err = os.Remove(filename)
	if err != nil {
		return err
	}

	filenameBak := filepath.Join(r.nest, innerFolder(id), id+".bak")
	_, err = os.Stat(filenameBak)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		return err
	}

	return os.Remove(filenameBak)
}

func innerFolder(id string) string {
	return id[len(id)-2:]
}

// GzippedText is a []byte which transparently gzips data being submitted to
// a database and ungzips data being Scanned from a database.
type GzippedText []byte

// Value implements the driver.Valuer interface, gzipping the raw value of
// this GzippedText.
func (g GzippedText) Value() (driver.Value, error) {
	b := make([]byte, 0, len(g))
	buf := bytes.NewBuffer(b)
	w := gzip.NewWriter(buf)
	w.Write(g)
	w.Close()
	return buf.Bytes(), nil

}

// Scan implements the sql.Scanner interface, ungzipping the value coming off
// the wire and storing the raw result in the GzippedText.
func (g *GzippedText) Scan(src interface{}) error {
	var source []byte
	switch src := src.(type) {
	case string:
		source = []byte(src)
	case []byte:
		source = src
	default:
		return errors.New("incompatible type for GzippedText")
	}
	reader, err := gzip.NewReader(bytes.NewReader(source))
	if err != nil {
		return err
	}
	defer reader.Close()
	b, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	*g = GzippedText(b)
	return nil
}
