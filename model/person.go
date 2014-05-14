package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Person struct {
	Id    int    `json:"id" bson:"_id"`
	Phone string `json:"phone" bson:"phone"`
	Name  string `json:"name" bson:"name"`
}

func NewPerson(phone string, name string) *Person {
	p := &Person{CreateId(), phone, name}
	p.SetPerson()
	return p
}

func getPersonKey(Id int) string {
	return fmt.Sprintf("p:%d", Id)
}

func getAutoIdKey() string {
	return "autoIncrId"
}

func (p *Person) SetPerson() error {
	conn := redisPool.Get()
	defer conn.Close()

	key := getPersonKey(p.Id)
	json, _ := json.Marshal(p)
	conn.Send("SET", key, json)
	if _, err := conn.Do(""); err != nil {
		return err
	}
	return nil
}

func ExistPerson() {

}

func GetPerson(Id int) (result Person, err error) {
	conn := redisPool.Get()
	defer conn.Close()
	key := getPersonKey(Id)
	// conn.Send("GET", key)
	// var bytes []byte
	bytes, err := redis.Bytes(conn.Do("GET", key))
	fmt.Println(bytes, err)
	if err != nil {
		return result, err
	}
	var person Person
	json.Unmarshal(bytes, &person)
	return person, nil

}

func (p *Person) name() {

}

func CreateId() (id int) {
	conn := redisPool.Get()
	defer conn.Close()

	key := getAutoIdKey()
	// conn.Send("INCR", key)
	id, err := redis.Int(conn.Do("INCR", key))
	if err != nil {
		fmt.Println(err.Error())
		return id
	}

	return id
}

func GetDegreeOne() {

}

func (p *Person) SetDegreeOne(ids []int) (err error) {

	conn := redisPool.Get()
	defer conn.Close()

	key := getDegreeOneKey(p.Id)
	for id := range ids {

		conn.Send("ZADD", key, id, id)
	}
	if _, err = conn.Do(""); err != nil {
		return err
	}
	return nil
}

func (p *Person) GetDegreeOne() (ids []int, err error) {

	conn := redisPool.Get()
	defer conn.Close()

	key := getDegreeOneKey(p.Id)
	values, err := redis.Values(conn.Do("ZRANGE", key, 0, -1))
	if err != nil {
		return nil, err
	}
	redis.ScanSlice(values, &ids)
	return ids, nil
}

func (p *Person) SetDegreeTwo(ids []int) (err error) {

	conn := redisPool.Get()
	defer conn.Close()

	key := getDegreeTwoKey(p.Id)
	for id := range ids {

		conn.Send("ZADD", key, id, id)
	}
	if _, err = conn.Do(""); err != nil {
		return err
	}
	return nil
}

func (p *Person) GetDegreeTwo() (ids []int, err error) {

	conn := redisPool.Get()
	defer conn.Close()

	key := getDegreeTwoKey(p.Id)
	values, err := redis.Values(conn.Do("ZRANGE", key, 0, -1))
	if err != nil {
		return nil, err
	}
	redis.ScanSlice(values, &ids)
	return ids, nil
}

func getDegreeOneKey(id int) string {
	return fmt.Sprintf("d1:%d", id)

}

func getDegreeTwoKey(id int) string {
	return fmt.Sprintf("d2:%d", id)

}
