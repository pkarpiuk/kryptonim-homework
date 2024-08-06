https://github.com/pkarpiuk/kryptonim-homework.git

docker build --no-cache=true --tag kryptonim-homework .
docker run --rm -p 8080:8080 -ti kryptonim-homework

curl 'http://localhost:8080/exchange?from=WBTC&to=USDT&amount=1.0'
curl 'http://localhost:8080/rates?currencies=USD,GBP,EUR'
