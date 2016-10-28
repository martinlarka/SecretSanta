#!/bin/bash
curl -X POST -H "Content-Type: application/json" -d '[{"name":"santa1", "phone":"123", "exclude":[1], "id":0},{"name":"santa2", "phone":"1233", "exclude":[0], "id":1},{"name":"santa3", "phone":"1243", "exclude":[1], "id":2}]' "http://localhost:3000/santa/"
