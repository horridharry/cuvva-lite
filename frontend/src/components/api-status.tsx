"use client";

import { useEffect, useState } from "react";

type HealthResponse = {
  status: string;
};

export function ApiStatus() {
  const [status, setStatus] = useState("Checking server health");

  useEffect(() => {
    async function checkApi() {
      try {
        const response = await fetch("/api/health");

        if (!response.ok) {
          throw new Error("Api request failed");
        }

        const data: HealthResponse = await response.json();

        setStatus(data.status);
      } catch {
        setStatus("Unavailable");
      }
    }

    checkApi();
  }, []);

  return (
    <p>
      API status: <strong>{status}</strong>
    </p>
  );
}