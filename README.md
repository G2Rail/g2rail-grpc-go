# G2Rail Ground Transportation APIs

With 5 APIs (search, book, confirm, download and refund) and domain models, [G2Rail](https://www.g2rail.com) as well as XMove App (Apple Store, Google Play) provides a very simple and unified experience to search and book tickets from oversea trains and buses operators. With the APIs, you have access to time table, prices, availabilities from operators in more 50+ countries 10, 000+ stations in Europe, Asia, North America, including Italy, Germany, France, Spain, UK, Belgium, Netherlands,  Luxembourg, Switzerland, Austria, Czech Republic, Poland, Russia, Norway, Finnland, Sweden, Japan, Korean, China, Hong Kong, Taiwan, USA, Canada and so on. [Proto files](https://github.com/G2Rail/g2rail-grpc-go/tree/master/proto)

The API suite is based on unified encoding of data such as stations, ticket tariffs, coach class. For example, only need to pass Milano as departure, you are able to qurey and book from carriers including Trenitalia, TGV, Thello, Flixbus, ÖBB etc.

In addition to typical search/booking API, there are also support APIs such as station departure, train update (delay/change).

[Frankfurt Central Station](http://help.g2rail.com/stations/frankfurt-hbf)

# G2Rail grpc go

This Repository Calls G2Rail Ground Transportation APIs. It is provided in Go lang.

# To generate the protobuf code

```bash
make gen
```

# How to run it?

```bash
ADDRESS=api.g2rail.com APIKEY=api_key_from_g2rail SECRET=api_secret_from_g2rail make client
```
