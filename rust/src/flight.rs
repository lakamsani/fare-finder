/// Represents a single flight option.
#[derive(Debug, Clone, PartialEq)]
pub struct Flight {
    pub airline: String,
    pub flight_number: String,
    pub price: u32,
    pub departure_time: String,
    pub arrival_time: String,
    pub duration: String,
}

impl std::fmt::Display for Flight {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "Flight {{ airline: {:?}, flight_number: {:?}, price: {}, \
             departure_time: {:?}, arrival_time: {:?}, duration: {:?} }}",
            self.airline,
            self.flight_number,
            self.price,
            self.departure_time,
            self.arrival_time,
            self.duration,
        )
    }
}
