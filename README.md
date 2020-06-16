# What is G2Rail?

[G2Rail](https://www.g2rail.com) provides a very simple and unified User Experience to make reservations to global train, but and other ground transportations. A user can book train and bus from Deutsche Bahn(DB), Trenitalia, Italo(NTV), SNCF, SNCB, Swiss Rail(SBB), Finland Rail(VR), The Great British(National Rail), Renfe, Japan Rail(JR), Korean Rail, China Rail and Amtrak, Flixbus in Europe and The USA. These [Proto files](https://github.com/G2Rail/g2rail-grpc-go/tree/master/proto) define a standard unifed API which has been adopt by major OTA and other travel businesses.

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
