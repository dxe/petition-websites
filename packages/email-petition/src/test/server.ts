import { setupServer } from "msw/node";

// Shared MSW server for all tests. No default handlers: each test installs its
// own with `server.use(...)` so it can both stub the response AND capture the
// outgoing request (URL, headers, body) to assert on what the component sent to
// the petition (/sign) and mailer (/message/create) endpoints.
export const server = setupServer();
