package main

import (
	"fmt"

	"github.com/influxdata/influxdb/tsdb"
	ts "github.com/influxdata/influxdb/tsdb/index/tsi1"
)

func main() {
	start()
}

var (
	data = []struct {
		Key  string
		Name string
		Tags map[string]string
	}{
		{"cpu,region=west,server=a", "cpu", map[string]string{"region": "west", "server": "a"}},
		{"cpu,region=west,server=b", "cpu", map[string]string{"region": "west", "server": "b"}},
		{"cpu,region=east,server=a", "cpu", map[string]string{"region": "east", "server": "a"}},
		{"cpu,region=north,server=c", "cpu", map[string]string{"region": "north", "server": "c"}},
		{"cpu,region=south,server=s", "cpu", map[string]string{"region": "south", "server": "s"}},
		{"mem,region=west,server=a", "mem", map[string]string{"region": "west", "server": "a"}},
		{"mem,region=west,server=b", "mem", map[string]string{"region": "west", "server": "b"}},
		{"mem,region=west,server=c", "mem", map[string]string{"region": "west", "server": "c"}},
		{"disk,region=east,server=a", "disk", map[string]string{"region": "east", "server": "a"}},
		{"disk,region=east,server=a", "disk", map[string]string{"region": "east", "server": "a"}},
		{"disk,region=north,server=c", "disk", map[string]string{"region": "north", "server": "c"}},
	}
)

func start() {
	var path = "./steplog/index"
	var sfpath = "./steplog/"
	var opt tsdb.EngineOptions
	var sfile = tsdb.NewSeriesFile(sfpath)
	ix := ts.NewIndex(sfile, "steplog", ts.WithPath(path), ts.WithMaximumLogFileSize(int64(opt.Config.MaxIndexLogFileSize)), ts.WithSeriesIDCacheSize(opt.Config.SeriesIDSetCacheSize))
	ix.PartitionN = 8
	err := ix.Open()
	if err != nil {
		fmt.Println(err)
	}
	//for _, pt := range data {
	//	if err := ix.CreateSeriesIfNotExists([]byte(pt.Key), []byte(pt.Name), models.NewTags(pt.Tags)); err != nil {
	//		fmt.Println(err)
	//	}
	//}
	bl, err := ix.MeasurementExists([]byte("cpu"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bl)

	ix.Close()
}
