"""Represents a single flight option."""


class Flight:
    """A single flight option with airline, flight number, price, times, and duration."""

    def __init__(
        self,
        airline: str,
        flight_number: str,
        price: int,
        departure_time: str,
        arrival_time: str,
        duration: str,
    ) -> None:
        self.airline = airline
        self.flight_number = flight_number
        self.price = price
        self.departure_time = departure_time
        self.arrival_time = arrival_time
        self.duration = duration

    def __repr__(self) -> str:
        return (
            f"Flight(airline={self.airline!r}, flight_number={self.flight_number!r}, "
            f"price={self.price}, departure_time={self.departure_time!r}, "
            f"arrival_time={self.arrival_time!r}, duration={self.duration!r})"
        )
