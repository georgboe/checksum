# Checksum

Generates SHA256 hashes for files and signs the hashes with your GPG key.    
[Inspired by the Fedora spins checksums](https://spins.fedoraproject.org/static/checksums/Fedora-Live-x86_64-23-CHECKSUM)

## Installation

    $ go get github.com/georgboe/checksum

## Usage

Run `checksum` with the filenames you want to be checksummed. A file `CHECKSUM` will be created with the clearsigned output.

    $ checksum ubuntu-16.04-beta2-desktop-amd64.iso
    Enter password: 
    -----BEGIN PGP SIGNED MESSAGE-----
    Hash: SHA256

    SHA256 (ubuntu-16.04-beta2-desktop-amd64.iso) = 643aa06bcb025c31aaf111247ff2d2608ce1ecaa1d3b1178929376b38f650be4
    -----BEGIN PGP SIGNATURE-----

    wsFcBAEBCAAQBQJXF/N3CRDPSrTfcRRoXQAAeg4QADn3igz+8cmOnAJcu3+cTdO1
    ZMbQxAKqm0qEn20gbygUEvUS/nG8Ieml/Uv2rbwxLsecwbuCzXjTqxAh2xTWx3in
    Vis/fJsOncB7wGmZksCYcJ+TDN+YrF7xt8bNxLlCmomyYgCFhlNG7GffbfMTmu+o
    X2kJXsrEySStiPCYroZOZloDMZWi8RrPRNCTsc/hhdhBBblgN8ZiBQItpOSQysyd
    qRG9488wr6BwvMG/IftNCnr9QyEJs0uxWRhkGdeCTtrWahYp+5I92WSm0ECbSdxy
    5AUIACwSjRQRLDoGFQt/sp5Wac368Zr8tWYHtMX3CT3BJUzvdg3prmTNtrBfTN28
    Fr0y1ZoHOSGFSk6UhOACPO2TY/KrxsDgJIzPPUiws0Uav4Aj66gGocpTIGzrkQQb
    hl+qwOr7Oq+b5jL4Iy9cC1MlWyPVbOd3JxMmKvYVst2fBMm8TpCUWTmQtxBJht6n
    daNjXnKmTuvSS5ozZ4wyVo+eQKGUwqjg09BfOkghb62Nu4X0/6KFIV9Tj5p6tQJ3
    yU39CqOWAHQlFOyic+tNCt6+5TKv57DoBPzaWzKoq1ZVYYttdOEKWjYuQIfewCKW
    BUk7TpsKqUizItpxLdgDn+uhk7sM2Qgg8YI4GEm7EfeFBn1zq6+yeVE91Wa8Tewp
    Ff/yn/44V4aN9Qpmp9hC
    =Ep6w
    -----END PGP SIGNATURE-----

## License

MIT License

Copyright (c) 2016 Georg Bøe

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

