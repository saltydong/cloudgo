package server
import (
    "github.com/codegangsta/martini" 
)
func NewServer(port string) {   
    m := martini.Classic()

    m.Get("/", func(params martini.Params) string {
        return "bingo"
    })

    m.RunOnAddr(":"+port)   
}