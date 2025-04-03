<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();

  type ProjectDisplay = {
    id: string;
    title: string;
    description: string;
    skills: string[];
    start_time: Date;
    end_time: Date;
    created_by_id: string;
    created_by_name: string;
    status: string;
    max_member: number;
    current_member: number;
    created_at: string;
  };

  const token = data.token;

  let projectsDisplay: ProjectDisplay[] = $state(data.projects);
  let errorLoadingProjects: string | null = $state(data.error);

  let filterState: string = $state("all");

  // Modal states
  let showEditModal: boolean = $state(false);
  let showDeleteModal: boolean = $state(false);
  let currentProject: ProjectDisplay | null = $state(null);

  // Modal functions
  function openEditModal(project: ProjectDisplay) {
    currentProject = { ...project };
    showEditModal = true;
  }

  function openDeleteModal(project: ProjectDisplay) {
    currentProject = project;
    showDeleteModal = true;
  }

  function closeEditModal() {
    showEditModal = false;
    currentProject = null;
  }

  function closeDeleteModal() {
    showDeleteModal = false;
    currentProject = null;
  }

  // Edit form state
  let editTitle: string = $state("");
  let editDescription: string = $state("");
  let editSkills: string[] = $state([]);
  let editStartDate: string = $state("");
  let editEndDate: string = $state("");
  let editMaxMember: number = $state(0);
  let editStatus: string = $state("");

  $effect(() => {
    if (currentProject && showEditModal) {
      editTitle = currentProject.title;
      editDescription = currentProject.description;
      editSkills = currentProject.skills;
      editStartDate = new Date(currentProject.start_time)
        .toISOString()
        .split("T")[0];
      editEndDate = new Date(currentProject.end_time)
        .toISOString()
        .split("T")[0];
      editMaxMember = currentProject.max_member;
      editStatus = currentProject.status;
    }
  });

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

  // Updated to use our server routes
  const deleteProject = async () => {
    try {
      if (!currentProject) return;

      const response = await fetch(`/api/projects/${currentProject.id}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to delete project");
      }

      // Remove the deleted project from the display list
      projectsDisplay = projectsDisplay.filter(
        (p) => p.id !== currentProject?.id
      );

      // Close the modal after successful deletion
      closeDeleteModal();
    } catch (error) {
      console.error("Error deleting project:", error);
      alert("Failed to delete project: " + error);
    }
  };

  // Updated to use our server routes
  const updateProject = async () => {
    try {
      if (!currentProject) return;

      const response = await fetch(`/api/projects/${currentProject.id}`, {
        method: "PUT",
        body: JSON.stringify({
          title: editTitle,
          description: editDescription,
          skills: editSkills,
          start_time: new Date(editStartDate),
          end_time: new Date(editEndDate),
          max_member: editMaxMember,
          status: editStatus,
        }),
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to update project");
      }

      // Update the local state with the edited project
      const updatedProject = await response.json();
      projectsDisplay = projectsDisplay.map((p) =>
        p.id === currentProject?.id ? updatedProject : p
      );

      // Close the modal after successful update
      closeEditModal();
    } catch (error) {
      console.error("Error updating project:", error);
      alert("Failed to update project: " + error);
    }
  };
</script>

<header class="flex justify-between items-center mb-4 ml-64 pr-4 pl-4 pt-4">
  <button class="btn-secondary">
    <a href="/project/application">
      <div class="flex items-center">
        <div
          class="flex items-center justify-center w-5 h-5 bg-red-500 text-white text-xs font-medium rounded-full mr-2"
        >
          5
        </div>
        Applications
      </div>
    </a>
  </button>

  {#if data.role === "business"}
    <button class="btn">
      <a href="/project/create">
        <div class="flex items-center">
          <svg
            class="w-4 h-4 mr-1"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 6v6m0 0v6m0-6h6m-6 0H6"
            ></path>
          </svg>
          Create Project
        </div>
      </a>
    </button>
  {/if}
</header>

<main class="flex-1 pr-4 pl-4 ml-64 pt-4">
  <div class="flex space-x-4">
    <!-- Phần danh sách dự án active -->
    <div class="w-3/5 space-y-3">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-semibold">Your Active Projects</h3>
        <div class="flex space-x-2">
          <button
            class={filterState === "all"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "all")}
          >
            All
          </button>
          <button
            class={filterState === "open"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "open")}
          >
            Open
          </button>
          <button
            class={filterState === "close"
              ? "text-sm text-[#6b48ff] border-b-2 hover:text-[#6b48ff]"
              : "text-sm text-gray-500 hover:text-[#6b48ff]"}
            onclick={() => (filterState = "close")}
          >
            Close
          </button>
        </div>
      </div>
      {#if projectsDisplay.length !== 0 && !errorLoadingProjects}
        {#each projectsDisplay as ProjectDisplay}
          <!-- Project Card demo -->
          <a
            href={"/project/" + ProjectDisplay.id}
            class="card p-3 block relative group"
          >
            <div class="flex justify-between items-center">
              <div>
                <h4 class="text-base font-medium">{ProjectDisplay.title}</h4>
                <p class="text-sm text-gray-500">
                  Skills: {ProjectDisplay.skills}
                </p>
                <p class="text-xs text-gray-400 mt-1">
                  Progress: <span class="accent-color">60%</span> | Start: {formatDate(
                    ProjectDisplay.start_time.toString()
                  )}
                  | End: {formatDate(ProjectDisplay.end_time.toString())}
                </p>
              </div>

              <div class="flex space-x-2">
                <div class="flex flex-row gap-1">
                  <p class="text-sm text-gray-500">
                    {ProjectDisplay.current_member}/{ProjectDisplay.max_member}
                  </p>
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="gray"
                    height="20px"
                    width="px"
                    version="1.1"
                    id="Layer_1"
                    viewBox="0 0 315.539 315.539"
                    xml:space="preserve"
                  >
                    <g>
                      <g>
                        <g>
                          <path
                            d="M38.877,38.884C17.44,38.884,0,56.325,0,77.761c0,21.436,17.44,38.877,38.877,38.877s38.877-17.44,38.877-38.877     S60.314,38.884,38.877,38.884z M38.877,101.638C25.711,101.638,15,90.927,15,77.761s10.711-23.877,23.877-23.877     s23.877,10.711,23.877,23.877S52.043,101.638,38.877,101.638z"
                          />
                          <path
                            d="M38.877,131.65c-20.073,0-36.404,16.331-36.404,36.404v101.101c0,4.143,3.358,7.5,7.5,7.5h57.808     c4.142,0,7.5-3.357,7.5-7.5V168.054C75.281,147.981,58.95,131.65,38.877,131.65z M60.281,261.655H17.473v-93.601     c0-11.802,9.602-21.404,21.404-21.404c11.803,0,21.404,9.601,21.404,21.404V261.655z"
                          />
                          <path
                            d="M157.77,38.884c-21.437,0-38.877,17.44-38.877,38.877s17.44,38.877,38.877,38.877s38.877-17.44,38.877-38.877     S179.207,38.884,157.77,38.884z M157.77,101.638c-13.166,0-23.877-10.711-23.877-23.877s10.711-23.877,23.877-23.877     s23.877,10.711,23.877,23.877S170.936,101.638,157.77,101.638z"
                          />
                          <path
                            d="M157.769,131.65c-20.073,0-36.404,16.331-36.404,36.404v101.101c0,4.143,3.358,7.5,7.5,7.5h57.808     c4.142,0,7.5-3.357,7.5-7.5V168.054C194.173,147.981,177.842,131.65,157.769,131.65z M179.174,261.655h-42.808v-93.601     c0-11.802,9.602-21.404,21.404-21.404s21.404,9.601,21.404,21.404V261.655z"
                          />
                          <path
                            d="M276.662,116.638c21.437,0,38.877-17.44,38.877-38.877s-17.44-38.877-38.877-38.877s-38.877,17.44-38.877,38.877     S255.225,116.638,276.662,116.638z M276.662,53.884c13.166,0,23.877,10.711,23.877,23.877s-10.711,23.877-23.877,23.877     c-13.166,0-23.877-10.711-23.877-23.877S263.496,53.884,276.662,53.884z"
                          />
                          <path
                            d="M276.662,131.65c-20.073,0-36.404,16.331-36.404,36.404v101.101c0,4.143,3.358,7.5,7.5,7.5h57.808     c4.142,0,7.5-3.357,7.5-7.5V168.054C313.066,147.981,296.735,131.65,276.662,131.65z M298.066,261.655h-42.808v-93.601     c0-11.802,9.602-21.404,21.404-21.404s21.404,9.601,21.404,21.404V261.655z"
                          />
                        </g>
                      </g>
                    </g>
                  </svg>
                </div>
                <span
                  class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded font-bold"
                  >{ProjectDisplay.status.toUpperCase()}</span
                >
                <svg
                  class="w-5 h-5 text-gray-400"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 5l7 7-7 7"
                  ></path>
                </svg>
              </div>
            </div>
          </a>
          <!-- Project action buttons - Only show for business users and only on hover -->
          {#if data.role === "business"}
            <div
              class="group-hover:opacity-100 transition-opacity flex space-x-2"
            >
              <button
                class="p-1.5 bg-blue-100 text-blue-600 rounded hover:bg-blue-200"
                onclick={() => openEditModal(ProjectDisplay)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
                  />
                </svg>
              </button>
              <button
                class="p-1.5 bg-red-100 text-red-600 rounded hover:bg-red-200"
                onclick={() => {
                  openDeleteModal(ProjectDisplay);
                }}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
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
          {/if}
        {/each}
      {:else}
        <div class="card p-4 text-center">
          <p class="text-gray-500">No active projects found.</p>
          <!-- <button class="btn mt-2">Create Your First Project</button> -->
        </div>
      {/if}
    </div>

    {#if data.role === "business"}
      <!-- Right sidebar with analytics and talent pool -->
      <div class="w-2/5 space-y-4">
        <!-- Business Analytics -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Project Analytics</h3>
          <div class="space-y-4">
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">Active Projects</span>
                <span class="text-sm font-medium">4</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-[#6b48ff]" style="width: 80%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">Talent Allocation</span>
                <span class="text-sm font-medium">12/15</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-[#ff6f61]" style="width: 75%"></div>
              </div>
            </div>
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm">Tasks Completion</span>
                <span class="text-sm font-medium">92%</span>
              </div>
              <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
                <div class="h-full bg-green-500" style="width: 92%"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Active Talent Pool -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Talent Pool</h3>
          <div class="space-y-2">
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/234567?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Sarah Johnson</p>
                  <p class="text-xs text-gray-500">React, Node.js</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">View</button>
            </div>
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/345678?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Michael Chen</p>
                  <p class="text-xs text-gray-500">Python, MongoDB</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">View</button>
            </div>
            <div
              class="flex items-center justify-between p-2 bg-gray-100 rounded"
            >
              <div class="flex items-center">
                <img
                  class="w-8 h-8 rounded-full mr-2"
                  src="https://avatars.githubusercontent.com/u/456789?v=4"
                  alt="Developer"
                />
                <div>
                  <p class="text-sm font-medium">Alex Rivera</p>
                  <p class="text-xs text-gray-500">UI/UX, Vue.js</p>
                </div>
              </div>
              <button class="btn-secondary text-xs">View</button>
            </div>
            <div class="flex justify-center mt-2">
              <button class="text-sm text-[#6b48ff] hover:underline">
                View All Talent
              </button>
            </div>
          </div>
        </div>
      </div>
    {/if}
  </div>

  <!-- Edit Project Modal -->
  {#if showEditModal && currentProject}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div
        class="bg-white rounded-lg p-6 w-full max-w-2xl max-h-[90vh] overflow-y-auto"
      >
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-xl font-semibold">Edit Project</h2>
          <button
            class="text-gray-500 hover:text-gray-700"
            onclick={closeEditModal}
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

        <form class="space-y-4">
          <div>
            <label for="title" class="block text-sm font-medium text-gray-700"
              >Project Title</label
            >
            <input
              type="text"
              id="title"
              bind:value={editTitle}
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            />
          </div>

          <div>
            <label
              for="description"
              class="block text-sm font-medium text-gray-700">Description</label
            >
            <textarea
              id="description"
              bind:value={editDescription}
              rows="4"
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            ></textarea>
          </div>

          <div>
            <label for="skills" class="block text-sm font-medium text-gray-700"
              >Skills (comma separated)</label
            >
            <input
              type="text"
              id="skills"
              bind:value={editSkills}
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            />
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label
                for="startDate"
                class="block text-sm font-medium text-gray-700"
                >Start Date</label
              >
              <input
                type="date"
                id="startDate"
                bind:value={editStartDate}
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              />
            </div>

            <div>
              <label
                for="endDate"
                class="block text-sm font-medium text-gray-700">End Date</label
              >
              <input
                type="date"
                id="endDate"
                bind:value={editEndDate}
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              />
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label
                for="maxMember"
                class="block text-sm font-medium text-gray-700"
                >Max Members</label
              >
              <input
                type="number"
                id="maxMember"
                bind:value={editMaxMember}
                min={currentProject.current_member}
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              />
            </div>

            <div>
              <label
                for="status"
                class="block text-sm font-medium text-gray-700">Status</label
              >
              <select
                id="status"
                bind:value={editStatus}
                class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              >
                <option value="open">Open</option>
                <option value="close">Closed</option>
              </select>
            </div>
          </div>

          <div class="flex justify-end space-x-2 pt-4">
            <button
              type="button"
              class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
              onclick={closeEditModal}
            >
              Cancel
            </button>
            <button
              type="button"
              class="px-4 py-2 bg-[#6b48ff] text-white rounded-md hover:bg-[#5a3dd3]"
              onclick={updateProject}
            >
              Save Changes
            </button>
          </div>
        </form>
      </div>
    </div>
  {/if}

  <!-- Delete Project Modal -->
  {#if showDeleteModal && currentProject}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 modal"
    >
      <div class="bg-white rounded-lg p-6 w-full max-w-md">
        <div class="text-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-12 w-12 mx-auto text-red-500"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
            />
          </svg>

          <h2 class="text-xl font-semibold mt-4">Delete Project</h2>

          <p class="mt-2 text-gray-600">
            Are you sure you want to delete "{currentProject.title}"? This
            action cannot be undone.
          </p>

          <div class="flex justify-center space-x-3 mt-6">
            <button
              class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300"
              onclick={closeDeleteModal}
            >
              Cancel
            </button>
            <button
              class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
              onclick={deleteProject}
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </div>
  {/if}
</main>

<style>
  [contenteditable]:empty:before {
    content: attr(placeholder);
    color: #9ca3af;
    display: block;
  }
</style>
