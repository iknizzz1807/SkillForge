<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();
  const project = data.project;

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

  // Format date and time function (with time)
  function formatDateTime(dateString: string): string {
    if (!dateString) return "N/A";

    const date = new Date(dateString);

    // Check if date is valid
    if (isNaN(date.getTime())) return "Invalid date";

    // Format: "Mar 24, 2025, 2:30 PM"
    return date.toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
      hour: "numeric",
      minute: "numeric",
      hour12: true,
    });
  }

  // Calculate project progress based on start and end dates
  function calculateProgress(start: string, end: string): number {
    const startDate = new Date(start);
    const endDate = new Date(end);
    const currentDate = new Date();

    // If dates are invalid, return 0
    if (isNaN(startDate.getTime()) || isNaN(endDate.getTime())) {
      return 0;
    }

    // If project hasn't started yet
    if (currentDate < startDate) {
      return 0;
    }

    // If project is completed
    if (currentDate > endDate) {
      return 100;
    }

    // Calculate progress percentage
    const totalDuration = endDate.getTime() - startDate.getTime();
    const elapsedDuration = currentDate.getTime() - startDate.getTime();
    const progress = Math.floor((elapsedDuration / totalDuration) * 100);

    return progress;
  }

  // Calculate progress percentage
  const progressPercentage = calculateProgress(
    project.start_time,
    project.end_time
  );

  // Task creation modal
  let showTaskModal = $state(false);

  let newTasks = $state([
    { name: "", description: "", note: "", assignedTo: "" },
  ]);
  let teamMembers = $state([
    { id: 1, name: "John Doe" },
    { id: 2, name: "Sarah Smith" },
    { id: 3, name: "Alex Johnson" },
    { id: 4, name: "Emma Wilson" },
  ]);

  function addAnotherTask() {
    newTasks.push({ name: "", description: "", note: "", assignedTo: "" });
  }

  function removeTask(index: number) {
    newTasks = newTasks.filter((_, i) => i !== index);
  }

  function openTaskModal() {
    showTaskModal = true;
  }

  function closeTaskModal() {
    showTaskModal = false;
    // Reset tasks form
    newTasks = [{ name: "", description: "", note: "", assignedTo: "" }];
  }

  function createTasksWithConfirm() {
    // Filter out empty tasks
    const validTasks = newTasks.filter((task) => task.name.trim() !== "");

    if (validTasks.length === 0) {
      alert("Please add at least one task with a name");
      return;
    }

    createTasks();
  }

  function createTasks() {
    // Here you would send the tasks to your backend
    // console.log("Creating tasks:", validTasks);

    closeTaskModal();
  }

  // Activities mock data
  let showActivitiesModal = $state(false);

  // Mock data for activities
  const mockActivities = $state([
    {
      id: 1,
      user: "Sarah",
      action: "moved",
      item: "Test Features",
      targetStatus: "Review",
      timestamp: "2025-04-24T14:32:00",
    },
    {
      id: 2,
      user: "John",
      action: "moved",
      item: "Code Backend",
      targetStatus: "In Progress",
      timestamp: "2025-04-24T13:45:00",
    },
    {
      id: 3,
      user: "Alex",
      action: "completed",
      item: "Login Screen UI",
      targetStatus: "Done",
      timestamp: "2025-04-24T11:20:00",
    },
    {
      id: 4,
      user: "Emma",
      action: "created",
      item: "New Feature Request",
      targetStatus: "To-do",
      timestamp: "2025-04-23T16:15:00",
    },
    {
      id: 5,
      user: "Michael",
      action: "commented on",
      item: "API Documentation",
      targetStatus: "",
      timestamp: "2025-04-23T14:32:00",
    },
    {
      id: 6,
      user: "Tran Le Minh Nhat",
      action: "assigned",
      item: "Wireframes",
      targetStatus: "To-do",
      timestamp: "2025-04-22T09:30:00",
    },
    {
      id: 7,
      user: "Nguyen My Thong",
      action: "updated",
      item: "Authentication Logic",
      targetStatus: "",
      timestamp: "2025-04-21T17:25:00",
    },
    {
      id: 8,
      user: "Tran Tuan Kiet",
      action: "completed",
      item: "Database Schema",
      targetStatus: "Done",
      timestamp: "2025-04-21T15:10:00",
    },
    {
      id: 9,
      user: "Tran Le Minh Nhat",
      action: "created",
      item: "Project Repository",
      targetStatus: "",
      timestamp: "2025-04-20T11:45:00",
    },
    {
      id: 10,
      user: "Sarah",
      action: "added",
      item: "New Team Member",
      targetStatus: "",
      timestamp: "2025-04-19T10:30:00",
    },
  ]);

  // Function to format the date for grouping
  function formatActivityDate(dateString: string): string {
    const date = new Date(dateString);

    // Get today and yesterday dates for comparison
    const today = new Date();
    today.setHours(0, 0, 0, 0);

    const yesterday = new Date(today);
    yesterday.setDate(yesterday.getDate() - 1);

    // Check if the date is today, yesterday, or earlier
    if (date >= today) {
      return "Today";
    } else if (date >= yesterday) {
      return "Yesterday";
    } else {
      return date.toLocaleDateString("en-US", {
        month: "short",
        day: "numeric",
        year: "numeric",
      });
    }
  }

  // Function to format the time
  function formatActivityTime(dateString: string): string {
    const date = new Date(dateString);
    return date.toLocaleTimeString("en-US", {
      hour: "2-digit",
      minute: "2-digit",
    });
  }

  // Function to get the grouped activities
  function getGroupedActivities() {
    const groups: Record<string, typeof mockActivities> = {};

    mockActivities.forEach((activity) => {
      const dateGroup = formatActivityDate(activity.timestamp);
      if (!groups[dateGroup]) {
        groups[dateGroup] = [];
      }
      groups[dateGroup].push(activity);
    });

    // Convert the groups object to an array of { date, activities } objects
    return Object.entries(groups).map(([date, activities]) => {
      return { date, activities };
    });
  }

  function openActivitiesModal() {
    showActivitiesModal = true;
  }

  function closeActivitiesModal() {
    showActivitiesModal = false;
  }
</script>

<header class="flex justify-between items-center ml-64 pr-4 pl-4 pt-4">
  <div class="flex items-center mb-4">
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
    <div>
      <h2 class="text-xl font-semibold">{project.title}</h2>
      <p class="text-sm text-gray-600">Manage your project</p>
    </div>
  </div>
</header>

<main class="flex-1 pr-4 pl-4 ml-64">
  <div class="space-y-4">
    <!-- Top section: Timeline, Progression, and Recent Activities -->
    <div class="flex space-x-4">
      <!-- Timeline and Progression -->
      <div class="flex-1 space-y-4">
        <!-- Timeline Bar -->
        <div class="card p-3 w-full">
          <h3 class="text-base font-semibold mb-2">Timeline</h3>
          <div class="flex items-center justify-between">
            <div class="text-center">
              <p class="text-sm font-medium">Start</p>
              <p class="text-xs text-gray-500">
                {formatDate(project.start_time.toString())}
              </p>
            </div>
            <div class="flex-1 h-2 bg-gray-200 rounded mx-2 relative">
              <div
                class="absolute h-2 bg-[#6b48ff] rounded"
                style="width: {progressPercentage}%"
              ></div>
            </div>
            <div class="text-center">
              <p class="text-sm font-medium">End</p>
              <p class="text-xs text-gray-500">
                {formatDate(project.end_time.toString())}
              </p>
            </div>
          </div>
          <p class="text-sm mt-2">
            Progress: <span class="accent-color">{progressPercentage}%</span>
          </p>
        </div>

        <!-- New Progression Bar -->
        <div class="card p-3 w-full">
          <h3 class="text-base font-semibold mb-2">Progression</h3>
          <div class="flex items-center justify-between">
            <div class="text-center">
              <p class="text-sm font-medium">Tasks</p>
              <p class="text-xs text-gray-500">Completed</p>
            </div>
            <div class="flex-1 h-2 bg-gray-200 rounded mx-2 relative">
              <div
                class="absolute h-2 bg-green-500 rounded"
                style="width: 35%"
              ></div>
            </div>
            <div class="text-center">
              <p class="text-sm font-medium">Total</p>
              <p class="text-xs text-gray-500">4/12</p>
            </div>
          </div>
          <p class="text-sm mt-2">
            Tasks completed: <span class="text-green-600">35%</span>
          </p>
        </div>
      </div>

      <!-- Right sidebar - Recent Activities -->
      <div class="w-72">
        <div class="card p-4 sticky top-4">
          <h3 class="text-base font-semibold mb-2">Recent Activities</h3>
          <!-- Timeline -->
          <div class="relative">
            <!-- Timeline line -->
            <div class="absolute left-4 top-0 bottom-0 w-0.5 bg-gray-200"></div>

            <!-- Activity items -->
            <div class="space-y-6">
              <!-- Today's activities -->
              <div class="relative pl-8">
                <!-- Timeline dot -->
                <div
                  class="absolute left-[13px] top-1 w-3 h-3 rounded-full bg-[#6b48ff] border-2 border-white"
                ></div>
                <span class="text-xs font-semibold text-gray-500 mb-1 block"
                  >18 Mar, 14:32</span
                >
                <p class="text-sm">
                  <span class="font-medium text-[#6b48ff]">Sarah</span>
                  moved <span class="font-medium">Test Features</span> to Review
                </p>
              </div>

              <div class="relative pl-8">
                <!-- Timeline dot -->
                <div
                  class="absolute left-[13px] top-1 w-3 h-3 rounded-full bg-[#6b48ff] border-2 border-white"
                ></div>
                <span class="text-xs font-semibold text-gray-500 mb-1 block"
                  >18 Mar, 13:45</span
                >
                <p class="text-sm">
                  <span class="font-medium text-[#6b48ff]">John</span> moved
                  <span class="font-medium">Code Backend</span> to In Progress
                </p>
              </div>
            </div>
          </div>

          <!-- View all button -->
          <button
            class="btn-secondary w-full mt-4 text-sm"
            onclick={openActivitiesModal}
          >
            View All Activities
          </button>
        </div>
      </div>
    </div>

    <!-- Kanban Board (full width, full height) -->
    <div class="card p-3 w-full min-h-[calc(100vh-380px)] mb-4">
      <h3 class="text-base font-semibold mb-2">Tasks</h3>
      <div class="grid grid-cols-4 gap-3 h-full">
        <!-- To-do -->
        <div class="flex flex-col">
          <div class="flex items-center justify-between mb-2">
            <div class="flex items-center space-x-2">
              <div class="p-1.5 rounded-md bg-gray-100">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4 text-gray-600"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <circle cx="12" cy="12" r="10"></circle>
                  <line x1="12" y1="8" x2="12" y2="16"></line>
                  <line x1="8" y1="12" x2="16" y2="12"></line>
                </svg>
              </div>
              <h4 class="text-sm font-medium">To-do</h4>
            </div>
            <span
              class="inline-flex items-center justify-center w-5 h-5 text-xs font-medium rounded-full bg-gray-100 text-gray-700"
              >3</span
            >
          </div>
          <button class="text-xs text-[#6b48ff] mb-2" onclick={openTaskModal}
            >+ Add tasks</button
          >
          <div class="bg-gray-50 rounded-lg p-2 flex-1 overflow-y-auto">
            <div class="card p-2 bg-white mb-2 shadow-sm">
              <div class="flex justify-between items-center">
                <div class="flex-1">
                  <p class="text-sm font-medium">Create Wireframes</p>
                  <p class="text-xs text-gray-500">
                    Assign to:
                    <a href="#" class="text-[#6b48ff]">Tran Le Minh Nhat</a>
                  </p>
                </div>
                <button
                  class="text-gray-400 hover:text-gray-600 ml-2 p-1 rounded-full hover:bg-gray-100 self-center flex-shrink-0"
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
                      d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"
                    />
                  </svg>
                </button>
              </div>
            </div>

            <div class="card p-2 bg-white mb-2 shadow-sm">
              <div class="flex justify-between items-center">
                <div class="flex-1">
                  <p class="text-sm font-medium">Create Wireframes</p>
                  <p class="text-xs text-gray-500">
                    Assign to:
                    <a href="#" class="text-[#6b48ff]">Tran Le Minh Nhat</a>
                  </p>
                </div>
                <button
                  class="text-gray-400 hover:text-gray-600 ml-2 p-1 rounded-full hover:bg-gray-100 self-center flex-shrink-0"
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
                      d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"
                    />
                  </svg>
                </button>
              </div>
            </div>

            <div class="card p-2 bg-white shadow-sm">
              <p class="text-sm font-medium">Research API</p>
              <p class="text-xs text-gray-500">
                Assign to:
                <a href="#" class="text-[#6b48ff]">Tran Tuan Kiet</a>
              </p>
            </div>
          </div>
        </div>

        <!-- In Progress -->
        <div class="flex flex-col">
          <div class="flex items-center justify-between mb-2">
            <div class="flex items-center space-x-2">
              <div class="p-1.5 rounded-md bg-yellow-100">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4 text-yellow-600"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <path
                    d="M12 2v4M12 18v4M4.93 4.93l2.83 2.83M16.24 16.24l2.83 2.83M2 12h4M18 12h4M4.93 19.07l2.83-2.83M16.24 7.76l2.83-2.83"
                  ></path>
                </svg>
              </div>
              <h4 class="text-sm font-medium">In Progress</h4>
            </div>
            <span
              class="inline-flex items-center justify-center w-5 h-5 text-xs font-medium rounded-full bg-yellow-100 text-yellow-700"
              >2</span
            >
          </div>
          <div class="bg-yellow-50 rounded-lg p-2 flex-1 overflow-y-auto">
            <div class="card p-2 bg-white mb-2 shadow-sm">
              <p class="text-sm font-medium">Code Backend</p>
              <p class="text-xs text-gray-500">
                Doing by:
                <a href="#" class="text-[#6b48ff]">Tran Tuan Kiet</a>
              </p>
            </div>
            <div class="card p-2 bg-white shadow-sm">
              <p class="text-sm font-medium">Implement Auth</p>
              <p class="text-xs text-gray-500">
                Doing by:
                <a href="#" class="text-[#6b48ff]">Nguyen My Thong</a>
              </p>
            </div>
          </div>
        </div>

        <!-- Review -->
        <div class="flex flex-col">
          <div class="flex items-center justify-between mb-2">
            <div class="flex items-center space-x-2">
              <div class="p-1.5 rounded-md bg-blue-100">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4 text-blue-600"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z"></path>
                  <circle cx="12" cy="12" r="3"></circle>
                </svg>
              </div>
              <h4 class="text-sm font-medium">Review</h4>
            </div>
            <span
              class="inline-flex items-center justify-center w-5 h-5 text-xs font-medium rounded-full bg-blue-100 text-blue-700"
              >1</span
            >
          </div>
          <div class="bg-blue-50 rounded-lg p-2 flex-1 overflow-y-auto">
            <div class="card p-2 bg-white shadow-sm">
              <p class="text-sm font-medium">Test Features</p>
              <p class="text-xs text-gray-500">
                Done by:
                <a href="#" class="text-[#6b48ff]">Nguyen My Thong</a>
              </p>
            </div>
          </div>
        </div>

        <!-- Done -->
        <div class="flex flex-col">
          <div class="flex items-center justify-between mb-2">
            <div class="flex items-center space-x-2">
              <div class="p-1.5 rounded-md bg-green-100">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4 text-green-600"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <path d="M20 6L9 17l-5-5"></path>
                </svg>
              </div>
              <h4 class="text-sm font-medium">Done</h4>
            </div>
            <span
              class="inline-flex items-center justify-center w-5 h-5 text-xs font-medium rounded-full bg-green-100 text-green-700"
              >6</span
            >
          </div>
          <div class="bg-green-50 rounded-lg p-2 flex-1 overflow-y-auto">
            <div class="card p-2 bg-white mb-2 shadow-sm">
              <p class="text-sm font-medium">Setup DB</p>
              <p class="text-xs text-gray-500">
                Completed by:
                <a href="#" class="text-[#6b48ff]">Nguyen My Thong</a>
              </p>
            </div>

            <div class="card p-2 bg-white mb-2 shadow-sm">
              <p class="text-sm font-medium">Project Setup</p>
              <p class="text-xs text-gray-500">
                Completed by:
                <a href="#" class="text-[#6b48ff]">Tran Le Minh Nhat</a>
              </p>
            </div>

            <div class="card p-2 bg-white mb-2 shadow-sm">
              <p class="text-sm font-medium">Create Repository</p>
              <p class="text-xs text-gray-500">
                Completed by:
                <a href="#" class="text-[#6b48ff]">Tran Tuan Kiet</a>
              </p>
            </div>

            <div class="card p-2 bg-white shadow-sm">
              <p class="text-sm font-medium">Initial Planning</p>
              <p class="text-xs text-gray-500">
                Completed by:
                <a href="#" class="text-[#6b48ff]">Nguyen My Thong</a>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Activities Modal -->
  {#if showActivitiesModal}
    <div
      class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center modal"
    >
      <div
        class="bg-white rounded-lg w-full max-w-2xl max-h-[90vh] flex flex-col overflow-hidden shadow-lg"
      >
        <div class="p-4 border-b border-gray-100 bg-white flex-shrink-0">
          <div class="flex justify-between items-center">
            <h3 class="text-lg font-semibold">Project Activities</h3>
            <button
              class="text-gray-500 hover:text-gray-700"
              onclick={closeActivitiesModal}
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

        <!-- Content - phần này sẽ cuộn -->
        <div class="p-4 pb-6 overflow-y-auto flex-grow">
          {#each getGroupedActivities() as group}
            <div class="mb-6">
              <h4 class="text-sm font-semibold text-gray-500 mb-3">
                {group.date}
              </h4>

              <!-- Timeline -->
              <div class="relative">
                <!-- Timeline line -->
                <div
                  class="absolute left-4 top-0 bottom-0 w-0.5 bg-gray-200"
                ></div>

                <!-- Activity items -->
                <div class="space-y-4">
                  {#each group.activities as activity}
                    <div class="relative pl-10">
                      <!-- Timeline dot -->
                      <div
                        class="absolute left-[13px] top-1 w-3 h-3 rounded-full bg-[#6b48ff] border-2 border-white"
                      ></div>

                      <div class="flex justify-between items-start">
                        <div>
                          <span
                            class="text-xs font-semibold text-gray-500 mb-1 block"
                          >
                            {formatActivityTime(activity.timestamp)}
                          </span>
                          <p class="text-sm">
                            <span class="font-medium text-[#6b48ff]"
                              >{activity.user}</span
                            >
                            {activity.action}
                            <span class="font-medium">{activity.item}</span>
                            {#if activity.targetStatus}
                              to <span class="font-medium"
                                >{activity.targetStatus}</span
                              >
                            {/if}
                          </p>
                        </div>
                      </div>
                    </div>
                  {/each}
                </div>
              </div>
            </div>
          {/each}
        </div>
      </div>
    </div>
  {/if}
</main>

<!-- Task Creation Modal -->
{#if showTaskModal}
  <div
    class="fixed inset-0 bg-black z-50 flex items-center justify-center modal"
  >
    <div
      class="bg-white w-full max-w-2xl max-h-[90vh] overflow-y-auto shadow-lg"
    >
      <div class="p-4 border-b border-gray-100">
        <div class="flex justify-between items-center">
          <h3 class="text-lg font-semibold">Create New Tasks</h3>
          <button
            class="text-gray-500 hover:text-gray-700"
            onclick={closeTaskModal}
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

      <div class="p-4">
        <!-- Task form repeater -->
        {#each newTasks as task, index}
          <div class="card p-3 mb-4 bg-white shadow-sm">
            <div class="flex justify-between items-center mb-2">
              <h4 class="font-medium">Task {index + 1}</h4>
              {#if newTasks.length > 1}
                <button
                  class="text-gray-500 hover:text-red-500"
                  onclick={() => removeTask(index)}
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
              {/if}
            </div>

            <div class="space-y-3">
              <div>
                <label class="block text-sm font-medium mb-1">Task Name*</label>
                <input
                  type="text"
                  class="w-full p-2 border border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                  placeholder="Enter task name"
                  bind:value={task.name}
                />
              </div>

              <div>
                <label class="block text-sm font-medium mb-1"
                  >Description*</label
                >
                <textarea
                  class="w-full p-2 bg-white border border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                  placeholder="Enter task description"
                  rows="2"
                  bind:value={task.description}
                ></textarea>
              </div>

              <div>
                <label class="block text-sm font-medium mb-1">Note</label>
                <textarea
                  class="w-full p-2 border bg-white border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                  placeholder="Note if there is"
                  rows="1"
                  bind:value={task.note}
                ></textarea>
              </div>

              <div class="grid grid-cols-2 gap-3">
                <div>
                  <label class="block text-sm font-medium mb-1"
                    >Assigned To</label
                  >
                  <select
                    class="w-full p-2 border bg-white border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                    bind:value={task.assignedTo}
                  >
                    <option value="">Unassigned</option>
                    {#each teamMembers as member}
                      <option value={member.id}>{member.name}</option>
                    {/each}
                  </select>
                </div>
              </div>
            </div>
          </div>
        {/each}

        <!-- Add another task button -->
        <div class="flex justify-center">
          <button
            class="p-2 card shadow-sm bg-white text-gray-600 hover:bg-gray-50 mb-4 flex items-center justify-center"
            onclick={addAnotherTask}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5 mr-1 text-[#6b48ff]"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 4v16m8-8H4"
              />
            </svg>
            Add Another Task
          </button>
        </div>
      </div>

      <div class="p-4 border-t border-gray-100 flex justify-end space-x-3">
        <button
          class="px-4 py-2 border border-gray-200 rounded-lg hover:bg-gray-50"
          onclick={closeTaskModal}
        >
          Cancel
        </button>
        <button
          class="px-4 py-2 bg-[#6b48ff] text-white rounded-lg hover:bg-[#5a3dd4]"
          onclick={createTasksWithConfirm}
        >
          Create Tasks
        </button>
      </div>
    </div>
  </div>
{/if}

<style>
  .modal {
    background-color: rgba(0, 0, 0, 0.5);
  }

  .card {
    border-radius: 0.5rem;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  }
</style>
