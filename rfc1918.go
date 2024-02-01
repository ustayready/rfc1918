package main

import (
    "bufio"
    "flag"
    "fmt"
    "net"
    "os"
)

func isRFC1918(ip net.IP) bool {
    return ip.IsPrivate()
}

func main() {
    inputFile := flag.String("file", "", "Input file containing a list of IP addresses")
    flag.Parse()

    if *inputFile == "" {
        fmt.Println("Please provide a valid input file using the --file flag.")
        return
    }

    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        ipAddress := scanner.Text()
        ip := net.ParseIP(ipAddress)

        if ip != nil && !isRFC1918(ip) {
            fmt.Println(ip)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
}

