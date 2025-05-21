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
		//TODO: Pass request object instead content

		if request.Path == "/" {
			controller.DefaultController(conn, content)

		} else if request.Path == "/user-agent" {
			controller.UserAgentController(conn, content)

		} else if strings.Contains(request.Path, "/echo/") {
			controller.EchoController(conn, content)

		} else if strings.Contains(request.Path, "/files/") {
			controller.FilesController(conn, content)

		} else {
			res.Status404(conn)
		}

		if request.RequestHeaders["Connection"] == "close" {
			defer conn.Close()
			break
		}
	}
}
