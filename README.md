IP Address Analyzer
This focuses on analyzing IPv4 CIDR blocks, calculating network details, and checking IP address membership—key and networking fundamentals.

Project Summary
This tool:
- Parses a given CIDR block
- Outputs the network address, broadcast address, subnet mask, and number of usable hosts
- Optionally, checks whether a given IP address belongs to the CIDR block

Build Instructions
To build the binary:
bash: go build -o bin/ipanalyzer src/ipanalyzer.go

Usage 
CIDR Analysis:./bin/ipanalyzer 192.168.1.0/24
IP Membership Check:./bin/ipanalyzer 192.168.1.0/24 192.168.1.25

Example Output 
Analyzing network: 192.168.1.0/24

Network address: 192.168.1.0
Broadcast address: 192.168.1.255
Subnet mask: 255.255.255.0
Number of usable hosts: 254

