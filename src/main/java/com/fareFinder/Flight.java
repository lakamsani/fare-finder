package com.fareFinder;

public record Flight(
    String airline,
    String flightNumber,
    int price,
    String departureTime,
    String arrivalTime,
    String duration
) {}
