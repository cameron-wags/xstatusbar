pkgname=xstatusbar
pkgver=1.1.2
pkgrel=1
epoch=
pkgdesc="CLI tool which manages packages from the Arch User Repository & other git repos."
arch=("any")
url="https://github.com/cameron-wags/xstatusbar"
license=("MIT")
depends=("acpi"
         "xorg-xsetroot"
         "xssstate"
         "pulseaudio")
makedepends=("go>=1.15"
             "make")
source=("$pkgname-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz")
md5sums=("SKIP")

build() {
    cd "$pkgname-$pkgver"
    make build
}

package() {
    cd "$pkgname-$pkgver"
    install -Dm 755 "$pkgname" -t "$pkgdir/usr/bin/"
}
