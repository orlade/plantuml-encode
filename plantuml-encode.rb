# This file was generated by GoReleaser. DO NOT EDIT.
class PlantumlEncode < Formula
  desc "Encodes PlantUML source to send to the server."
  homepage ""
  version "0.1.0"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/orlade/plantuml-encode/releases/download/v0.1.0/plantuml-encode_0.1.0_darwin-amd64.tar.gz"
    sha256 "ba0e7ceb481516ce062de889f332725653ac69aa59cba2bcd4fc36c52fe09f11"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/orlade/plantuml-encode/releases/download/v0.1.0/plantuml-encode_0.1.0_linux-amd64.tar.gz"
      sha256 "b43b7bbb489a3823dddaf426c7da8119c46afee2a21421486aea2445e4762761"
    end
  end

  def install
    bin.install "plantuml-encode"
  end

  test do
    system "#{bin}/plantuml-encode --version"
  end
end
