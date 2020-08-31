package main
import (
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/wsw364321644/go-botil"
	"net"
	"os"
	"strconv"
	"time"
)
var host = flag.String("host", "", "host")
var port = flag.String("port", "30001", "port")

type ModeType int
const (
	Mode_Client ModeType = iota
	Mode_Server
	Mode_End
)
func GetModeStr(mt ModeType) string{
	switch (mt){
	case Mode_Client:
		return "client"
	case Mode_Server:
		return "server"
	default:
		return ""
	}
}
func main() {
	flag.Parse()
	for mode:=Mode_Client;mode<Mode_End;mode++{
		modestr:=GetModeStr(mode)
		if(modestr!="") {
			fmt.Printf("%d-%s\n", mode,modestr )
		}
	}
	indexstr:=botil.CheckedScanfln("choose plat index:",func(input string)bool{
		i,err:=strconv.ParseInt(input,10,64)
		if(err==nil&& GetModeStr(ModeType(i))!=""){
			return true
		}
		return false
	})
	index,_:=strconv.ParseInt(indexstr,10,64)
	mode:=ModeType(index)
	switch(mode){
	case Mode_Client:
		StartClient()
	case Mode_Server:
		StartServer()
	}

}

func StartClient(){
	if *host == "" {
		fmt.Println("empty host")
		os.Exit(1)
	}
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil,addr)
	if err != nil {
		fmt.Println("Error connect:", err)
		os.Exit(1)
	}
	defer conn.Close()
	go func(){
		for {
			buffer := make([]byte, 1024)
			n, addr, err := conn.ReadFromUDP(buffer)
			if (err != nil) {
				fmt.Print(err)
			}
			fmt.Println("UDP Server : ", addr)

			fmt.Printf("Received from UDP server :%d-%s\n", binary.BigEndian.Uint32(buffer[:4]),string(buffer[4:n]))
		}
	}()
	for {
		str:=botil.ReadLine()
		conn.Write([]byte(str))
	}
}

func StartServer(){
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err.Error())
		return
	}
	daytime := time.Now().Unix()
	fmt.Println(n, remoteAddr)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(daytime))
	b=append(b, []byte(fmt.Sprintf("%d byte", n))...)
	conn.WriteToUDP(b, remoteAddr)
}