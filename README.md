<p align="center">
  <a href="https://www.fathom5.co/" target="_blank"><img src="https://static.wixstatic.com/media/3d35e8_264823a2b86b409e9615664cd09d57da~mv2.png/v1/fill/w_218,h_45,al_c,q_85,usm_0.66_1.00_0.01/3d35e8_264823a2b86b409e9615664cd09d57da~mv2.webp"/></a>
</p>

<br>

<h1 align="center">ryerGenerator</h1>

<p align="center">
  <a href="https://docs.python.org/3/index.html" target="_blank"><img src="https://img.shields.io/badge/Golang-1.15-green"/></a>
  <a href="https://docs.python.org/3/index.html" target="_blank"><img src="https://img.shields.io/badge/Bash-yellow"/></a>
  <a href="https://travis-ci.org/mingrammer/pyreportcard" target="_blank"><img src="https://travis-ci.org/mingrammer/pyreportcard.svg?branch=master"/></a>
</p>

<h2>Table of contents</h2>
<p>
  <ul>
    <li><a href="https://github.com/FATHOM5/ryerGenerator/blob/master/README.md#overview">Overview</a></li>
    <li><a href="https://github.com/FATHOM5/ryerGenerator/blob/master/README.md#getting-started">Getting Started</a></li>
  </ul>
</p>

<h2>Overview</h2>

<p>
  This ryerGenerator tool allows for one to rapidly create microservices by generating boilerplate code in the ryer pattern. 
  This tool streamlines how developers make microservices. This is in an effort to homogenize the codebase across the ecosystem, 
  allowing for a better sense of familiarity when managing multiple microservices. 
</p>

<h2>Getting Started</h2>

<p>
  CD into this project directory and install the ryer pattern tool to your computer. It will put the tool in <strong>/usr/local/bin</strong>
</p>

```bash
chmod +x install.sh
sudo ./install.sh
```

<p>
  Now you can generate your own ryer microservices. Supply a name and a port to expose the service on. For example, let's make a new project called <strong>my-new-app</strong> and expose it to port <strong>8090</strong>
</p>

```bash
ryerGenerator -name=my-new-app -port=8090
```

<p>
  This will create a directory in your current location called <strong>my-new-app</strong>. The directory structure is as follows:
</p>

```bash
.
├── Dockerfile
├── README.md
├── main.go
├── routes.go
└── routes_test.go
```

<p>
  The golang files follow the ryer pattern, and any instance of <strong>name</strong> and <strong>port</strong> are replaced 
  accordingly. Out of the box, the code base can handle incoming requests. Run some tests to verify:
</p>

```bash
go test -v
```

<p>
  You should see output similar to the following
</p>

```bash
=== RUN   Test_handleVersion
=== RUN   Test_handleVersion/successful_request
 - - [27/May/2020:21:05:54 -0500] "GET /version HTTP/1.1" 200 20
--- PASS: Test_handleVersion (0.00s)
    --- PASS: Test_handleVersion/successful_request (0.00s)
PASS
ok      github.com/ryerGenerator/my-new-app    0.046s
```

<p>
  Now run the web server:
</p>

```bash
go run main.go routes.go
```
<p>
  You should see output similar to the following
</p>

```bash
2020/05/27 20:54:12 my-new-app service listening on localhost port 8090
```
<p>
  In another terminal, make a curl request to the <strong>/version</strong> endpoint
</p>

```bash
curl http://localhost:8090/version
```

<p>
  You should see the following output
</p>

```bash
{"version":"v0.0.1"}
```
