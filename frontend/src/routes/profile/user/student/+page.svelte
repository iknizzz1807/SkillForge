<script lang="ts">
  import { onMount } from "svelte";
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  // Người xem hiện tại (viewer) - có thể lấy từ store hoặc params
  // Đây là mock data, thực tế nên lấy từ API
  let currentUser = {
    id: "current-user-123",
    role: "business", // hoặc "student"
    name: "Your Name",
  };

  let studentXP = $state(750);
  let currentLevel = $state(3);
  let xpForNextLevel = $state(250);
  let xpProgress = $state(75); // percentage to next level

  // Mock data cho profile đang xem (profileData)
  // Thực tế sẽ lấy từ API dựa trên ID trong URL
  let profileData = {
    id: "profile-456",
    name: "John Smith",
    email: "john.smith@example.com",
    title: "Full Stack Developer",
    role: "student", // hoặc "business"
    avatarUrl: "https://randomuser.me/api/portraits/men/44.jpg",
    shortBio:
      "Passionate developer with 3 years of experience in web technologies.",
    skills: [
      { name: "React", level: "Advanced", percentage: 85 },
      { name: "Node.js", level: "Intermediate", percentage: 65 },
      { name: "JavaScript", level: "Advanced", percentage: 90 },
      { name: "MongoDB", level: "Intermediate", percentage: 60 },
      { name: "Python", level: "Beginner", percentage: 30 },
    ],

    certificates: [
      {
        name: "React Development",
        issuer: "TechCorp",
        date: "18/03/2025",
      },
      {
        name: "Web Development Basics",
        issuer: "SKILLFORGE",
        date: "10/03/2025",
      },
    ],
    badges: [
      {
        name: "First Project",
        icon: "🏆",
        description: "Completed first project",
        achieved: true,
        date: "10/02/2025",
      },
      {
        name: "Team Player",
        icon: "👥",
        description: "Collaborated with 3+ students",
        achieved: true,
        date: "01/03/2025",
      },
    ],
    stats: {
      completedProjects: 12,
      certificates: 8,
      onTimeCompletion: "98%",
      rating: 4.9,
    },
    businessInfo: {
      companyType: "Technology Solutions",
      founded: "2018",
      companySize: "15-50 employees",
      website: "www.techsolutions.com",
      aboutUs:
        "Tech Solutions Inc. is a forward-thinking technology company specializing in developing innovative web and mobile applications.",
    },
  };

  // State cho các thao tác UI
  let isInTalentPool = $state(false);
  let talentPoolSuccess = $state(false);

  // Hàm chức năng - thực tế sẽ gọi API
  const addToTalentPool = async () => {
    // API Call: POST /api/talent-pool
    // Body: { studentId: profileData.id }
    console.log("Adding to talent pool:", profileData.id);

    // Mock successful API response
    setTimeout(() => {
      isInTalentPool = true;
      talentPoolSuccess = true;

      // Auto-hide success message after 3 seconds
      setTimeout(() => {
        talentPoolSuccess = false;
      }, 3000);
    }, 800);
  };

  onMount(() => {
    // API call to fetch profile data
    // GET /api/profile/{id}
    console.log("Fetching profile data for:", data.id);

    // API call to check if student is already in talent pool (if viewer is business)
    // GET /api/talent-pool/check/{studentId}
    if (currentUser.role === "business" && profileData.role === "student") {
      console.log("Checking if student is in talent pool");
      // Mock response - would come from API
      isInTalentPool = Math.random() > 0.5;
    }
  });
</script>

<main class="flex-1 pr-4 pl-4 ml-64 pt-4">
  <!-- Header section -->
  <div class="flex justify-between items-center mb-6">
    <div>
      <h1 class="text-2xl font-bold flex items-center">
        {profileData.name}'s Profile
        {#if profileData.role === "student"}
          <span
            class="ml-2 text-xs bg-purple-100 text-purple-800 px-2 py-1 rounded"
            >Student</span
          >
        {:else}
          <span class="ml-2 text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded"
            >Business</span
          >
        {/if}
      </h1>
      <p class="text-gray-500">{profileData.title}</p>
    </div>

    <!-- Action buttons based on viewer role and viewed profile role -->
    <div class="flex space-x-2">
      {#if currentUser.role === "business" && profileData.role === "student"}
        {#if isInTalentPool}
          <button
            class="bg-gray-100 text-gray-600 px-4 py-2 rounded flex items-center"
            disabled
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4 mr-1"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clip-rule="evenodd"
              />
            </svg>
            In Talent Pool
          </button>
        {:else}
          <button class="btn px-4 py-2" onclick={addToTalentPool}>
            Add to Talent Pool
          </button>
        {/if}
      {/if}
    </div>
  </div>

  {#if talentPoolSuccess}
    <div class="bg-green-50 border border-green-200 rounded-md p-4 mb-4">
      <div class="flex">
        <svg
          class="h-5 w-5 text-green-500 mr-2"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M5 13l4 4L19 7"
          ></path>
        </svg>
        <span class="text-green-700"
          >Added to your talent pool successfully!</span
        >
      </div>
    </div>
  {/if}

  <!-- Profile Content -->
  <div class="grid grid-cols-3 gap-4">
    <!-- Left Column - Basic Info -->
    <div class="space-y-4">
      <!-- Profile Card -->
      <div class="card p-4">
        <div class="flex flex-col items-center mb-4">
          <div class="relative">
            <img
              class="w-24 h-24 rounded-full object-cover border-2 border-[#6b48ff]"
              src={profileData.avatarUrl}
              alt="{profileData.name}'s Avatar"
            />
            <div
              class="absolute -bottom-1 -right-1 bg-white rounded-full p-0.5 shadow-sm"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4 text-green-500"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>
          </div>
          <h3 class="text-lg font-semibold text-center mt-3">
            {profileData.name}
          </h3>
          <p
            class="text-sm text-gray-500 text-center flex items-center justify-center"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4 mr-1 text-gray-400"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M6 6V5a3 3 0 013-3h2a3 3 0 013 3v1h2a2 2 0 012 2v3.57A22.952 22.952 0 0110 13a22.95 22.95 0 01-8-1.43V8a2 2 0 012-2h2zm2-1a1 1 0 011-1h2a1 1 0 011 1v1H8V5zm1 5a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1z"
                clip-rule="evenodd"
              />
              <path
                d="M2 13.692V16a2 2 0 002 2h12a2 2 0 002-2v-2.308A24.974 24.974 0 0110 15c-2.796 0-5.487-.46-8-1.308z"
              />
            </svg>
            {profileData.title}
          </p>
        </div>

        <!-- Add level system UI -->
        <div class="mt-4 p-3 bg-gray-50 rounded-lg border border-gray-100">
          <div class="flex justify-between items-center mb-1">
            <div class="flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4 mr-1 text-[#6b48ff]"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path d="M13.5 3a1.5 1.5 0 013 0v7.5a1.5 1.5 0 01-3 0V3z" />
                <path d="M3.5 3a1.5 1.5 0 013 0v7.5a1.5 1.5 0 01-3 0V3z" />
                <path
                  d="M3.5 10.5a1.5 1.5 0 013 0v6a1.5 1.5 0 01-3 0v-6zM13.5 10.5a1.5 1.5 0 013 0v6a1.5 1.5 0 01-3 0v-6z"
                />
              </svg>
              <span class="text-sm font-medium">Level 3</span>
            </div>
            <span class="text-xs bg-[#6b48ff] text-white px-2 py-1 rounded-full"
              >750 XP</span
            >
          </div>
          <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
            <div class="bg-[#6b48ff] h-full" style="width: 75%"></div>
          </div>
          <div class="flex justify-between mt-1 text-xs text-gray-500">
            <div class="flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-3 w-3 mr-1"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-8.707l-3-3a1 1 0 00-1.414 0l-3 3a1 1 0 001.414 1.414L9 9.414V13a1 1 0 102 0V9.414l1.293 1.293a1 1 0 001.414-1.414z"
                  clip-rule="evenodd"
                />
              </svg>
              Current: Level 3
            </div>
            <div class="flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-3 w-3 mr-1"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"
                />
              </svg>
              Next: 250 XP for Level 4
            </div>
          </div>
        </div>

        <div class="space-y-3 mt-4">
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4 mr-2 text-gray-400"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z"
                />
                <path
                  d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z"
                />
              </svg>
              Email:
            </span>
            <span class="text-sm text-gray-600">{profileData.email}</span>
          </div>
          {#if profileData.shortBio}
            <div>
              <span class="text-sm font-medium flex items-center mb-1">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4 mr-2 text-gray-400"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                    clip-rule="evenodd"
                  />
                </svg>
                About:
              </span>
              <p class="text-sm text-gray-600">{profileData.shortBio}</p>
            </div>
          {/if}
        </div>

        {#if profileData.role === "student"}
          <!-- Stats for student with improved icons -->
          <div class="grid grid-cols-2 gap-2 mt-4">
            <div class="bg-gray-50 rounded p-2 text-center">
              <div class="flex justify-center mb-1">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5 text-[#6b48ff]"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M6 6V5a3 3 0 013-3h2a3 3 0 013 3v1h2a2 2 0 012 2v3.57A22.952 22.952 0 0110 13a22.95 22.95 0 01-8-1.43V8a2 2 0 012-2h2zm2-1a1 1 0 011-1h2a1 1 0 011 1v1H8V5zm1 5a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1z"
                    clip-rule="evenodd"
                  />
                </svg>
              </div>
              <p class="text-xl font-bold text-[#6b48ff]">
                {profileData.stats.completedProjects}
              </p>
              <p class="text-xs text-gray-500">Completed projects</p>
            </div>
            <div class="bg-gray-50 rounded p-2 text-center">
              <div class="flex justify-center mb-1">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5 text-[#6b48ff]"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
                  />
                </svg>
              </div>
              <p class="text-xl font-bold text-[#6b48ff]">
                {profileData.stats.rating}
              </p>
              <p class="text-xs text-gray-500">Rating</p>
            </div>
          </div>
        {/if}
      </div>

      {#if profileData.role === "business"}
        <!-- Business Info -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Company Information</h3>
          <div class="space-y-3">
            <div>
              <p class="text-sm font-medium">Company Type</p>
              <p class="text-sm text-gray-600">
                {profileData.businessInfo.companyType || "Not provided"}
              </p>
            </div>
            <div>
              <p class="text-sm font-medium">Founded</p>
              <p class="text-sm text-gray-600">
                {profileData.businessInfo.founded || "Not provided"}
              </p>
            </div>
            <div>
              <p class="text-sm font-medium">Size</p>
              <p class="text-sm text-gray-600">
                {profileData.businessInfo.companySize || "Not provided"}
              </p>
            </div>
            <div>
              <p class="text-sm font-medium">Website</p>
              {#if profileData.businessInfo.website}
                <a
                  href={profileData.businessInfo.website.startsWith("http")
                    ? profileData.businessInfo.website
                    : `https://${profileData.businessInfo.website}`}
                  target="_blank"
                  class="text-sm text-[#6b48ff] hover:underline"
                >
                  {profileData.businessInfo.website}
                </a>
              {:else}
                <p class="text-sm text-gray-600">Not provided</p>
              {/if}
            </div>
            <div>
              <p class="text-sm font-medium">About Us</p>
              <p class="text-sm text-gray-600">
                {profileData.businessInfo.aboutUs || "Not provided"}
              </p>
            </div>
          </div>
        </div>
      {/if}

      {#if profileData.role === "student"}
        <!-- Student Badges & Achievements -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Badges & Achievements</h3>
          <div class="grid grid-cols-2 gap-2">
            {#each profileData.badges as badge}
              <div
                class="text-center p-2 rounded border border-[#6b48ff] bg-purple-50"
              >
                <div class="text-2xl mb-1">{badge.icon}</div>
                <p class="text-sm font-medium">{badge.name}</p>
                <p class="text-xs text-gray-500">{badge.description}</p>
                <p class="text-xs text-[#6b48ff] mt-1">Earned {badge.date}</p>
              </div>
            {/each}
          </div>
        </div>
      {/if}
    </div>

    <!-- Middle Column - Skills & Projects -->
    <div class="space-y-4">
      {#if profileData.role === "student"}
        <!-- Student Skills -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Skills & Expertise</h3>
          <div class="space-y-4">
            {#each profileData.skills as skill}
              <div>
                <div class="flex justify-between mb-1">
                  <span class="text-sm font-medium">{skill.name}</span>
                  <span class="text-xs text-[#6b48ff]">{skill.level}</span>
                </div>
                <div class="progress-bar">
                  <div
                    class="progress-fill"
                    style="width: {skill.percentage}%"
                  ></div>
                </div>
              </div>
            {/each}
          </div>
        </div>

        <!-- Student Certificates -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Certificates</h3>
          <div class="space-y-3">
            {#each profileData.certificates as certificate}
              <div class="flex justify-between items-center">
                <div>
                  <p class="text-sm font-medium">{certificate.name}</p>
                  <p class="text-xs text-gray-500">
                    Issued: {certificate.date} by {certificate.issuer}
                  </p>
                </div>
                <div class="flex items-center">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-5 w-5 text-[#6b48ff]"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                    <path
                      fill-rule="evenodd"
                      d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                      clip-rule="evenodd"
                    />
                  </svg>
                </div>
              </div>
            {/each}
          </div>
        </div>
      {:else if profileData.role === "business"}
        <!-- Business Active Projects -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Current Projects</h3>
          <div class="space-y-3">
            <div class="p-3 bg-gray-50 rounded border border-gray-200">
              <div class="flex justify-between mb-1">
                <p class="text-sm font-medium">Mobile Banking App</p>
                <span
                  class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded"
                  >In Progress</span
                >
              </div>
              <p class="text-xs text-gray-500 mb-2">
                Skills: React Native, Node.js | Due: 25/05/2025
              </p>
              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-2">
                <div
                  class="bg-blue-500 h-1.5 rounded-full"
                  style="width: 65%"
                ></div>
              </div>
            </div>

            <div class="p-3 bg-gray-50 rounded border border-gray-200">
              <div class="flex justify-between mb-1">
                <p class="text-sm font-medium">E-commerce Dashboard</p>
                <span
                  class="text-xs bg-yellow-100 text-yellow-800 px-2 py-1 rounded"
                  >Just Started</span
                >
              </div>
              <p class="text-xs text-gray-500 mb-2">
                Skills: React, MongoDB, Express | Due: 10/06/2025
              </p>
              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-2">
                <div
                  class="bg-yellow-500 h-1.5 rounded-full"
                  style="width: 15%"
                ></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Business Completed Projects -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Completed Projects</h3>
          <div class="space-y-3">
            <div
              class="flex justify-between items-center p-2 bg-gray-50 rounded border border-gray-200"
            >
              <div>
                <p class="text-sm font-medium">Company Website Redesign</p>
                <p class="text-xs text-gray-500">
                  Completed: 10/03/2025 • 5 students
                </p>
              </div>
              <div class="flex items-center">
                <span
                  class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded mr-2"
                  >4.8 ★</span
                >
              </div>
            </div>

            <div
              class="flex justify-between items-center p-2 bg-gray-50 rounded border border-gray-200"
            >
              <div>
                <p class="text-sm font-medium">CRM Integration</p>
                <p class="text-xs text-gray-500">
                  Completed: 25/02/2025 • 3 students
                </p>
              </div>
              <div class="flex items-center">
                <span
                  class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded mr-2"
                  >4.9 ★</span
                >
              </div>
            </div>
          </div>
        </div>
      {/if}
    </div>

    <!-- Right Column - Additional Info -->
    <div class="space-y-4">
      {#if profileData.role === "student"}
        <!-- Student Rating & Feedback -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Rating & Feedback</h3>
          <div class="text-center mb-4">
            <p class="text-3xl font-bold text-[#6b48ff]">
              {profileData.stats.rating}
            </p>
            <div class="flex justify-center text-yellow-400 text-lg">
              <span>★</span><span>★</span><span>★</span><span>★</span><span
                >★</span
              >
            </div>
            <p class="text-xs text-gray-500 mt-1">
              Overall Rating ({profileData.stats.completedProjects} projects)
            </p>
          </div>

          <!-- Sample Feedbacks -->
          <div class="space-y-3">
            <div class="p-3 bg-gray-50 rounded">
              <div class="flex items-center mb-1">
                <p class="text-sm font-medium flex-1">Tech Solutions Inc.</p>
                <div class="text-yellow-400 text-xs">★★★★★</div>
              </div>
              <p class="text-xs text-gray-600">
                "John delivered exceptional work on our e-commerce project.
                Great communication and technical skills."
              </p>
              <p class="text-xs text-gray-400 mt-1">
                Project: E-commerce Dashboard
              </p>
            </div>

            <div class="p-3 bg-gray-50 rounded">
              <div class="flex items-center mb-1">
                <p class="text-sm font-medium flex-1">InnovateCorp</p>
                <div class="text-yellow-400 text-xs">
                  ★★★★<span class="text-gray-300">★</span>
                </div>
              </div>
              <p class="text-xs text-gray-600">
                "Very skilled developer who delivered on time. Would work with
                again."
              </p>
              <p class="text-xs text-gray-400 mt-1">
                Project: Build a Chat App
              </p>
            </div>
          </div>
        </div>

        <!-- Student Stats -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Performance Stats</h3>
          <div class="grid grid-cols-2 gap-3">
            <div class="text-center p-2 bg-gray-50 rounded">
              <p class="text-xl font-bold text-[#6b48ff]">
                {profileData.stats.completedProjects}
              </p>
              <p class="text-xs text-gray-500">Completed Projects</p>
            </div>
            <div class="text-center p-2 bg-gray-50 rounded">
              <p class="text-xl font-bold text-[#6b48ff]">
                {profileData.stats.certificates}
              </p>
              <p class="text-xs text-gray-500">Certificates</p>
            </div>

            <div class="text-center p-2 bg-gray-50 rounded">
              <p class="text-xl font-bold text-[#6b48ff]">
                {profileData.stats.rating}
              </p>
              <p class="text-xs text-gray-500">Client Rating</p>
            </div>
          </div>
        </div>
      {:else if profileData.role === "business"}
        <!-- Business Student Feedback -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Student Feedback</h3>
          <div class="flex justify-center items-center py-3">
            <div class="text-center mr-6">
              <p class="text-3xl font-bold text-[#6b48ff]">4.8</p>
              <div class="flex text-yellow-400 text-lg">
                <span>★</span><span>★</span><span>★</span><span>★</span><span
                  >★</span
                >
              </div>
              <p class="text-xs text-gray-500 mt-1">Overall Rating</p>
            </div>
            <div class="space-y-1">
              <div class="flex items-center text-xs">
                <span class="w-12">5 ★</span>
                <div class="w-32 bg-gray-200 rounded-full h-2 mx-2">
                  <div
                    class="bg-yellow-400 h-2 rounded-full"
                    style="width: 80%"
                  ></div>
                </div>
                <span>80%</span>
              </div>
              <div class="flex items-center text-xs">
                <span class="w-12">4 ★</span>
                <div class="w-32 bg-gray-200 rounded-full h-2 mx-2">
                  <div
                    class="bg-yellow-400 h-2 rounded-full"
                    style="width: 15%"
                  ></div>
                </div>
                <span>15%</span>
              </div>
              <div class="flex items-center text-xs">
                <span class="w-12">3 ★</span>
                <div class="w-32 bg-gray-200 rounded-full h-2 mx-2">
                  <div
                    class="bg-yellow-400 h-2 rounded-full"
                    style="width: 5%"
                  ></div>
                </div>
                <span>5%</span>
              </div>
            </div>
          </div>

          <!-- Sample Feedbacks -->
          <div class="space-y-3 mt-4">
            <div class="p-3 bg-gray-50 rounded">
              <div class="flex items-center mb-1">
                <p class="text-sm font-medium flex-1">Sarah Johnson</p>
                <div class="text-yellow-400 text-xs">★★★★★</div>
              </div>
              <p class="text-xs text-gray-600">
                "Great company to work with! Clear requirements and timely
                feedback."
              </p>
              <p class="text-xs text-gray-400 mt-1">
                Project: Company Website Redesign
              </p>
            </div>
          </div>
        </div>

        <!-- Business Engagement Stats -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Engagement Statistics</h3>
          <div class="grid grid-cols-2 gap-3">
            <div class="text-center p-2 bg-gray-50 rounded">
              <p class="text-xl font-bold text-[#6b48ff]">15</p>
              <p class="text-xs text-gray-500">Projects Posted</p>
            </div>
            <div class="text-center p-2 bg-gray-50 rounded">
              <p class="text-xl font-bold text-[#6b48ff]">42</p>
              <p class="text-xs text-gray-500">Students Engaged</p>
            </div>
            <div class="text-center p-2 bg-gray-50 rounded">
              <p class="text-xl font-bold text-[#6b48ff]">12</p>
              <p class="text-xs text-gray-500">Completed Projects</p>
            </div>
            <div class="text-center p-2 bg-gray-50 rounded">
              <p class="text-xl font-bold text-[#6b48ff]">4.8</p>
              <p class="text-xs text-gray-500">Avg. Student Rating</p>
            </div>
          </div>
        </div>
      {/if}
    </div>
  </div>
</main>

<style>
  .modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 50;
  }
  .progress-bar {
    height: 8px;
    background-color: #e2e8f0;
    border-radius: 4px;
    overflow: hidden;
  }
  .progress-fill {
    height: 100%;
    background-color: #6b48ff;
  }
  .btn {
    background-color: #6b48ff;
    color: white;
    border-radius: 0.375rem;
    font-weight: 500;
    transition-property: background-color, border-color, color, fill, stroke;
    transition-duration: 150ms;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
  .btn:hover {
    background-color: #5a3cd9;
  }
  .btn:disabled {
    background-color: #a78eff;
    cursor: not-allowed;
  }
  .btn-secondary {
    background-color: white;
    color: #4b5563;
    border: 1px solid #d1d5db;
    border-radius: 0.375rem;
    font-weight: 500;
    transition-property: background-color, border-color, color, fill, stroke;
    transition-duration: 150ms;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
  .btn-secondary:hover {
    background-color: #f3f4f6;
  }
  .btn-secondary:disabled {
    background-color: #f9fafb;
    color: #9ca3af;
    cursor: not-allowed;
  }
  .card {
    background-color: white;
    border-radius: 0.5rem;
    box-shadow:
      0 1px 3px 0 rgba(0, 0, 0, 0.1),
      0 1px 2px 0 rgba(0, 0, 0, 0.06);
  }
</style>
