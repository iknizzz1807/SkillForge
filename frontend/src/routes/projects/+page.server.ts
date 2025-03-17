import type { PageServerLoad, Actions } from "./$types";
import { type Cookies, fail } from "@sveltejs/kit";

type Project = {
  title: string;
  description: string;
  category: string;
  timeline: string;
};

export const load = (async ({ fetch, cookies }) => {
  try {
    // Get auth token from cookies
    const token = cookies.get("auth_token");

    // Make API request with proper URL and headers
    const response = await fetch("http://localhost:8080/api/projects", {
      headers: {
        Authorization: token ? `Bearer ${token}` : "",
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      console.error("Error fetching projects:", response.statusText);
      return {
        projects: [],
        error: `Failed to load projects: ${response.statusText}`,
      };
    }

    // Parse the response data
    const projectsData: Project[] = await response.json();

    // Map the returned data to our frontend format
    const projects: Project[] = projectsData.map((project: any) => ({
      title: project.title,
      description: project.description,
      // skills: project.skills || [],
      timeline: project.timeline,
      // status: project.status,
      category: project.category,
    }));

    return {
      projects,
      error: null,
    };
  } catch (error) {
    console.error("Error loading projects:", error);
    return {
      projects: [],
      error: "Failed to load projects. Please try again later.",
    };
  }
}) satisfies PageServerLoad;

export const actions = {
  default: async ({
    request,
    cookies,
  }: {
    request: Request;
    cookies: Cookies;
  }) => {
    let message: string = "";
    let error: string = "";
    const projects: Project[] = [];

    try {
      const data = await request.formData();

      const title = data.get("title");
      const description = data.get("description");
      const category = data.get("category");
      const timeline = data.get("timeline");

      const token = cookies.get("auth_token");
      if (!token) {
        return fail(401, {
          success: false,
          error: "Authentication required",
        });
      }

      if (!title || !description || !timeline || !category) {
        return fail(400, {
          success: false,
          error: "All fields are required",
          data: { title, description, timeline, category },
        });
      }

      // Send data to API
      const response = await fetch("http://localhost:8080/api/projects", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          title,
          description,
          timeline,
          // Additional fields required by your API
          status: "open",
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        return fail(response.status, {
          success: false,
          error: errorData.error || "Failed to create project",
          data: { title, description, timeline, category },
        });
      }

      // Success
      return {
        success: true,
        message: "Project created successfully!",
      };
    } catch (error) {
      console.error("Error creating project:", error);
      return fail(500, {
        success: false,
        error: "An unexpected error occurred",
      });
    }
  },
} satisfies Actions;
