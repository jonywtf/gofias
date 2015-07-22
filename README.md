Базовый код для чтения КЛАДР из формата DBF (формат старой версии, упрощенной структуры).
Архив Base.7z на странице http://www.gnivc.ru/inf_provision/classifiers_reference/kladr/
Код примера чтения в файле __kladr.go__ можно вызвать следующим образом:
```Go
package main

import (
    "github.com/jonywtf/gofias"
)

func main() {
    err := fias.Import()
    if err != nil {
        println(err)
    }
}

```
