import LeadForm from "@/components/lead-form";

type ThunderEvent = {
  id: string;
  title: string;
  subtitle: string;
  description: string;
  location: string;
  venue: string;
  startDate: string;
  endDate: string;
  status: string;
  primaryImage?: string;
  tags: string[];
};

const fallbackEvents: ThunderEvent[] = [
  {
    id: "grand-prix-lounge",
    title: "Grand Prix Command Lounge",
    subtitle: "Trackside strategy suites for high-velocity networking",
    description:
      "Private race control set-ups with biometric concierge, hospitality, and immersive data walls so your guests feel part of the pit wall.",
    location: "Dubai",
    venue: "Autodrome VIP Grid",
    startDate: "2025-02-14",
    endDate: "2025-02-16",
    status: "Now booking",
    primaryImage: "https://images.unsplash.com/photo-1512436991641-6745cdb1723f?auto=format&fit=crop&w=1600&q=80",
    tags: ["F1", "Hospitality", "Data"]
  },
  {
    id: "skyline-court",
    title: "Skyline Court Classics",
    subtitle: "Pop-up glass court suspended over iconic skylines",
    description:
      "Host C-level matchplay above the city with kinetic seating, Michelin-led tasting menus, and AR match analytics.",
    location: "Singapore",
    venue: "Marina Vista Sky Deck",
    startDate: "2025-05-05",
    endDate: "2025-05-10",
    status: "Concept release",
    primaryImage: "https://images.unsplash.com/photo-1461896836934-ffe607ba8211?auto=format&fit=crop&w=1600&q=80",
    tags: ["Tennis", "Pop-up", "AR"]
  },
  {
    id: "arena-storm",
    title: "Arena Storm Week",
    subtitle: "Immersive e-sports storyworld built for product drops",
    description:
      "Seven-day residency with programmable LED canyon, modular broadcast pods, and Keycloak-secured creator labs.",
    location: "Berlin",
    venue: "Thunder Vault",
    startDate: "2025-07-01",
    endDate: "2025-07-07",
    status: "Limited",
    primaryImage: "https://images.unsplash.com/photo-1508609349937-5ec4ae374ebf?auto=format&fit=crop&w=1600&q=80",
    tags: ["Esports", "Launch", "Web3"]
  }
];

async function getEvents(): Promise<ThunderEvent[]> {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080";
  try {
    const res = await fetch(`${baseUrl}/api/events`, {
      next: { revalidate: 60 }
    });
    if (!res.ok) {
      return fallbackEvents;
    }
    const payload = await res.json();
    return payload?.data ?? fallbackEvents;
  } catch {
    return fallbackEvents;
  }
}

export default async function Home() {
  const events = await getEvents();

  return (
    <main className="space-y-20 pb-32">
      <section className="section-padding pt-12">
        <div className="glass rounded-3xl p-10 md:p-16 relative overflow-hidden">
          <div className="absolute inset-0 opacity-40 bg-gradient-to-r from-thunder-600 via-fuchsia-500 to-amber-400 blur-3xl" />
          <div className="relative z-10 grid gap-10 lg:grid-cols-[1.2fr,0.8fr]">
            <div className="space-y-8">
              <p className="inline-flex items-center gap-2 rounded-full border border-white/20 px-4 py-2 text-xs uppercase tracking-[0.3em]">
                Thunder • Sporting Event Architects
              </p>
              <h1 className="font-display text-4xl leading-tight md:text-6xl lg:text-7xl">
                Re-engineer sporting experiences for{" "}
                <span className="text-transparent bg-clip-text bg-gradient-to-r from-thunder-300 via-white to-amber-200">
                  decisive brands
                </span>
              </h1>
              <p className="text-lg text-white/80 md:text-xl">
                From hyper-exclusive paddock briefings to sky courts suspended above global capitals, Thunder designs
                event systems that move faster than the season calendar.
              </p>
              <div className="flex flex-wrap gap-4">
                <a
                  href="#experiences"
                  className="rounded-full bg-white px-6 py-3 font-semibold text-black transition hover:bg-thunder-100"
                >
                  Explore Experiences
                </a>
                <a
                  href="#contact"
                  className="rounded-full border border-white/40 px-6 py-3 font-semibold text-white transition hover:border-white"
                >
                  Book a Strategy Call
                </a>
              </div>
              <div className="grid grid-cols-2 gap-6 pt-6 md:grid-cols-4">
                {[
                  { label: "Events delivered", value: "180+" },
                  { label: "Cities activated", value: "32" },
                  { label: "Guest NPS", value: "92" },
                  { label: "Avg. lead time", value: "21 days" }
                ].map((stat) => (
                  <div key={stat.label}>
                    <p className="text-2xl font-semibold">{stat.value}</p>
                    <p className="text-sm text-white/70">{stat.label}</p>
                  </div>
                ))}
              </div>
            </div>
            <div className="glass rounded-2xl border-white/5 p-6 space-y-6">
              <p className="text-sm uppercase text-white/70 tracking-[0.2em]">Trusted by</p>
              <div className="grid grid-cols-2 gap-4 text-2xl font-display text-white/80">
                <span>NEOM</span>
                <span>Oracle Red Bull</span>
                <span>Adidas Global</span>
                <span>Extreme E</span>
              </div>
              <div className="rounded-2xl bg-white/10 p-4">
                <p className="text-white/70 text-sm">Latest concept drop</p>
                <p className="text-xl font-semibold">Thunder x Keycloak secure guest graphs</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section id="experiences" className="section-padding space-y-12">
        <div className="flex flex-col gap-4 md:flex-row md:items-end md:justify-between">
          <div>
            <p className="text-sm uppercase tracking-[0.3em] text-white/70">Flagship systems</p>
            <h2 className="font-display text-3xl md:text-4xl">Curated sporting blueprints</h2>
          </div>
          <p className="max-w-2xl text-white/70">
            Built with modular staging, Keycloak-secured guest graphs, and live PocketBase intelligence. Deployable in
            21 days or custom engineered for your calendar.
          </p>
        </div>
        <div className="grid gap-6 lg:grid-cols-3">
          {events.map((event) => (
            <article key={event.id} className="glass rounded-3xl flex flex-col overflow-hidden">
              {event.primaryImage && (
                <div
                  className="h-52 bg-cover bg-center"
                  style={{ backgroundImage: `linear-gradient(180deg, rgba(8,8,12,0.1), rgba(8,8,12,0.9)), url(${event.primaryImage})` }}
                />
              )}
              <div className="p-8 space-y-4 flex-1 flex flex-col">
                <div className="flex items-center justify-between text-sm text-white/70">
                  <span>{event.location}</span>
                  <span>{event.status}</span>
                </div>
                <h3 className="font-display text-2xl">{event.title}</h3>
                <p className="text-white/70 flex-1">{event.description}</p>
                <div className="flex flex-wrap gap-2">
                  {event.tags?.map((tag) => (
                    <span key={tag} className="rounded-full bg-white/5 px-3 py-1 text-xs uppercase tracking-widest text-white/70">
                      {tag}
                    </span>
                  ))}
                </div>
                <div className="text-sm text-white/70">
                  {event.venue} • {new Date(event.startDate).toLocaleDateString(undefined, { month: "short", day: "numeric" })} -{" "}
                  {new Date(event.endDate).toLocaleDateString(undefined, { month: "short", day: "numeric" })}
                </div>
              </div>
            </article>
          ))}
        </div>
      </section>

      <section className="section-padding space-y-12">
        <div className="grid gap-6 lg:grid-cols-3">
          {[
            {
              title: "Intelligence sprint",
              body: "PocketBase pipelines sync athlete, sponsor, and guest intelligence so your team programs with live data."
            },
            {
              title: "Identity-first security",
              body: "Keycloak-secured APIs gate every admin console, RSVP portal, and broadcast tool with enterprise-grade auth."
            },
            {
              title: "Echo-driven reliability",
              body: "Golang + Echo services keep latency low, unlocking real-time screen control, telemetry, and guest comms."
            }
          ].map((item) => (
            <div key={item.title} className="glass rounded-3xl p-8 space-y-4">
              <h3 className="font-display text-2xl">{item.title}</h3>
              <p className="text-white/70">{item.body}</p>
            </div>
          ))}
        </div>
      </section>

      <section className="section-padding space-y-8" id="contact">
        <div className="max-w-3xl space-y-4">
          <p className="text-sm uppercase tracking-[0.3em] text-white/70">Engage Thunder</p>
          <h2 className="font-display text-4xl">Engineer your next sporting moment</h2>
          <p className="text-white/70">
            Share your objectives and our producers will return a calibrated event system inside two business days.
          </p>
        </div>
        <div className="glass rounded-3xl p-8">
          <LeadForm />
        </div>
      </section>

      <footer className="section-padding border-t border-white/10 text-sm text-white/60">
        <div className="flex flex-col gap-3 md:flex-row md:items-center md:justify-between">
          <p>© {new Date().getFullYear()} Thunder Event Systems. Crafted for decisive brands.</p>
          <div className="flex gap-6 text-white/70">
            <a href="mailto:hello@thunderevents.io">hello@thunderevents.io</a>
            <a href="https://suffix.events" target="_blank" rel="noreferrer">
              Inspiration
            </a>
          </div>
        </div>
      </footer>
    </main>
  );
}
