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

      <div className="flex gap-2">
        <div className="flex-1">
          <input
            className="border rounded p-2 border-white/50 w-full"
            id="registration"
            type="text"
            value={registration}
            onChange={(event) => setRegistration(event.target.value)}
            placeholder="FH18 UKU"
            required
          />

   
        </div>

        <button
          className="bg-white text-black rounded-md px-3 text-xs"
          type="submit"
          disabled={loading}
        >
          {loading ? "Searching..." : "Find vehicle"}
        </button>
      </div>
             <p className="text-xs text-white/80 mt-2">
            Try: FH18 UKU or GK21LNP
          </p>
      </form>

      {error && <p>{error}</p>}
    </section>
  );
}