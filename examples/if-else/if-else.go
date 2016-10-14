// شاخه بندی به `if` و  `else` در گو بسیار راحت است.

package main

import "fmt"

func main() {

    // یک مثال ساده.
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // حتی شما میتونید از `if` بدون else استفاده کنید.
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // یک شرایط میتونه بر دیگر وضعیت‌ها مقدم باشه؛
    // هر متغیری در این شرایط تعریف شود
    // در تمامی حالت‌ها در دسترس است.
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

// حواستون باشه که نیازی به پرانتز‌های مزخرف
// در اطراف حالت‌ها در گو نیست، اما براکت‌ها الزامی‌ست.
