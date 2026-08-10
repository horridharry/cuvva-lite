import { ApiStatus } from "@/components/api-status";
import { VehicleSearch } from "@/components/vehicle-search";
import { QuoteFlow } from "@/components/quote-flow";

export default function Home() {
  return (
    <main className="">
      <nav className="p-5 border-b border-white/20">
        <h1 className='text-sm font-bold monospace tracking-widest'>CUVVA LITE</h1>
      </nav>
      <div className="p-3">

      <h1 className="text-5xl opacity-80 font-bold">Cuvva Lite</h1>

      <p className="pt-5">Cuvva Lite is a demo car insurance simulator - get started by getting your own virtual insurance.</p>

      <div className="mt-4">
        <QuoteFlow/>
      </div>
      </div>
    </main>
  );
}
