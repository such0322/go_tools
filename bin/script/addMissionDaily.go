package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	TM "odin_tool/models/twod/master"
	"runtime"
	"strings"
	"time"
)

//addMissionDaily 追加日常任务
//params toDate string 追加到指定的日期
func addMissionDaily(toDate string) {
	var err error
	var toTime time.Time
	if strings.ContainsRune(toDate, '/') {
		toTime, err = time.Parse("2006/01/02", toDate)
	} else if strings.ContainsRune(toDate, '-') {
		toTime, err = time.Parse("2006-01-02", toDate)
	} else {
		log.Fatal(errors.New("无效的时间参数"))
	}
	if err != nil {
		fmt.Println(err)
	}

	ml := TM.MissionLetter{}
	mls := ml.GetLastDaily()
	mls.LoadReward()
	if !mls[0].SessionFrom.Valid {
		log.Fatal(errors.New("无有效的复制对象"))
	}
	lastTime, err := time.Parse("2006-01-02 15:04:05", mls[0].SessionFrom.String)
	if err != nil {
		log.Fatal(err)
	}
	copyDay := int(math.Ceil(toTime.Sub(lastTime).Hours() / 24))
	ch := make(chan int, copyDay)
	defer close(ch)
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 1; i <= copyDay; i++ {
		go mls.CopyAndInsert(i, ch)
	}
	for i := 0; i < copyDay; i++ {
		<-ch
	}

}
