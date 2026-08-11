"use client";

import { SubmitEvent, useState } from "react";
import type { Vehicle } from "@/types/vehicle";

type Props = {
  onVehicleFound: (vehicle: Vehicle) => void;
};

export function VehicleSearch({ onVehicleFound }: Props) {
  const [registration, setRegistration] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  async function handleSubmit(event: SubmitEvent<HTMLFormElement>) {
    event.preventDefault();

    setError("");
    setLoading(true);

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

      const vehicle: Vehicle = await response.json();

      onVehicleFound(vehicle);
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

        <button className='bg-white text-black rounded-md p-2 px-3 text-xs' type="submit" disabled={loading}>
          {loading ? "Searching..." : "Find vehicle"}
        </button>
      </form>

      {error && <p>{error}</p>}
    </section>
  );
}