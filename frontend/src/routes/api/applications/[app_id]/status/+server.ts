import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

export const PUT: RequestHandler = async ({ params, request, locals }) => {
  const user = locals.user;
  const applicationId = params.app_id;
  const token = locals.token;

  if (!applicationId) {
    return json({ error: "Missing application ID" }, { status: 400 });
  }

  if (!user || !token) {
    return json({ error: "Unauthorized" }, { status: 401 });
  }

  // Only business users can update application status
  if (user.role !== "business") {
    return json(
      { error: "Only businesses can update application status" },
      { status: 403 }
    );
  }

  try {
    const body = await request.json();

    // Validate the status
    if (
      !body.status ||
      !["approved", "rejected", "pending"].includes(body.status)
    ) {
      return json({ error: "Invalid status value" }, { status: 400 });
    }

    const response = await fetch(
      `http://backend:8080/api/applications/status/${applicationId}`,
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          status: body.status,
        }),
      }
    );

    if (!response.ok) {
      const errorText = await response.text();
      console.error(
        `Error updating application ${applicationId} status:`,
        errorText
      );

      return json(
        {
          error: `Failed to update application status: ${response.statusText}`,
          details: errorText,
        },
        { status: response.status }
      );
    }

    const result = await response.json();
    return json(result);
  } catch (err) {
    console.error(
      `Error processing application ${applicationId} status update:`,
      err
    );
    return json(
      { error: "Failed to process application status update" },
      { status: 500 }
    );
  }
};
