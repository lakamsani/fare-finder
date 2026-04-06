export class Flight {
  constructor(
    public readonly airline: string,
    public readonly flightNumber: string,
    public readonly price: number,
    public readonly departureTime: string,
    public readonly arrivalTime: string,
    public readonly duration: string,
  ) {}

  toString(): string {
    return `Flight{airline='${this.airline}', flightNumber='${this.flightNumber}', price=${this.price}, departureTime='${this.departureTime}', arrivalTime='${this.arrivalTime}', duration='${this.duration}'}`;
  }
}
