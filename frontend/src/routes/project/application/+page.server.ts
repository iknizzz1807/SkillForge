import type { PageServerLoad } from "./$types";
import { env } from "$env/dynamic/private";

export const load = (async (event) => {
  const user = event.locals.user;
  const token = event.locals.token;

  if (!user) {
    return {
      applications: [],
      role: null,
      name: null,
      email: null,
    };
  }

  try {
    // Get applications based on user role
    // The endpoint returns different data for business vs student users
    const response = await fetch(`http://backend:8080/api/applications/me`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      console.error("Error fetching applications:", await response.text());
      return {
        applications: [],
        error: "Failed to fetch applications",
        role: user.role,
        name: user.name,
        email: user.email,
      };
    }

    const applications = await response.json();

    return {
      applications,
      role: user.role,
      name: user.name,
      email: user.email,
    };
  } catch (error) {
    console.error("Error in application load function:", error);
    return {
      applications: [],
      error: "An error occurred while fetching applications",
      role: user.role,
      name: user.name,
      email: user.email,
    };
  }
}) satisfies PageServerLoad;
