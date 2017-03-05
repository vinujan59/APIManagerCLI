package main

import (
	"os"

	"github.com/urfave/cli"
	"gopkg.in/resty.v0"
	"time"
	"fmt"
	"net/http"
	"encoding/json"
)

type LoginResponse struct {
	error bool `json:"error"`
}

func main() {
	var username string
	var password string
	var cookies []*http.Cookie

	app := cli.NewApp()
	app.Name = "WSO2 API Manager CLI"
	app.Usage = "[API] ---> Import || Export || Creation || Publishing || Subcription ||  Key Generation"
	app.UsageText = app.Usage
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Vinujan Shanagr",
			Email: "vinujan59@gmail.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "login",
			Aliases: []string{"l"},
			Usage:   "User login",
			Action:  func(c *cli.Context) error {
				//if(username)
				fmt.Println("User login : ", c.Args().First(), c.String("username"), " ", c.String("password"))
				resp, err := resty.R().SetQueryParams(map[string]string{
					"action":"login",
					"username":username,
					"password":password,
				})./*SetHeader("accept","").*/
					Post("http://localhost:9763/publisher/site/blocks/user/login/ajax/login.jag")
				fmt.Printf("\nResponse Body: %v", resp)
				loginResponse := LoginResponse{}
				json.Unmarshal(resp.Body(), loginResponse)
				if (loginResponse.error) {
					fmt.Println("login Failed!")
				}else {
					cookies = resp.Cookies()
					fmt.Println("login success")
				}
				fmt.Printf("\nError: %v", err)
				return nil
			},
			Flags : []cli.Flag{
				cli.StringFlag{
					Name:"username,u",
					Value:"",
					Usage:"Credential for api manager",
					Destination:&username,
				}, cli.StringFlag{
					Name:"password,p",
					Value:"",
					Usage:"Credential for api manager",
					Destination:&password, //hidden is used for hiding from help
				},
			},

		},
		{
			Name:    "import",
			Aliases: []string{"i"},
			Usage:   "add a task to the list",
			Action:  func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		//{
		//	Name:        "template",
		//	Aliases:     []string{"t"},
		//	Usage:       "options for task templates",
		//	Subcommands: []cli.Command{
		//		{
		//			Name:  "add",
		//			Usage: "add a new template",
		//			Action: func(c *cli.Context) error {
		//				fmt.Println("new task template: ", c.Args().First())
		//				return nil
		//			},
		//		},
		//		{
		//			Name:  "remove",
		//			Usage: "remove an existing template",
		//			Action: func(c *cli.Context) error {
		//				fmt.Println("removed task template: ", c.Args().First())
		//				return nil
		//			},
		//		},
		//	},
		//},
	}

	app.Run(os.Args)
}