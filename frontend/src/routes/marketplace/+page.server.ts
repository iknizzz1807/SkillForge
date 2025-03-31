import type { PageServerLoad } from "./$types";

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
    const role = parentData.role || "user";

    // Make API request with proper URL and headers
    const response = await fetch("http://backend:8080/api/projects", {
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
    const responseData = await response.json();

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
    };
  } catch (error) {
    console.error("Error loading projects:", error);
    return {
      projects: [],
      error: "Failed to load projects. Please try again later.",
    };
  }
}) satisfies PageServerLoad;
