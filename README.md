# G2Rail Ground Transportatiion APIs

With 5 standard APIs (search, book, confirm, download and refund) and domain models, [G2Rail](https://www.g2rail.com) as well as XMove App (Apple Store, Google Play) provides a very simple and unified experience to search and book tickets from oversea trains and buses operators. With the APIs, you have access to time table, prices, availabilities from operators in more 50+ countries 10, 000+ stations in Europe, Asia, North America, including Deutsche Bahn(DB), Trenitalia, Italo(NTV), SNCF, SNCB, Swiss Rail(SBB), Finland Rail(VR), National Rail (UK), Renfe (Spain), ÖBB (Austria), Japan Rail(JR), Korean Rail, PKP (Poland), NSB (Norway), China Rail and Amtrak, Taiwan High Speed Rail, Flixbus, RZD (Russia), SJ (Sweden) in Europe, Asia and North America. [Proto files](https://github.com/G2Rail/g2rail-grpc-go/tree/master/proto) 

The API suite is based on unified encoding of data such as stations, ticket tariffs, coach class. For example, only need to pass Milano as departure, you are able to qurey and book from carriers including Trenitalia, TGV, Thello, Flixbus, ÖBB. 

In addition to typical search/booking API, there are also support APIs such as station departure, train update (delay/change). 

[Frankfurt Central Station](http://help.g2rail.com/stations/frankfurt-hbf)

# G2Rail grpc go

This Repository Calls G2Rail Global Rail/Bus and Other Ground Transportation API To Make Search and Reservation. The API is based on the G2Rail API, and is provided in Go lang.

# To generate the protobuf code

```bash
make gen
```

# How to run it?

```bash
ADDRESS=api.g2rail.com APIKEY=api_key_from_g2rail SECRET=api_secret_from_g2rail make client
```
