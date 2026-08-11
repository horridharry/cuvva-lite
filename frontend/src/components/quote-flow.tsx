"use client";

import { useState } from "react";

import { VehicleSearch } from "./vehicle-search";
import { DriverDetailsForm } from "./driver-details-form";

import type { Vehicle } from "@/types/vehicle";

export function QuoteFlow() {
  const [vehicle, setVehicle] = useState<Vehicle | null>(null);

  if (!vehicle) {
    return (
      <VehicleSearch onVehicleFound={setVehicle} />
    );
  }

  return (
    <>
      <section>
        <p>Your vehicle</p>

        <h2>
          {vehicle.make} {vehicle.model}
        </h2>

        <p>{vehicle.registration}</p>

        <button
         className='bg-white text-black rounded-md p-2 px-3 text-xs'
          type="button"
          onClick={() => setVehicle(null)}
        >
          Change vehicle
        </button>
      </section>

      <DriverDetailsForm vehicle={vehicle} />
    </>
  );
}