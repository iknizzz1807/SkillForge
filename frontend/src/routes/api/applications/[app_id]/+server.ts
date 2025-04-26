import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

export const DELETE: RequestHandler = async ({
  params,
  fetch,
  request,
  locals,
}) => {
  try {
    const applicationId = params.app_id;

    const token = locals.token;

    if (!token) {
      return json({ error: "Unauthorized" }, { status: 401 });
    }

    const response = await fetch(
      `http://backend:8080/api/applications/${applicationId}`,
      {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      }
    );

    if (!response.ok) {
      const errorData = await response.json();
      return json(errorData, { status: response.status });
    }

    return json({ success: true });
  } catch (error) {
    console.error("Error deleting application:", error);
    return json({ error: "Failed to delete application" }, { status: 500 });
  }
};
