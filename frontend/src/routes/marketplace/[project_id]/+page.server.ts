import type { PageServerLoad } from "./$types";
import { error } from "@sveltejs/kit";

export const load = (async ({ params, fetch, cookies }) => {
  const project_id = params.project_id;
  const token = cookies.get("auth_token"); // For auth

  try {
    // Fetch project based on project_id
    const response = await fetch(
      `http://backend:8080/api/projects/${project_id}`,
      {
        headers: {
          Authorization: token ? `Bearer ${token}` : "",
          "Content-Type": "application/json",
        },
      }
    );

    if (!response.ok) {
      throw new Error(`Failed to fetch project: ${response.statusText}`);
    }

    const project = await response.json();

    // You can fetch additional related data here if needed
    // For example, project tasks, team members, etc.

    // console.log(project);

    return {
      project,
      token,
      // Add any other data you want to return
    };
  } catch (err) {
    console.error("Error loading project:", err);
    throw error(404, {
      message: "Project not found or failed to load",
    });
  }
}) satisfies PageServerLoad;
