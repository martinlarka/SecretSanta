#!/bin/bash
curl -X POST -H "Content-Type: application/json" -d '
[{"name":"santa1", "phone":"123", "exclude":[1,5], "id":0},{"name":"santa2", "phone":"1233", "exclude":[0,4], "id":1},{"name":"santa3", "phone":"1233", "exclude":[3,0], "id":2},{"name":"santa4", "phone":"1233", "exclude":[2,1], "id":3},{"name":"santa5", "phone":"1233", "exclude":[5,2], "id":4},{"name":"santa6", "phone":"1243", "exclude":[4,3], "id":5}]' "http://localhost:3000/santa/"
