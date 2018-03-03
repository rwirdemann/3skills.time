# Go-Tracker

Beispiel für einen einfachen Time-Tracker in Go.

## Ressourcen
```
curl http://localhost:8080/projects

curl -d '{"Name":"Go-Tracker"}' -H "Content-Type: application/json" -X POST http://localhost:8080/projects

curl http://localhost:8080/projects/1

curl -d '{"Name":"Add MySQL-Repository"}' -H "Content-Type: application/json" -X POST http://localhost:8080/projects/1/activities

curl http://localhost:8080/projects/1/activities
```
