// Code generated by Thrift Compiler (0.18.0). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"shared"
	"strconv"
	"strings"
	"tutorial"
)

var _ = shared.GoUnusedProtection__
var _ = tutorial.GoUnusedProtection__

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  void ping()")
	fmt.Fprintln(os.Stderr, "  i32 add(i32 num1, i32 num2)")
	fmt.Fprintln(os.Stderr, "  i32 calculate(i32 logid, Work w)")
	fmt.Fprintln(os.Stderr, "  void zip()")
	fmt.Fprintln(os.Stderr, "  SharedStruct getStruct(i32 key)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
	var m map[string]string = h
	return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
	parts := strings.Split(value, ": ")
	if len(parts) != 2 {
		return fmt.Errorf("header should be of format 'Key: Value'")
	}
	h[parts[0]] = parts[1]
	return nil
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	headers := make(httpHeaders)
	var parsedUrl *url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
	flag.Parse()

	if len(urlString) > 0 {
		var err error
		parsedUrl, err = url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	var cfg *thrift.TConfiguration = nil
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
		if len(headers) > 0 {
			httptrans := trans.(*thrift.THttpClient)
			for key, value := range headers {
				httptrans.SetHeader(key, value)
			}
		}
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans = thrift.NewTSocketConf(net.JoinHostPort(host, portStr), cfg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransportConf(trans, cfg)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := tutorial.NewCalculatorClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "ping":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Ping requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Ping(context.Background()))
		fmt.Print("\n")
		break
	case "add":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Add requires 2 args")
			flag.Usage()
		}
		tmp0, err17 := (strconv.Atoi(flag.Arg(1)))
		if err17 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		tmp1, err18 := (strconv.Atoi(flag.Arg(2)))
		if err18 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.Add(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "calculate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Calculate requires 2 args")
			flag.Usage()
		}
		tmp0, err19 := (strconv.Atoi(flag.Arg(1)))
		if err19 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		arg20 := flag.Arg(2)
		mbTrans21 := thrift.NewTMemoryBufferLen(len(arg20))
		defer mbTrans21.Close()
		_, err22 := mbTrans21.WriteString(arg20)
		if err22 != nil {
			Usage()
			return
		}
		factory23 := thrift.NewTJSONProtocolFactory()
		jsProt24 := factory23.GetProtocol(mbTrans21)
		argvalue1 := tutorial.NewWork()
		err25 := argvalue1.Read(context.Background(), jsProt24)
		if err25 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Calculate(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "zip":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Zip requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Zip(context.Background()))
		fmt.Print("\n")
		break
	case "getStruct":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetStruct requires 1 args")
			flag.Usage()
		}
		tmp0, err26 := (strconv.Atoi(flag.Arg(1)))
		if err26 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetStruct(context.Background(), value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
