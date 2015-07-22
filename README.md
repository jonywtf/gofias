Базовый код для чтения КЛАДР из формата DBF (формат старой версии, упрощенной структуры).
Архив Base.7z на странице http://www.gnivc.ru/inf_provision/classifiers_reference/kladr/
Код примера чтения в файле __kladr.go__ можно вызвать следующим образом:
```Go
package main

import (
    "github.com/jonywtf/gofias"
)

func main() {
	_, err := fias.Import("C:\\temp\\kladr\\base\\")
	if err != nil {
		println(err)
		return
	}
	println("Считываение *.DBF файлов успешно завершено")
}
```
