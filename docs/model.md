

# model
`import "Smilo-blackbox/src/server/model"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type PeerNode](#PeerNode)


#### <a name="pkg-files">Package files</a>
[p2p_headers.go](/src/Smilo-blackbox/src/server/model/p2p_headers.go) [peernode.go](/src/Smilo-blackbox/src/server/model/peernode.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    //NETWORK_LIST p2p message type
    NETWORK_LIST = "NETWORK_LIST"
    //NETWORK_STATE p2p message type
    NETWORK_STATE = "NETWORK_STATE"
    //REQUEST_NET_STATE p2p message type
    REQUEST_NET_STATE = "REQUEST_NET_STATE"
    //PEER p2p message type
    PEER = "PEER"
    //GET_PEER p2p message type
    GET_PEER = "GET_PEER"
    //PEER_LIST p2p message type
    PEER_LIST = "PEER_LIST"
    //GET_PEER_LIST p2p message type
    GET_PEER_LIST = "GET_PEER_LIST"
    //MESSAGE p2p message type
    MESSAGE = "MESSAGE"

    //COMMIT p2p message type
    COMMIT = "COMMIT"
    //APPROVE p2p message type
    APPROVE = "APPROVE"
    //DECLINE p2p message type
    DECLINE = "DECLINE"
    //BLOCK p2p message type
    BLOCK = "BLOCK"
    //GET_BLOCK p2p message type
    GET_BLOCK = "GET_BLOCK"
    //TRANSACTION p2p message type
    TRANSACTION = "TRANSACTION"
)
```



## <a name="PeerNode">type</a> [PeerNode](/src/target/peernode.go?s=883:1094#L24)
``` go
type PeerNode struct {
    ID            string    `json:"id" storm:"id"`
    LastSeen      time.Time `json:"last_seen"`
    NetworkStatus string    `json:"network_status"`
    RemoteAddr    string    `json:"remote_addr"`
}

```
PeerNode Peers contains the peer info














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
