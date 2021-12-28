class Gostocks < Formula
    homepage "https://github.com/krishpranav/gostocks"
    url "https://github.com/krishpranav/gostocks/releases/download/v1/gostocks.zip"
    sha256 "e66f8ed63d038e9da952e36afea137ad984f77a77cefe61dbe9509b10e1391c2"

    def install
        system "./gostocks"
        puts "Installed successfully :)"
    end
end
