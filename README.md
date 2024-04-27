RestTail
--------

RestTail is a Go client library for interacting with the [TestRail](http://www.gurock.com/testrail/) API


Example usage
-------------

```go
  package main

  import "github.com/nicholasarvelo/RestTail"

  func main(){

    url := flag.String(
		"testrail-url",
		"",
		"(Required) TrailRail Account URL",
	)

	username := flag.String(
		"testrail-username",
		"",
		"(Required) The TestRail username",
	)

	password := flag.String(
		"testrail-password",
		"",
		"(Required) The TestRail password",
	)

	flag.Parse()

	if *username == "" || *password == "" {
		flag.Usage()
		os.Exit(1)
	}

	apiClient := resttail.NewClient(*url, *username, *password)
  }
```