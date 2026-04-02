class Chewiecli < Formula
  desc "A CLI Tool to interface with the chewie github app (github.com/swirldslabs/chewie)"
  homepage "https://github.com/PandasWhoCode/chewie-cli"
  url "SUB_URL"
  sha256 "SUB_SHA256"
  license "Apache-2.0"

  depends_on "node"

  def install
    system "npm", "install", *std_npm_args
    bin.install_symlink Dir["#{libexec}/bin/*"]
  end

  test do
     assert_match "chewie", shell_output("#{bin}/chewie --help 2>&1", 0)
  end
end
