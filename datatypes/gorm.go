package datatypes

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// var my_logger = logging.GetLogger()
/*
 GORM data type
*/

// var myLogger = logging.GetLogger()

type IRFTime struct {
	time.Time
}

const (
	CCL_TIME_FORMART      = "2006-01-02T15:04:05"
	CCL_TIME_FROM_FORMART = "2006-01-02 15:04:05"
	CCL_DATE_FORMART      = "2006-01-02"
)

func NewFromDateString(dateString string) (*IRFTime, error) {
	// t, err := time.Parse("2006-01-01T00:00:00", timeString)
	t, err := time.Parse(CCL_DATE_FORMART, dateString)
	if err != nil {
		// my_logger.Error(err)
		return nil, err
	}
	return &IRFTime{t}, err
}

func NewFromTimeString(dateString string) (*IRFTime, error) {
	// t, err := time.Parse("2006-01-01T00:00:00", timeString)
	t, err := time.Parse(CCL_TIME_FORMART, dateString)
	if err != nil {
		// my_logger.Error(err)
		return nil, err
	}
	return &IRFTime{t}, err
}

func (ct *IRFTime) UnmarshalJSON(b []byte) (err error) {
	// myLogger.Info("UnmarshalJSON=", b)
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(CCL_TIME_FROM_FORMART, s)
	// myLogger.Info("ct.Time=", ct.Time)
	return
}

func (ct IRFTime) MarshalJSON() ([]byte, error) {
	// stamp := ct.Time.Format(CCL_TIME_FORMART)
	// myLogger.Info("MarshalJSON=")
	stamp := fmt.Sprintf("\"%s\"", ct.Time.Format(CCL_TIME_FORMART))
	// myLogger.Info("\nstamp=", stamp)
	return []byte(stamp), nil
}

func (ct *IRFTime) Scan(value interface{}) error {
	// myLogger.Info("CCLTime.Scan=", value)
	ct.Time = value.(time.Time)
	return nil
}

// 实现 driver.Valuer 接口, Value 接口不使用指针
func (ct IRFTime) Value() (driver.Value, error) {
	xx := ct.Format(CCL_TIME_FORMART)
	// myLogger.Error("ctkkkkk=", xx)
	return xx, nil
}
