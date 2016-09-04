# Detect BOMB

Detects files which have BOMB (byte order mark bytes) as UTF-8 encoded text.

    # Install
    $ go install github.com/koron/detect-bomb

    # How to run
    $ cd $GOPATH/github.com/koron/detect-bomb
    $ detect-bomb
    BOMB: testdata/bomb.txt
