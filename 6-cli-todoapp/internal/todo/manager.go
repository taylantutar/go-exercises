package todo

import (
    "fmt"
)

var tasks []Task
var nextID = 1

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
    fmt.Println("Görev eklendi.")
}

func DeleteTask(id int) {
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            fmt.Println("Görev silindi.")
            return
        }
    }
    fmt.Println("Görev bulunamadı.")
}
