"use client";

import { useState } from "react";
import { VehicleSearch } from "./vehicle-search";
import { DriverDetailsForm } from "./driver-details-form";
import type { Vehicle } from "@/types/vehicle";

export function QuoteFlow() {
  const [vehicle, setVehicle] = useState<Vehicle | null>(null);

  if (!vehicle) {
    return (
      <VehicleSearch
        onVehicleFound={setVehicle}
      />
    );
  }

  return (
    <>
      <section>
        <h2>
          {vehicle.make} {vehicle.model}
        </h2>

        <p>{vehicle.registration}</p>

        <button onClick={() => setVehicle(null)}>
          Change vehicle
        </button>
      </section>

      <DriverDetailsForm vehicle={vehicle} />
    </>
  );
}