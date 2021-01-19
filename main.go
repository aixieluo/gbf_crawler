package main

import (
	"crawler/database"
	"crawler/gbf"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func main() {
	db := database.InitDB()
	defer func() {
		_ = db.Close()
	}()
	pages := make(chan int, 100)
	ll := make(gbf.List, 0, 10000)
	now := time.Now()
	go func() {
		for {
			select {
			case page := <-pages:
				list := gbf.GetPage(page)
				ll = append(ll, list...)
				if len(ll) >= 10000 {
					fmt.Println("1w data: ", time.Since(now))
					SyncData(ll, db)
					ll = ll[0:0]
					now = time.Now()
				}
			}
		}
	}()
	now1 := time.Now()
	for i := 1; i <= 25000; i++ {
		pages <- i
	}
	fmt.Println(time.Since(now1))
}

func Addslashes(str string) string {
	var tmpRune []rune
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

func SyncData(list gbf.List, db *sql.DB) {
	now := time.Now()
	s := "insert into person_ranks_55_2 (`rank`, user_id, level, name, point, created_at) values"
	for _, item := range list {
		s = fmt.Sprintf("%s ('%s', '%s', '%s', '%s', '%s', '%s'),", s, item.Rank, item.UserID, item.Level, Addslashes(item.Name), item.Point, database.GetDateTime())
	}
	s = strings.TrimRight(s, ",")
	_, err := db.Exec(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("db insert:", time.Since(now))
}
