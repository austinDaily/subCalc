# CLI Subnet Calculator (Go)

A simple command-line Subnet Calculator written in Go.  
It takes a CIDR notation as input (`192.168.1.0/24`) and outputs useful subnetting information.

## ✨ Features

- Parses IPv4 CIDR blocks
- Displays:
  - Network Address
  - Subnet Mask
  - Prefix Length
  - Broadcast Address
  - Number of Usable Hosts

## 📦 Usage

```bash
go run main.go <CIDR>

