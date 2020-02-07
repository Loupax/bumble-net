A proof of concept on how we can create a 'soft' parental lock, which slows downs network traffic that comes from specific internet hosts. 

In this example I slow down all incoming traffic from facebook.com:443

Running:
Compile `main.go` and place the WinDivert files in the same folder as the executable according to the [WinDivert documentation](https://www.reqrypt.org/windivert-doc-1.4.html#installing)

Run the executable and notice how traffic related to facebook is slowed down to a single packer per 255 milliseconds

Known problems: 

1. We need to make a DNS lookup per new IP address being checked. This could be resolved with a library that uses BPF filters that allow for packet filtering by the hostname directly.
    
2. Needs elevated priviledges in order to run. [gopacket](https://github.com/google/gopacket) seems to be able to sniff packets without them but couldn't figure out how to actually block the traffic yet.
    
3. Works with version 1.4 of WebDivert. Not exactly a problem, just a caveat.
    
4. Only works on Windows.

Must see if I can recreate this functionality using gopacket instead
