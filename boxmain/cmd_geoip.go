package boxmain

import (
	E "github.com/sagernet/sing/common/exceptions"

	"github.com/oschwald/maxminddb-golang"
)

var (
	geoipReader          *maxminddb.Reader
	commandGeoIPFlagFile string
)

func geoipPreRun() error {
	reader, err := maxminddb.Open(commandGeoIPFlagFile)
	if err != nil {
		return err
	}
	if reader.Metadata.DatabaseType != "sing-geoip" {
		reader.Close()
		return E.New("incorrect database type, expected sing-geoip, got ", reader.Metadata.DatabaseType)
	}
	geoipReader = reader
	return nil
}
