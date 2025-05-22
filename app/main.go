package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"
	"regexp"

	"http-server/app/controller"
	"http-server/app/request"
	"http-server/app/response"
	"http-server/app/utils"
)

type Controller func(request request.Request) response.Response

type Route map[string]Controller

var routes Route = map[string]Controller{
	"/":           controller.DefaultController,
	"/user-agent": controller.UserAgentController,
	"/echo/*":     controller.EchoController,
	"/files/*":    controller.FilesController,
}

// simple validation, enhance in the future
func selectRoutePath(path string) (Controller, error) {
	regexRule, _ := regexp.Compile("^" + path)
	for route, ctrl := range routes {
		if regexRule.MatchString(route) {
			return ctrl, nil
		}
	}
	return nil, fmt.Errorf("route not found")
}

func handleConnection(conn net.Conn) {
	for {
		content, err := utils.ReadRequestContent(conn)
		if err != nil {
			slog.Error("error reading request content: %v", "error", err)
			os.Exit(1)
		}

		// in case "Connection: close" header is not set -> if there's nothing else, close conn
		if len(content) == 1 {
			defer conn.Close()
			break
		}

		var req request.Request
		var res response.Response

		request := req.ParseRequest(content)

		ctrl, err := selectRoutePath(request.Path)
		if err != nil {
			res.StatusCode = response.HTTP404
		}

		if ctrl != nil {
			res = ctrl(request)
		}

		conn.Write(res.ParseReponse())

		if request.Header["Connection"] == "close" {
			defer conn.Close()
			break
		}
	}
}

func main() {
	port := flag.String("port", "4221", "HTTP server port")
	flag.Parse()

	address := "0.0.0.0:" + *port

	slog.Info("Server started on 0.0.0.0", "port", *port)

	listn, err := net.Listen("tcp", address)
	if err != nil {
		slog.Error("Failed to bind to", "address", address)
		os.Exit(1)
	}
	defer listn.Close()

	for {
		conn, err := listn.Accept()
		if err != nil {
			slog.Error("Error accepting connection: %v", "error", err)
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}
