from dataclasses import dataclass


@dataclass
class Flight:
    airline: str
    flight_number: str
    price: int
    departure_time: str
    arrival_time: str
    duration: str

    def __str__(self) -> str:
        return (
            f"Flight{{airline='{self.airline}', flightNumber='{self.flight_number}', "
            f"price={self.price}, departureTime='{self.departure_time}', "
            f"arrivalTime='{self.arrival_time}', duration='{self.duration}'}}"
        )
