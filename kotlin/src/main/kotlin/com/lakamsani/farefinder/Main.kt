package com.lakamsani.farefinder

import java.time.LocalDate

/**
 * CLI entry point for fare-finder.
 */
fun main(args: Array<String>) {
    if (args.size != 4) {
        System.err.println("Usage: fare-finder <city1> <state1> <city2> <state2>")
        System.err.println("Example: fare-finder \"San Francisco\" CA \"New York\" NY")
        System.exit(1)
    }

    val city1 = titleCase(args[0].trim())
    val state1 = args[1].trim().uppercase()
    val city2 = titleCase(args[2].trim())
    val state2 = args[3].trim().uppercase()

    val origin = try {
        Airport.lookupAirport(city1, state1)
    } catch (e: Exception) {
        System.err.println("Error: ${e.message}")
        System.exit(1)
        return
    }

    val dest = try {
        Airport.lookupAirport(city2, state2)
    } catch (e: Exception) {
        System.err.println("Error: ${e.message}")
        System.exit(1)
        return
    }

    val tomorrow = LocalDate.now().plusDays(1).toString()

    println("Searching for flights $origin -> $dest on $tomorrow...\n")

    val flights = try {
        Searcher.searchFlights(origin, dest, tomorrow)
    } catch (e: Exception) {
        System.err.println("Error: ${e.message}")
        System.exit(1)
        return
    }

    if (flights.isEmpty()) {
        println("No flights found for this route and date. Try a different date or city pair.")
        System.exit(0)
    }

    val cheapest = flights[0]
    println("Cheapest flight: $origin -> $dest")
    println("${cheapest.airline} ${cheapest.flightNumber}")
    println("$${cheapest.price}")
    println("Departs ${cheapest.departureTime} | Arrives ${cheapest.arrivalTime}")
    println("Duration: ${cheapest.duration}")
}

/**
 * Converts a string so each word starts with an uppercase letter and remaining letters are lowercase.
 */
fun titleCase(s: String): String {
    if (s.isEmpty()) return s
    var capitalizeNext = true
    val sb = StringBuilder()
    for (c in s) {
        when {
            c.isWhitespace() -> {
                capitalizeNext = true
                sb.append(c)
            }
            capitalizeNext -> {
                sb.append(c.uppercaseChar())
                capitalizeNext = false
            }
            else -> sb.append(c.lowercaseChar())
        }
    }
    return sb.toString()
}
