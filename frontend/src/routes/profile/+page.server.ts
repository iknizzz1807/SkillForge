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
        // Get the business info directly from response
        const data = await response.json();

        businessInfo = {
          companyType: data.company_type || "",
          founded: data.founded || "",
          companySize: data.company_size || "",
          website: data.website || "",
          aboutUs: data.about_us || "",
        };
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
    businessInfo,
  };
}) satisfies PageServerLoad;
