import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

export const PUT: RequestHandler = async ({ params, request, cookies, fetch }) => {
  try {
    const invitationId = params.invitation_id;
    if (!invitationId) {
      return json({ error: "Missing invitation ID" }, { status: 400 });
    }

    const token = cookies.get("auth_token");
    if (!token) {
      return json({ error: "Unauthorized" }, { status: 401 });
    }

    const body = await request.json();

    const response = await fetch(
      `http://backend:8080/api/invitations/${invitationId}/respond`,
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(body),
      }
    );

    if (!response.ok) {
      const errorText = await response.text();
      return json(
        { error: errorText || "Failed to respond to invitation" },
        { status: response.status }
      );
    }

    const data = await response.json();
    return json(data);
  } catch (err) {
    console.error("Error responding to invitation:", err);
    return json({ error: "Failed to respond to invitation" }, { status: 500 });
  }
};
