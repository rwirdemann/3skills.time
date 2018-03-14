# Go-Tracker

Eine Zeiterfassungs-API für Freiberufler und kleinere Firmen.

## Ressourcen

### GET /projects
```
curl http://localhost:8080/projects
```

### POST /projects
```
curl -d '{"Name":"Go-Tracker"}' \
    -H "Authorization: Bearer mytoken123" \
    -H "Content-Type: application/json" -X POST \
    http://localhost:8080/projects
```

### POST /projects/project_id/bookings
```
curl -d '{"Name":"Add MySQL-Repository"}' \
    -H "Authorization: Bearer mytoken123" \
    -H "Content-Type: application/json" -X POST \
    http://localhost:8080/projects/1/activities
```
