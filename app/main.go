package main

import (
	"log/slog"
	"net"
	"os"
	"strings"

	"http-server/app/controller"
	"http-server/app/request"
	"http-server/app/response"
	"http-server/app/utils"
)

func main() {
	address := "0.0.0.0:4221"

	slog.Info("Server started", "address", address)

	listn, err := net.Listen("tcp", address)
	if err != nil {
		slog.Error("Failed to bind to address %s", "error", address)
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

		if request.Path == "/" {
			res = controller.DefaultController(request)

		} else if request.Path == "/user-agent" {
			res = controller.UserAgentController(request)

		} else if strings.Contains(request.Path, "/echo/") {
			res = controller.EchoController(request)

		} else if strings.Contains(request.Path, "/files/") {
			res = controller.FilesController(request)

		} else {
			res.StatusCode = "404 Not Found"
		}

		conn.Write(res.ParseReponse())

		if request.Header["Connection"] == "close" {
			defer conn.Close()
			break
		}
	}
}
