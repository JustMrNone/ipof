# IPOF (IP of Domain)

IPOF is a simple Go project that retrieves the IPv4 and IPv6 addresses of a given domain name. It consists of two modules:
1. **domaintoip**: A library that provides functionality to resolve domain names to their IP addresses.
2. **ipof**: A command-line application that uses the `domaintoip` library to interact with users and display the resolved IP addresses.

## Features
- Resolves both IPv4 and IPv6 addresses for a given domain
- Interactive mode for entering domain names
- Command-line argument support for non-interactive usage

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/JustMrNone/ipof.git
   cd ipof/ipof
   ```

2. Build and Run the application:
   ```bash
   go build ipof.go
   
   # Interactive mode
   .\ipof.exe

   # With domain argument
   .\ipof.exe example.com
   ```

3. Install globally:
   ```bash
   set PATH=%PATH%;C:\path\to\your\install\directory
   # Adds the specified directory to the system PATH environment variable
   
   OR
   
   go env -w GOBIN=C:\path\to\your\bin
   # Sets the GOBIN environment variable to the specified directory

   go install
   
   ipof example.com
   ```

## Requirements
- Go 1.21 or higher

Enjoy!
