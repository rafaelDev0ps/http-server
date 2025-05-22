group "default" {
    targets = ["server"]
}

target "server" {
    context = "."
    dockerfile = "Dockerfile"
    args = {
        PORT = "8000"
    }
    platforms = ["linux/amd64", "linux/arm64"]
    tags = ["rafaeldev0ps/laboratory:http-server"]
}