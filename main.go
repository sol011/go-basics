package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"sol011/go-basics/my-lib/lib1"
	"sort"
	"strings"

	"github.com/jackc/pgx/v4"
)

func fsToPostgres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://hoppscotch.io")
	fmt.Println(r.Method)
	fmt.Println(r.RemoteAddr)
	switch r.Method {
	case http.MethodPost:
		{
			var params = r.URL.Query()
			fmt.Println(params)
			var filePath = params["filePath"][0]
			var filename = strings.Split(filePath, ".")[0]
			var file, _ = os.Open("/mnt/games/go-web-app/" + filePath)

			var fileScanner = bufio.NewScanner(file)
			var fileContents []string
			for fileScanner.Scan() {
				fileContents = append(fileContents, fileScanner.Text())
			}
			fmt.Println(fileContents)
			sort.Strings(fileContents)
			fmt.Println(fileContents)

			var ctx = context.Background()
			var pgConn, _ = pgx.Connect(ctx, "postgres-url")
			defer pgConn.Close(ctx)

			var sqlStmt = "create table first_schema." + filename + "(contents text)"
			var rows, err = pgConn.Query(ctx, sqlStmt)
			if err != nil {
				fmt.Println(err.Error())
			}
			rows.Close()
			for _, line := range fileContents {
				sqlStmt = "insert into first_schema." + filename + " values ('" + line + "')"
				fmt.Println(sqlStmt)
				var rows, err = pgConn.Query(ctx, sqlStmt)
				rows.Close()
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("wrote " + strings.Join(fileContents, ".") + "to table"))
		}
	case http.MethodOptions:
		{
			w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, PUT")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusNoContent)
		}
	default:
		{
			w.Write([]byte("method " + r.Method + " does nothing at the moment"))
		}
	}
}

func startServer() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { rw.Write([]byte("index page")) })
	http.HandleFunc("/readandwritetopostgres", fsToPostgres)
	http.ListenAndServe(":8080", nil)
}

func main() {
	var sq = lib1.Square{S: 4}
	var geoSq = &sq
	var sqPrice = lib1.GetPrice(geoSq)
	fmt.Println(sqPrice)

	var circ = lib1.Circle{R: 10}
	var geoCirc = &circ
	var sqCirc = lib1.GetPrice(geoCirc)
	fmt.Println(sqCirc)

	fmt.Println("starting server")
	startServer()
}
