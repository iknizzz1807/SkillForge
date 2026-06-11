import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";
import { BACKEND_URL } from '$lib/backend';

export const POST: RequestHandler = async ({ request, fetch, locals }) => {
  const token = locals.token;
  if (!token) {
    return json({ error: "Unauthorized" }, { status: 401 });
  }

  const body = await request.json();
  const response = await fetch(BACKEND_URL + "/api/portfolios", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(body),
  });

  const data = await response.json().catch(() => ({}));
  return json(data, { status: response.status });
};
