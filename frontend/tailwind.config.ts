import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}"
  ],
  theme: {
    extend: {
      colors: {
        thunder: {
          50: "#f5f7ff",
          100: "#e5ebff",
          200: "#c6ceff",
          300: "#9aa2ff",
          400: "#6c6bff",
          500: "#4d3fff",
          600: "#3a2ced",
          700: "#2b20c7",
          800: "#231e9d",
          900: "#1e1b7d"
        },
        carbon: "#0b0b10"
      },
      fontFamily: {
        display: ["var(--font-space-grotesk)", "sans-serif"],
        body: ["var(--font-inter)", "sans-serif"]
      }
    }
  },
  plugins: []
};

export default config;
