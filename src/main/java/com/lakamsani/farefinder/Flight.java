package com.lakamsani.farefinder;

/**
 * Represents a single flight option.
 */
public class Flight {
    private final String airline;
    private final String flightNumber;
    private final int price;
    private final String departureTime;
    private final String arrivalTime;
    private final String duration;

    public Flight(String airline, String flightNumber, int price,
                  String departureTime, String arrivalTime, String duration) {
        this.airline = airline;
        this.flightNumber = flightNumber;
        this.price = price;
        this.departureTime = departureTime;
        this.arrivalTime = arrivalTime;
        this.duration = duration;
    }

    public String getAirline() { return airline; }
    public String getFlightNumber() { return flightNumber; }
    public int getPrice() { return price; }
    public String getDepartureTime() { return departureTime; }
    public String getArrivalTime() { return arrivalTime; }
    public String getDuration() { return duration; }

    @Override
    public String toString() {
        return String.format("Flight{airline='%s', flightNumber='%s', price=%d, " +
                "departureTime='%s', arrivalTime='%s', duration='%s'}",
                airline, flightNumber, price, departureTime, arrivalTime, duration);
    }
}
