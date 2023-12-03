# WASAPhoto

## TODO list

### OpenAPI

- controlla le response di errore se possono effettivamente essere ritornate

### Backend

- controlla che chi fa l'operazione non sia stato bannato quando
  - follow
  - remove follow
  - get followers list
  - get following list
  - get user profile
  - like
  - remove like
  - get photo like count
  - get photo comment count
  - get photo count
  - comment
  - remove comment
  - get database stream
- generali
  - gestisci le PUT bene
  - ricontrollare
    - response
    - errori
    - ogni possibile cosa che può fallire di ogni request
  - metti la data al commento (anche nell'api)
  - riformatta tutte le query
- errori
  - gli errori vanno stampati nelle response o nel logger?
    - se no, cosa va messo nella response? un messaggio standard?
- fare l'handler della liveness

## Domande

- quale licenza va usata?
