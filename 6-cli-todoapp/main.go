package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"

    "github.com/taylantutar/6-cli-todoapp/internal/todo"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\n--- TODO MENU ---")
        fmt.Println("1. Görevleri Listele")
        fmt.Println("2. Yeni Görev Ekle")
        fmt.Println("3. Görev Sil")
        fmt.Println("0. Çıkış")
        fmt.Print("Seçim: ")

        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            todo.ListTasks()
        case "2":
            fmt.Print("Görev Başlığı: ")
            scanner.Scan()
            title := scanner.Text()
            todo.AddTask(title)
        case "3":
            fmt.Print("Silinecek Görev ID: ")
            scanner.Scan()
            idStr := scanner.Text()
            id, _ := strconv.Atoi(idStr)
            todo.DeleteTask(id)
        case "0":
            fmt.Println("Çıkılıyor...")
            return
        default:
            fmt.Println("Geçersiz seçim.")
        }
    }
}
