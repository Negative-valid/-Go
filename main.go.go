package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    lowercase := "abcdefghijklmnopqrstuvwxyz"
    uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    digits := "0123456789"
    specials := "!@#$%^&*()-_=+[]{};:,.<>?/|"

    var length int
    var useLower, useUpper, useDigits, useSpecials string

    fmt.Print("Введіть довжину пароля: ")
    fmt.Scan(&length)

    fmt.Print("Використовувати малі літери? (Yes/No): ")
    fmt.Scan(&useLower)

    fmt.Print("Використовувати великі літери? (Yes/No): ")
    fmt.Scan(&useUpper)

    fmt.Print("Використовувати цифри? (Yes/No): ")
    fmt.Scan(&useDigits)

    fmt.Print("Використовувати спецсимволи? (Yes/No): ")
    fmt.Scan(&useSpecials)

    charset := ""
    if useLower == "Yes" {
        charset += lowercase
    }
    if useUpper == "Yes" {
        charset += uppercase
    }
    if useDigits == "Yes" {
        charset += digits
    }
    if useSpecials == "Yes" {
        charset += specials
    }

    password := ""
    for i := 0; i < length; i++ {
        index := rand.Intn(len(charset))
        password += string(charset[index])
    }

    fmt.Println("Ваш пароль:", password)
}