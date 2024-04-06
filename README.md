### Go implementation of LRU Cache with React App to track it

### Usage
Clone
> git clone https://github.com/SpoonBuoy/golru.git
>
Setup Backend
> cd golru <br> cd server <br>
go mod tidy <br>
go run main.go <br>
>
Setup Frontend
>cd golru <br> cd app <br>  npm install <br> npm start
>

Config (Change as per your convenience)

| Property | Defined In | Defined For |
|----------|----------|----------|
| SIZE    | server/main.go     | Capacity of LRU  | 
| PORT    | server/main.go    | Port on which server runs     |
|port | app/store.js | Port of backend | 
|api | app/store.js | Backend host url
