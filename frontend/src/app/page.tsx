import { ApiStatus } from "@/components/api-status";
import { VehicleSearch } from "@/components/vehicle-search";

export default function Home() {
  return (
    <main>
      <h1>Cuvva Lite</h1>

      <p>Temporary car insurance simulator.</p>

      <ApiStatus />
      <VehicleSearch/>
    </main>
  );
}
