# Mensajeria con formato ISO-8583 con GOLANG

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

_Aplicación Servidor para manejo y procesamiento de transferencias electrónicas de POS a Server Gateway con uso de ISO-8583_

_La presente aplicación actua como intermediaria entre las terminales POS y la Banca (Server Gateway) para mensajeria_

_El servidor recive JSON de terminales o puntos de venta, procesa dicha data y la envia en formato ISO-8583 hacia Server Gateway_

### Requisitos para correr el proyecto 📋

_El presente proyecto corre con la version go version go1.19.1_

_Comandos DOCKER_

```shell
// Para levantar la imagne de docker
docker-compose up

// Para eliminar la imagen de docker
docker-compose down
```

_Comando correr servidor_

```shell
go run .
```

Simular mensajes hacia Servidor desde Cliente HTTP mediante HTTP requests como "POSTMAN"(Terminal)\_
_Para simular requestPayment se debe realizar mediante el metodo POST a la URL "localhost:3000/requestPayment" con el siguiente JSON de prueba_

```json
{
  "entry_mode": "magnetic_stripe",
  "device": {
    "serialnr": "PB04204S60977",
    "version": "100",
    "counter": 1
  },
  "key": {
    "check_value": "CAA9BO",
    "crc32": "BF8425A0",
    "name": "A000BZPY72",
    "rsa": "8481B7D10576D49CE1AACACF284B13256D8313A104F9C68434E3A931759F659917BD7434198F5A358DCEF0F615FD6D84332710C30ABCD050C5D96752658AB02AFAAF053B3F0B9997DD02D12472B06EC9F0A8F5E740486E875F467572E39C0FC386EDBD882D624273B1AA44945942BAC597CC333CACB6C334743E5495E708A9B10D1A3461ED58F5A48000A1862DAB658B8B4A6F0BB45617E5AAA0538A12F776A9CC752BC8070802ECB5388A3C14D811D318378E13639DB7E96BD3824C127BC5B224A137470DCF7547961EA344B5138898302F04327B3846AA586E036AF8F17FEFF061A431223B15A8E185B85343D1A2387B27"
  },
  "localtime": "2022-04-22 14:13:00",
  "amount": "100.00"
}
```

_De la petición anterior el Server responde el siguiente JSON_

```json
{
  "servertime": "Thu Sep 15 2022 20:11:42 GMT-0300 (hora estándar de Argentina)",
  "rc": -1,
  "rcmessage": "bla bla bla",
  "id": 12,
  "workkey": {
    "ksn": "00000101403388200001",
    "key": "AEA40693C054BE3A65E14D572C58EAB9",
    "crc32": "BF8425A0",
    "check_value": "CAA9B0"
  }
}
```

_Simular mensajes hacia Servidor desde Cliente HTTP mediante HTTP requests como "POSTMAN"(Terminal)_
_Para simular executePayment se debe realizar mediante el metodo POST a la URL "localhost:3000/executePayment" con el siguiente JSON de prueba_
_Notar que el "id" del JSON es el mismo "id" del JSON de respuesta del Servidor_

```json
{
  "id": 12,
  "device": {
    "authentication": "pin",
    "failcounter": 0,
    "realcounter": 8,
    "ksn": "00000045467101200008"
  },
  "cardInformation": {
    "bin": "425982",
    "cvv_length": 0,
    "cvv_present": false,
    "emv_tags": "9F2701809F26088AD7DDD181D667D29F3704F100D6A39F360200759C0100820239009F3303E0B8C89F34034403029A032101145F2A0204849F02060000000010009F03060000000000009F3501225F3401019F10120310A74005020000000000000000000000FF8407A00000000410109F090200029F1A020484950500000080009F1E0830303030303930359F6E009F0607A00000000410109F3901055A08520416567447099657105204165674470996D2505201000002269F21031603569F6C005F201a50455246494C455320415A554C2F5052472020202020202020205F24032505318F01069F5D009F4104000000509F53009F7C004F07A000000004101050104465626974204D6173746572636172649B02E8008E1400000000000000004201440341031E0342031F03",
    "cardholder": "NOMBRE DEL TARJETAHABIENTE ",
    "last4": "1111",
    "track2": "DDD0BD616C1DCA0B24DD0D71FF10A78D680BABD638037F3F",
    "track2_crc": "BF4B379D",
    "track2_length": 32
  }
}
```

_De la petición anterior el Server responde el siguiente JSON_

```json
{
  "id": 11,
  "timestamp": "Thu Sep 15 2022",
  "rc": 0,
  "rcdatetime": "0803",
  "rcmessage": "APROBADA",
  "ticket": 11,
  "authorization": "552285",
  "keys_expired": false
}
```

## Construido con 🛠️

_Herramientas usadas en el proyecto:_

- Golang
- gorilla/mux
- gorm.io
- Sockets
- MySQL
- Docker

## Wiki 📖

Puedes encontrar mucho más de cómo se usa el protocolo ISO-8583 este proyecto en [Wiki](https://es.wikipedia.org/wiki/ISO_8583)
