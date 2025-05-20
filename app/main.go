package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"http-server/app/controller"
	"http-server/app/request"
	"http-server/app/response"
	"http-server/app/utils"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")

	address := "0.0.0.0:4221"

	listn, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer listn.Close()

	for {
		conn, err := listn.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		content, err := utils.ReadRequestContent(conn)
		if err != nil {
			fmt.Printf("error reading request content: %v\n", err)
			os.Exit(1)
		}

		// in case "Connection: close" header is not set -> if there's nothing else, close conn
		if len(content) == 1 {
			defer conn.Close()
			break
		}

		var req request.Request
		var res response.Response

		path := req.GetPath(content)

		if path == "/" {
			controller.DefaultController(conn, content)

		} else if path == "/user-agent" {
			controller.UserAgentController(conn, content)

		} else if strings.Contains(path, "/echo/") {
			controller.EchoController(conn, content)

		} else if strings.Contains(path, "/files/") {
			controller.FilesController(conn, content)

		} else {
			res.Status404(conn)
		}

		if req.GetHeaderValue(req.GetHeaders(content), "Connection") == "close" {
			defer conn.Close()
			break
		}
	}
}
