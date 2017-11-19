package main

import (
	"time"
	"fmt"
	"strconv"
)

func timenow() time.Time {
	return time.Now()
}

// format time by layout (formatter)
func format() [][]string {
	t := time.Now()
	result := make([][]string, 24)
	i := 0
	makeformatdata(result, i, "UnixDate: ", t.Format(time.UnixDate))
	i++
	makeformatdata(result, i, "ANSIC:", t.Format(time.ANSIC))
	i++
	makeformatdata(result, i, "RFC3339:", t.Format(time.RFC3339))
	i++
	makeformatdata(result, i, "StampMilli:", t.Format(time.StampMilli))
	i++
	makeformatdata(result, i, "Common:", t.Format("2006-01-02 15:04:05"))
	i++
	makeformatdata(result, i, "CommonSimple:", t.Format("20060102150405"))
	i++
	makeformatdata(result, i, "CommonMilli:", t.Format("2006-01-02 15:04:05.000"))
	i++
	millifmt := strconv.Itoa(1000 + t.Nanosecond()/1000000)[1:]
	makeformatdata(result, i, "CommonMilliWithoutPeriod:", t.Format("2006-01-02 15:04:05")+" "+millifmt)
	i++
	makeformatdata(result, i, "CommonMilliSimple:", t.Format("20060102150405.000"))
	i++
	makeformatdata(result, i, "CommonMilliSimpleWithoutPeriod:", t.Format("20060102150405")+millifmt)
	return result
}

func makeformatdata(result [][]string, i int, l, ls string) {
	result[i] = make([]string, 2)
	result[i][0] = l
	result[i][1] = ls
}

// examples of parsing time from string by layout
func parse() map[string]time.Time{
	result := make(map[string]time.Time)

	// parse local time
	timestr:="2017-11-18 23:00:08"
	l:="2006-01-02 15:04:05"
	t,_:=time.ParseInLocation(l, timestr, time.Local)
	result["layout: " +l+"; local time: "+timestr]=t
	timestr="2017-11-18 23:00:08.666"
	l="2006-01-02 15:04:05"
	t,_=time.ParseInLocation(l, timestr, time.Local)
	result["layout: " +l+"; local time: "+timestr]=t
	return result

	// parse utc time
}

// local time from timestamp (millis)
func parseLocaltimeOfTimestamp(millis int64) time.Time {
	return time.Unix(millis/1000, 1000000*(millis%1000))
}

// recommend using this method to get milliseconds
func toTimestampMillis1(t time.Time) int64 {
	sec := t.Unix()
	millis := int64(t.Nanosecond()) / int64(time.Millisecond)
	if sec > 0 {
		return sec*1000 + millis
	} else if sec < 0 {
		return sec*1000 - millis
	} else if t.UnixNano() >= 0 {
		return millis
	} else {
		return -millis
	}
}

// this retrieve millis, possible exceeding int64, see time.UnixNano
func toTimestampMillis2(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func main() {

	print := fmt.Println

	print("time now:")
	tnow := timenow()
	print(tnow)

	print()
	print("date time format:")
	for _, fmttime := range format() {
		if len(fmttime) > 0 {
			print("\t" + fmttime[0] + "\t" + fmttime[1])
		}
	}

	print()
	print("date time parse:")
	for k,v:=range parse(){
		print(k+" ---> "+v.String())
	}

	print()
	print("currrent time millis:")
	print("\tmethod1: " + strconv.FormatInt(toTimestampMillis1(tnow), 10))
	print("\tmehtod2: " + strconv.FormatInt(toTimestampMillis2(tnow), 10))

	print("date time millis:")
	longt, _ := time.ParseInLocation("2006-01-02 15:04:05", "1978-12-20 11:00:12.888", time.Local)
	print("\tmethod1: " + strconv.FormatInt(toTimestampMillis1(longt), 10))
	print("\tmehtod2: " + strconv.FormatInt(toTimestampMillis2(longt), 10))

	print(time.Unix(282970812, 888000000))

}
