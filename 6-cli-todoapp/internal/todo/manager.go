package todo

import (
    "encoding/json"
    "fmt"
    "os"
)

var tasks []Task
var nextID = 1
const dataFile = "tasks.json"

func init() {
    loadTasks()
}

func ListTasks() {
    fmt.Println("\n--- GÖREVLER ---")
    if len(tasks) == 0 {
        fmt.Println("Henüz görev yok.")
        return
    }
    for _, task := range tasks {
        status := " "
        if task.Done {
            status = "X"
        }
        fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
    }
}

func AddTask(title string) {
    task := Task{
        ID:    nextID,
        Title: title,
        Done:  false,
    }
    tasks = append(tasks, task)
    nextID++
    saveTasks()
    fmt.Println("Görev eklendi.")
}

func DeleteTask(id int) {
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            saveTasks()
            fmt.Println("Görev silindi.")
            return
        }
    }
    fmt.Println("Görev bulunamadı.")
}

func saveTasks() {
    file, err := os.Create(dataFile)
    if err != nil {
        fmt.Println("Dosya oluşturulamadı:", err)
        return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    err = encoder.Encode(tasks)
    if err != nil {
        fmt.Println("Dosyaya yazılamadı:", err)
    }
}

func loadTasks() {
    file, err := os.Open(dataFile)
    if err != nil {
        // Dosya yoksa sorun değil
        return
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&tasks)
    if err != nil {
        fmt.Println("Görevler yüklenemedi:", err)
    }

    // Son ID'yi bulalım
    for _, task := range tasks {
        if task.ID >= nextID {
            nextID = task.ID + 1
        }
    }
}

func MarkDone(id int) {
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Done = true
            saveTasks()
            fmt.Println("Görev tamamlandı.")
            return
        }
    }
    fmt.Println("Görev bulunamadı.")
}

