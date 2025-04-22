import type { LayoutServerLoad } from "./$types";

// Modify your +layout.server.ts file
export const load = (async (event) => {
  // Get auth token from cookies
  const token = event.cookies.get("auth_token");

  // Get current path and origin
  const path = event.url.pathname;
  const origin = event.url.origin;

  // Default empty user data
  let userData: {
    role: string | null;
    id: string | null;
    userName: string | null;
    avatarUrl: string | null;
    title: string | null;
  } = {
    role: null,
    id: null,
    userName: null,
    avatarUrl: null,
    title: null,
  };

  // Only fetch user data if we have an auth token
  if (token) {
    try {
      // Make request to backend API
      const response = await fetch("http://backend:8080/api/user", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        const user = await response.json();

        // Update userData with fetched data
        userData = {
          role: user.role,
          id: user.id,
          userName: user.name,
          title: user.title, // Include title from backend
          // Keep using the SvelteKit API endpoint for avatar
          avatarUrl: `${origin}/api/avatars`,
        };
      }
    } catch (error) {
      console.error("Error fetching user data:", error);
    }
  }

  // Check if we need to redirect
  if (userData.id && path === "/") {
    return {
      ...userData,
      // Uncomment if you want to redirect
      // redirect: {
      //   status: 302,
      //   location: "/dashboard",
      // },
    };
  }

  // Return user data without redirecting
  return userData;
}) satisfies LayoutServerLoad;
