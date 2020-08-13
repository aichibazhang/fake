package util

import (
	"math/rand"
	"time"
)

func Seed(seed int64)  {
	if seed==0{
		rand.Seed(time.Now().UnixNano())
	}else{
		rand.Seed(seed)
	}
}