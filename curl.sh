
readonly Bearer="Bearer"
readonly TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjMxMjMsInVzZXJuYW1lIjoiMTIzMTIzIiwicm9sZSI6MSwiZXhwIjoxNzI1NjI4ODg0fQ.becHdr4eQxky7-_tCPg6Kczc0gp2X6MmYOpkeXCVhZM"
curl --location --request GET 'http://localhost:8888/api/users/info' \
--header 'Authorization: $Bearer $TOKEN' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Accept: */*' \
--header 'Host: localhost:8888' \
--header 'Connection: keep-alive'