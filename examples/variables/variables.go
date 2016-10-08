// در گو، _متغییرها_ به صراحت تعریف میشوند و توسط کامپایلر استفاده می‌شوند.

package main

import "fmt"

func main() {

    // با `var` یک یا چنیدین متغییر تعریف کنید.
    var a string = "initial"
    fmt.Println(a)

    // شما میتونید چند متغیر رو یک مرتبه تعریف کنید.
    var b, c int = 1, 2
    fmt.Println(b, c)

    // گو حتی نوع متغیر تعریف شده را حدس میزند.
    var d = true
    fmt.Println(d)

    // متغیر‌ها بدون مقدار اولیه نیز قابل تعریف هستند.
    // برای مثال مقدار اولیه `اعداد صحیح` `0` است.
    var e int
    fmt.Println(e)

    // عبارت `:=` راهی برای کوتاه کردن تعریف و مقداردهی اولیه برای یک متغیر می‌باشد،
    // برای مثال در این کیس `var f string = "short"`
    f := "short"
    fmt.Println(f)
}
