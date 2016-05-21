package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

func main() {
    templatePath := flag.String("template-path", "gotpl.tpl", "[required] Path to the template file")
    destinationFolder := flag.String("destination-folder", "current-dir", "[required] Destination folder")
    destinationName := flag.String(
        "destination-name",
        "",
        "[optional] Destination name. Defaults to template-path's filename.")
    parametersFile := flag.String(
        "params-file",
        "params.json",
        "[optional] Json file with template params.")

    flag.Parse()

    params_data, params_errors := ioutil.ReadFile(templatePath)
    if params_errors != nil {
        
    }

    fmt.Println("templatePath:", *templatePath)
    fmt.Println("destinationFolder:", *destinationFolder)
    fmt.Println("destinationName:", *destinationName)
    fmt.Println("parametersFile:", *parametersFile)
}
