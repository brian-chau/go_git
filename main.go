package main

import (
    "fmt"
    "log"
    "os/exec"

    "github.com/manifoldco/promptui"
)

func menu_prompt(prompt_text string, options []string) string {
    prompt := promptui.Select{
        Label: prompt_text,
        Items: options,
    }

    _, result, err := prompt.Run()
    if err != nil {
        log.Fatalf("Could not accept menu prompt response!\n")
    }
    return result
}

func text_prompt(prompt_text string) string {
    prompt := promptui.Prompt{
        Label:    prompt_text,
    }

    result, err := prompt.Run()
    if err != nil {
        log.Fatalf("Could not accept text prompt response!\n")
    }
    return result
}

func create_branch(branch_name string) bool {
    string_cmd := fmt.Sprintf("git checkout -b %s", branch_name)
    cmd        := exec.Command(string_cmd)
    err        := cmd.Run()
    if err != nil {
        log.Fatalf("Could not create a branch!")
    }
    return true
}

func switch_to_branch(branch_name string) bool {
    string_cmd := fmt.Sprintf("git checkout %s", branch_name)
    cmd        := exec.Command(string_cmd)
    err        := cmd.Run()
    if err != nil {
        log.Fatalf("Could not switch to the branch!")
    }
    return true
}

func main() {
    menu_choice := menu_prompt("Git actions", []string{"Create a branch", "Switch branches", "List all branches", "Commit changes"})

    switch menu_choice {
        case "Create a branch":
            branch_name := text_prompt("Enter a branch name")
            create_branch(branch_name)
        case "Switch branches":
            branch_name := text_prompt("Enter a branch name")
            switch_to_branch(branch_name)
        case "List all branches":
        default:
            fmt.Printf("Invalid option: %v\n", menu_choice);
    }
}

