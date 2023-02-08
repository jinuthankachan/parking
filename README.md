# parking

## How to use

### test
```
make test
```

### run 
```
make run
```

## Description
A parking lot is a dedicated area that is intended for parking vehicles. Parking lots are
present in every city and suburban area. Shopping malls, stadiums, airports, train stations,
and similar venues often feature a parking lot with a large capacity. A parking lot can spread
across multiple buildings with multiple floors or can be in a large open area.

* The parking lot will allow different types of vehicles to be parked:
  * Motorcycles/Scooters
  * Cars/SUVs
  * Buses/Trucks
* Each vehicle will occupy a single spot and the spot size will be different for different
vehicles.
* The number of spots per vehicle type will be different for different parking lots. For
example
  * Motorcycles/scooters: 100 spots
  * Cars/SUVs: 80 spots
  * Buses/Trucks: 40 spots
* When a vehicle is parked, a parking ticket should be generated with the spot number
and the entry date-time.
* When a vehicle is unparked, a receipt should be generated with the entry date-time,
exit date-time, and the applicable fees to be paid.
* Different locations have different fee models.

Given a parking lot with details about the vehicle types that can be parked, the number of
spots, and the fee model for the parking lot; compute the fees to be paid for the parked
vehicles when the vehicle is unparked.
