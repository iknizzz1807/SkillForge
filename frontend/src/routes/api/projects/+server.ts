import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

export const POST: RequestHandler = async ({ request, cookies, fetch }) => {
  try {
    const token = cookies.get("auth_token");
    if (!token) {
      throw new Error("Authentication required");
    }

    const data = await request.json();
    const { title, description, skills, max_member, start_time, end_time } =
      data;

    // Validate required fields
    if (
      !title ||
      !description ||
      !start_time ||
      !skills ||
      skills.length === 0 ||
      !end_time ||
      !max_member
    ) {
      throw new Error("All fields are required");
    }

    // Forward request to backend API
    const response = await fetch("http://backend:8080/api/projects", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        title,
        description,
        skills,
        max_member,
        start_time: new Date(start_time).toISOString(),
        end_time: new Date(end_time).toISOString(),
      }),
    });

    if (!response.ok) {
      let errorMessage = "Failed to create project";
      try {
        const errorData = await response.json();
        errorMessage = errorData.error || errorMessage;
      } catch (e) {
        // Ignore JSON parsing errors
      }

      throw new Error(errorMessage);
    }

    // Return the created project
    const createdProject = await response.json();
    return json(createdProject);
  } catch (err: any) {
    console.error("Error creating project:", err);
    return json(
      { error: err.message || "An unexpected error occurred" },
      { status: err.status || 500 }
    );
  }
};
