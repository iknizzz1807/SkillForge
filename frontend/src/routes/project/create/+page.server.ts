import { redirect } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { type Cookies, fail } from "@sveltejs/kit";

// type ProjectInput = {
//   title: string;
//   description: string;
//   skills: string[];
//   timeline: string;
// };

// type ProjectDisplay = {
//   id: string;
//   title: string;
//   description: string;
//   skills: string[];
//   timeline: string;
//   created_by: string;
//   status: string;
//   repo_url: string;
//   created_at: string;
// };

export const load = (async ({ parent }) => {
  const parentData = await parent();
  if (parentData.role !== "business") throw redirect(302, "/project");
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
      const max_member = parseInt(
        data.get("max-member")?.toString() || "0",
        10
      ); // Convert to number
      const start_time = data.get("start-time")?.toString();
      const end_time = data.get("end-time")?.toString();

      const token = cookies.get("auth_token");
      if (!token) {
        return fail(401, {
          success: false,
          error: "Authentication required",
        });
      }

      if (
        !title ||
        !description ||
        !start_time ||
        skills.length === 0 ||
        !end_time ||
        !max_member
      ) {
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
          max_member,
          start_time: new Date(start_time).toISOString(), // Đảm bảo gửi đúng định dạng ISO
          end_time: new Date(end_time).toISOString(), // Đảm bảo gửi đúng định dạng ISO
        }),
      });

      if (!response.ok) {
        let errorMessage = "Failed to create project";
        try {
          const errorData = await response.json();
          errorMessage = errorData.error || errorMessage;
        } catch (e) {
          //
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
