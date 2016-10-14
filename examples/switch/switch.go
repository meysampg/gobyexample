// _حالت سوییچ_ بین حالت‌های مختلف
// یک وضعیت مناسب را بر میگزیند.

package main

import "fmt"
import "time"

func main() {

    // Here's a basic .
    // یک `سوییچ` ساده
    i := 2
    fmt.Print("write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    // میتونید با استفاده از کاما در یک `case`،
    // اکسپرشن‌های مختلفی را پوشش دهید.
    // ما از کیس `default` به صورت دل‌خواه نیز در این مثال
    // برای درک بهتر استفاده کردیم.
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("it's the weekend")
    default:
        fmt.Println("it's a weekday")
    }

    // از `switch` بدون اکسپرشن می‌توان بجای منطق if/else
    // استفاده نمود. در این مثال ما حتی نشان
    // می‌دهیم که چگونه `case` اکسپرشن‌ها
    // می‌توانند ناثابت باشند.
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("it's before noon")
    default:
        fmt.Println("it's after noon")
    }
}

// todo: type switches
