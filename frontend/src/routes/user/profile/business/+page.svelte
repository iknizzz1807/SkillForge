<script lang="ts">
  import { onMount } from "svelte";
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  // Người xem hiện tại (viewer) - có thể lấy từ store hoặc params
  // Đây là mock data, thực tế nên lấy từ API
  let currentUser = {
    id: "current-user-123",
    role: "student", // hoặc "business"
    name: "Your Name",
  };

  // Mock data cho business profile
  let businessData = {
    id: "business-456",
    name: "Tech Solutions Inc.",
    email: "contact@techsolutions.com",
    title: "Technology Solutions Provider",
    logo: "https://randomuser.me/api/portraits/lego/1.jpg",
    businessInfo: {
      companyType: "Technology Consulting",
      founded: "2018",
      companySize: "15-50 employees",
      website: "www.techsolutions.com",
      location: "San Francisco, CA",
      aboutUs:
        "Tech Solutions Inc. is a forward-thinking technology company specializing in developing innovative web and mobile applications for businesses of all sizes. We focus on creating modern, scalable solutions that help our clients achieve their digital transformation goals.",
    },
    stats: {
      projectsPosted: 38,
      studentsEngaged: 120,
      completedProjects: 25,
      avgRating: 4.8,
    },
    activeProjects: [
      {
        id: "proj1",
        name: "Mobile Banking App",
        skills: "React Native, Node.js",
        dueDate: "25/05/2025",
        progress: 65,
        status: "In Progress",
        openPositions: 1,
      },
      {
        id: "proj2",
        name: "E-commerce Dashboard",
        skills: "React, MongoDB, Express",
        dueDate: "10/06/2025",
        progress: 15,
        status: "Just Started",
        openPositions: 2,
      },
      {
        id: "proj3",
        name: "AI Content Generator",
        skills: "Python, ML/AI, JavaScript",
        dueDate: "30/07/2025",
        progress: 0,
        status: "Recruiting",
        openPositions: 4,
      },
    ],
    completedProjects: [
      {
        id: "comp1",
        name: "Company Website Redesign",
        skills: "HTML, CSS, React",
        completedDate: "10/03/2025",
        students: 5,
        rating: 4.8,
      },
      {
        id: "comp2",
        name: "CRM Integration",
        skills: "Salesforce, JavaScript, API",
        completedDate: "25/02/2025",
        students: 3,
        rating: 4.9,
      },
      {
        id: "comp3",
        name: "Payment Gateway API",
        skills: "Node.js, Express, Stripe",
        completedDate: "15/01/2025",
        students: 4,
        rating: 4.7,
      },
    ],
    testimonials: [
      {
        studentName: "Sarah Johnson",
        studentAvatar: "https://randomuser.me/api/portraits/women/44.jpg",
        project: "Company Website Redesign",
        rating: 5,
        text: "Great company to work with! The team provided clear requirements and timely feedback throughout the project.",
      },
      {
        studentName: "David Chen",
        studentAvatar: "https://randomuser.me/api/portraits/men/32.jpg",
        project: "CRM Integration",
        rating: 4,
        text: "Excellent learning experience. The mentors were supportive and helped me expand my Salesforce knowledge.",
      },
    ],
  };

  // State vars
  let showContactInfoModal = $state(false);
  let applyingToProject = $state(false);
  let selectedProject = $state("");
  let applicationSuccess = $state(false);
  let showProjectDetailsModal = $state(false);
  let currentProjectDetails = $state(null);

  const toggleContactInfoModal = () => {
    showContactInfoModal = !showContactInfoModal;
    if (!showContactInfoModal) {
      selectedProject = "";
      applicationSuccess = false;
    }
  };

  const applyToProject = async () => {
    if (!selectedProject) return;
    applyingToProject = true;

    // API Call: POST /api/projects/apply
    // Body: { projectId: selectedProject, studentId: currentUser.id }
    console.log("Applying to project:", selectedProject);

    // Mock successful API response
    setTimeout(() => {
      applyingToProject = false;
      applicationSuccess = true;

      // Auto-close modal after success message
      setTimeout(() => {
        toggleContactInfoModal();
      }, 2000);
    }, 1200);
  };

  const viewProjectDetails = (project) => {
    currentProjectDetails = project;
    showProjectDetailsModal = true;
  };

  const closeProjectDetailsModal = () => {
    showProjectDetailsModal = false;
    currentProjectDetails = null;
  };

  onMount(() => {
    // API call to fetch business profile data
    // GET /api/business/{id}
    console.log("Fetching business profile data for:", data.id);

    // Additional API calls as needed
  });
</script>

<!-- Contact Info Modal -->
{#if showContactInfoModal}
  <div class="modal">
    <div class="bg-white rounded-lg w-full max-w-md shadow-xl mx-4 p-5">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-semibold">Contact {businessData.name}</h3>
        <button
          on:click={toggleContactInfoModal}
          class="text-gray-500 hover:text-gray-700 transition-colors"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
              clip-rule="evenodd"
            />
          </svg>
        </button>
      </div>

      {#if applicationSuccess}
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
              >Application submitted successfully!</span
            >
          </div>
        </div>
      {:else}
        <div class="mb-4">
          <p class="text-gray-600 mb-2">Business contact information:</p>
          <div class="bg-gray-50 p-3 rounded border border-gray-200">
            <p class="text-sm">
              <span class="font-medium">Email:</span>
              {businessData.email}
            </p>
            {#if businessData.businessInfo.website}
              <p class="text-sm mt-1">
                <span class="font-medium">Website:</span>
                <a
                  href={`https://${businessData.businessInfo.website}`}
                  target="_blank"
                  class="text-[#6b48ff] hover:underline"
                >
                  {businessData.businessInfo.website}
                </a>
              </p>
            {/if}
            {#if businessData.businessInfo.location}
              <p class="text-sm mt-1">
                <span class="font-medium">Location:</span>
                {businessData.businessInfo.location}
              </p>
            {/if}
          </div>
        </div>

        {#if currentUser.role === "student"}
          <div class="mb-4">
            <label class="block text-sm font-medium mb-1" for="project">
              Apply to open project
            </label>
            <select
              id="project"
              bind:value={selectedProject}
              class="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none"
            >
              <option value="">Select a project...</option>
              {#each businessData.activeProjects as project}
                {#if project.openPositions > 0}
                  <option value={project.id}
                    >{project.name} ({project.openPositions} positions)</option
                  >
                {/if}
              {/each}
            </select>
          </div>
        {/if}

        <div class="flex justify-end space-x-3">
          <button
            class="btn-secondary px-4 py-2"
            on:click={toggleContactInfoModal}
          >
            Close
          </button>

          {#if currentUser.role === "student"}
            <button
              class="btn px-4 py-2"
              on:click={applyToProject}
              disabled={!selectedProject || applyingToProject}
            >
              {#if applyingToProject}
                <div class="flex items-center">
                  <svg
                    class="animate-spin -ml-1 mr-2 h-4 w-4 text-white"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                  >
                    <circle
                      class="opacity-25"
                      cx="12"
                      cy="12"
                      r="10"
                      stroke="currentColor"
                      stroke-width="4"
                    ></circle>
                    <path
                      class="opacity-75"
                      fill="currentColor"
                      d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                    ></path>
                  </svg>
                  Submitting...
                </div>
              {:else}
                Submit Application
              {/if}
            </button>
          {/if}
        </div>
      {/if}
    </div>
  </div>
{/if}

<!-- Project Details Modal -->
{#if showProjectDetailsModal && currentProjectDetails}
  <div class="modal">
    <div class="bg-white rounded-lg w-full max-w-2xl shadow-xl mx-4 p-5">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-semibold">{currentProjectDetails.name}</h3>
        <button
          on:click={closeProjectDetailsModal}
          class="text-gray-500 hover:text-gray-700 transition-colors"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
              clip-rule="evenodd"
            />
          </svg>
        </button>
      </div>

      <div class="mb-4">
        <div class="flex justify-between mb-2">
          <p class="text-sm font-medium">Status</p>
          <span class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded">
            {currentProjectDetails.status}
          </span>
        </div>
        <div class="mb-2">
          <p class="text-sm font-medium">Required Skills</p>
          <p class="text-sm text-gray-600">{currentProjectDetails.skills}</p>
        </div>
        <div class="mb-2">
          <p class="text-sm font-medium">Due Date</p>
          <p class="text-sm text-gray-600">{currentProjectDetails.dueDate}</p>
        </div>
        {#if currentProjectDetails.openPositions}
          <div class="mb-2">
            <p class="text-sm font-medium">Open Positions</p>
            <p class="text-sm text-gray-600">
              {currentProjectDetails.openPositions}
            </p>
          </div>
        {/if}
        {#if currentProjectDetails.progress !== undefined}
          <div class="mb-2">
            <p class="text-sm font-medium">Progress</p>
            <div class="w-full bg-gray-200 rounded-full h-2 mt-1">
              <div
                class="bg-blue-500 h-2 rounded-full"
                style="width: {currentProjectDetails.progress}%"
              ></div>
            </div>
            <p class="text-xs text-gray-500 text-right mt-1">
              {currentProjectDetails.progress}% complete
            </p>
          </div>
        {/if}

        {#if currentProjectDetails.rating}
          <div class="mt-4">
            <p class="text-sm font-medium">Project Rating</p>
            <div class="flex items-center mt-1">
              <span class="text-lg font-bold text-[#6b48ff] mr-2"
                >{currentProjectDetails.rating}</span
              >
              <div class="flex text-yellow-400">
                {#each Array(5) as _, i}
                  {#if i < Math.floor(currentProjectDetails.rating)}
                    <span>★</span>
                  {:else if i < Math.ceil(currentProjectDetails.rating) && currentProjectDetails.rating % 1 !== 0}
                    <span class="text-yellow-300">★</span>
                  {:else}
                    <span class="text-gray-300">★</span>
                  {/if}
                {/each}
              </div>
            </div>
          </div>
        {/if}
      </div>

      <div class="flex justify-end space-x-3 mt-6">
        <button
          class="btn-secondary px-4 py-2"
          on:click={closeProjectDetailsModal}
        >
          Close
        </button>

        {#if currentUser.role === "student" && currentProjectDetails.openPositions > 0}
          <button
            class="btn px-4 py-2"
            on:click={() => {
              selectedProject = currentProjectDetails.id;
              closeProjectDetailsModal();
              toggleContactInfoModal();
            }}
          >
            Apply for Project
          </button>
        {/if}
      </div>
    </div>
  </div>
{/if}

<main class="flex-1 pr-4 pl-4 ml-64 pt-4">
  <!-- Header section -->
  <div class="flex justify-between items-center mb-6">
    <div>
      <h1 class="text-2xl font-bold flex items-center">
        {businessData.name}
        <span class="ml-2 text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded"
          >Business</span
        >
      </h1>
      <p class="text-gray-500">{businessData.title}</p>
    </div>

    <!-- Action buttons -->
    <div class="flex space-x-2">
      <button class="btn-secondary px-4 py-2" on:click={toggleContactInfoModal}>
        Contact Business
      </button>
    </div>
  </div>

  <!-- Profile Content -->
  <div class="grid grid-cols-3 gap-4">
    <!-- Left Column - Basic Info -->
    <div class="space-y-4">
      <!-- Business Card -->
      <div class="card p-4">
        <div class="flex flex-col items-center mb-4">
          <div
            class="w-24 h-24 rounded-full bg-gray-200 flex items-center justify-center mb-3 overflow-hidden"
          >
            <img
              class="w-full h-full object-cover"
              src={businessData.logo}
              alt="{businessData.name} Logo"
            />
          </div>
          <h3 class="text-lg font-semibold text-center">{businessData.name}</h3>
          <p class="text-sm text-gray-500 text-center">{businessData.title}</p>
        </div>

        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium">Email:</span>
            <span class="text-sm text-gray-600">{businessData.email}</span>
          </div>
          {#if businessData.businessInfo.website}
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium">Website:</span>
              <a
                href={`https://${businessData.businessInfo.website}`}
                target="_blank"
                class="text-sm text-[#6b48ff] hover:underline"
              >
                {businessData.businessInfo.website}
              </a>
            </div>
          {/if}
          {#if businessData.businessInfo.location}
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium">Location:</span>
              <span class="text-sm text-gray-600"
                >{businessData.businessInfo.location}</span
              >
            </div>
          {/if}
        </div>

        <!-- Stats summary -->
        <div class="grid grid-cols-2 gap-2 mt-4">
          <div class="bg-gray-50 rounded p-2 text-center">
            <p class="text-xl font-bold text-[#6b48ff]">
              {businessData.stats.completedProjects}
            </p>
            <p class="text-xs text-gray-500">Projects Completed</p>
          </div>
          <div class="bg-gray-50 rounded p-2 text-center">
            <p class="text-xl font-bold text-[#6b48ff]">
              {businessData.stats.avgRating}
            </p>
            <p class="text-xs text-gray-500">Avg. Rating</p>
          </div>
        </div>
      </div>

      <!-- Business Info -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">Company Information</h3>
        <div class="space-y-3">
          <div>
            <p class="text-sm font-medium">Company Type</p>
            <p class="text-sm text-gray-600">
              {businessData.businessInfo.companyType || "Not provided"}
            </p>
          </div>
          <div>
            <p class="text-sm font-medium">Founded</p>
            <p class="text-sm text-gray-600">
              {businessData.businessInfo.founded || "Not provided"}
            </p>
          </div>
          <div>
            <p class="text-sm font-medium">Size</p>
            <p class="text-sm text-gray-600">
              {businessData.businessInfo.companySize || "Not provided"}
            </p>
          </div>
          <div>
            <p class="text-sm font-medium">About Us</p>
            <p class="text-sm text-gray-600">
              {businessData.businessInfo.aboutUs || "Not provided"}
            </p>
          </div>
        </div>
      </div>

      <!-- Business Stats -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">Engagement Statistics</h3>
        <div class="grid grid-cols-2 gap-3">
          <div class="text-center p-2 bg-gray-50 rounded">
            <p class="text-xl font-bold text-[#6b48ff]">
              {businessData.stats.projectsPosted}
            </p>
            <p class="text-xs text-gray-500">Projects Posted</p>
          </div>
          <div class="text-center p-2 bg-gray-50 rounded">
            <p class="text-xl font-bold text-[#6b48ff]">
              {businessData.stats.studentsEngaged}
            </p>
            <p class="text-xs text-gray-500">Students Engaged</p>
          </div>
          <div class="text-center p-2 bg-gray-50 rounded">
            <p class="text-xl font-bold text-[#6b48ff]">
              {businessData.stats.completedProjects}
            </p>
            <p class="text-xs text-gray-500">Completed Projects</p>
          </div>
          <div class="text-center p-2 bg-gray-50 rounded">
            <p class="text-xl font-bold text-[#6b48ff]">
              {businessData.stats.avgRating}
            </p>
            <p class="text-xs text-gray-500">Avg. Student Rating</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Middle Column - Active Projects -->
    <div class="space-y-4">
      <!-- Open Projects -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">Open Projects</h3>
        <div class="space-y-3">
          {#each businessData.activeProjects.filter((p) => p.openPositions > 0) as project}
            <div class="p-3 bg-gray-50 rounded border border-gray-200">
              <div class="flex justify-between mb-1">
                <p class="text-sm font-medium">{project.name}</p>
                <span
                  class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded"
                >
                  {project.openPositions} position{project.openPositions > 1
                    ? "s"
                    : ""}
                </span>
              </div>
              <p class="text-xs text-gray-500 mb-2">
                Skills: {project.skills} | Due: {project.dueDate}
              </p>

              <div class="flex justify-between mt-2">
                <button
                  class="text-xs text-[#6b48ff] hover:underline"
                  on:click={() => viewProjectDetails(project)}
                >
                  View Details
                </button>

                {#if currentUser.role === "student"}
                  <button
                    class="text-xs bg-[#6b48ff] text-white px-2 py-1 rounded"
                    on:click={() => {
                      selectedProject = project.id;
                      toggleContactInfoModal();
                    }}
                  >
                    Apply
                  </button>
                {/if}
              </div>
            </div>
          {/each}

          {#if businessData.activeProjects.filter((p) => p.openPositions > 0).length === 0}
            <p class="text-sm text-gray-500 text-center py-2">
              No open positions available at the moment.
            </p>
          {/if}
        </div>
      </div>

      <!-- In Progress Projects -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">In Progress Projects</h3>
        <div class="space-y-3">
          {#each businessData.activeProjects.filter((p) => p.progress > 0) as project}
            <div class="p-3 bg-gray-50 rounded border border-gray-200">
              <div class="flex justify-between mb-1">
                <p class="text-sm font-medium">{project.name}</p>
                <span
                  class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded"
                >
                  {project.status}
                </span>
              </div>
              <p class="text-xs text-gray-500 mb-2">
                Skills: {project.skills} | Due: {project.dueDate}
              </p>

              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-2">
                <div
                  class="bg-blue-500 h-1.5 rounded-full"
                  style="width: {project.progress}%"
                ></div>
              </div>
              <div class="flex justify-between mt-2">
                <p class="text-xs text-gray-500">
                  {project.progress}% complete
                </p>
                <button
                  class="text-xs text-[#6b48ff] hover:underline"
                  on:click={() => viewProjectDetails(project)}
                >
                  View Details
                </button>
              </div>
            </div>
          {/each}

          {#if businessData.activeProjects.filter((p) => p.progress > 0).length === 0}
            <p class="text-sm text-gray-500 text-center py-2">
              No projects currently in progress.
            </p>
          {/if}
        </div>
      </div>

      <!-- Completed Projects -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">Completed Projects</h3>
        <div class="space-y-3">
          {#each businessData.completedProjects as project}
            <div
              class="flex justify-between items-center p-2 bg-gray-50 rounded border border-gray-200"
            >
              <div>
                <p class="text-sm font-medium">{project.name}</p>
                <p class="text-xs text-gray-500">
                  Completed: {project.completedDate} • {project.students} students
                </p>
              </div>
              <div class="flex items-center">
                <span
                  class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded flex items-center"
                >
                  <span class="mr-1">{project.rating}</span>
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-3 w-3"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
                    />
                  </svg>
                </span>
                <button
                  class="text-xs text-[#6b48ff] hover:underline ml-2"
                  on:click={() => viewProjectDetails(project)}
                >
                  Details
                </button>
              </div>
            </div>
          {/each}

          {#if businessData.completedProjects.length > 5}
            <div class="text-center mt-3">
              <a
                href="/projects?business={businessData.id}"
                class="text-sm text-[#6b48ff] hover:underline"
              >
                View All Completed Projects
              </a>
            </div>
          {/if}
        </div>
      </div>
    </div>

    <!-- Right Column - Reviews & Ratings -->
    <div class="space-y-4">
      <!-- Student Testimonials & Reviews -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">Student Feedback</h3>
        <div class="flex justify-center items-center py-3">
          <div class="text-center mr-6">
            <p class="text-3xl font-bold text-[#6b48ff]">
              {businessData.stats.avgRating}
            </p>
            <div class="flex text-yellow-400 text-lg">
              {#each Array(5) as _, i}
                {#if i < Math.floor(businessData.stats.avgRating)}
                  <span>★</span>
                {:else if i < Math.ceil(businessData.stats.avgRating) && businessData.stats.avgRating % 1 !== 0}
                  <span class="text-yellow-300">★</span>
                {:else}
                  <span class="text-gray-300">★</span>
                {/if}
              {/each}
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

        <!-- Testimonials -->
        <div class="space-y-4 mt-4">
          {#each businessData.testimonials as testimonial}
            <div class="p-4 bg-gray-50 rounded border border-gray-100">
              <div class="flex items-start mb-2">
                <img
                  src={testimonial.studentAvatar}
                  alt={testimonial.studentName}
                  class="w-8 h-8 rounded-full mr-2"
                />
                <div>
                  <p class="text-sm font-medium">{testimonial.studentName}</p>
                  <div class="flex text-yellow-400 text-xs">
                    {#each Array(5) as _, i}
                      {#if i < testimonial.rating}
                        <span>★</span>
                      {:else}
                        <span class="text-gray-300">★</span>
                      {/if}
                    {/each}
                  </div>
                </div>
              </div>
              <p class="text-sm text-gray-600">"{testimonial.text}"</p>
              <p class="text-xs text-gray-500 mt-2">
                Project: {testimonial.project}
              </p>
            </div>
          {/each}

          {#if businessData.testimonials.length > 2}
            <div class="text-center mt-2">
              <a
                href="/reviews?business={businessData.id}"
                class="text-sm text-[#6b48ff] hover:underline"
              >
                See All Reviews
              </a>
            </div>
          {/if}
        </div>
      </div>

      <!-- Industry Focus -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">Industry Focus</h3>
        <div class="flex flex-wrap gap-2">
          <span class="bg-blue-50 text-blue-700 text-xs px-2 py-1 rounded-full"
            >Fintech</span
          >
          <span class="bg-blue-50 text-blue-700 text-xs px-2 py-1 rounded-full"
            >E-commerce</span
          >
          <span class="bg-blue-50 text-blue-700 text-xs px-2 py-1 rounded-full"
            >Healthcare</span
          >
          <span class="bg-blue-50 text-blue-700 text-xs px-2 py-1 rounded-full"
            >SaaS</span
          >
          <span class="bg-blue-50 text-blue-700 text-xs px-2 py-1 rounded-full"
            >Mobile Apps</span
          >
        </div>

        <h4 class="text-sm font-medium mt-4 mb-2">Technology Stack</h4>
        <div class="flex flex-wrap gap-2">
          <span
            class="bg-purple-50 text-purple-700 text-xs px-2 py-1 rounded-full"
            >React</span
          >
          <span
            class="bg-purple-50 text-purple-700 text-xs px-2 py-1 rounded-full"
            >Node.js</span
          >
          <span
            class="bg-purple-50 text-purple-700 text-xs px-2 py-1 rounded-full"
            >MongoDB</span
          >
          <span
            class="bg-purple-50 text-purple-700 text-xs px-2 py-1 rounded-full"
            >AWS</span
          >
          <span
            class="bg-purple-50 text-purple-700 text-xs px-2 py-1 rounded-full"
            >Python</span
          >
          <span
            class="bg-purple-50 text-purple-700 text-xs px-2 py-1 rounded-full"
            >Flutter</span
          >
        </div>
      </div>

      <!-- Success Metrics -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">Success Metrics</h3>
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <p class="text-sm font-medium">Project Completion Rate</p>
            <p class="text-sm text-[#6b48ff] font-semibold">96%</p>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-[#6b48ff] h-2 rounded-full" style="width: 96%"></div>
          </div>

          <div class="flex justify-between items-center mt-3">
            <p class="text-sm font-medium">On-time Delivery</p>
            <p class="text-sm text-[#6b48ff] font-semibold">92%</p>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-[#6b48ff] h-2 rounded-full" style="width: 92%"></div>
          </div>

          <div class="flex justify-between items-center mt-3">
            <p class="text-sm font-medium">Student Satisfaction</p>
            <p class="text-sm text-[#6b48ff] font-semibold">94%</p>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-[#6b48ff] h-2 rounded-full" style="width: 94%"></div>
          </div>
        </div>

        <div class="mt-4 pt-3 border-t border-gray-100">
          <p class="text-xs text-gray-500">
            Data based on {businessData.stats.completedProjects} completed projects
            with {businessData.stats.studentsEngaged} students
          </p>
        </div>
      </div>
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
