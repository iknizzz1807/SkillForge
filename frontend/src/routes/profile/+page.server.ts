import type { PageServerLoad } from "./$types";

export const load = (async ({ locals, fetch, url, parent }) => {
  const user = locals.user;
  const origin = url.origin;
  const token = locals.token;
  const parentData = await parent();
  const title = parentData.title;

  // Initialize businessInfo with default empty values
  let businessInfo = {
    companyType: "",
    founded: "",
    companySize: "",
    website: "",
    aboutUs: "",
  };

  let userData = {};

  try {
    const response = await fetch("/api/profile", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (response.ok) {
      // Kiểm tra content-type trước khi parse JSON
      const contentType = response.headers.get("content-type");
      if (contentType && contentType.includes("application/json")) {
        userData = await response.json();
      } else {
        console.error("API didn't return JSON content:", await response.text());
      }
    } else {
      console.error(
        "Profile API error:",
        response.status,
        await response.text()
      );
    }
  } catch (error) {
    console.error("Failed to fetch user data:", error);
  }

  // Only fetch business info if user is logged in and the role is business
  if (token && user?.role === "business") {
    try {
      const response = await fetch("/api/business-info", {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        // Kiểm tra content-type trước khi parse JSON
        const contentType = response.headers.get("content-type");
        if (contentType && contentType.includes("application/json")) {
          const data = await response.json();
          businessInfo = {
            companyType: data.company_type || "",
            founded: data.founded || "",
            companySize: data.company_size || "",
            website: data.website || "",
            aboutUs: data.about_us || "",
          };
        } else {
          console.error(
            "Business API didn't return JSON:",
            await response.text()
          );
        }
      } else {
        console.error(
          "Business API error:",
          response.status,
          await response.text()
        );
      }
    } catch (error) {
      console.error("Failed to fetch business info:", error);
    }
  }

  return {
    id: user?.id,
    name: user?.name,
    email: user?.email,
    role: user?.role,
    title: title,
    avatarUrl: `${origin}/api/avatars` || null,
    userData,
    businessInfo,
  };
}) satisfies PageServerLoad;
