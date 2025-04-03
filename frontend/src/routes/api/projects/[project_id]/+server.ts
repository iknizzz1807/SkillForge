import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

// DELETE endpoint - for deleting a project
export const DELETE: RequestHandler = async ({ params, fetch, request }) => {
  const id = params.project_id;
  const authHeader = request.headers.get("Authorization");

  if (!authHeader) {
    return json({ error: "Authorization header missing" }, { status: 401 });
  }

  try {
    const response = await fetch(`http://backend:8080/api/projects/${id}`, {
      method: "DELETE",
      headers: {
        Authorization: authHeader,
      },
    });

    if (!response.ok) {
      const errorData = await response.json();
      return json(
        { error: errorData.error || "Failed to delete project" },
        { status: response.status }
      );
    }

    return json({ success: true });
  } catch (error) {
    console.error("Error deleting project:", error);
    return json(
      { error: "Server error while deleting project" },
      { status: 500 }
    );
  }
};

// PUT endpoint - for updating a project
export const PUT: RequestHandler = async ({ params, fetch, request }) => {
  const id = params.project_id;
  const authHeader = request.headers.get("Authorization");

  if (!authHeader) {
    return json({ error: "Authorization header missing" }, { status: 401 });
  }

  try {
    const requestData = await request.json();

    const response = await fetch(`http://backend:8080/api/projects/${id}`, {
      method: "PUT",
      body: JSON.stringify(requestData),
      headers: {
        "Content-Type": "application/json",
        Authorization: authHeader,
      },
    });

    if (!response.ok) {
      const errorData = await response.json();
      return json(
        { error: errorData.error || "Failed to update project" },
        { status: response.status }
      );
    }

    const updatedProject = await response.json();
    return json(updatedProject);
  } catch (error) {
    console.error("Error updating project:", error);
    return json(
      { error: "Server error while updating project" },
      { status: 500 }
    );
  }
};
