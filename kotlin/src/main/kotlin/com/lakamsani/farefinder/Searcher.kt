package com.lakamsani.farefinder

import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import java.net.URI
import java.net.http.HttpClient
import java.net.http.HttpRequest
import java.net.http.HttpResponse

/**
 * Handles flight search via SerpAPI and JSON parsing.
 */
object Searcher {

    /** SerpAPI endpoint. Package-private so tests can override. */
    @JvmField
    var baseUrl: String = "https://serpapi.com/search"

    /**
     * API key override.
     * null  = read from SERPAPI_KEY environment variable
     * ""    = simulate missing key (triggers error)
     * other = use this key directly
     */
    @JvmField
    var apiKeyOverride: String? = null

    private val mapper = jacksonObjectMapper()

    /**
     * Fetches flights from SerpAPI for the given origin, destination, and date.
     *
     * @param origin 3-letter IATA origin code
     * @param dest   3-letter IATA destination code
     * @param date   departure date in yyyy-MM-dd format
     * @return list of flights sorted by price ascending
     * @throws Exception if the API key is missing, the HTTP request fails, or JSON parsing fails
     */
    fun searchFlights(origin: String, dest: String, date: String): List<Flight> {
        val apiKey: String = if (apiKeyOverride != null) {
            apiKeyOverride!!
        } else {
            System.getenv("SERPAPI_KEY") ?: ""
        }

        if (apiKey.isEmpty()) {
            throw Exception("SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com")
        }

        val url = "$baseUrl?engine=google_flights&departure_id=$origin&arrival_id=$dest" +
                "&outbound_date=$date&currency=USD&hl=en&api_key=$apiKey&type=2"

        val client = HttpClient.newHttpClient()
        val request = HttpRequest.newBuilder()
            .uri(URI.create(url))
            .GET()
            .build()

        val response = client.send(request, HttpResponse.BodyHandlers.ofString())

        if (response.statusCode() != 200) {
            throw Exception("API returned status ${response.statusCode()}")
        }

        return parseFlights(response.body())
    }

    /**
     * Parses a SerpAPI JSON response and returns flights sorted by price ascending.
     *
     * @param json the raw JSON string from SerpAPI
     * @return list of flights sorted by price ascending
     * @throws Exception if JSON parsing fails
     */
    fun parseFlights(json: String): List<Flight> {
        val root: JsonNode = try {
            mapper.readTree(json)
        } catch (e: Exception) {
            throw Exception("failed to parse JSON: ${e.message}", e)
        }

        val flights = mutableListOf<Flight>()

        for (groupKey in listOf("best_flights", "other_flights")) {
            val groups = root.path(groupKey)
            if (groups.isMissingNode || !groups.isArray) continue

            for (group in groups) {
                val flightLegs = group.path("flights")
                if (!flightLegs.isArray || flightLegs.size() == 0) continue

                val price = group.path("price").asInt(0)
                val firstLeg = flightLegs.get(0)
                val lastLeg = flightLegs.get(flightLegs.size() - 1)

                val airline = firstLeg.path("airline").asText("")
                val flightNumber = firstLeg.path("flight_number").asText("")
                val departureTime = firstLeg.path("departure_airport").path("time").asText("")
                val arrivalTime = lastLeg.path("arrival_airport").path("time").asText("")

                var totalDuration = 0
                for (leg in flightLegs) {
                    totalDuration += leg.path("duration").asInt(0)
                }

                flights.add(
                    Flight(airline, flightNumber, price, departureTime, arrivalTime, formatDuration(totalDuration))
                )
            }
        }

        flights.sortBy { it.price }
        return flights
    }

    /**
     * Converts total minutes into a human-readable "Xh Ym" string.
     */
    private fun formatDuration(minutes: Int): String = "${minutes / 60}h ${minutes % 60}m"
}
