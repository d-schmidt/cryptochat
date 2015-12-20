package main

import (
    "bytes"
    "golang.org/x/crypto/sha3"
    "golang.org/x/net/websocket"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/rsa"
    "encoding/base64"
    "encoding/binary"
    "encoding/json"
    "log"
    "math/big"
    "bufio"
    "os"
    "time"
    "hash/crc32"
    "flag"
    "github.com/toqueteos/webbrowser"
    "net/http"
)

var (
    username string
    partnerName string
    serverWS *websocket.Conn
    privateKey *rsa.PrivateKey
    publicKeyCRC uint32
    partnerKey *rsa.PublicKey
    partnerKeyCRC uint32
    clientWS *websocket.Conn
    httpModeEnabled = false
    stop = true
    gotPKCh = make(chan bool)
    couldRegCh = make(chan bool)
)

const (
    CMD_MESSAGE = "msg"
    CMD_ERROR = "error"
    CMD_REGISTER = "reg"
    CMD_CONNECT = "con"
    CMD_SUCCESS = "success"
    CMD_PUBLICKEY = "pk"
    CMD_LIST = "list"
    CMD_USER_NOT_FOUND = "user not found"
    CMD_USER_DC = "user disconnected"
    CMD_NAME_IN_USE = "name in use"
    CMD_PING = "ping"
    CMD_PONG = "pong"
)

type Message struct {
    Target string `json:"Target,omitempty"`
    Name string `json:"Name,omitempty"`
    CMD string `json:"CMD,omitempty"`
    Key string `json:"Key,omitempty"`
    IV string `json:"IV,omitempty"`
    CT string `json:"CT,omitempty"`
    PK string `json:"PK,omitempty"`
    Part int `json:"Part,omitempty"`
    CRC uint32 `json:"CRC,omitempty"`
    ID int `json:"ID,omitempty"`
}


type FEMessage struct {
    Name string `json:"name,omitempty"`
    Target string `json:"target,omitempty"`
    Text string `json:"text,omitempty"`
    Type string `json:"type,omitempty"`
    Callback_id int `json:"callback_id,omitempty"`
}


func connectWS() {
    origin := "http://0.0.0.0/"
    url := "ws://0.0.0.0:17846/events/"
    if serverWS == nil {
        var err error
        serverWS, err = websocket.Dial(url, "", origin)
        if err != nil {
            log.Fatalf("connectWS() could not connect, error: %v", err)
        }
    }
}

func sendWS(data []byte) {
    log.Printf("sendWS() sending server message: %s", string(data))
    if _, err := serverWS.Write(data); err != nil {
        log.Fatalf("sendWS() could not send message to server, error: %v", err)
    }
}

func readWS() []byte {
    log.Printf("readWS() read started")
    var msg = make([]byte, 30 * 1024)
    n, err := serverWS.Read(msg)
    if err != nil {
        log.Fatalf("readWS() error while reading from server: %v", err)
    }
    raw := msg[:n]
    log.Printf("readWS() server message received: %s", string(raw))
    return raw
}


func baseEncode(data []byte) string {
    return string(base64.URLEncoding.EncodeToString(data))
}

func baseDecode(data string) []byte {
    result, _ := base64.URLEncoding.DecodeString(data)
    return result
}


func read_int32(data []byte) (ret int32) {
    buf := bytes.NewBuffer(data)
    binary.Read(buf, binary.LittleEndian, &ret)
    return
}

func write_int32(i int32) []byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, i)
    return buf.Bytes()
}


func encryptMessage(rsaKey *rsa.PublicKey, msg []byte) (encryptedAESKey []byte, nonce []byte, encryptedMsg []byte){

    random := rand.Reader
    // generate new aes key
    aeskey := make([]byte, 16)
    random.Read(aeskey)
    // generate new nonce
    nonce = make([]byte, 12)
    random.Read(nonce)
    // encrypt aes key
    encryptedAESKey, _ = rsa.EncryptOAEP(sha3.New512(), random, rsaKey, aeskey, nil)
    // encrypt msg
    aesblock, _ := aes.NewCipher(aeskey)
    aesgcm, _ := cipher.NewGCM(aesblock)
    encryptedMsg = aesgcm.Seal(nil, nonce, msg, nil)
    return
}

func decryptMessage(rsaKey *rsa.PrivateKey, encryptedAESKey []byte, nonce []byte, encryptedMsg []byte) (plaintext []byte) {

    random := rand.Reader

    aeskey, _ := rsa.DecryptOAEP(sha3.New512(), random, rsaKey, encryptedAESKey, nil)

    aesblock, _ := aes.NewCipher(aeskey)
    aesgcm, _ := cipher.NewGCM(aesblock)
    plaintext, _ = aesgcm.Open(nil, nonce, encryptedMsg, nil)
    return
}


func encodePublicKey(rsaKey *rsa.PublicKey) string {
    return baseEncode(append(write_int32(int32(rsaKey.E)), rsaKey.N.Bytes() ...))
}

func decodePublicKey(rsaKey string) *rsa.PublicKey {
    keyBytes := baseDecode(rsaKey)

    modu := new(big.Int).SetBytes(keyBytes[4:])
    newExpo := int(read_int32(keyBytes[0:4]))

    return &rsa.PublicKey{ modu, newExpo }
}


func encodeComMessage(encryptedAESKey []byte, nonce []byte, encryptedMsg []byte, name string, target string, crc uint32) []byte {
    msg := Message{}
    msg.Target = target
    msg.Name = name
    msg.CMD = CMD_MESSAGE
    msg.IV = baseEncode(nonce)
    msg.Key = baseEncode(encryptedAESKey)
    msg.CT = baseEncode(encryptedMsg)
    msg.CRC = crc
    b, _ := json.Marshal(msg)

    return b
}

func encodeRegMessage(name string, pubKey string, crc uint32, id int) []byte {
    msg := Message{}
    msg.Name = name
    msg.CMD = CMD_REGISTER
    msg.PK = pubKey
    msg.CRC = crc
    msg.ID = id
    b, _ := json.Marshal(msg)

    return b
}

func encodeConnectMessage(target string, id int) []byte {
    msg := Message{}
    msg.CMD = CMD_CONNECT
    msg.Target = target
    msg.ID = id
    b, _ := json.Marshal(msg)

    return b
}

func encodePingMessage() []byte {
    msg := Message{}
    msg.CMD = CMD_PING
    b, _ := json.Marshal(msg)

    return b
}

func encodeListMessage() []byte {
    msg := Message{}
    msg.CMD = CMD_LIST
    b, _ := json.Marshal(msg)

    return b
}


func decodeMessage(data []byte) Message {
    var m Message
    err := json.Unmarshal(data, &m)
    if err != nil {
        log.Fatalf("error while parsing json from server: %v", err)
        return Message{}
    }
    return m
}


func decodeFEMessage(data []byte) FEMessage {
    var m FEMessage
    err := json.Unmarshal(data, &m)
    if err != nil {
        log.Fatalf("error while parsing json from client: %v", err)
        return FEMessage{}
    }
    return m
}


func sendClientMessage(message FEMessage) {
    if (httpModeEnabled && clientWS != nil) {
        log.Printf("sendClientMessage() sending client message: %v", message)
        b, _ := json.Marshal(message)
        clientWS.Write(b)
    }
}


func asyncRead(ch chan []byte) {
    ch <- readWS()
}


func asyncWSReader() {
    if (! stop) {
        return
    }

    artSearchCh := make(chan []byte)
    stop = false
    defer func() {stop = true
    log.Print("asyncWSReader() stopped")
    }()

    go asyncRead(artSearchCh)

    for {
        select {
            case <-time.After(10 * time.Second):
                sendWS(encodePingMessage())
            case r := <- artSearchCh:
                msg := decodeMessage(r)
                log.Printf("asyncWSReader() server message received: %v", msg)

                if msg.CMD == CMD_MESSAGE {
                    if msg.CRC == publicKeyCRC {
                        data := decryptMessage(privateKey, baseDecode(msg.Key), baseDecode(msg.IV), baseDecode(msg.CT))
                        log.Printf("%s: %s", msg.Name, string(data))
                        sendClientMessage(FEMessage{
                            Name: msg.Name,
                            Text: string(data),
                            Type: CMD_MESSAGE})
                    } else {
                        log.Printf("ERROR %s: tried to send you a message encrypted with an old key", msg.Name)
                    }

                } else if msg.CMD == CMD_PUBLICKEY {
                    partnerKey = decodePublicKey(msg.PK)
                    partnerKeyCRC = crc32.ChecksumIEEE([]byte(msg.PK))
                    sendClientMessage(FEMessage{
                        Name: "local server",
                        Text: "found user, you can now chat",
                        Callback_id: msg.ID,
                        Type: CMD_SUCCESS})

                    if ! httpModeEnabled { gotPKCh <- true }

                } else if msg.CMD == CMD_USER_NOT_FOUND {
                    sendClientMessage(FEMessage{
                        Name: "local server",
                        Text: "chat partner is offline",
                        Type: CMD_USER_NOT_FOUND,
                        Callback_id: msg.ID })

                    if ! httpModeEnabled { gotPKCh <- false }

                } else if msg.CMD == CMD_USER_DC {
                    sendClientMessage(FEMessage{
                        Name: "www server",
                        Text: "chat partner went offline",
                        Type: CMD_USER_DC })

                    if ! httpModeEnabled { gotPKCh <- false }
                    
                } else if msg.CMD == CMD_NAME_IN_USE {
                    sendClientMessage(FEMessage{
                        Name: "www server",
                        Text: CMD_NAME_IN_USE,
                        Type: CMD_ERROR,
                        Callback_id: msg.ID })

                    if ! httpModeEnabled { couldRegCh <- false }

                }else if msg.CMD == CMD_REGISTER {
                    sendClientMessage(FEMessage{
                        Name: "www server",
                        Text: "registration successful",
                        Type: CMD_SUCCESS,
                        Callback_id: msg.ID })

                    if ! httpModeEnabled { couldRegCh <- true }

                } else if msg.CMD == CMD_PONG {
                    // silent
                } else {
                    log.Printf("asyncWSReader() unexpected server message: %s", msg.CMD)
                }

                go asyncRead(artSearchCh)
        }
    }
}


func cmdMode() {
    log.Print("starting command line mode")
    log.Print("what is your name?")
    bio := bufio.NewReader(os.Stdin)
    name, _, _ := bio.ReadLine()
    if len(name) == 0 {
        log.Fatal("you have to enter a name...")
        os.Exit(2)
    }
    username = string(name)

    log.Printf("welcome %s, trying to register you now", username)
    connectWS()
    go asyncWSReader()
    log.Print("connected with server")
    myPubKey := encodePublicKey(&privateKey.PublicKey)
    publicKeyCRC = crc32.ChecksumIEEE([]byte(myPubKey))
    sendWS(encodeRegMessage(username, myPubKey, publicKeyCRC, 0))
    couldRegister := <- couldRegCh

    if couldRegister {
        log.Print("you are now registered\n\nWho do you want to talk to?")
        name, _, _ = bio.ReadLine()
        if len(name) == 0 {
            log.Fatal("you have to enter a name...")
            os.Exit(2)
        }
        partnerName = string(name)

        sendWS(encodeConnectMessage(partnerName, 0))
        gotPK := <- gotPKCh

        for ! gotPK {
            log.Print("user not found\n\nWho do you want to talk to?")
            name, _, _ = bio.ReadLine()
            if len(name) == 0 {
                log.Fatal("you have to enter a name...")
                os.Exit(2)
            }
            partnerName = string(name)

            sendWS(encodeConnectMessage(partnerName, 0))
            gotPK = <- gotPKCh
        }

        log.Printf("got the key for %s, you can now talk", partnerName)
        var text, out, nonce, ct, data []byte

        for {
            text, _, _ = bio.ReadLine()
            if len(text) != 0 {
                out, nonce, ct = encryptMessage(partnerKey, text)
                data = encodeComMessage(out, nonce, ct, username, partnerName, partnerKeyCRC)
                sendWS(data)
            }
        }
    } else {
        log.Fatal("register failed")
    }
}


func WebsocketServer(ws *websocket.Conn) {

    if clientWS != nil {
        clientWS.Close()
        clientWS = nil
    }

    username = ""
    partnerName = ""

    clientWS = ws

    log.Print("WebsocketServer() new local connection established")
    buf := make([]byte, 32*1024)
    for {
        nr, er := clientWS.Read(buf)
        if nr > 0 {
            reqMsg := decodeFEMessage(buf[:nr])
            log.Printf("WebsocketServer() client message received: %v", reqMsg)

            if reqMsg.Type == CMD_REGISTER {
                sendClientMessage(FEMessage{
                    Name: "local server",
                    Text: "trying to register you",
                    Type: CMD_MESSAGE})

                username = reqMsg.Name
                connectWS()
                go asyncWSReader()

                myPubKey := encodePublicKey(&privateKey.PublicKey)
                publicKeyCRC = crc32.ChecksumIEEE([]byte(myPubKey))
                sendWS(encodeRegMessage(username, myPubKey, publicKeyCRC, reqMsg.Callback_id))
            } else if reqMsg.Type == CMD_CONNECT {
                sendClientMessage(FEMessage{
                    Name: "local server",
                    Text: "trying to find target user",
                    Type: CMD_MESSAGE})

                partnerName = reqMsg.Target
                sendWS(encodeConnectMessage(partnerName, reqMsg.Callback_id))
            } else if reqMsg.Type == CMD_MESSAGE {
                out, nonce, ct := encryptMessage(partnerKey, []byte(reqMsg.Text))
                data := encodeComMessage(out, nonce, ct, username, partnerName, partnerKeyCRC)
                sendWS(data)
            }
        }

        if er != nil {
            log.Printf("WebsocketServer() an error occured: %v", er)
            clientWS = nil
            break
        }
    }
}

func WebsocketServerFunc(w http.ResponseWriter, req *http.Request) {
    s := websocket.Server{Handler: websocket.Handler(WebsocketServer)}
    s.ServeHTTP(w, req)
}


func ChatHtmlServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("content-type", "text/html")
	w.Write(baseDecode(INDEX_HTML))
}

func ScrollGlueServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("content-type", "text/javascript")
	w.Write(baseDecode(SCROLLGLUE_JS))
}

func ChatAppJSServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("content-type", "text/javascript")
	w.Write(baseDecode(CHATAPP_JS))
}

func ChatAppCSSServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("content-type", "text/css")
	w.Write(baseDecode(CHATAPP_CSS))
}


func startServer(serverCh chan bool) {
	http.HandleFunc("/", ChatHtmlServer)
	http.HandleFunc("/scrollglue.js", ScrollGlueServer)
	http.HandleFunc("/chatapp.js", ChatAppJSServer)
	http.HandleFunc("/chatapp.css", ChatAppCSSServer)
	http.HandleFunc("/ws", WebsocketServerFunc)
    //http.Handle("/ws", websocket.Handler(WebsocketServer))

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Printf("startServer() ListenAndServe() error: %v", err)
        serverCh <- false
	}
    serverCh <- true
}


func httpMode() {
    log.Printf("httpMode() starting browser mode: http://127.0.0.1:12345/")
    serverCh := make(chan bool)
    go startServer(serverCh)
    webbrowser.Open("http://127.0.0.1:12345/")
    <- serverCh
}


func main() {

    tada := flag.Bool("cmd", false, "enable command line")
    flag.Parse()

    log.Print("main() starting chat")
    log.Print("main() generating rsa keys")
    privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
    privateKey.Precompute()

    if *tada {
        cmdMode()
    } else {
        httpModeEnabled = true
        httpMode()
    }
}