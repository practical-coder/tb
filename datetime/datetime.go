package datetime

import (
	"fmt"
	"time"
)

func TheTimeIsNow() {
	t := time.Now()
	fmt.Printf("NOW!:\t\t%s\n", t)
}

func Format() {
	// t := time.Date(2021, time.March, 21, 7, 39, 6, 0, time.Local)
	t := time.Now()
	fmt.Printf("Date\t\t%s\n", t.Format("2006/1/2"))
	fmt.Printf("Time\t\t%s\n", t.Format("15:04"))
	fmt.Printf("Kitchen time\t%s\n", t.Format(time.Kitchen))
	fmt.Printf("ANSIC\t\t%s\n", t.Format(time.ANSIC))
	fmt.Printf("RFC822\t\t%s\n", t.Format(time.RFC822))
	fmt.Printf("RFC822Z\t\t%s\n", t.Format(time.RFC822Z))
	fmt.Printf("RFC850\t\t%s\n", t.Format(time.RFC850))
	fmt.Printf("RFC1123\t\t%s\n", t.Format(time.RFC1123))
	fmt.Printf("RFC1123Z\t%s\n", t.Format(time.RFC1123Z))
	fmt.Printf("RFC3339\t\t%s\n", t.Format(time.RFC3339))
	fmt.Printf("RFC3339Nano\t%s\n", t.Format(time.RFC3339Nano))
	fmt.Printf("RubyDate\t%s\n", t.Format(time.RubyDate))
	fmt.Printf("UnixDate\t%s\n", t.Format(time.UnixDate))
	fmt.Printf("Stamp\t\t%s\n", t.Format(time.Stamp))
	fmt.Printf("StampMilli\t%s\n", t.Format(time.StampMilli))
	fmt.Printf("StampMicro\t%s\n", t.Format(time.StampMicro))
	fmt.Printf("StampNano\t%s\n", t.Format(time.StampNano))
}

func Epoch() {
	startEpoch := time.Unix(0, 0)
	fmt.Printf("\nStart of the Epoch: %s\n", startEpoch)
	t := time.Now().Unix()
	fmt.Printf("Seconds since Epoch: %d\n", t)
	mt := time.Now().UnixMilli()
	fmt.Printf("Milliseconds since Epoch: %d\n", mt)
	ut := time.Now().UnixMicro()
	fmt.Printf("Microseconds since Epoch: %d\n", ut)
	nt := time.Now().UnixNano()
	fmt.Printf("Nanoseconds since Epoch: %d\n", nt)
}
