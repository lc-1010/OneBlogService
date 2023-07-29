package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lc-1010/OneBlogService/chatroom/server"
)

var (
	addr   = ":2023"
	banner = `
	 
                                    : o o ;    
                                    : \_'/       
                                   '_|__|_'       
                                   /_____\          
                                   |||||||       
                                 _'' '  ' ''_       
                            _.-'_:______/:_'-._    
                     _..-'    [______/]    '-.._       
                 _-'                   [ ]              '-_       
            _.='  [#]                   [#]                 '=._          
                          [#]                     [#]           
                               |                      |     
                                \                    /           
                                 \    \__/\__/    /             
                                   '-._|##|##|_.-'            
                                        |##|##|              
                                      _/##|##|\_          
                                     /____/|__|\           
                                    |    /___\|##|      
                                   ||   \#|##/  ||   
                                  |||   \#/   |||
	`
)

func main() {

	fmt.Printf(banner+"\n", addr)
	server.RegisterHandle()
	log.Fatal(http.ListenAndServe(addr, nil))
}
