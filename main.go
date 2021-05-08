package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("/Users/igorpecenikin/Spam базы/icann/ru_ns.csv")
	if err != nil {
		log.Fatalln(err)
	}
	tpl := []string{}
	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			continue
		}
		count := parts[0]
		label := strings.ToTitle(strings.TrimRight(parts[1], "."))
		tpl = append(tpl, fmt.Sprintf(`{ label: "%s", y: %s }`, label, count))
	}
	r := strings.NewReplacer("${tpl}", strings.Join(tpl[len(tpl)-100:], ",\n"))

	index, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile("result.html", []byte(r.Replace(string(index))), 0655)
}
