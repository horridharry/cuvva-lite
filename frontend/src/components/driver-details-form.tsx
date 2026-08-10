"use client";

import { SubmitEvent, useState } from "react";
import type { Vehicle } from "@/types/vehicle";

type Props = {
  vehicle: Vehicle;
};

type Quote = {
  id: number;
  vehicleId: number;
  durationMinutes: number;
  pricePence: number;
  expiresAt: string;
};

export function DriverDetailsForm({ vehicle }: Props) {
  const [age, setAge] = useState(24);
  const [yearsLicensed, setYearsLicensed] = useState(4);
  const [penaltyPoints, setPenaltyPoints] = useState(0);
  const [duration, setDuration] = useState(180);

  const [quote, setQuote] = useState<Quote | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  async function handleSubmit(event: SubmitEvent<HTMLFormElement>) {
    event.preventDefault();

    setLoading(true);
    setError("");

    try {
      const response = await fetch("/api/quotes", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          vehicleId: vehicle.id,
          driverAge: age,
          yearsLicensed,
          penaltyPoints,
          durationMinutes: duration,
        }),
      });

      if (!response.ok) {
        throw new Error("Quote request failed");
      }

      const data: Quote = await response.json();

      setQuote(data);
    } catch {
      setError("Unable to create quote.");
    } finally {
      setLoading(false);
    }
  }

  if (quote) {
    return (
      <section>
        <p>Your quote</p>

        <h2>
          £{(quote.pricePence / 100).toFixed(2)}
        </h2>

        <p>
          {vehicle.make} {vehicle.model}
        </p>

        <p>
          Valid until{" "}
          {new Date(quote.expiresAt).toLocaleTimeString()}
        </p>
      </section>
    );
  }

  return (
    <form onSubmit={handleSubmit}>
      <h2>About the driver</h2>

      <label>
        Age
        <input
          type="number"
          value={age}
          min={18}
          max={80}
          onChange={(event) =>
            setAge(Number(event.target.value))
          }
        />
      </label>

      <label>
        Years licence held
        <input
          type="number"
          value={yearsLicensed}
          min={0}
          onChange={(event) =>
            setYearsLicensed(Number(event.target.value))
          }
        />
      </label>

      <label>
        Penalty points
        <input
          type="number"
          value={penaltyPoints}
          min={0}
          onChange={(event) =>
            setPenaltyPoints(Number(event.target.value))
          }
        />
      </label>

      <label>
        Cover length
        <select
          value={duration}
          onChange={(event) =>
            setDuration(Number(event.target.value))
          }
        >
          <option value={60}>1 hour</option>
          <option value={180}>3 hours</option>
          <option value={360}>6 hours</option>
          <option value={1440}>24 hours</option>
        </select>
      </label>

      {error && <p>{error}</p>}

      <button disabled={loading}>
        {loading ? "Getting quote..." : "Get quote"}
      </button>
    </form>
  );
}