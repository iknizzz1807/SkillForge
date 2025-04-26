<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  const role = data.role;

  // Update Application type to match backend model
  type Application = {
    id: string;
    project_id: string;
    project_name?: string;
    user_id: string;
    user_name?: string;
    user_avatar?: string;
    skills?: string[];
    experience?: string;
    motivation: string;
    detailed_proposal: string;
    status: "pending" | "approved" | "rejected";
    created_at: string;
    updated_at?: string;
    match_score?: number; // AI matching score (0-100)
  };

  // Initialize applications with data from server
  let applications: Application[] = $state(data.applications || []);

  // Add any error message from the server
  let errorMessage = $state(data.error || null);

  // Filter applications
  let filterStatus: "all" | "pending" | "approved" | "rejected" = $state("all");
  let filterProject: string = $state("all");

  // Unique project list for filter dropdown
  let projects = $derived(
    [...new Set(applications.map((app) => app.project_name || ""))].filter(
      (name) => name !== ""
    )
  );

  // Filtered applications
  let filteredApplications = $derived(
    applications.filter((app) => {
      const matchStatus = filterStatus === "all" || app.status === filterStatus;
      const matchProject =
        filterProject === "all" || app.project_name === filterProject;
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
      // Call API to approve application
      const response = await fetch(`/api/applications/${appId}/status`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ status: "approved" }),
      });

      if (!response.ok) {
        throw new Error(
          `Failed to approve application: ${response.statusText}`
        );
      }

      // Update local state
      applications = applications.map((app) =>
        app.id === appId ? { ...app, status: "approved" } : app
      );

      // Close modal if open
      if (showModal && selectedApplication?.id === appId) {
        closeModal();
      }
    } catch (error) {
      console.error("Error approving application:", error);
      alert("Failed to approve application. Please try again.");
    }
  }

  async function rejectApplication(appId: string) {
    try {
      // Call API to reject application
      const response = await fetch(`/api/applications/${appId}/status`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ status: "rejected" }),
      });

      if (!response.ok) {
        throw new Error(`Failed to reject application: ${response.statusText}`);
      }

      // Update local state
      applications = applications.map((app) =>
        app.id === appId ? { ...app, status: "rejected" } : app
      );

      // Close modal if open
      if (showModal && selectedApplication?.id === appId) {
        closeModal();
      }
    } catch (error) {
      console.error("Error rejecting application:", error);
      alert("Failed to reject application. Please try again.");
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

  // Thêm vào phần khai báo state
  let showDeleteConfirmModal = $state(false);
  let applicationToDelete: string | null = $state(null);

  // Thêm hàm để hiển thị modal xác nhận xóa
  function confirmDeleteApplication(appId: string) {
    applicationToDelete = appId;
    showDeleteConfirmModal = true;
  }

  // Thêm hàm để xóa application
  async function deleteApplication() {
    if (!applicationToDelete) return;

    try {
      // Gọi API để xóa application
      const response = await fetch(`/api/applications/${applicationToDelete}`, {
        method: "DELETE",
      });

      if (!response.ok) {
        throw new Error(`Failed to delete application: ${response.statusText}`);
      }

      // Cập nhật state local sau khi xóa thành công
      applications = applications.filter(
        (app) => app.id !== applicationToDelete
      );

      // Đóng modal confirm
      showDeleteConfirmModal = false;
      applicationToDelete = null;

      // Đóng modal chi tiết nếu đang mở và đúng application bị xóa
      if (showModal && selectedApplication?.id === applicationToDelete) {
        closeModal();
      }
    } catch (error) {
      console.error("Error deleting application:", error);
      alert("Failed to delete application. Please try again.");
    }
  }
</script>

<svelte:head>
  <title>Project Applications</title>
</svelte:head>

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
      {#if role === "business" && projects.length > 0}
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
  {#if errorMessage}
    <div
      class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4"
      role="alert"
    >
      <strong class="font-bold">Error:</strong>
      <span class="block sm:inline"> {errorMessage}</span>
    </div>
  {/if}

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
                src={app.user_avatar || "/default-avatar.png"}
                alt={app.user_name || "User"}
                class="w-12 h-12 rounded-full object-cover border-2 border-gray-100"
              />

              <div>
                <h4 class="text-base font-medium">
                  <a
                    href={`/profile/${app.user_id}`}
                    class="hover:text-[#6b48ff] hover:underline"
                  >
                    {app.user_name || "User"}
                  </a>
                </h4>
                <div
                  class="flex items-center space-x-2 text-sm text-gray-500 mt-1"
                >
                  <span
                    >{role === "student" ? "Your application to" : "Applied to"}
                    <span class="text-[#6b48ff]">{app.project_name}</span></span
                  >
                  <span>•</span>
                  <span>{formatDate(app.created_at)}</span>
                </div>
              </div>
            </div>

            <div class="flex items-center space-x-2">
              <!-- AI Match Score Badge (if available) -->
              {#if app.match_score !== undefined}
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
              {:else}
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
                  <span>Match 76%</span>
                </div>
              {/if}

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

              <!-- Thêm nút Delete -->
              <button
                class="ml-2 p-2 text-gray-400 hover:text-red-600 hover:bg-gray-100 rounded-full"
                onclick={() => confirmDeleteApplication(app.id)}
                title="Delete application"
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
                    d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                  />
                </svg>
              </button>
            </div>
          </div>

          <div class="mt-3">
            {#if app.skills && app.skills.length > 0}
              <div class="flex flex-wrap gap-1 mb-2">
                {#each app.skills as skill}
                  <span
                    class="px-2 py-1 bg-gray-100 text-gray-700 text-xs rounded"
                    >{skill}</span
                  >
                {/each}
              </div>
            {/if}

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
            src={selectedApplication.user_avatar || "/default-avatar.png"}
            alt={selectedApplication.user_name || "User"}
            class="w-16 h-16 rounded-full object-cover border-2 border-gray-200"
          />
          <div>
            <h4 class="text-xl font-semibold">
              <a
                href={`/profile/${selectedApplication.user_id}`}
                class="hover:text-[#6b48ff] hover:underline"
              >
                {selectedApplication.user_name || "User"}
              </a>
            </h4>
            <p class="text-gray-500">
              Applied on {formatDate(selectedApplication.created_at)}
            </p>
          </div>

          <!-- AI Match Score (if available) -->
          <div class="ml-auto flex items-center space-x-3">
            {#if selectedApplication.match_score !== undefined}
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
            {/if}

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
          <p class="text-gray-700">{selectedApplication.project_name}</p>
        </div>

        {#if selectedApplication.skills && selectedApplication.skills.length > 0}
          <div class="mb-6">
            <h5 class="text-base font-medium mb-2">Skills</h5>
            <div class="flex flex-wrap gap-2">
              {#each selectedApplication.skills as skill}
                <span
                  class="px-3 py-1 bg-gray-100 text-gray-700 text-sm rounded"
                  >{skill}</span
                >
              {/each}
            </div>
          </div>
        {/if}

        {#if selectedApplication.experience}
          <div class="mb-6">
            <h5 class="text-base font-medium mb-2">Experience</h5>
            <p class="text-gray-700">{selectedApplication.experience}</p>
          </div>
        {/if}

        <div class="mb-6">
          <h5 class="text-base font-medium mb-2">Motivation</h5>
          <div class="bg-gray-50 p-4 rounded">
            <p class="text-gray-700">{selectedApplication.motivation}</p>
          </div>
        </div>

        <!-- Detailed proposal section -->
        <div class="mb-10">
          <h5 class="text-base font-medium mb-2">Detailed Proposal</h5>
          <div class="bg-gray-50 p-4 rounded border border-gray-200">
            <div class="prose prose-sm max-w-none">
              <pre
                class="whitespace-pre-wrap text-gray-700">{selectedApplication.detailed_proposal}</pre>
            </div>
          </div>
        </div>

        <!-- Add extra space at the bottom -->
        <div class="h-16"></div>
      </div>

      <!-- Fixed footer with action buttons (only for business with pending applications) -->
      {#if selectedApplication.status === "pending" && role === "business"}
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

{#if showDeleteConfirmModal}
  <div
    class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4 modal"
  >
    <div class="bg-white rounded-lg w-full max-w-md shadow-xl p-6">
      <div class="text-center">
        <div
          class="mx-auto flex items-center justify-center h-16 w-16 rounded-full bg-red-100 mb-6"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-8 w-8 text-red-600"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
            />
          </svg>
        </div>
        <h3 class="text-lg font-medium text-gray-900 mb-2">
          Delete Application
        </h3>
        <p class="text-sm text-gray-500 mb-6">
          Are you sure you want to delete this application? This action cannot
          be undone.
        </p>
        <div class="flex justify-center space-x-3">
          <button
            class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300 transition-colors"
            onclick={() => {
              showDeleteConfirmModal = false;
              applicationToDelete = null;
            }}
          >
            Cancel
          </button>
          <button
            class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors"
            onclick={deleteApplication}
          >
            Delete
          </button>
        </div>
      </div>
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
