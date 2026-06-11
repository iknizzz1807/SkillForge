import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";
import { BACKEND_URL } from '$lib/backend';

export const GET: RequestHandler = async ({ fetch, locals }) => {
  try {
    const token = locals.token;

    const headers: Record<string, string> = {};
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }

    const response = await fetch(BACKEND_URL + "/api/invitations/me", {
      headers,
    });

    if (!response.ok) {
      const errorText = await response.text();
      return json(
        { error: errorText || "Failed to fetch invitations" },
        { status: response.status }
      );
    }

    const data = await response.json();
    return json(data);
  } catch (err) {
    console.error("Error fetching invitations:", err);
    return json({ error: "Failed to fetch invitations" }, { status: 500 });
  }
};
