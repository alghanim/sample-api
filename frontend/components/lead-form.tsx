"use client";

import { FormEvent, useMemo, useState, useTransition } from "react";

type FormFields = {
  fullName: string;
  email: string;
  company: string;
  eventType: string;
  budget: string;
  message: string;
};

const initialState: FormFields = {
  fullName: "",
  email: "",
  company: "",
  eventType: "",
  budget: "",
  message: ""
};

const eventTypes = ["Hospitality Suite", "Pop-up Arena", "Tournament Ownership", "Esports Launch", "Other"];
const budgets = ["< $250k", "$250k - $500k", "$500k - $1M", "$1M+", "Undisclosed"];

export default function LeadForm() {
  const [fields, setFields] = useState<FormFields>(initialState);
  const [status, setStatus] = useState<"idle" | "success" | "error">("idle");
  const [message, setMessage] = useState("");
  const [isPending, startTransition] = useTransition();

  const apiBase = useMemo(() => process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080", []);

  const updateField = (name: keyof FormFields, value: string) => {
    setFields((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setStatus("idle");
    setMessage("");

    startTransition(async () => {
      try {
        const res = await fetch(`${apiBase}/api/leads`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(fields)
        });

        if (!res.ok) {
          throw new Error("Unable to submit. Please try again.");
        }

        setFields(initialState);
        setStatus("success");
        setMessage("Request received. A Thunder producer will reply shortly.");
      } catch (error) {
        setStatus("error");
        setMessage(error instanceof Error ? error.message : "Something went wrong.");
      }
    });
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      <div className="grid gap-6 md:grid-cols-2">
        <Field label="Full name" required>
          <input
            required
            value={fields.fullName}
            onChange={(e) => updateField("fullName", e.target.value)}
            className="w-full rounded-2xl border border-white/20 bg-white/5 px-4 py-3 focus:border-white focus:outline-none"
            placeholder="Alex Thunder"
          />
        </Field>
        <Field label="Work email" required>
          <input
            required
            type="email"
            value={fields.email}
            onChange={(e) => updateField("email", e.target.value)}
            className="w-full rounded-2xl border border-white/20 bg-white/5 px-4 py-3 focus:border-white focus:outline-none"
            placeholder="you@brand.com"
          />
        </Field>
      </div>
      <div className="grid gap-6 md:grid-cols-2">
        <Field label="Company / Collective">
          <input
            value={fields.company}
            onChange={(e) => updateField("company", e.target.value)}
            className="w-full rounded-2xl border border-white/20 bg-white/5 px-4 py-3 focus:border-white focus:outline-none"
            placeholder="Northern Circuit Labs"
          />
        </Field>
        <Field label="Event archetype">
          <select
            value={fields.eventType}
            onChange={(e) => updateField("eventType", e.target.value)}
            className="w-full rounded-2xl border border-white/20 bg-white/5 px-4 py-3 focus:border-white focus:outline-none"
          >
            <option value="">Select</option>
            {eventTypes.map((option) => (
              <option key={option} value={option}>
                {option}
              </option>
            ))}
          </select>
        </Field>
      </div>
      <div className="grid gap-6 md:grid-cols-2">
        <Field label="Budget window">
          <select
            value={fields.budget}
            onChange={(e) => updateField("budget", e.target.value)}
            className="w-full rounded-2xl border border-white/20 bg-white/5 px-4 py-3 focus:border-white focus:outline-none"
          >
            <option value="">Select</option>
            {budgets.map((option) => (
              <option key={option} value={option}>
                {option}
              </option>
            ))}
          </select>
        </Field>
        <Field label="What are you launching?">
          <input
            value={fields.message}
            onChange={(e) => updateField("message", e.target.value)}
            className="w-full rounded-2xl border border-white/20 bg-white/5 px-4 py-3 focus:border-white focus:outline-none"
            placeholder="Global hospitality retrofit, Miami GP"
          />
        </Field>
      </div>
      <button
        type="submit"
        disabled={isPending}
        className="w-full rounded-2xl bg-white px-6 py-4 font-semibold text-black transition hover:bg-thunder-100 disabled:cursor-not-allowed disabled:opacity-70"
      >
        {isPending ? "Sending..." : "Send the blueprint"}
      </button>
      {message && (
        <p className={`text-sm ${status === "error" ? "text-red-300" : "text-thunder-200"}`}>
          {message}
        </p>
      )}
    </form>
  );
}

function Field({
  label,
  required,
  children
}: {
  label: string;
  required?: boolean;
  children: React.ReactNode;
}) {
  return (
    <label className="space-y-2 text-sm font-medium text-white/80">
      <span>
        {label}
        {required && <sup className="text-red-400 pl-0.5">*</sup>}
      </span>
      {children}
    </label>
  );
}
