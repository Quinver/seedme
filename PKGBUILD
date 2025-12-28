pkgname=seedme
pkgver=0.1.0
pkgrel=1
pkgdesc="CLI Streaming tool"
arch=('x86_64')
license=('MIT')
source=("git+https://github.com/Quinver/seedme.git")
sha256sums=('SKIP')

makedepends=('go')

build() {
  cd "$srcdir/$pkgname"
  go build -o seedme ./cmd/seedme
}

package() {
  cd "$srcdir/$pkgname"
  install -Dm755 seedme "$pkgdir/usr/bin/seedme"
}

