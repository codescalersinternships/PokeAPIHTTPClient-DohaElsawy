
# Poke API HTTP Client Package 
Create an HTTP client in Go that consumes the Poke APIs. https://pokeapi.co/docs/v2#info

## Installation 
- 1. Download project
```golang
   git get https://github.com/codescalersinternships/PokeAPIHTTPClient-DohaElsawy.git
```
- 2. import package :
```golang
   import pokeClient "github.com/codescalersinternships/PokeAPIHTTPClient-DohaElsawy.git"
```
### Functions
- 1 -  get a resource list:
     - directly passing your values 
       ```golang
       //                           "/pokemon"    3       2
       resource, err := GetResource(endpoint, offset, limit) 
       ```
     - load from .env, use following env valraibles:
       ```
       ENDPOINT=/pokemon
       OFFSET=4
       LIMIT=3
       ```
       ```golang
       endpoint , params , err := LoadConfigFromENV("../testdata/.env")
       //                           "/pokemon" offset=3  limit=4
       resource, err := GetResource(endpoint, params[0], params[1])
       ```
> [!NOTE]
> `endpoint` is the only mandatiry field

- 2 - get pokemon by id:
  ```golang
   pokemon, err := GetPokemon(id)
  ```

### Test
- to run all tests
```golang
  make test
```
### Format
- format all files inside project
```golang
  make test
```
