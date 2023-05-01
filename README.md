# demo-go-rpc
Playground to test how to structure a project with net/rpc from std lib.


# example


        $> ./run.sh 

        Launching server
        Setting initial ticker at  10s
        Started server
        Listening on  :1234
        Launching client 1
        Launching client 2
        xavier@ubuntuXPS:~/Bureau/demo-go-rpc$ Start Client
        Connecting to  :1234
        {7 8}  --->  {0 7}
        Changing ticker to  10µs
        .....{7 8}  --->  {0 7}
        ..2023/05/01 11:11:03 Hello world !
        Did you see the print ?
        .{50 8}  --->  {6 2}
        ...........Changing ticker to  100ms
        killing server in  2s
        Start Client
        Connecting to  :1234
        {7 8}  --->  {0 7}
        Changing ticker to  10µs
        ......................................{7 8}  --->  {0 7}
        2023/05/01 11:11:03 Hello world !
        ....Did you see the print ?
        ............{50 8}  --->  {6 2}
        ...Changing ticker to  100ms
        .killing server in  2s
        ...................Server stopped :  accept tcp [::]:1234: use of closed network connection

