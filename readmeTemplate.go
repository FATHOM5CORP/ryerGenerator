package main

type readmeDefiner struct {
	Name	string
}

var readmeTemplate = &Readme {
	`<p align="center">
	<a href="https://www.fathom5.co/" target="_blank"><img src="https://static.wixstatic.com/media/3d35e8_264823a2b86b409e9615664cd09d57da~mv2.png/v1/fill/w_218,h_45,al_c,q_85,usm_0.66_1.00_0.01/3d35e8_264823a2b86b409e9615664cd09d57da~mv2.webp"/></a>
  </p>
  
  <br>
  
  <h1 align="center">{{.Name}}</h1>
  
  <p align="center">
	<a href="https://docs.python.org/3/index.html" target="_blank"><img src="https://img.shields.io/badge/Golang-1.15-green"/></a>
	<a href="https://travis-ci.org/mingrammer/pyreportcard" target="_blank"><img src="https://travis-ci.org/mingrammer/pyreportcard.svg?branch=master"/></a>
  </p>
  
  <h2>Table of contents</h2>
  <p>
	<ul>
	  <li><a href="https://github.com/FATHOM5/{{.Name}}/blob/master/README.md#overview">Overview</a></li>
	  <li><a href="https://github.com/FATHOM5/{{.Name}}/blob/master/README.md#getting-started">Getting Started</a></li>
	  <li><a href="https://github.com/FATHOM5/{{.Name}}/blob/master/README.md#routes">Routes</a></li>
	</ul>
  </p>
  
  <h2>Overview</h2>
  
  <p>
	  Describe the project. What is the business objective? What are you trying to solve?
  </p>
  
  <h2>Getting Started</h2>

	  How do you run your project?

  <h2>Routes</h2>
  <table>
  <tr>
    <th>Route</th>
    <th>HTTP Method</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>/exampleRoute</td>
    <td>GET</td>
    <td>Example description</td>
  </tr>
</table>
  `,
}