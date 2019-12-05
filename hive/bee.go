package hive

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Bee struct {
	url string
	id  int
}

func NewBee(url string, id int) *Bee {
	b := new(Bee)
	b.url = url
	b.id = id
	return b
}

func (b *Bee) Id() int { return b.id }

func (b *Bee) FindNectar() *Dance {
	req, err := http.NewRequest("GET", b.url, nil)
	//req.SetBasicAuth(USERNAME, PASSWORD)
	cli := &http.Client{}
	rs, err := cli.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer rs.Body.Close()

	nectar, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}
	return NewDance(b, len(nectar), rs.StatusCode)
}
