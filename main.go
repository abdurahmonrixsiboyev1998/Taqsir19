// package main

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "14022014"
// 	dbname   = "demo"
// )

// type Record struct {
// 	Id        int
// 	Generated int
// }

// func main() {
// 	num := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)
// 	n, err := sql.Open("postgres", num)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer n.Close()

// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		fmt.Println("Operation complete")
// 	}()

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println(ctx.Err())
// 	}

// 	x, err := n.QueryContext(ctx, "SELECT ID, generated FROM large_dataset")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer x.Close()

// 	var nums []Record
// 	for x.Next() {
// 		var r Record
// 		if err := x.Scan(&r.Id, &r.Generated); err != nil {
// 			log.Fatal(err)
// 		}
// 		nums = append(nums, r)
// 	}

// 	if err := x.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Record retrieved: ", len(nums))

// }

package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	x, err := sql.Open("postgres", "postgres://postgres:14022014@localhost/demo?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer x.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `SELECT id, generated FROM large_dataset`

	number, err := x.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer number.Close()

	for number.Next() {
		var id, generated int
		err := number.Scan(&id, &generated)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Generated: %d\n", id, generated)
	}

	if err = number.Err(); err != nil {
		panic(err)
	}
}
