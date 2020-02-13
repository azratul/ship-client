package ship

import(
    "bytes"
    "net"
    "strconv"
)

/*
Client : Estructura básica de un cliente para el
         protocolo SHIP de Salto Server
Attrs :
    Socket : Conexión activa a Salto Server
    Proto : Protocolo TCP o UDP
    IP : IP o DNS de Salto Server
    Port : Puerto en escucha de Salto Server
*/
type Client struct {
    Socket net.Conn
    Proto  string
    IP     string
    Port   string
}

/*
Connect : Función para Conectar a Salto.
*/
func (client *Client) Connect() {
    conn, error := net.Dial(client.Proto, client.IP + ":" + client.Port)
    if error != nil {
        panic(error)
    }

    client.Socket = conn
}

/*
Write : Función escribir en el socket server de Salto.
*/
func (client *Client) Write(request string) {
    length := strconv.Itoa(len(request))
    data := "STP/00/" + length + "/" + request
    client.Socket.Write([]byte(data))
}

/*
Read : Función para leer la respuesta del socket server de Salto.
*/
func (client *Client) Read() string {
    bs := make([]byte, 1)

    cont := 0
    size := ""

    for {
        _, err := client.Socket.Read(bs)
        if err != nil {
            panic(err)
        } else {
            if string(bs) == "/" {
                cont++

                if cont == 2 {
                    continue
                }

                if cont == 3 {
                    break
                }
            }

            if cont == 2 {
                size += string(bs)
            }
        }
    }

    s, _ := strconv.Atoi(size)
    brk  := 0
    xml  := ""

    bs = make([]byte, s)

    for {
        ln, err := client.Socket.Read(bs)

        if err != nil {
            panic(err)
        }

        brk += ln
        xml += string(bytes.Trim(bs, "\x00"))

        if brk == s {
            break
        }

        bs = make([]byte, (len(bs)-ln))
    }

    return xml
}