# Assumptions

1. The Server and Client are connected on the same local network.
2. The port info is easy for both server and client to know.
3. Only Client and server are present on the local network. (not necessary but does make configuring easier)

> NOTE: I have only tested this code with an Android Phone and a Linux Computer as server and client (and vice-versa) respectively.

# Steps

1. Use Phone Hotspot and connect computer to it. (Making them both on the same local network)
2. Use USB Debugging on Phone and connect phone to Computer.
3. Write your go code for client and server.
4. The ip address to use can be found using ip address of Phone or Computer (use ip of whichever device you are using as server).
	4.1. To get ip address of both phone and computer. Write `ifconfig` on computer to get computer's ip
	4.2. To get phone's ip addr use `arp -a` to get phone's ip address. (Assuming no other device is connected to the mobile hotspot).
	4.3 Use a port between 1025 and 60000. These are all unpriviledged ports and hence shoudnt cause problems in using.
5. Use `GOARCH=arm64 GOOS=linux go build -o logproxyservice main.go` to compile the golang code which needs to go to the Android Phone.
6. Use `adb devices` in terminal to see if the phone is connected via USB to computer or not. [This will only work if USB Debugging is turned ON in the Android phone]
7. Use `adb push <<NAME_OF_BINARY_GENERATED_BY_GOLANG>> /data/local/tmp/` to push the binary to the Android Phone.
8. Use `adb shell` to open remote shell of Android Phone on Computer.
9. In shell write commands `cd /data/local/tmp/` then `ls` then `./<<NAME_OF_BINARY_GENERATED_BY_GOLANG>>`
10. In Computer side run the other script (Server/Client) by using command `go run .` [Do this in a DIFFERENT terminal]
