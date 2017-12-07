// Go's offers extensive support for times and durations;
// here are some examples.

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // We'll start by getting the current time.
    now := time.Now()
    p(now)

    // You can build a `time` struct by providing the
    // year, month, day, etc. Times are always associated
    // with a `Location`, i.e. time zone.
    then := time.Date(
        2014, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    // You can extract the various components of the time
    // value as expected.
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    // The Monday-Sunday `Weekday` is also available.
    p(then.Weekday())

    // These methods compare two times, testing if the
    // first occurs before, after, or at the same time
    // as the second, respectively.
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    // The `Sub` methods returns a `Duration` representing
    // the interval between two times.
    diff := now.Sub(then)
    p(diff)

    // We can compute the length of the duration in
    // various units.
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    // You can use `Add` to advance a time by a given
    // duration, or with a `-` to move backwards by a
    // duration.
    p(then.Add(diff))
    p(then.Add(-diff))
    t := time.Now()
    p(t.Format("2006-01-02 15:04:05"))
    fmt.Printf("%T",t.Format("2006-01-02 15:04:05"))
    formatedTime := t.Format(time.RFC1123)
    fmt.Println(formatedTime)
    text := "168h0m0s"
    duration, _ := time.ParseDuration(string(text))
    fmt.Printf("%T",duration)
     t = time.Date(2016, 06, 28, 10, 34, 58, 651387237, time.UTC)
     fmt.Println("*******",t)
    var creation time.Time 
    fmt.Println("creation add and sub",t.Add(duration).Sub(time.Now()))
    p("start date ",t.Format(time.RFC1123))
   
    	if creation.Add(24*time.Hour).Sub(time.Now()) < time.Minute {
		fmt.Println("errror")
	}
	p("format  :",t.Format("2016-07-05 17:43:57.718000022 +0530 IST"))
	tokenExpire := 7 * 24 * time.Hour
	t1 := time.Now().String()
	p("-----------",t1)
      t2, _ := time.Parse(time.RFC1123,t1)
      p("-----------",t2)
      time.Sleep(100)
      ram := t2.Add(tokenExpire).Sub(time.Now())
      p("remaining :",ram)
      
}
