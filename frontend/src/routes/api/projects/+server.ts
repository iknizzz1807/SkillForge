import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

export const POST: RequestHandler = async ({ request, locals, fetch }) => {
  try {
    const token = locals.token;
    if (!token) {
      return json({ error: "Authentication required" }, { status: 401 });
    }

    const data = await request.json();
    const {
      title,
      description,
      skills,
      max_member,
      start_time,
      end_time,
      difficulty,
    } = data;

    if (
      !title ||
      !description ||
      !start_time ||
      !skills ||
      skills.length === 0 ||
      !end_time ||
      !max_member ||
      !difficulty
    ) {
      return json({ error: "All fields are required" }, { status: 400 });
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
        difficulty,
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

      return json({ error: errorMessage }, { status: response.status });
    }

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
