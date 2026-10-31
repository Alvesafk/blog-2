// @ts-check
import { defineConfig, fontProviders } from "astro/config";

import tailwindcss from "@tailwindcss/vite";

// https://astro.build/config
export default defineConfig({
  vite: {
    plugins: [tailwindcss()],
  },

  fonts: [
    {
      provider: fontProviders.local(),
      name: "Libron",
      cssVariable: "--font-libron",
      options: {
        variants: [
          {
            src: ["./src/assets/fonts/Libron-Regular.woff2"],
            weight: "normal",
            style: "normal",
          },
          {
            src: ["./src/assets/fonts/Libron-Bold.woff2"],
            weight: "bold",
            style: "normal",
          },
          {
            src: ["./src/assets/fonts/Libron-Italic.woff2"],
            weight: "normal",
            style: "italic",
          },
          {
            src: ["./src/assets/fonts/Libron-BoldItalic.woff2"],
            weight: "bold",
            style: "italic",
          },
        ],
      },
    },
    {
      provider: fontProviders.local(),
      name: "Advercase",
      cssVariable: "--font-advercase",
      options: {
        variants: [
          {
            src: ["./src/assets/fonts/Advercase-Regular.otf"],
          },
        ],
      },
    },
  ],
});
