
![RestTail Logo](https://raw.githubusercontent.com/nicholasarvelo/resttail-sdk/main/.github/logo.png)

RestTail is an SDK written in Go for interacting with the [TestRail's API](http://www.gurock.com/testrail/)


Example usage
-------------

```go
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/nicholasarvelo/resttail"
)

func main() {
    var projectNameList string
    setupFlags(&projectNameList)

    flag.Parse()

    if err := validateFlags(); err != nil {
        fmt.Println(err)
        flag.Usage()
        os.Exit(1)
    }

    projectNames := strings.Split(projectNameList, ",")
    if len(projectNames) == 0 || projectNames[0] == "" {
        fmt.Println("Error: no project names provided")
        flag.Usage()
        os.Exit(1)
    }

    apiClient := resttail.NewClient(*url, *username, *password)
    activeProjects, err := apiClient.GetProjects(false)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Project Name: %s\nProject ID: %d\n Project URL: %s\n",
        activeProjects.Projects[0].Name,
        activeProjects.Projects[0].ID,
        activeProjects.Projects[0].URL,
    )
}

func setupFlags(projectNameList *string) {
    flag.StringVar(projectNameList, "project-names", "", "A comma separated list of project names")
    flag.StringVar(url, "testrail-url", "", "The TestRail Account URL")
    flag.StringVar(username, "testrail-username", "", "The TestRail username")
    flag.StringVar(password, "testrail-password", "", "The TestRail password")
}

func validateFlags() error {
    if *username == "" || *password == "" || *url == "" {
        return fmt.Errorf("missing required flags: username, password, or url")
    }
    if !strings.HasPrefix(*url, "https://") {
        *url = "https://" + *url
    }
    return nil
}
```