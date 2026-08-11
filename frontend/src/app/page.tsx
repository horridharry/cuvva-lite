import { ApiStatus } from "@/components/api-status";
import { VehicleSearch } from "@/components/vehicle-search";
import { QuoteFlow } from "@/components/quote-flow";

export default function Home() {
  return (
    <main className="min-h-screen px-4 py-12">
      <div className="mx-auto max-w-lg">
        <header className="mb-8">
          <p className="text-sm font-semibold uppercase tracking-wider">
            Cuvva Lite
          </p>

          <h1 className="mt-2 text-4xl font-bold tracking-tight">
            Temporary cover in minutes.
          </h1>

          <p className="mt-3">
            Get a demo quote for short-term car cover.
          </p>
        </header>

        <QuoteFlow />
      </div>
    </main>
  );
}
