package com.lakamsani.farefinder

/**
 * Represents a single flight option.
 */
data class Flight(
    val airline: String,
    val flightNumber: String,
    val price: Int,
    val departureTime: String,
    val arrivalTime: String,
    val duration: String,
)
