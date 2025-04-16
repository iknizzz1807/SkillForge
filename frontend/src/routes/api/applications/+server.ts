import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

export const POST: RequestHandler = async ({ request, locals }) => {
  const user = locals.user;
  const token = locals.token;

  if (!user || !token) {
    return json({ error: "Unauthorized" }, { status: 401 });
  }

  // Only students can submit applications
  if (user.role !== "student") {
    return json(
      { error: "Only students can apply to projects" },
      { status: 403 }
    );
  }

  try {
    const body = await request.json();

    // Validate required fields
    if (!body.project_id || !body.motivation || !body.detailed_proposal) {
      return json({ error: "Missing required fields" }, { status: 400 });
    }

    // Forward the request to the backend
    const response = await fetch(`http://backend:8080/api/applications`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        project_id: body.project_id,
        motivation: body.motivation,
        detailed_proposal: body.detailed_proposal,
      }),
    });

    if (!response.ok) {
      const errorData = await response.text();
      console.error("Backend application submission error:", errorData);

      try {
        const parsedError = JSON.parse(errorData);
        return json(
          { error: parsedError.error || "Failed to submit application" },
          { status: response.status }
        );
      } catch {
        return json(
          { error: "Failed to submit application" },
          { status: response.status }
        );
      }
    }

    const result = await response.json();
    return json(result);
  } catch (err) {
    console.error("Error processing application submission:", err);
    return json(
      { error: "Failed to process application submission" },
      { status: 500 }
    );
  }
};
