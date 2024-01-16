import { Outlet, RootRoute } from "@tanstack/react-router";
import "../index.css";

// TODO: add GA

export const Route = new RootRoute({
  component: () => (
    <>
      <Outlet />
    </>
  ),
});
