# Detect BOMB

[![CircleCI](https://circleci.com/gh/koron/detectbomb.svg?style=svg)](https://circleci.com/gh/koron/detectbomb)

Detects files which have BOMB (byte order mark bytes) as UTF-8 encoded text.

    # Install
    $ go install github.com/koron/detectbomb

    # How to run
    $ cd $GOPATH/github.com/koron/detectbomb
    $ ls testdata
    bomb.txt  nobomb.txt

    $ detectbomb
    BOMB: testdata/bomb.txt
    $
