#Stargopher - the golang starbound server

This project will be an attempt to make a more functional server for starbound. This project is similar in spirit to StarryPy which is a fantastic server and you should definitely check it out here on github!

##TODO

* ~~make functional pass-through TCP proxy~~
* ~~Organize data stream into packets (High Priority)~~
* ~~Add zlib decompression for packets with zlib headers (High Priority)~~
* Determine packet behavior (Medium Priority)
* Cleanup handling of packets with compression
  * refactor code to accommodate re-compression if data in compressed packet
    was manipulated
