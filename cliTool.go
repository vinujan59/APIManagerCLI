package main

import (
	"os"

	"github.com/urfave/cli"
	"gopkg.in/resty.v0"
	"time"
	"fmt"
	"net/http"
	"encoding/json"
	"crypto/tls"
)

type LoginResponse struct {
	Error bool
	Message string
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
				//fmt.Println("User login : ", c.Args().First(), c.String("username"), " ", c.String("password"))
				resp, err := resty.R().SetQueryParams(map[string]string{
					"action":"login",
					"username":username,
					"password":password,
				}).Post("http://localhost:9763/publisher/site/blocks/user/login/ajax/login.jag")

				if err!=nil {fmt.Printf("\nError: %v \n", err)}

				//fmt.Println(resp)
				var loginResponse LoginResponse
				unMarshalError := json.Unmarshal([]byte(resp.String()), &loginResponse)
				if unMarshalError != nil{
					fmt.Println(unMarshalError)
				}else{
					//fmt.Println(loginResponse)
				}
				if loginResponse.Error {
					fmt.Println("login Failed!")
				}else {
					cookies = resp.Cookies()
					fmt.Println("login success")
				}
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
			Name:    "export",
			Aliases: []string{"e"},
			Usage:   "exporting api as zip file",
			Action:  func(c *cli.Context) error {
				fmt.Println("Exporting api....!")
				resty.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
				resp,err := resty.R().SetQueryParams(map[string]string{
					"name":c.String("apiname"),
					"version":c.String("apiversion"),
					"provider":c.String("apiprovider"),
				}).SetBasicAuth("admin","admin").
				Get("https://localhost:9443/api-import-export-2.0.0-v0/export-api")


				if err!=nil {
					fmt.Printf("Error is %v",err)
				}else{
					fmt.Printf("Resposne %v \n",resp)
				}
				return nil
			},
			Flags:[]cli.Flag{
				cli.StringFlag{
					Name:"apiname",
					Value:"",
					Usage:"Name of the API",
				},
				cli.StringFlag{
					Name:"apiprovider",
					Value:"",
					Usage:"API provider username",
				},
				cli.StringFlag{
					Name:"apiversion",
					Value:"",
					Usage:"API version",
				},
				cli.StringFlag{
					Name:"exportedapiname",
					Value:"",
					Usage:"Name of exporting file with path",
				},
			},
		},
		{
			Name:    "import",
			Aliases: []string{"i"},
			Usage:   "importing api as zip file",
			Action:  func(c *cli.Context) error {
				fmt.Println("Importing api....!")
				resty.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
				resp,err := resty.R().SetBasicAuth("admin","admin").
				SetFile(c.String("filePath")).
				Post("https://localhost:9443/api-import-export-2.0.0-v0/import-api")


				if err!=nil {
					fmt.Printf("Error is %v",err)
				}else{
					fmt.Printf("Resposne %v \n",resp)
				}
				return nil
			},
			Flags:[]cli.Flag{
				cli.StringFlag{
					Name:"filePath",
					Value:"",
					Usage:"importing file path",
				},
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