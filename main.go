package main

import (
        "context"
        "log"
        "bufio"
        "os"
        "sync"
        "flag"
        "github.com/chromedp/chromedp"
)


func main() {

var userJS string
flag.StringVar(&userJS, "j", "", "js to run all pages")
flag.Parse()

sc := bufio.NewScanner(os.Stdin)
var wg sync.WaitGroup

urls :=make(chan string)

for i:=1; i < 5; i++ {
        wg.Add(1)

        go func(){

        for u:= range urls {

        ctx, cancel := chromedp.NewContext(context.Background())
        defer cancel()

        // run task list
        var res string
        err := chromedp.Run(ctx,
                chromedp.Navigate(u),
                chromedp.Evaluate(userJS, &res),
        )
        cancel()
        if err != nil {
                log.Printf("Error on %s : %s", u, err)
                continue
        }

        log.Printf("%s %v", u, res)
         
       }
       wg.Done()
}()

}

for sc.Scan(){
  u:=sc.Text()
       urls <- u
}
wg.Wait()
}
