import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import Main from "@/pages/main.tsx";
import { ElementsProvider } from "./store/context";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ElementsProvider>
      <Main />
    </ElementsProvider>
  </StrictMode>
);
