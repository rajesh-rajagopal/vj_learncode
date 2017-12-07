package main 
 
 import (
 "time"
  "fmt"
  )
  
  func main() {
   p := fmt.Println
   pf := fmt.Printf
  tokenExpire := 7 * 24 * time.Hour
	at := time.Now()
	end := at.Add(tokenExpire)
	p("start",at,"end",end)
	pf("\n %T - %T",at,end)
        t1 := at.Format(time.RFC3339)
        t2 := end.Format(time.RFC3339)
        p("start",t1,"end",t2)
        pf("\n %T - %T",t1,t2)
        t3, _ := time.Parse(time.RFC3339,t1)
        t4, _ := time.Parse(time.RFC3339,t2)
        p(t3)
        pf("\n%T",t3)
        ram := t3.Add(tokenExpire).Sub(time.Now())
        p("ram :",ram)
        p(t4.Sub(t3))
        p(time.Unix(123456789,0).String())
        t5,err := time.Parse("yyyy-mm-dd hh:mm:ss ZZ",time.Now().String())
        p(" error :",err)
        p(t5)
}
