import { StrictMode } from "react";
import ReactDOM from "react-dom/client";
import { Router, RouterProvider } from "@tanstack/react-router";
import { routeTree } from "./routeTree.gen";

const router = new Router({ routeTree });

// Registers the router instance for type safety.
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const rootElement = document.getElementById("root")!;
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <StrictMode>
      <RouterProvider router={router} />
    </StrictMode>,
  );
}
