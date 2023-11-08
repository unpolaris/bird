package gconst

import "time"

var (
	BirdListPrefix = "BIRD_LIST"
	BirdListExpire = time.Minute * 10
)

var (
	BirdInfoPrefix = "BIRD_INFO"
	BirdInfoExpire = time.Minute * 10
)
