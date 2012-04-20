This is a TCP/IP debugger/proxy allowing to intercept the network traffic.
It can serve multiple connections in parallel.

Prerequisites
=============

This program is written in Go and requires at least the release 1.

Usage
=====

For example:

    go run gotcpspy.go -host=ftp.idsoftware.com -port 21 -listen_port=8081

It starts running with the message:

    Start listening on port 8081 and forwarding data to ftp.idsoftware.com:21

In a separate console you can run:

    telnet localhost 8081

and enter something, for example, `USER test` `<ENTER>`,
`PASS test@test.org` `<ENTER>` and finally `QUIT` `<ENTER>`.

It should produce three logs.

Bidirectional dump
------------------

`log-2012.04.20-12.10.12-0001-10.44.2.21-26311-192.246.40.185-21.log`

    Connected to ftp.idsoftware.com:21 at 2012.04.20-12.10.12
    Received (#0, 00000000) 24 bytes from 10.44.2.21-26311
    00000000  32 32 30 2d 66 74 70 2e  69 64 73 6f 66 74 77 61  |220-ftp.idsoftwa|
    00000010  72 65 2e 63 6f 6d 0d 0a                           |re.com..|
    Sent (#0) to 127.0.0.1-8081
    Received (#1, 00000018) 353 bytes from 10.44.2.21-26311
    00000000  32 32 30 2d 2d 2d 2d 2d  2d 2d 2d 2d 2d 2d 2d 2d  |220-------------|
    00000010  2d 2d 2d 2d 2d 2d 2d 2d  2d 2d 2d 2d 2d 2d 2d 2d  |----------------|
    00000020  2d 0d 0a 32 32 30 2d 57  65 6c 63 6f 6d 65 20 74  |-..220-Welcome t|
    00000030  6f 20 66 74 70 2e 69 64  73 6f 66 74 77 61 72 65  |o ftp.idsoftware|
    00000040  2e 63 6f 6d 0d 0a 32 32  30 2d 2d 2d 2d 2d 2d 2d  |.com..220-------|
    00000050  2d 2d 2d 2d 2d 2d 2d 2d  2d 2d 2d 2d 2d 2d 2d 2d  |----------------|
    00000060  2d 2d 2d 2d 2d 2d 2d 0d  0a 32 32 30 2d 0d 0a 32  |-------..220-..2|
    00000070  32 30 2d 43 6f 6e 6e 65  63 74 69 6f 6e 20 66 72  |20-Connection fr|
    00000080  6f 6d 20 38 30 2e 31 36  39 2e 33 34 2e 31 39 34  |om 80.169.34.194|
    00000090  20 6c 6f 67 67 65 64 0d  0a 32 32 30 2d 59 6f 75  | logged..220-You|
    000000a0  20 61 72 65 20 75 73 65  72 20 32 30 20 6f 66 20  | are user 20 of |
    000000b0  31 35 30 20 61 76 61 69  6c 61 62 6c 65 20 63 6f  |150 available co|
    000000c0  6e 6e 65 63 74 69 6f 6e  73 2e 0d 0a 32 32 30 2d  |nnections...220-|
    000000d0  0d 0a 32 32 30 2d 41 76  65 72 61 67 65 20 74 68  |..220-Average th|
    000000e0  72 6f 75 67 68 70 75 74  20 66 6f 72 20 74 68 69  |roughput for thi|
    000000f0  73 20 73 65 72 76 65 72  20 69 73 20 32 34 38 2e  |s server is 248.|
    00000100  38 38 37 20 4b 42 70 73  2e 0d 0a 32 32 30 2d 35  |887 KBps...220-5|
    00000110  35 37 34 20 70 65 6f 70  6c 65 20 68 61 76 65 20  |574 people have |
    00000120  76 69 73 69 74 65 64 20  74 68 69 73 20 73 69 74  |visited this sit|
    00000130  65 20 69 6e 20 74 68 65  20 6c 61 73 74 20 32 34  |e in the last 24|
    00000140  20 68 6f 75 72 73 2e 0d  0a 32 32 30 2d 0d 0a 32  | hours...220-..2|
    00000150  32 30 2d 0d 0a 32 32 30  2d 0d 0a 32 32 30 20 0d  |20-..220-..220 .|
    00000160  0a                                                |.|
    Sent (#1) to 127.0.0.1-8081
    Received (#0, 00000000) 1 bytes from 127.0.0.1-8081
    00000000  55                                                |U|
    Sent (#0) to 10.44.2.21-26311
    Received (#1, 00000001) 1 bytes from 127.0.0.1-8081
    00000000  53                                                |S|
    Sent (#1) to 10.44.2.21-26311
    Received (#2, 00000002) 10 bytes from 127.0.0.1-8081
    00000000  45 52 20 61 6e 6f 6e 79  6d 6f                    |ER anonymo|
    Sent (#2) to 10.44.2.21-26311
    Received (#3, 0000000C) 2 bytes from 127.0.0.1-8081
    00000000  75 73                                             |us|
    Sent (#3) to 10.44.2.21-26311
    Received (#4, 0000000E) 2 bytes from 127.0.0.1-8081
    00000000  0d 0a                                             |..|
    Sent (#4) to 10.44.2.21-26311
    Received (#2, 00000179) 70 bytes from 10.44.2.21-26311
    00000000  33 33 31 20 55 73 65 72  20 6e 61 6d 65 20 6f 6b  |331 User name ok|
    00000010  61 79 2c 20 70 6c 65 61  73 65 20 73 65 6e 64 20  |ay, please send |
    00000020  63 6f 6d 70 6c 65 74 65  20 45 2d 6d 61 69 6c 20  |complete E-mail |
    00000030  61 64 64 72 65 73 73 20  61 73 20 70 61 73 73 77  |address as passw|
    00000040  6f 72 64 2e 0d 0a                                 |ord...|
    Sent (#2) to 127.0.0.1-8081
    Received (#5, 00000010) 1 bytes from 127.0.0.1-8081
    00000000  50                                                |P|
    Sent (#5) to 10.44.2.21-26311
    Received (#6, 00000011) 3 bytes from 127.0.0.1-8081
    00000000  41 53 53                                          |ASS|
    Sent (#6) to 10.44.2.21-26311
    Received (#7, 00000014) 2 bytes from 127.0.0.1-8081
    00000000  20 74                                             | t|
    Sent (#7) to 10.44.2.21-26311
    Received (#8, 00000016) 1 bytes from 127.0.0.1-8081
    00000000  65                                                |e|
    Sent (#8) to 10.44.2.21-26311
    Received (#9, 00000017) 1 bytes from 127.0.0.1-8081
    00000000  73                                                |s|
    Sent (#9) to 10.44.2.21-26311
    Received (#10, 00000018) 2 bytes from 127.0.0.1-8081
    00000000  74 40                                             |t@|
    Sent (#10) to 10.44.2.21-26311
    Received (#11, 0000001A) 6 bytes from 127.0.0.1-8081
    00000000  6e 61 6d 65 2e 6f                                 |name.o|
    Sent (#11) to 10.44.2.21-26311
    Received (#12, 00000020) 2 bytes from 127.0.0.1-8081
    00000000  72 67                                             |rg|
    Sent (#12) to 10.44.2.21-26311
    Received (#13, 00000022) 2 bytes from 127.0.0.1-8081
    00000000  0d 0a                                             |..|
    Sent (#13) to 10.44.2.21-26311
    Received (#3, 000001BF) 30 bytes from 10.44.2.21-26311
    00000000  32 33 30 20 55 73 65 72  20 6c 6f 67 67 65 64 20  |230 User logged |
    00000010  69 6e 2c 20 70 72 6f 63  65 65 64 2e 0d 0a        |in, proceed...|
    Sent (#3) to 127.0.0.1-8081
    Received (#14, 00000024) 1 bytes from 127.0.0.1-8081
    00000000  51                                                |Q|
    Sent (#14) to 10.44.2.21-26311
    Received (#15, 00000025) 3 bytes from 127.0.0.1-8081
    00000000  55 49 54                                          |UIT|
    Sent (#15) to 10.44.2.21-26311
    Received (#16, 00000028) 2 bytes from 127.0.0.1-8081
    00000000  0d 0a                                             |..|
    Sent (#16) to 10.44.2.21-26311
    Received (#4, 000001DD) 14 bytes from 10.44.2.21-26311
    00000000  32 32 31 20 47 6f 6f 64  62 79 65 21 0d 0a        |221 Goodbye!..|
    Sent (#4) to 127.0.0.1-8081
    Disconnected from 10.44.2.21-26311
    Disconnected from 127.0.0.1-8081
    Finished at 2012.04.20-12.10.12, duration 16.5769481s
    
Outgoing binary log
-------------------

`log-binary-2012.04.20-12.10.12-0001-10.44.2.21-26311.log`

    USER anonymous
    PASS test@name.org
    QUIT

Incoming binary log
-------------------

`log-binary-2012.04.20-12.10.12-0001-192.246.40.185-21.log`

    220-ftp.idsoftware.com
    220------------------------------
    220-Welcome to ftp.idsoftware.com
    220------------------------------
    220-
    220-Connection from 80.169.34.194 logged
    220-You are user 20 of 150 available connections.
    220-
    220-Average throughput for this server is 248.887 KBps.
    220-5574 people have visited this site in the last 24 hours.
    220-
    220-
    220-
    220 
    331 User name okay, please send complete E-mail address as password.
    230 User logged in, proceed.
    221 Goodbye!
