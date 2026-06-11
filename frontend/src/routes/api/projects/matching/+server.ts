import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";
import { BACKEND_URL } from '$lib/backend';

export const GET: RequestHandler = async ({ fetch, locals }) => {
  const token = locals.token;

  try {
    const response = await fetch(BACKEND_URL + "/api/matches", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      console.error(`Backend responded with ${response.status}`);
      return json([], { status: response.status });
    }

    const matchedProjects = await response.json();
    return json(matchedProjects);
  } catch (error) {
    console.error("Error fetching project matches:", error);
    return json([], { status: 500 });
  }
};
