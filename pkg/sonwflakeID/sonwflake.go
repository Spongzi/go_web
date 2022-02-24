package sonwflakeID

import (
	"fmt"
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

const TimeFormat = "2006-01-02"

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse(TimeFormat, startTime)
	if err != nil {
		fmt.Println("parse time failed!", err)
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
