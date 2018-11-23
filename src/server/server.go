package server

import (
	"net"
	"net/http"
	"os"

	"Smilo-blackbox/src/server/api"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

var (
	privateAPI *mux.Router
	publicAPI  *mux.Router

	sockPath = os.TempDir() + "/blackbox.sock"

	serverStatus               = status{false, false}
	log          *logrus.Entry = logrus.WithField("package", "server")
)

type status struct {
	httpServer bool
	unixServer bool
}

// SetLogger set the logger
func SetLogger(loggers *logrus.Entry) {
	log = loggers.WithFields(log.Data)

	filenameHook := filename.NewHook()

	logrus.AddHook(filenameHook)

}

func NewServer(Port string) (*http.Server, *http.Server) {
	publicAPI, privateAPI = InitRouting()

	return &http.Server{
			Addr:    ":" + Port,
			Handler: publicAPI,
		},
		&http.Server{
			Handler: privateAPI,
		}

}

func StartServer(Port string, sockFilePath string) {

	if sockFilePath != "" {
		sockPath = sockFilePath
	}
	defer os.Remove(sockPath)

	glog.Info("Starting server")
	pub, priv := NewServer(Port)
	glog.Info("Server starting --> " + Port)
	sock, _ := net.Listen("unix", sockPath)
	glog.Info("Unix Domain Socket Up --> " + sockPath)
	go func() {
		err := gracehttp.Serve(
			pub)
		if err != nil {
			glog.Error("Error: %v", err)
			os.Exit(1)
		}
	}()
	err := priv.Serve(sock)
	if err != nil {
		glog.Error("Error: %v", err)
		os.Exit(1)
	}
}

func InitRouting() (*mux.Router, *mux.Router) {

	publicAPI := mux.NewRouter()
	privateAPI := mux.NewRouter()

	publicAPI.HandleFunc("/version", api.GetVersion).Methods("GET")
	publicAPI.HandleFunc("/upcheck", api.Upcheck).Methods("GET")

	publicAPI.HandleFunc("/push", api.Push).Methods("POST")
	publicAPI.HandleFunc("/resend", api.Resend).Methods("POST")
	publicAPI.HandleFunc("/partyinfo", api.GetPartyInfo).Methods("GET")
	publicAPI.HandleFunc("/delete", api.Delete).Methods("POST")
	publicAPI.HandleFunc("/receiveraw", api.ReceiveRaw).Methods("GET")
	publicAPI.HandleFunc("/transaction/{key:.*}", api.TransactionDelete).Methods("DELETE")
	publicAPI.HandleFunc("/config/peers", api.ConfigPeersPut).Methods("PUT")
	publicAPI.HandleFunc("/config/peers/{index:.*}", api.ConfigPeersGet).Methods("GET")
	publicAPI.HandleFunc("/metrics", api.Metrics).Methods("GET")

	privateAPI.HandleFunc("/version", api.GetVersion).Methods("GET")
	privateAPI.HandleFunc("/upcheck", api.Upcheck).Methods("GET")
	privateAPI.HandleFunc("/sendraw", api.SendRaw).Methods("POST")
	privateAPI.HandleFunc("/send", api.Send).Methods("POST")
	privateAPI.HandleFunc("/receive", api.Receive).Methods("GET")
	privateAPI.HandleFunc("/transaction/{hash:.*}", api.TransactionGet).Methods("GET")

	return publicAPI, privateAPI
}
