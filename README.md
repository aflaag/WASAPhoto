# WASAPhoto

## TODO list

### OpenAPI

- aggiungi getPhotos

### Backend

- FIXARE IL PROBLEMA DELLE PERSONE CHE VEDONO LE COSE CHE HANNO FATTO LORO STESSE
  - se ho messo like ad una foto, devo vederlo indipendentemente se l'autore della foto mi ha bannato o meno
  - se ho commentato una foto, devo vedere il mio commento indipendentemente se l'autore della foto mi ha bannato o meno
  - se seguo una persona, devo vederlo tra i miei seguiti indipendentemente se la persona che seguo mi ha bannato o meno
- vieta la possibilità di fare self-follow e self-ban
- ricontrolla components con struct
- ricontrolla tabelle del DB
- FIXA LA COSA DELL'UPDATE DELLO USERNAME CON LA UNIQUENESS
- generali
  - RILEGGI TUTTO QUANTO E RICONTROLLA TUTTO
  - commenti
- errori
  - gli errori vanno stampati nelle response o nel logger?
    - se no, cosa va messo nella response? un messaggio standard?
- fare l'handler della liveness

## Domande

- quale licenza va usata?
