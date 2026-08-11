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
    <section className="mt-6 rounded-2xl border border-gray-200 bg-white p-6 shadow-sm">
      <p className="text-sm font-medium text-gray-500">
        Your quote
      </p>

      <h2 className="mt-2 text-4xl font-bold tracking-tight text-gray-900">
        £{(quote.pricePence / 100).toFixed(2)}
      </h2>

      <div className="mt-6 border-t border-gray-100 pt-4">
        <p className="font-semibold text-gray-900">
          {vehicle.make} {vehicle.model}
        </p>

        <p className="mt-1 text-sm text-gray-500">
          {vehicle.registration}
        </p>
      </div>

      <p className="mt-4 text-sm text-gray-500">
        Quote valid until{" "}
        <span className="font-medium text-gray-700">
          {new Date(quote.expiresAt).toLocaleTimeString([], {
            hour: "2-digit",
            minute: "2-digit",
          })}
        </span>
      </p>

      <button
        type="button"
        className="mt-6 w-full rounded-xl bg-black px-4 py-3 font-semibold text-white transition hover:bg-gray-800"
      >
        Buy cover
      </button>
    </section>
  );
}

return (
  <form
    onSubmit={handleSubmit}
    className="mt-6 space-y-5 rounded-2xl border border-gray-200 bg-white p-6 shadow-sm"
  >
    <div>
      <p className="text-sm font-medium text-gray-500">
        Step 2
      </p>

      <h2 className="mt-1 text-2xl font-bold text-gray-900">
        About the driver
      </h2>

      <p className="mt-1 text-sm text-gray-500">
        We&apos;ll use these details to calculate your demo quote.
      </p>
    </div>

    <div>
      <label
        htmlFor="age"
        className="mb-2 block text-sm font-medium text-gray-700"
      >
        Age
      </label>

      <input
        id="age"
        type="number"
        value={age}
        min={18}
        max={80}
        onChange={(event) =>
          setAge(Number(event.target.value))
        }
        className="w-full rounded-xl border border-gray-300 px-4 py-3 text-gray-900 outline-none transition focus:border-black"
      />
    </div>

    <div>
      <label
        htmlFor="yearsLicensed"
        className="mb-2 block text-sm font-medium text-gray-700"
      >
        Years licence held
      </label>

      <input
        id="yearsLicensed"
        type="number"
        value={yearsLicensed}
        min={0}
        onChange={(event) =>
          setYearsLicensed(Number(event.target.value))
        }
        className="w-full rounded-xl border border-gray-300 px-4 py-3 text-gray-900 outline-none transition focus:border-black"
      />
    </div>

    <div>
      <label
        htmlFor="penaltyPoints"
        className="mb-2 block text-sm font-medium text-gray-700"
      >
        Penalty points
      </label>

      <input
        id="penaltyPoints"
        type="number"
        value={penaltyPoints}
        min={0}
        onChange={(event) =>
          setPenaltyPoints(Number(event.target.value))
        }
        className="w-full rounded-xl border border-gray-300 px-4 py-3 text-gray-900 outline-none transition focus:border-black"
      />
    </div>

    <div>
      <label
        htmlFor="duration"
        className="mb-2 block text-sm font-medium text-gray-700"
      >
        Cover length
      </label>

      <select
        id="duration"
        value={duration}
        onChange={(event) =>
          setDuration(Number(event.target.value))
        }
        className="w-full rounded-xl border border-gray-300 bg-white px-4 py-3 text-gray-900 outline-none transition focus:border-black"
      >
        <option value={60}>1 hour</option>
        <option value={180}>3 hours</option>
        <option value={360}>6 hours</option>
        <option value={1440}>24 hours</option>
      </select>
    </div>

    {error && (
      <div className="rounded-xl bg-red-50 px-4 py-3 text-sm text-red-700">
        {error}
      </div>
    )}

    <button
      type="submit"
      disabled={loading}
      className="w-full rounded-xl bg-black px-4 py-3 font-semibold text-white transition hover:bg-gray-800 disabled:cursor-not-allowed disabled:opacity-50"
    >
      {loading ? "Getting quote..." : "Get quote"}
    </button>
  </form>
);
}