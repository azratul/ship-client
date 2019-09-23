# SHIP GO

Client for Salto Systems Servers via SHIP protocol.

# Usage

```golang
package main

import(
    "github.com/azratul/ship-client"
    "fmt"
)

func main() {
    salto := &ship.Client{
        Proto: "tcp",
        IP: "192.168.0.1",
        Port: "6500",
    }

    salto.Connect()
    salto.Write("<?xml version=\"1.0\" encoding=\"UTF-8\"?><RequestCall><RequestName> SaltoDBUserList.Read </RequestName><Params><MaxCount> 10 </MaxCount><ReturnAccessPermissions_User_Door> 1 </ReturnAccessPermissions_User_Door><ReturnAccessPermissions_User_Zone> 1 </ReturnAccessPermissions_User_Zone><ReturnAccessPermissions_User_Location> 1 </ReturnAccessPermissions_User_Location><ReturnAccessPermissions_User_Function> 1 </ReturnAccessPermissions_User_Function><ReturnAccessPermissions_User_Output> 1 </ReturnAccessPermissions_User_Output><ReturnMembership_User_Group> 1 </ReturnMembership_User_Group></Params></RequestCall>")
    fmt.Println(salto.Read())
}
```
