package db

//db global storage in memory
var db = make(map[int64]Beer)

//id auto incremental value
var index int64 = 0
