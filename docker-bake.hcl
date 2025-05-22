group "default" {
    targets = ["server"]
}

target "server" {
    context = "."
    dockerfile = "Dockerfile"
    args = {
        PORT = "8000"
    }
    tags = ["rafaeldev0ps/laboratory:http-server"]
}