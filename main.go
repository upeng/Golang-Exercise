package main
import (
  "fmt"
  "time"
  "strconv"
  "github.com/robfig/cron"
)

func main() {
  
	c := cron.New()
	//定时任务 cron表达式，每10s一次  精确到秒
	c.AddFunc("*/10 * * * * *", func() {
		start := time.Now().Add(-30*time.Second).UnixNano() / int64(time.Second)
		end := time.Now().Add(-1*time.Second).UnixNano() / int64(time.Second)

		s := strconv.FormatInt(start, 10) + "000"
		e := strconv.FormatInt(end, 10) + "999"
		fmt.Println(s,e)
	})

	c.Start()
	defer c.Stop()
	select {} //阻塞主线程停止

  // just test file

}
// import (
// 	"encoding/json"
// 	"fmt"
// )

// type User struct {
//   Name string 
//   Age int64
// }

// func main() {
//   user := User{Name:"taylr", Age:15}
//   str, _ := json.Marshal(user)
// 	fmt.Println("Hello, World!", string(str))
// }
