<script lang="ts">
  import type { PageData } from "../application/$types";

  let { data }: { data: PageData } = $props();

  const role = data.role;

  console.log(role);

  // Định nghĩa kiểu dữ liệu cho applications
  // Update Application type to include match_score
  type Application = {
    id: string;
    project_id: string;
    project_title: string;
    user_id: string;
    user_name: string;
    user_avatar: string;
    skills: string[];
    experience: string;
    motivation: string;
    status: "pending" | "approved" | "rejected";
    created_at: string;
    match_score: number; // AI matching score (0-100)
    proposal: string; // Detailed proposal from the applicant
  };

  // Mẫu dữ liệu - trong thực tế sẽ lấy từ API
  let applications: Application[] = $state([
    {
      id: "app1",
      project_id: "proj1",
      project_title: "E-commerce Platform",
      user_id: "user1",
      user_name: "Sarah Johnson",
      user_avatar: "https://avatars.githubusercontent.com/u/234567?v=4",
      skills: ["React", "Node.js", "MongoDB"],
      experience:
        "5 years as full-stack developer with focus on e-commerce solutions",
      motivation:
        "I'm interested in this project because I have extensive experience in building e-commerce platforms and I believe I can contribute significantly to the project's success.",
      status: "pending",
      created_at: "2025-03-15T09:30:00Z",
      match_score: 92,
      proposal:
        "## Project Approach\n\nI plan to implement a scalable e-commerce solution using React for the frontend and Node.js with Express for the backend. The architecture will follow microservices principles to ensure modularity.\n\n## Timeline\n\n- Week 1-2: Requirements analysis and system design\n- Week 3-4: Database schema design and API development\n- Week 5-8: Frontend development with component-based architecture\n- Week 9-10: Integration, testing and deployment\n\n## Technical Stack\n\n- Frontend: React, Redux, Tailwind CSS\n- Backend: Node.js, Express, MongoDB\n- DevOps: Docker, GitHub Actions, AWS\n\nI've previously built similar platforms for clients in retail and fashion industries with 50,000+ monthly active users.",
    },
    // Update remaining mock data similarly
    {
      id: "app2",
      project_id: "proj1",
      project_title: "E-commerce Platform",
      user_id: "user2",
      user_name: "Michael Chen",
      user_avatar: "https://avatars.githubusercontent.com/u/345678?v=4",
      skills: ["Python", "Django", "PostgreSQL"],
      experience: "3 years working with large-scale web applications",
      motivation:
        "The project aligns perfectly with my expertise in Python and database design. I'm looking to apply my skills in a meaningful way.",
      status: "approved",
      created_at: "2025-03-16T10:45:00Z",
      match_score: 85,
      proposal:
        "## Implementation Strategy\n\nI propose using Django with PostgreSQL for a robust, scalable e-commerce backend. Django's built-in admin panel and ORM will accelerate development time.\n\n## Features\n\n- Custom user authentication system with multi-factor authentication\n- Advanced product filtering and recommendation engine\n- Real-time inventory management with webhooks\n- Built-in SEO optimization tools\n\n## Technical Specifications\n\nI will implement a layered architecture with clearly separated concerns. Database optimization will be a primary focus to ensure fast query response times even with large product catalogs.\n\nI'm prepared to start immediately and can commit 30+ hours weekly to this project.",
    },
    {
      id: "app3",
      project_id: "proj2",
      project_title: "AI Learning Platform",
      user_id: "user3",
      user_name: "Alex Rivera",
      user_avatar: "https://avatars.githubusercontent.com/u/456789?v=4",
      skills: ["UI/UX", "Vue.js", "Figma"],
      experience:
        "4 years in frontend and UI/UX design for educational platforms",
      motivation:
        "I'm passionate about education and creating intuitive interfaces. This project is an excellent opportunity to combine these interests.",
      status: "pending",
      created_at: "2025-03-18T14:20:00Z",
      match_score: 78,
      proposal:
        "## Design Philosophy\n\nMy approach for this AI Learning Platform focuses on user-centered design with accessibility as a core principle. I believe educational interfaces should be intuitive, engaging, and reduce cognitive load.\n\n## Proposed Deliverables\n\n1. Complete UI kit with component library\n2. Interactive prototypes for key user flows\n3. Design system documentation\n4. Frontend implementation using Vue.js\n\n## User Experience Vision\n\nI envision an adaptive interface that personalizes the learning journey based on user progress and preferences. The UI will incorporate microinteractions that make the learning process more engaging while maintaining focus on educational content.\n\nMy previous work on EduLearn resulted in a 40% improvement in course completion rates.",
    },
    {
      id: "app4",
      project_id: "proj3",
      project_title: "Healthcare Management System",
      user_id: "user4",
      user_name: "Emily Zhang",
      user_avatar: "https://avatars.githubusercontent.com/u/567890?v=4",
      skills: ["Java", "Spring Boot", "MySQL"],
      experience:
        "7 years developing healthcare software and compliance systems",
      motivation:
        "I have extensive experience with healthcare systems and HIPAA compliance, which I believe will be valuable for this project.",
      status: "rejected",
      created_at: "2025-03-20T08:15:00Z",
      match_score: 95,
      proposal:
        "## Healthcare System Architecture\n\nI propose developing a HIPAA-compliant healthcare management system using Spring Boot and a secure microservices architecture. My experience with medical data systems informs my approach to security and compliance.\n\n## Key Components\n\n- Patient records management with detailed access controls\n- Appointment scheduling with intelligent conflict resolution\n- HL7 FHIR compliant API for interoperability\n- Comprehensive audit logging system\n\n## Security Measures\n\nThe system will implement end-to-end encryption, role-based access control, and detailed audit trails to ensure HIPAA compliance. I have implemented similar systems for three major hospitals in the region.\n\n## Timeline\n\nI estimate a 6-month development cycle with monthly deliverables and regular security reviews.",
    },
    {
      id: "app5",
      project_id: "proj2",
      project_title: "AI Learning Platform",
      user_id: "user5",
      user_name: "David Wilson",
      user_avatar: "https://avatars.githubusercontent.com/u/678901?v=4",
      skills: ["Machine Learning", "Python", "TensorFlow"],
      experience: "4 years in ML model development and data science",
      motivation:
        "I want to contribute my expertise in ML to create an adaptive learning platform that can personalize education.",
      status: "pending",
      created_at: "2025-03-21T13:10:00Z",
      match_score: 88,
      proposal:
        "## AI Implementation Strategy\n\nI propose developing an adaptive learning system using a combination of supervised and reinforcement learning techniques. The system will personalize content delivery based on user performance and learning patterns.\n\n## Technical Implementation\n\n- Content recommendation engine using collaborative filtering\n- Learning path optimization using knowledge graphs\n- Performance prediction models to identify potential learning gaps\n- Sentiment analysis for feedback evaluation\n\n## Data Management\n\nI'll implement a robust data pipeline that ensures user privacy while collecting sufficient information to train our models effectively. All algorithms will be designed to explain their decisions, avoiding black-box recommendations.\n\n## Evaluation Methodology\n\nI propose an A/B testing framework to continuously measure the effectiveness of our AI systems against traditional learning approaches.",
    },
  ]);

  // Filter applications
  let filterStatus: "all" | "pending" | "approved" | "rejected" = $state("all");
  let filterProject: string = $state("all");

  // Unique project list for filter dropdown
  let projects = $derived([
    ...new Set(applications.map((app) => app.project_title)),
  ]);

  // Filtered applications
  let filteredApplications = $derived(
    applications.filter((app) => {
      const matchStatus = filterStatus === "all" || app.status === filterStatus;
      const matchProject =
        filterProject === "all" || app.project_title === filterProject;
      return matchStatus && matchProject;
    })
  );

  // Modal state
  let selectedApplication: Application | null = $state(null);
  let showModal = $state(false);

  function openApplicationDetails(app: Application) {
    selectedApplication = app;
    showModal = true;
  }

  function closeModal() {
    showModal = false;
    selectedApplication = null;
  }

  // Application actions
  async function approveApplication(appId: string) {
    try {
      // In real implementation: call API to approve application
      // const response = await fetch(`/api/applications/${appId}/approve`, { method: 'PUT' });

      // For demo: update local state
      applications = applications.map((app) =>
        app.id === appId ? { ...app, status: "approved" } : app
      );

      // Close modal if open
      if (showModal && selectedApplication?.id === appId) {
        closeModal();
      }
    } catch (error) {
      console.error("Error approving application:", error);
    }
  }

  async function rejectApplication(appId: string) {
    try {
      // In real implementation: call API to reject application
      // const response = await fetch(`/api/applications/${appId}/reject`, { method: 'PUT' });

      // For demo: update local state
      applications = applications.map((app) =>
        app.id === appId ? { ...app, status: "rejected" } : app
      );

      // Close modal if open
      if (showModal && selectedApplication?.id === appId) {
        closeModal();
      }
    } catch (error) {
      console.error("Error rejecting application:", error);
    }
  }

  // Format date function
  function formatDate(dateString: string): string {
    if (!dateString) return "N/A";

    const date = new Date(dateString);

    // Check if date is valid
    if (isNaN(date.getTime())) return "Invalid date";

    // Format: "Mar 24, 2025"
    return date.toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  }
</script>

<header class="flex justify-between items-center mb-6 ml-64 pr-4 pl-4 pt-4">
  <div class="flex items-center">
    <a href="/project">
      <button class="text-gray-500 hover:text-gray-700 mr-3">
        <svg
          class="w-5 h-5"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M10 19l-7-7m0 0l7-7m-7 7h18"
          ></path>
        </svg>
      </button>
    </a>
    <div class="flex flex-col">
      <h2 class="text-xl font-semibold">Project Applications</h2>
      <p class="text-sm text-gray-600">
        {role === "business"
          ? "Review and manage applications for your projects"
          : "Review and manage your applications to projects"}
      </p>
    </div>
  </div>
  <div class="flex space-x-3">
    <div class="relative">
      {#if role === "student"}
        <select
          bind:value={filterProject}
          class="appearance-none bg-white border border-gray-200 rounded px-3 py-2 pr-8 text-sm focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
        >
          <option value="all">All Projects</option>
          {#each projects as project}
            <option value={project}>{project}</option>
          {/each}
        </select>
        <div
          class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-500"
        >
          <svg
            class="w-4 h-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M19 9l-7 7-7-7"
            ></path>
          </svg>
        </div>
      {/if}
    </div>
    <div class="flex border border-gray-200 rounded overflow-hidden">
      <button
        class={`px-3 py-2 text-sm ${filterStatus === "all" ? "bg-[#6b48ff] text-white" : "bg-white text-gray-700 hover:bg-gray-50"}`}
        onclick={() => (filterStatus = "all")}
      >
        All
      </button>
      <button
        class={`px-3 py-2 text-sm ${filterStatus === "pending" ? "bg-[#6b48ff] text-white" : "bg-white text-gray-700 hover:bg-gray-50"}`}
        onclick={() => (filterStatus = "pending")}
      >
        Pending
      </button>
      <button
        class={`px-3 py-2 text-sm ${filterStatus === "approved" ? "bg-[#6b48ff] text-white" : "bg-white text-gray-700 hover:bg-gray-50"}`}
        onclick={() => (filterStatus = "approved")}
      >
        Approved
      </button>
      <button
        class={`px-3 py-2 text-sm ${filterStatus === "rejected" ? "bg-[#6b48ff] text-white" : "bg-white text-gray-700 hover:bg-gray-50"}`}
        onclick={() => (filterStatus = "rejected")}
      >
        Rejected
      </button>
    </div>
  </div>
</header>

<main class="flex-1 pr-4 pl-4 ml-64">
  {#if filteredApplications.length === 0}
    <div class="card p-6 text-center">
      <svg
        class="w-16 h-16 text-gray-300 mx-auto mb-4"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
        ></path>
      </svg>
      <h3 class="text-lg font-medium text-gray-700 mb-2">
        No applications found
      </h3>
      <p class="text-gray-500 mb-4">
        There are no applications matching your current filters
      </p>
      <button
        class="btn-secondary"
        onclick={() => {
          filterStatus = "all";
          filterProject = "all";
        }}
      >
        Clear Filters
      </button>
    </div>
  {:else}
    <div class="space-y-4">
      {#each filteredApplications as app}
        <div class="card p-4 hover:shadow-md transition-shadow">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-4">
              <img
                src={app.user_avatar}
                alt={app.user_name}
                class="w-12 h-12 rounded-full object-cover border-2 border-gray-100"
              />

              <div>
                <h4 class="text-base font-medium">
                  <a
                    href={`/profile/${app.user_id}`}
                    class="hover:text-[#6b48ff] hover:underline"
                  >
                    {app.user_name}
                  </a>
                </h4>
                <div
                  class="flex items-center space-x-2 text-sm text-gray-500 mt-1"
                >
                  <span
                    >{role === "student"
                      ? "Reply to your application to"
                      : "Apply to"}
                    <span class="text-[#6b48ff]">{app.project_title}</span
                    ></span
                  >
                  <span>•</span>
                  <span>{formatDate(app.created_at)}</span>
                </div>
              </div>
            </div>

            <div class="flex items-center space-x-2">
              <!-- AI Match Score Badge -->
              <div
                class="px-2 py-1 bg-purple-100 text-purple-800 text-xs font-medium rounded-full flex items-center space-x-1"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-3 w-3"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z"
                    clip-rule="evenodd"
                  />
                </svg>
                <span>Match {app.match_score}%</span>
              </div>

              {#if app.status === "pending"}
                <span
                  class="px-3 py-1 bg-yellow-100 text-yellow-800 text-xs font-medium rounded-full"
                  >Pending</span
                >
                {#if role === "business"}
                  <div class="flex space-x-2">
                    <button
                      class="p-1 rounded-full bg-green-100 text-green-600 hover:bg-green-200"
                      title="Approve"
                      onclick={() => approveApplication(app.id)}
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="h-5 w-5"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                      >
                        <path
                          stroke-linecap="round"
                          stroke-linejoin="round"
                          stroke-width="2"
                          d="M5 13l4 4L19 7"
                        />
                      </svg>
                    </button>
                    <button
                      class="p-1 rounded-full bg-red-100 text-red-600 hover:bg-red-200"
                      title="Reject"
                      onclick={() => rejectApplication(app.id)}
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="h-5 w-5"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                      >
                        <path
                          stroke-linecap="round"
                          stroke-linejoin="round"
                          stroke-width="2"
                          d="M6 18L18 6M6 6l12 12"
                        />
                      </svg>
                    </button>
                  </div>
                {/if}
              {:else if app.status === "approved"}
                <span
                  class="px-3 py-1 bg-green-100 text-green-800 text-xs font-medium rounded-full"
                  >Approved</span
                >
              {:else}
                <span
                  class="px-3 py-1 bg-red-100 text-red-800 text-xs font-medium rounded-full"
                  >Rejected</span
                >
              {/if}

              <button
                class="ml-2 p-2 text-gray-400 hover:text-[#6b48ff] hover:bg-gray-100 rounded-full"
                onclick={() => openApplicationDetails(app)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                  />
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
                  />
                </svg>
              </button>
            </div>
          </div>

          <div class="mt-3">
            <div class="flex flex-wrap gap-1 mb-2">
              {#each app.skills as skill}
                <span
                  class="px-2 py-1 bg-gray-100 text-gray-700 text-xs rounded"
                  >{skill}</span
                >
              {/each}
            </div>

            <p class="text-sm text-gray-600 line-clamp-2">{app.motivation}</p>

            <button
              class="text-sm text-[#6b48ff] mt-2 hover:underline"
              onclick={() => openApplicationDetails(app)}
            >
              View Details
            </button>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</main>

<!-- Application Detail Modal -->
{#if showModal && selectedApplication}
  <div
    class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4 modal"
  >
    <div
      class="bg-white rounded-lg w-full max-w-3xl max-h-[90vh] flex flex-col shadow-xl"
    >
      <!-- Fixed header -->
      <div
        class="p-6 border-b border-gray-200 sticky top-0 bg-white z-10 rounded-t-lg"
      >
        <div class="flex justify-between items-center">
          <h3 class="text-lg font-semibold">Application Details</h3>
          <button
            class="text-gray-400 hover:text-gray-600"
            onclick={closeModal}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>
      </div>

      <!-- Scrollable content -->
      <div class="p-6 overflow-y-auto flex-grow">
        <div class="flex items-center space-x-4 mb-6">
          <img
            src={selectedApplication.user_avatar}
            alt={selectedApplication.user_name}
            class="w-16 h-16 rounded-full object-cover border-2 border-gray-200"
          />
          <div>
            <h4 class="text-xl font-semibold">
              <a
                href={`/profile/${selectedApplication.user_id}`}
                class="hover:text-[#6b48ff] hover:underline"
              >
                {selectedApplication.user_name}
              </a>
            </h4>
            <p class="text-gray-500">
              Applied on {formatDate(selectedApplication.created_at)}
            </p>
          </div>

          <!-- AI Match Score -->
          <div class="ml-auto flex items-center space-x-3">
            <div
              class="px-3 py-1.5 bg-purple-100 text-purple-800 rounded-lg flex items-center space-x-2"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z"
                  clip-rule="evenodd"
                />
              </svg>
              <div>
                <div class="text-xs font-medium">AI Match</div>
                <div class="text-base font-bold">
                  {selectedApplication.match_score}%
                </div>
              </div>
            </div>

            {#if selectedApplication.status === "pending"}
              <span
                class="px-3 py-1 bg-yellow-100 text-yellow-800 text-sm font-medium rounded-full"
                >Pending Review</span
              >
            {:else if selectedApplication.status === "approved"}
              <span
                class="px-3 py-1 bg-green-100 text-green-800 text-sm font-medium rounded-full"
                >Approved</span
              >
            {:else}
              <span
                class="px-3 py-1 bg-red-100 text-red-800 text-sm font-medium rounded-full"
                >Rejected</span
              >
            {/if}
          </div>
        </div>

        <div class="mb-6">
          <h5 class="text-base font-medium mb-2">Project</h5>
          <p class="text-gray-700">{selectedApplication.project_title}</p>
        </div>

        <div class="mb-6">
          <h5 class="text-base font-medium mb-2">Skills</h5>
          <div class="flex flex-wrap gap-2">
            {#each selectedApplication.skills as skill}
              <span class="px-3 py-1 bg-gray-100 text-gray-700 text-sm rounded"
                >{skill}</span
              >
            {/each}
          </div>
        </div>

        <div class="mb-6">
          <h5 class="text-base font-medium mb-2">Experience</h5>
          <p class="text-gray-700">{selectedApplication.experience}</p>
        </div>

        <div class="mb-6">
          <h5 class="text-base font-medium mb-2">Motivation</h5>
          <div class="bg-gray-50 p-4 rounded">
            <p class="text-gray-700">{selectedApplication.motivation}</p>
          </div>
        </div>

        <!-- Add the new detailed proposal section -->
        <div class="mb-10">
          <!-- Added extra bottom margin to ensure content doesn't get hidden behind the fixed footer -->
          <h5 class="text-base font-medium mb-2">Detailed Proposal</h5>
          <div class="bg-gray-50 p-4 rounded border border-gray-200">
            <div class="prose prose-sm max-w-none">
              <!-- Render the proposal with Markdown formatting -->
              <!-- In a real app, you would use a Markdown renderer component here -->
              <pre
                class="whitespace-pre-wrap text-gray-700">{selectedApplication.proposal}</pre>
            </div>
          </div>
        </div>

        <!-- Add extra space at the bottom to prevent content from being hidden behind fixed footer -->
        <div class="h-16"></div>
      </div>

      <!-- Fixed footer with action buttons -->
      {#if selectedApplication.status === "pending"}
        <div
          class="p-6 border-t border-gray-200 sticky bottom-0 bg-white z-10 rounded-b-lg"
        >
          <div class="flex justify-end space-x-3">
            <button
              class="btn-secondary"
              onclick={() =>
                selectedApplication &&
                rejectApplication(selectedApplication.id)}
            >
              Reject Application
            </button>
            <button
              class="btn"
              onclick={() =>
                selectedApplication &&
                approveApplication(selectedApplication.id)}
            >
              Approve Application
            </button>
          </div>
        </div>
      {/if}
    </div>
  </div>
{/if}

<style>
  .line-clamp-2 {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  .modal {
    background-color: rgba(0, 0, 0, 0.5);
  }
</style>
