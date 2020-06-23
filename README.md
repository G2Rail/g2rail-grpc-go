# G2Rail Ground Transportation APIs

[G2Rail](https://www.g2rail.com) as well as [XMove App Apple Store](https://apps.apple.com/us/app/id1478629552) & [XMove App Google Play](https://play.google.com/store/apps/details?id=com.g2rail.mobile) provides straightforward and unified experience to search and book tickets from oversea trains and buses operators. With the APIs, you have access to time table, prices, availabilities from operators in more 50+ countries 10, 000+ stations in Europe, Asia, North America, including Italy, Germany, France, Spain, UK, Belgium, Netherlands,  Luxembourg, Switzerland, Austria, Czech Republic, Poland, Russia, Norway, Finnland, Sweden, Japan, Korean, China, Hong Kong, Taiwan, USA, Canada and so on.

## Booking API

The API suite unified booking process of various train/bus operators into 5 APIs including search, book, confirm, download and refund. [Proto files](https://github.com/G2Rail/g2rail-grpc-go/tree/master/proto) 

In addition, the API suite is built upon a consolidated ground transportation domain model as well as standardized data such as stations, ticket tariffs, coach class. 

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

## Supporting API

In addition to booking API, there are also support APIs such as station departure, train update.

[Frankfurt Central Station](http://help.g2rail.com/stations/frankfurt-hbf)

[Realtime update of DB ICE 1088 from Kassel to Frankfurt Hbf](https://help.g2rail.com/zh-cn/railways/db/carriers/db-intercity-express/trains/ICE1088)

Support API include:

### Suggestion

Based on departure city or station, return hot destination cities/destinations.

### Station update

Return real time update of station departure boards data. The data is consolidated from multiple train operators. For example, Roma Termini station departure board shows update from both Trenitalia and Italo. 

### Train update

Return record of train status based on carrier, train number and date. If date is not specified, will return most recent one.

## Partners

Alibaba - This API suite and data become data infrastructure of Alipay and Fliggy. It is now integrated seamlessly with airline booking, hotel booking, oversea railway booking and oversea bus booking. 
