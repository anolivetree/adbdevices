adbdevices
==========

It's cumbersome to set `ANDROID_SERIAL` if you have multiple android devices connected.
You can easily switch `ANDROID_SERIAL` with this program.

Setup
============

This program is written in Go. You need to install Go first.

1. install Go
2. $ go build
3. Write following line in your .bashrc.

    alias ad='export ANDROID_SERIAL=`adbdevices`'

How to use
===========

    user@ubuntu:~/project/go/adbdevices$ ad
    1) device 0146910xxxxxxxxx
    which device?: 1
    user@ubuntu:~/project/go/adbdevices$ 

