# Tuya Server API SDK for golang

## Initialize

```golang
import (
    tuyasdk "github.com/iot-eco-system/tuya-iot-service-sdk"
)

// Create API client
client := tuyasdk.NewTuyaAPIClient(tuyasdk.NewTuyaAPIClientOptions{
    Host: "https://openapi.tuyacn.com",
    ClientID: "<YOUR CLIENT ID>",
    Secret: "<YOUR SECRET>",
})

// Start the API client
client.Start()
```

## Invoke API

```golang
import (
    "github.com/iot-eco-system/tuya-iot-service-sdk/model"
)

// Sync user
response, err := client.UserSync(&model.UserSyncRequest{
    Schema: "<YOUR SCHEMA>",
    CountryCode: "China",
    Username: "Username",
    Password: "Password",
    UsernameType: 1,
    Nickname: "Nickname",
    TimeZoneID: "TimeZoneID",
})
if err != nil {
    log.Println(err)
}
```

## Stop the client

```golang
client.Stop()
```
