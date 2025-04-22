import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

// GET endpoint to retrieve business info
export const GET: RequestHandler = async ({ fetch, locals }) => {
  try {
    if (!locals.token) {
      return new Response(JSON.stringify({ error: "Unauthorized" }), {
        status: 401,
        headers: { "Content-Type": "application/json" },
      });
    }

    // Forward the request to the backend API
    const response = await fetch(`http://backend:8080/api/business-info`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${locals.token}`,
      },
    });

    const data = await response.json();

    if (!response.ok) {
      return new Response(
        JSON.stringify({
          error: data.error || "Failed to fetch business info",
        }),
        {
          status: response.status,
          headers: { "Content-Type": "application/json" },
        }
      );
    }

    return json(data);
  } catch (error) {
    console.error("Error fetching business info:", error);
    return new Response(JSON.stringify({ error: "Internal server error" }), {
      status: 500,
      headers: { "Content-Type": "application/json" },
    });
  }
};

// PUT endpoint to update business info
export const PUT: RequestHandler = async ({ request, fetch, locals }) => {
  try {
    if (!locals.token) {
      return new Response(JSON.stringify({ error: "Unauthorized" }), {
        status: 401,
        headers: { "Content-Type": "application/json" },
      });
    }

    // Get form data from request
    const formData = await request.formData();

    // Forward the request to the backend API
    const response = await fetch(`http://backend:8080/api/business-info`, {
      method: "PUT",
      headers: {
        Authorization: `Bearer ${locals.token}`,
      },
      body: formData,
    });

    const data = await response.json();

    if (!response.ok) {
      return new Response(
        JSON.stringify({
          error: data.error || "Failed to update business info",
        }),
        {
          status: response.status,
          headers: { "Content-Type": "application/json" },
        }
      );
    }

    return json({
      success: true,
      businessInfo: data,
    });
  } catch (error) {
    console.error("Error updating business info:", error);
    return new Response(JSON.stringify({ error: "Internal server error" }), {
      status: 500,
      headers: { "Content-Type": "application/json" },
    });
  }
};
