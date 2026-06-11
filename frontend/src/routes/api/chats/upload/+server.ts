import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";
import { BACKEND_URL } from '$lib/backend';

export const POST: RequestHandler = async ({ request, fetch, locals }) => {
  const token = locals.token;
  if (!token) {
    return json({ error: "Unauthorized" }, { status: 401 });
  }

  const formData = await request.formData();
  const response = await fetch(BACKEND_URL + "/api/chats/upload", {
    method: "POST",
    headers: { Authorization: `Bearer ${token}` },
    body: formData,
  });

  const data = await response.json().catch(() => ({}));
  return json(data, { status: response.status });
};
