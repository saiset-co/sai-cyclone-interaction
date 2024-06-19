# saiCycloneInteraction

Http proxy to create transaction in the CYCLONE SDK based applications.

## Configurations
**config.yml** - common saiService config file.

### Common block
- `http` - http server section
    - `enabled` - enable or disable http handlers
    - `port`    - http server port

### Cyclone block
- `node_address` - Cyclone node address for API calls
- `wallet` - wallet address
- `private` - wallet private key to sign transactions

### Other block
- `crypto_address` - sai-btc service address to sign transactions

## How to run
`make build`: rebuild and start service  
`make up`: start service  
`make down`: stop service  
`make logs`: display service logs

## API
### Send transaction
```json lines
{
  "method": "send_tx",
  "data": "$message"
}
```
#### Params
`$message` <- raw transaction message
