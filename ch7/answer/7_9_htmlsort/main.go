package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length int
}

var tracks = []Track{
	{Title: "Song A", Artist: "Artist A", Album: "Album A", Year: 2021, Length: 180},
	{Title: "Song B", Artist: "Artist B", Album: "Album B", Year: 2020, Length: 240},
	{Title: "Song C", Artist: "Artist C", Album: "Album C", Year: 2022, Length: 200},
}

func main() {
	http.HandleFunc("/", handleTracks)
	log.Fatal(http.ListenAndServe(":8080", nil)) //浏览器中访问 ：http://localhost:8080/
}

// 只是按照最近点击的字段进行排序
func handleTracks(w http.ResponseWriter, r *http.Request) {
	sortBy := r.URL.Query().Get("sort")

	switch sortBy {
	case "title":
		sort.Slice(tracks, func(i, j int) bool {
			return tracks[i].Title < tracks[j].Title
		})
	case "artist":
		sort.Slice(tracks, func(i, j int) bool {
			return tracks[i].Artist < tracks[j].Artist
		})
	case "album":
		sort.Slice(tracks, func(i, j int) bool {
			return tracks[i].Album < tracks[j].Album
		})
	case "year":
		sort.Slice(tracks, func(i, j int) bool {
			return tracks[i].Year < tracks[j].Year
		})
	case "length":
		sort.Slice(tracks, func(i, j int) bool {
			return tracks[i].Length < tracks[j].Length
		})
	}
	dir, _ := os.Getwd()
	println(dir)
	tmpl := template.Must(template.ParseFiles(dir + "/ch7/answer/7_9_htmlsort/tracks.html"))
	err := tmpl.Execute(w, tracks)
	if err != nil {
		log.Println(err)
	}
}
