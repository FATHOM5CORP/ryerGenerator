package main

import (
	"text/template"
	"os"
	"log"
	"fmt"
	"flag"
)

func main() {
	projectName := flag.String("name", "new-project", "Specifies the name of your project")
	port := flag.Int("port", 8080, "Port that your microservice will listen on")
	
	flag.Parse()

	var main_definer = &mainDefiner { *projectName, *port, }
	var routes_definer = &routesDefiner {}
	var routes_test_definer = &routesTestDefiner {}
	var dockerfile_definer = &dockerfileDefiner { *projectName, *port, }
	var readme_definer = &readmeDefiner { *projectName }

	// Create project directory based off of name in the main definer
	if err := os.Mkdir(main_definer.Name, 0777); err != nil {
		log.Printf("Could not make directory: %s", err)
	}

	// Create main file
	f, err := os.Create(fmt.Sprintf("%s/main.go", main_definer.Name))
	if err != nil {
    	log.Fatal(fmt.Sprintf("Could not create the main file: %s", err))
	}

	// Create the main template and execute on the definer
	main := template.Must(template.New("main").Parse(mainTemplate.Template))
	main.Execute(f, main_definer)

	// Create the routes file
	f, err = os.Create(fmt.Sprintf("%s/routes.go", main_definer.Name))
	if err != nil {
    	log.Fatal(fmt.Sprintf("Could not create the routes file: %s", err))
	}

	// Create the routes template and execute on the definer
	routes := template.Must(template.New("routes").Parse(routesTemplate.Template))
	routes.Execute(f, routes_definer)

	// Create the routes test file
	f, err = os.Create(fmt.Sprintf("%s/routes_test.go", main_definer.Name))
	if err != nil {
    	log.Fatal(fmt.Sprintf("Could not create the routes test file: %s", err))
	}

	// Create the routes template and execute on the definer
	routesTest := template.Must(template.New("routesTest").Parse(routesTestTemplate.Template))
	routesTest.Execute(f, routes_test_definer)

	// Create the routes test file
	f, err = os.Create(fmt.Sprintf("%s/Dockerfile", main_definer.Name))
	if err != nil {
    	log.Fatal(fmt.Sprintf("Could not create the dockerfile: %s", err))
	}

	// Create the routes template and execute on the definer
	dockerfile := template.Must(template.New("dockerfile").Parse(dockerfileTemplate.Template))
	dockerfile.Execute(f, dockerfile_definer)

	// Create the READEME
	f, err = os.Create(fmt.Sprintf("%s/README.md", main_definer.Name))
	if err != nil {
    	log.Fatal(fmt.Sprintf("Could not create the README: %s", err))
	}

	// Create the routes template and execute on the definer
	readme := template.Must(template.New("readme").Parse(readmeTemplate.Template))
	readme.Execute(f, readme_definer)
}