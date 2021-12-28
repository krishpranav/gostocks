class Gostocks < Formula
    homepage "https://github.com/krishpranav/gostocks"
    url "https://github.com/krishpranav/gostocks.git"

    def install
        system "go mod tidy", "go build -o gostocks main.go"
        puts "Installed successfully :)"
    end
end
