package main

import (
     "fmt"
     "github.com/PuerkitoBio/goquery"
	 "github.com/BurntSushi/toml"
)

type Config struct {
    Server ServerConfig
    Uri    UriConfig
}

type ServerConfig struct {
    Host  string        `toml:"host"`
    Port  string        `toml:"port"`
}

type UriConfig struct {
     Path[]string `toml:"path"`
}

func GetPage(no int, url string) {
     doc, _ := goquery.NewDocument(url)
	 if doc != nil {
		 title := ""
		 h1 := ""
		 h2 := ""
		 if doc.Find("title") != nil {
			title = doc.Find("title").Text()
		 }
		 if doc.Find("h1") != nil {
			h1 = doc.Find("h1").Text()
		 }
		 if doc.Find("h2") != nil {
			h2 = doc.Find("h2").Text()
		 }
		 fmt.Printf("%d\t%s\t%s\t%s\t%s\n", 
			no + 1, url, title, h1, h2)
	}
}

func main() {
	var config Config
    _, err := toml.DecodeFile("config.tml", &config)
    if err != nil {
        panic(err)
    }

	for i, v := range config.Uri.Path {
		url := "http://" + config.Server.Host + ":" + config.Server.Port + v
		GetPage(i, url)
	}
}
