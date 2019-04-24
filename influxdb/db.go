package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/branthz/utarrow/lib/log"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	client "github.com/influxdata/influxdb/client/v2"
)

var (
	Session *mgo.Session
	DB      *mgo.Database
)

func dbinit() {
	var err error
	mongoHost := "192.168.29.89:27017"
	Session, err = mgo.Dial(mongoHost)
	if err != nil {
		panic("mongo connection failed")
	}

	Session.SetMode(mgo.Eventual, false)
	Session.SetPoolLimit(200)

	DB = Session.DB("hello")
}

func ping() error {
	return Session.Ping()
}

func sessionAdd(data interface{}) error {
	err := DB.C("session").Insert(data)
	if err != nil {
		return err
	}
	return nil
}

const yestday = -24 * 3600
const oneday = 24 * 3600
const weekSeconds = 2 * 24 * 3600

type session02 struct {
	Tunnel_id         int    `bson:"tunnel_id"`
	Step_sn           string `bson:"step_sn"`
	Sip               string `bson:"sip"`
	Dip               string `bson:"dip"`
	Dport             int    `bson:"dport"`
	Protocol          string `bson:"protocol"`
	Rx_bytes          int    `bson:"rx_bytes"`
	Tx_bytes          int    `bson:"tx_bytes"`
	Start_seconds     int    `bson:"-"`
	Duration_seconds  int    `bson:"duration_seconds"`
	Tx_cmpz_bytes_raw int    `bson:"tx_cmpz_raw"`
	Rx_cmpz_bytes_raw int    `bson:"rx_cmpz_raw"`
	Tx_cmpz_bytes     int    `bson:"tx_cmpz"`
	Rx_cmpz_bytes     int    `bson:"rx_cmpz"`
	Rx_dedup_raw      int    `bson:"rx_dedup_raw"`
	Tx_dedup_raw      int    `bson:"tx_dedup_raw"`
	Tx_dedup          int    `bson:"tx_dedup"`
	Rx_dedup          int    `bson:"rx_dedup"`
	Band              int    `bson:"band"`
	Nowtime           int    `bson:"nowtime"`
	Period            int    `bson:"period"`
	Session_idstr     string `bson:"sid"`
	Cid               int    `bson:"cid"`
}

func getSessions(st, end int) client.BatchPoints {
	//var st = 1545235500
	//var end = 1545238800
	m := bson.M{
		"nowtime": bson.M{
			"$gte": st,
			"$lt":  end,
		},
	}
	result := make([]session02, 0)
	err := DB.C("session").Find(m).All(&result)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	var count int = 0
	for i := 0; i < len(result); i++ {
		tags := map[string]string{
			"sip":   result[i].Sip,
			"dip":   result[i].Dip,
			"dport": strconv.Itoa(result[i].Dport),
			"pro":   result[i].Protocol,
			"tid":   strconv.Itoa(result[i].Tunnel_id),
			"sn":    result[i].Step_sn,
			"cid":   strconv.Itoa(result[i].Cid),
		}
		fields := map[string]interface{}{
			"rx_bytes":     result[i].Rx_bytes,
			"tx_bytes":     result[i].Tx_bytes,
			"tx_dedup":     result[i].Tx_dedup,
			"rx_dedup":     result[i].Rx_dedup,
			"rx_dedup_raw": result[i].Rx_dedup_raw,
			"tx_dedup_raw": result[i].Tx_dedup_raw,
			"rx_cmpz":      result[i].Rx_cmpz_bytes,
			"tx_cmpz":      result[i].Tx_cmpz_bytes,
			"rx_cmpz_raw":  result[i].Rx_cmpz_bytes_raw,
			"tx_cmpz_raw":  result[i].Tx_cmpz_bytes_raw,
			"band":         result[i].Band,
			"period":       result[i].Period,
		}

		pt, _ := client.NewPoint(
			"session",
			tags,
			fields,
			time.Unix(int64(result[i].Nowtime), 0),
		)
		bp.AddPoint(pt)
		count++
	}
	fmt.Printf("-------%d\n", count)
	return bp
}
