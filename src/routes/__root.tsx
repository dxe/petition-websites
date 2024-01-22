import { Outlet, RootRoute } from "@tanstack/react-router";
import "../index.css";

export const Route = new RootRoute({
  component: () => (
    <>
      <Outlet />
    </>
  ),
});
