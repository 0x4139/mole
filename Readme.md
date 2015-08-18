# Mole

Mole is an http/s proxy server written in go that isn't transparent (doesn't send any kind of headers to the destination)

# Usage

```go
./mole -h
-ip="127.0.0.1": ip that the proxy server should use
-port="5445": the port that the proxy server should use
-verbose=true: if the proxy server should be verbose or not
```

#Example

```go
./mole -ip=8.8.8.8 -port=5445 -verbose=false // will start a proxy on the ip 8.8.8.8 and the port 5445
```

#Licence
MIT Vali Malinoiu
