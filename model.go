package main

// 人的结构体
type Person struct {
    Name string `form:"firstname" binding:"required"`
    Lastname  string `form:"lastname"`
}

