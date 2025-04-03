import type { PageServerLoad, Actions } from "./$types";
import { type Cookies, fail } from "@sveltejs/kit";

type ProjectDisplay = {
  id: string;
  title: string;
  description: string;
  skills: string[];
  start_time: Date;
  end_time: Date;
  created_by_id: string;
  created_by_name: string;
  max_member: number;
  current_member: number;
  status: string;
  created_at: string;
};

export const load = (async ({ fetch, cookies, parent }) => {
  try {
    // Get auth token from cookies
    const token = cookies.get("auth_token");

    const parentData = await parent();
    const role = parentData.role || "";
    const id = parentData.id || "";
    const name = parentData.userName || "";

    let response;

    if (role === "business") {
      // Make get request to the business API
      // Make API request with proper URL and headers
      response = await fetch("http://backend:8080/api/projects", {
        headers: {
          Authorization: token ? `Bearer ${token}` : "",
          "Content-Type": "application/json",
        },
      });
    } else if (role === "student") {
      // Make get request to the student API
      // Make API request with proper URL and headers
      response = await fetch("http://backend:8080/api/projects", {
        headers: {
          Authorization: token ? `Bearer ${token}` : "",
          "Content-Type": "application/json",
        },
      });
    }

    if (response) {
      if (!response.ok) {
        console.error("Error fetching projects:", response.statusText);
        return {
          projects: [],
          error: `Failed to load projects: ${response.statusText}`,
        };
      }
    }

    // Parse the response data
    const responseData = await response?.json();

    // Handle null response - treat as empty array
    if (responseData === null) {
      console.log("API returned null, treating as empty array");
      return {
        projects: [],
        error: null,
      };
    }

    // Check if the response data is an array
    if (!Array.isArray(responseData)) {
      console.error("API returned invalid data format:", responseData);
      return {
        projects: [],
        error: "Received invalid data format from server",
      };
    }

    // Map the returned data to our frontend format
    const projects: ProjectDisplay[] = responseData.map((project: any) => ({
      id: project.id || "",
      title: project.title || "",
      description: project.description || "",
      skills: Array.isArray(project.skills) ? project.skills : [],
      start_time: project.start_time || new Date().toISOString(),
      end_time: project.end_time || new Date().toISOString(),
      created_by_id: project.created_by_id || "",
      created_by_name: project.created_by_name || "",
      max_member: project.max_member || 0,
      current_member: project.current_member || 0,
      status: project.status || "",
      created_at: project.created_at || new Date().toISOString(),
    }));

    // Success
    return {
      role: role,
      projects,
      error: null,
      token: token,
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
    try {
      const data = await request.formData();

      const title = data.get("title")?.toString();
      const description = data.get("description")?.toString();
      const skills = data.getAll("skills").map((skill) => skill.toString());
      const timeline = data.get("timeline")?.toString();

      const token = cookies.get("auth_token");
      if (!token) {
        return fail(401, {
          success: false,
          error: "Authentication required",
        });
      }

      if (!title || !description || !timeline || skills.length === 0) {
        return fail(400, {
          success: false,
          error: "All fields are required",
        });
      }

      // Send data to API
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
          timeline,
        }),
      });

      if (!response.ok) {
        let errorMessage = "Failed to create project";
        try {
          const errorData = await response.json();
          errorMessage = errorData.error || errorMessage;
        } catch (e) {
          // Ignore JSON parsing error
        }

        return fail(response.status, {
          success: false,
          error: errorMessage,
        });
      }

      // Parse the created project
      const createdProject = await response.json();

      // Success
      return {
        success: true,
        project: createdProject,
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
