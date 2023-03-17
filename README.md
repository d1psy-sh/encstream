# encstream

Stream encrypting/encoding in Go (a PoC)

## TODO

- there is one char at the end of a client message that gets stuck in the stream
    - client "hello" -> last char gets stuck -> server "hell"
    - next enter server "o"
    - fix this!!!
