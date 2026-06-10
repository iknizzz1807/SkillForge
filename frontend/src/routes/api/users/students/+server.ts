import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

export const GET: RequestHandler = async ({ fetch, locals, url }) => {
  const token = locals.token;
  if (!token) {
    return json({ error: "Unauthorized" }, { status: 401 });
  }

  const response = await fetch(
    `http://backend:8080/api/users/students?${url.searchParams}`,
    {
      headers: { Authorization: `Bearer ${token}` },
    }
  );

  const data = await response.json().catch(() => ({}));
  return json(data, { status: response.status });
};
