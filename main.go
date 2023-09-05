package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"sync"
)

func main() {
	users := readExcel()
	url := "https://cloud.life.ccb.com/index_u.jhtml?param=4C4A647F371B43F11DBA91BA7799EE5886B689F42EBD4C9F151019803D21BFECB23869E598446758A237FF682F4BD607D1AABB8CEA42084E4CE8C5ADAC6473BDC47749CBBD30B2A06C82607E4D2CF04BD9ADC269DEFD0FE7D61058311BDAABD73F7CAFA0C143779B602AF2A9BE9743569C714A5ABD64A6560B60B92AFEA11F2C0D75995186952173"
	season := Season{Code: "202302", Name: "2023年二季度"}

	//demo := User{Id: "513424198510160066", Name: "龚黎"}
	//result := scanner(url, demo, season)
	//fmt.Printf("%s\t%s:\t%f\n", demo.Id, demo.Name, result)
	wg := sync.WaitGroup{}
	async := 6

	broadcast := make(chan User, async)
	f := func() {
		user := <-broadcast
		if len(user.Id) == 0 {
			return
		}
		wg.Add(1)
		resultTemp := scanner(url, user, season)
		defer func() {
			wg.Done()
			fmt.Printf("【结果】%s\t%s:\t%f\n", user.Id, user.Name, resultTemp)
		}()
	}
	for i := 0; i < async; i++ {
		go func() {
			for {
				f()
			}
		}()
	}
	for _, user := range users {
		broadcast <- user
	}
	wg.Wait()
}

func readExcel() []User {
	excelFileName := "23-2.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}

	var persons []User
	for _, sheet := range xlFile.Sheets {
		for index, row := range sheet.Rows {
			if len(row.Cells) < 2 || index == 0 {
				continue
			}
			person := User{
				Name: row.Cells[0].Value,
				Id:   row.Cells[1].Value,
			}
			//fmt.Printf("%d:%v\n", index, person)
			persons = append(persons, person)
		}
	}
	//fmt.Println(persons)
	return persons
}
