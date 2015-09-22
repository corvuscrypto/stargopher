#Packet Notes
##Chat
For chat received (id: 5) server -> client
message has two parts:

1. name of sender [a,b...] a is number of chars, b... are chars
2. message [a,b...] a is number of chars, b... are chars

chat sent (id:14) client -> server

Chat channel seems to be in the message as well.

##VLQ
VLQ also seems to be a bit weirdly done. Probably Unity variant?
* negative VLQs seem to point to compressed data and the payload seems to have a proper zlib header...
