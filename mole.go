package main
import "flag"
import "fmt"
import "github.com/elazarl/goproxy"
import "log"
import "net/http"

func main()  {
  proxyIP := flag.String("ip", "127.0.0.1", "ip that the proxy server should use")
  proxyPort := flag.String("port", "5445", "the port that the proxy server should use")
  proxyVerbose:= flag.Bool("verbose",true,"if the proxy server should be verbose or not")
  flag.Parse()
  proxy := goproxy.NewProxyHttpServer()
  proxy.Verbose = *proxyVerbose
  fmt.Println("Starting server on:",fmt.Sprint(*proxyIP,":",*proxyPort))
  log.Fatal(http.ListenAndServe(fmt.Sprint(*proxyIP,":",*proxyPort), proxy))
}
