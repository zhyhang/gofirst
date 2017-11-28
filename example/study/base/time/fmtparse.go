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
	makeresult(result, i, "UnixDate: ", t.Format(time.UnixDate))
	i++
	makeresult(result, i, "ANSIC:", t.Format(time.ANSIC))
	i++
	makeresult(result, i, "RFC3339:", t.Format(time.RFC3339))
	i++
	makeresult(result, i, "StampMilli:", t.Format(time.StampMilli))
	i++
	makeresult(result, i, "Common:", t.Format("2006-01-02 15:04:05"))
	i++
	makeresult(result, i, "CommonPM:", t.Format("2006-01-02 03:04:05PM"))
	i++
	makeresult(result, i, "CommonPMOffset:", t.Format("2006-01-02 03:04:05PMZ0700"))
	i++
	makeresult(result, i, "CommonPMOffsetLoc:", t.Format("2006-01-02 03:04:05PM -0700 MST"))
	i++
	makeresult(result, i, "CommonSimple:", t.Format("20060102150405"))
	i++
	makeresult(result, i, "CommonMilli:", t.Format("2006-01-02 15:04:05.000"))
	i++
	millifmt := strconv.Itoa(1000 + t.Nanosecond()/1000000)[1:]
	makeresult(result, i, "CommonMilliWithoutPeriod:", t.Format("2006-01-02 15:04:05")+" "+millifmt)
	i++
	makeresult(result, i, "CommonMilliSimple:", t.Format("20060102150405.000"))
	i++
	makeresult(result, i, "CommonMilliSimpleWithoutPeriod:", t.Format("20060102150405")+millifmt)
	return result
}

func makeresult(result [][]string, i int, l, ls string) {
	result[i] = make([]string, 2)
	result[i][0] = l
	result[i][1] = ls
}

// examples of parsing time from string by layout
func parse() [][]string {
	result := make([][]string, 128)
	i := 0

	// parse local time
	// when time string include location info, will using it
	timestr := "2017-11-18 23:00:08"
	l := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(l, timestr, time.Local)
	i++
	makeresult(result, i, "layout: "+l+"; local time: "+timestr, t.String())

	timestr = "2017-11-18 23:00:08.666"
	l = "2006-01-02 15:04:05" // layout can put .000 or not put .000, for parsing millis
	t, _ = time.ParseInLocation(l, timestr, time.Local)
	i++
	makeresult(result, i, "layout: "+l+"; local time: "+timestr, t.String())

	timestr = "20171118230008.666"
	l = "20060102150405.000" // layout can put .000 or not put .000, for parsing millis
	t, _ = time.ParseInLocation(l, timestr, time.Local)
	i++
	makeresult(result, i, "layout: "+l+"; local time: "+timestr, t.String())

	timestr = "20171118230008666"
	timestr1 := timestr[0:14] + "." + timestr[14:]
	l = "20060102150405.000"
	t, _ = time.ParseInLocation(l, timestr1, time.Local)
	i++
	makeresult(result, i, "layout: "+l+"; local time: "+timestr, t.String())

	timestr = "Sat Nov 18 2017 23:00:08.666 CST +0800"
	l = "Mon Jan _2 2006 15:04:05.000 MST Z0700"
	t, _ = time.ParseInLocation(l, timestr, time.Local)
	i++
	makeresult(result, i, "layout: "+l+"; local time: "+timestr, t.String())

	timestr = "Sat Nov 18 2017 23:00:08.666 MST" //include location info using it
	l = "Mon Jan _2 2006 15:04:05.000 MST"
	t, _ = time.ParseInLocation(l, timestr, time.Local)
	i++
	makeresult(result, i, "layout: "+l+"; local time: "+timestr, t.String())
	i++
	makeresult(result, i, "layout: "+l+"; local time(to local): "+timestr, t.Local().String())

	// parse utc time [parse() method default logic]
	// when time string include location info, will using it
	timestr = "2017-11-18 23:00:08"
	l = "2006-01-02 15:04:05"
	t, _ = time.Parse(l, timestr)
	i++
	makeresult(result, i, "layout: "+l+"; utc time: "+timestr, t.String())

	timestr = "2017-11-18 23:00:08.666"
	l = "2006-01-02 15:04:05" // layout can put .000 or not put .000, for parsing millis
	t, _ = time.Parse(l, timestr)
	i++
	makeresult(result, i, "layout: "+l+"; utc time: "+timestr, t.String())

	timestr = "20171118230008.666"
	l = "20060102150405.000" // layout can put .000 or not put .000, for parsing millis
	t, _ = time.Parse(l, timestr)
	i++
	makeresult(result, i, "layout: "+l+"; utc time: "+timestr, t.String())

	timestr = "20171118230008666"
	timestr1 = timestr[0:14] + "." + timestr[14:]
	l = "20060102150405.000"
	t, _ = time.Parse(l, timestr1)
	i++
	makeresult(result, i, "layout: "+l+"; utc time: "+timestr, t.String())

	timestr = "Sat Nov 18 2017 23:00:08.666 CST +0800" //include location info, using it
	l = "Mon Jan _2 2006 15:04:05.000 MST Z0700"
	t, _ = time.Parse(l, timestr)
	i++
	makeresult(result, i, "layout: "+l+"; utc time: "+timestr, t.String())
	i++
	makeresult(result, i, "layout: "+l+"; utc time(to utc): "+timestr, t.UTC().String())

	// time string is error
	timestr = "20171118230008.abc"
	l = "20060102150405.000"
	t, _ = time.ParseInLocation(l, timestr, time.Local)
	i++
	makeresult(result, i, "layout: "+l+"; local time: "+timestr, t.String())

	return result
}

// local time from timestamp (millis)
func parseLocaltimeOfTimestamp(millis int64) time.Time {
	sec := millis / 1000;
	nsec := millis % 1000 * 1000000
	if sec < 0 {
		sec--
		nsec = -nsec
	}
	return time.Unix(sec, nsec)
}

// recommend using this method to get milliseconds
func toTimestampMillis1(t time.Time) int64 {
	return t.Unix()*1000 + int64(t.Nanosecond())/int64(time.Millisecond)
}

// this retrieve millis, possible exceeding int64, see time.UnixNano
func toTimestampMillis2(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func main() {

	pln := fmt.Println

	pln("time now:")
	tnow := timenow()
	pln("\t" + tnow.String())

	pln()
	pln("date time format:")
	for i, v := range format() {
		if len(v) > 0 {
			pln("\t" + strconv.Itoa(i+1) + "." + v[0] + "\t" + v[1])
		}
	}

	pln()
	pln("date time parse:")
	for i, v := range parse() {
		if len(v) > 0 {
			pln("\t" + strconv.Itoa(i) + "." + v[0] + " ---> " + v[1])
		}
	}

	pln()
	pln("currrent time millis:")
	pln("\tmethod1: " + strconv.FormatInt(toTimestampMillis1(tnow), 10))
	pln("\tmehtod2: " + strconv.FormatInt(toTimestampMillis2(tnow), 10))

	pln()
	pln("date time millis:")
	longt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-11-17 18:28:38.888", time.Local)
	pln("\ttime is " + longt.String())
	pln("\t\tmethod1: " + strconv.FormatInt(toTimestampMillis1(longt), 10))
	pln("\t\tmehtod2: " + strconv.FormatInt(toTimestampMillis2(longt), 10))
	longt, _ = time.ParseInLocation("2006-01-02 15:04:05", "0025-10-28 08:28:38.666", time.Local)
	pln("\ttime is " + longt.String())
	pln("\t\tmethod1: " + strconv.FormatInt(toTimestampMillis1(longt), 10))
	pln("\t\tmehtod2: " + strconv.FormatInt(toTimestampMillis2(longt), 10))
	longt, _ = time.ParseInLocation("2006-01-02 15:04:05", "2888-02-21 08:18:28.666", time.Local)
	pln("\ttime is " + longt.String())
	pln("\t\tmethod1: " + strconv.FormatInt(toTimestampMillis1(longt), 10))
	pln("\t\tmehtod2: " + strconv.FormatInt(toTimestampMillis2(longt), 10))
	longt, _ = time.ParseInLocation("2006-01-02 15:04:05", "1970-01-01 08:00:00.100", time.Local)
	pln("\ttime is " + longt.String())
	pln("\t\tmethod1: " + strconv.FormatInt(toTimestampMillis1(longt), 10))
	pln("\t\tmehtod2: " + strconv.FormatInt(toTimestampMillis2(longt), 10))
	longt, _ = time.ParseInLocation("2006-01-02 15:04:05", "1970-01-01 07:59:59.900", time.Local)
	pln("\ttime is " + longt.String())
	pln("\t\tmethod1: " + strconv.FormatInt(toTimestampMillis1(longt), 10))
	pln("\t\tmehtod2: " + strconv.FormatInt(toTimestampMillis2(longt), 10))

	pln()
	pln("data time parse from millis timestamp:")
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", "0025-10-28 08:28:38.666", time.Local)
	ts := toTimestampMillis1(tt)
	pln("\ttimestamp millis: " + strconv.FormatInt(ts, 10) + "; to time: " + parseLocaltimeOfTimestamp(ts).String())
	ts = int64(1510914518888) //2017-11-17 18:28:38.888
	pln("\ttimestamp millis: " + strconv.FormatInt(ts, 10) + "; to time: " + parseLocaltimeOfTimestamp(ts).String())
	ts = int64(1511164435560)
	pln("\ttimestamp millis: " + strconv.FormatInt(ts, 10) + "; to time: " + parseLocaltimeOfTimestamp(ts).String())
	ts = int64(28973722708666) //2888-02-21 08:18:28.666
	pln("\ttimestamp millis: " + strconv.FormatInt(ts, 10) + "; to time: " + parseLocaltimeOfTimestamp(ts).String())
	ts = 0
	pln("\ttimestamp millis: " + strconv.FormatInt(ts, 10) + "; to time: " + parseLocaltimeOfTimestamp(ts).String())
	ts = 100
	pln("\ttimestamp millis: " + strconv.FormatInt(ts, 10) + "; to time: " + parseLocaltimeOfTimestamp(ts).String())
	ts = -100
	pln("\ttimestamp millis: " + strconv.FormatInt(ts, 10) + "; to time: " + parseLocaltimeOfTimestamp(ts).String())

}
