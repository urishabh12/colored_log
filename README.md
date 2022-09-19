# Colored Log
🟥🟧🟨🟩🟦🟪 Colored logging in golang

## Example
```
import (
	log "github.com/urishabh12/colored_log"
)

func main() {
    log.Println("Default Log")
    log.Success("Success")
    log.Panic("Panic/Fatal Log")
}
```