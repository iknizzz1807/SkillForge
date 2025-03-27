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
</script>

<header class="flex justify-between items-center ml-64 pr-4 pl-4 pt-4">
  <div class="flex items-center mb-6">
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
                {formatDate(project.start_time)}
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
                {formatDate(project.end_time)}
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

              <!-- <div class="relative pl-8">
                <div
                  class="absolute left-[13px] top-1 w-3 h-3 rounded-full bg-[#6b48ff] border-2 border-white"
                ></div>
                <span class="text-xs font-semibold text-gray-500 mb-1 block"
                  >18 Mar, 10:17</span
                >
                <p class="text-sm">
                  <span class="font-medium text-[#6b48ff]">Adam</span>
                  created
                  <span class="font-medium">Make a button style</span> in To-do
                </p>
              </div>

              <div class="relative pl-8">
                <div
                  class="absolute left-[13px] top-1 w-3 h-3 rounded-full bg-gray-400 border-2 border-white"
                ></div>
                <span class="text-xs font-semibold text-gray-500 mb-1 block"
                  >17 Mar, 16:23</span
                >
                <p class="text-sm">
                  <span class="font-medium">Alex</span> completed
                  <span class="font-medium">Setup DB</span>
                </p>
              </div>

              <div class="relative pl-8">
                <div
                  class="absolute left-[13px] top-1 w-3 h-3 rounded-full bg-gray-400 border-2 border-white"
                ></div>
                <span class="text-xs font-semibold text-gray-500 mb-1 block"
                  >17 Mar, 11:05</span
                >
                <p class="text-sm">
                  <span class="font-medium">Emma</span> created
                  <span class="font-medium">Test Features</span> task
                </p>
              </div> -->
            </div>
          </div>

          <!-- View all button -->
          <button class="btn-secondary w-full mt-4 text-sm">
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
          <button class="text-xs text-[#6b48ff] mb-2">+ Add tasks</button>
          <div class="bg-gray-50 rounded-lg p-2 flex-1 overflow-y-auto">
            <div class="card p-2 bg-white mb-2 shadow-sm">
              <p class="text-sm font-medium">Design UI</p>
              <p class="text-xs text-gray-500">Due: 20/03</p>
            </div>
            <div class="card p-2 bg-white mb-2 shadow-sm">
              <p class="text-sm font-medium">Create Wireframes</p>
              <p class="text-xs text-gray-500">Due: 22/03</p>
            </div>
            <div class="card p-2 bg-white shadow-sm">
              <p class="text-sm font-medium">Research API</p>
              <p class="text-xs text-gray-500">Due: 25/03</p>
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
              <p class="text-xs text-gray-500">Due: 22/03</p>
            </div>
            <div class="card p-2 bg-white shadow-sm">
              <p class="text-sm font-medium">Implement Auth</p>
              <p class="text-xs text-gray-500">Due: 24/03</p>
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
              <p class="text-xs text-gray-500">Due: 25/03</p>
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
              <p class="text-xs text-gray-500">Completed</p>
            </div>

            <div class="card p-2 bg-white mb-2 shadow-sm">
              <p class="text-sm font-medium">Project Setup</p>
              <p class="text-xs text-gray-500">Completed</p>
            </div>

            <div class="card p-2 bg-white mb-2 shadow-sm">
              <p class="text-sm font-medium">Create Repository</p>
              <p class="text-xs text-gray-500">Completed</p>
            </div>

            <div class="card p-2 bg-white shadow-sm">
              <p class="text-sm font-medium">Initial Planning</p>
              <p class="text-xs text-gray-500">Completed</p>
            </div>

            <div class="card p-2 bg-white shadow-sm">
              <p class="text-sm font-medium">Initial Planning</p>
              <p class="text-xs text-gray-500">Completed</p>
            </div>

            <div class="card p-2 bg-white shadow-sm">
              <p class="text-sm font-medium">Initial Planning</p>
              <p class="text-xs text-gray-500">Completed</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</main>
