"use client";

import { SubmitEvent, useState } from "react";
import type { Vehicle } from "@/types/vehicle";

type Props = {
  onVehicleFound: (vehicle: Vehicle) => void;
};

export function VehicleSearch({ onVehicleFound }: Props) {
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const [registration, setRegistration] = useState("");
  const [vehicle, setVehicle] = useState<Vehicle | null>(null);

  async function handleSubmit(event: SubmitEvent<HTMLFormElement>) {
    event.preventDefault();

    setError("");
    setLoading(true);
    setVehicle(null);

    try {
      const normalisedRegistration = registration.trim();

      const response = await fetch(
        `/api/vehicles/${encodeURIComponent(normalisedRegistration)}`
      );

      if (response.status === 404) {
        setError("Vehicle not found");
        return;
      }

      if (!response.ok) {
        throw new Error("Vehicle lookup failed");
      }

      const data: Vehicle = await response.json();

      setVehicle(data);
      onVehicleFound(data);
    } catch {
      setError("Something went wrong. Please try again.");
    } finally {
      setLoading(false);
    }
  }

  return (
    <section>
      <form onSubmit={handleSubmit}>
        <label htmlFor="registration">
          Vehicle registration
        </label>

        <input
          id="registration"
          type="text"
          value={registration}
          onChange={(event) => setRegistration(event.target.value)}
          placeholder="FH18 UKU"
          required
        />

        <button type="submit" disabled={loading}>
          {loading ? "Searching..." : "Find vehicle"}
        </button>
      </form>

      {error && <p>{error}</p>}

      {vehicle && (
        <div>
          <h2>
            {vehicle.make} {vehicle.model}
          </h2>

          <p>Registration: {vehicle.registration}</p>
          <p>Year: {vehicle.year}</p>
          <p>Fuel: {vehicle.fuelType}</p>
          <p>Engine: {vehicle.engineSizeCc}cc</p>
        </div>
      )}
    </section>
  );
}