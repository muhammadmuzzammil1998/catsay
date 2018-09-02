class Catsay < Formula
  desc "A program that generates pictures of a cat holding a sign with a message."
  homepage "https://github.com/muhammadmuzzammil1998/catsay"
  url "https://codeload.github.com/muhammadmuzzammil1998/catsay/tar.gz/v1.0"
  sha256 "3b31056aa635b7209df84c900af2800292f9626d56c74bfc61f260708e7a0ebd"
  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin / "catsay"
  end

  test do
    system "catsay"
  end
end
