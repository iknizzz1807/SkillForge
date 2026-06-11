import type { PageServerLoad } from "./$types";
import { BACKEND_URL } from '$lib/backend';

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
  difficulty: string;
  status: string;
  created_at: string;
};

export const load = (async ({ fetch, locals, parent }) => {
  try {
    const token = locals.token;
    const parentData = await parent();
    const role = parentData.role || "user";

    // Make API request with proper URL and headers
    const response = await fetch(BACKEND_URL + "/api/projects", {
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

    const responseData = await response.json();

    if (responseData === null) {
      return {
        projects: [],
        error: null,
      };
    }

    if (!Array.isArray(responseData)) {
      return {
        projects: [],
        error: "Received invalid data format from server",
      };
    }

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
      difficulty: project.difficulty || "",
      created_at: project.created_at || new Date().toISOString(),
    }));

    // Fetch user's applications (student only) to show status badges on marketplace
    let userApplications: { project_id: string; status: string }[] = [];
    if (role === "student" && token) {
      try {
        const appsResponse = await fetch(BACKEND_URL + "/api/applications/me", {
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        });
        if (appsResponse.ok) {
          const appsData = await appsResponse.json();
          if (Array.isArray(appsData)) {
            userApplications = appsData.map((app: any) => ({
              project_id: app.project_id || "",
              status: app.status || "pending",
            }));
          }
        }
      } catch (appErr) {
        console.error("Error fetching user applications:", appErr);
      }
    }

    return {
      role: role,
      projects,
      userApplications,
      error: null,
    };
  } catch (error) {
    console.error("Error loading projects:", error);
    return {
      projects: [],
      userApplications: [],
      error: "Failed to load projects. Please try again later.",
    };
  }
}) satisfies PageServerLoad;
