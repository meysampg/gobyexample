// `for` تنها ساختار لوپ گو است.
// سه مثال اولیه از لوپ‌های `for` را ببینیم.
package main

import "fmt"

func main() {

    // اولین مدل استفاده، با تک وضعیت.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // یک مثال از اجرا/وضعیت/بعد  برای لوپ `for`
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // `for` بدون وضعیت تا زمانی که شما از آن وضعیت خارج شید به صورت
    // مکرر تکرار میشود مگر با `break` یا `return`(در یک تابع).
    for {
        fmt.Println("loop")
        break
    }
}
