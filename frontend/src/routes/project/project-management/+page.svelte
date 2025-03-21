<main class="flex-1 pr-4 pl-4 ml-64">
  <div class="space-y-4">
    <!-- Timeline Bar -->
    <div class="card p-3 w-full">
      <h3 class="text-base font-semibold mb-2">Timeline</h3>
      <div class="flex items-center justify-between">
        <div class="text-center">
          <p class="text-sm font-medium">Start</p>
          <p class="text-xs text-gray-500">15/03/2025</p>
        </div>
        <div class="flex-1 h-2 bg-gray-200 rounded mx-2 relative">
          <div
            class="absolute h-2 bg-[#6b48ff] rounded"
            style="width: 60%"
          ></div>
        </div>
        <div class="text-center">
          <p class="text-sm font-medium">End</p>
          <p class="text-xs text-gray-500">29/03/2025</p>
        </div>
      </div>
      <p class="text-sm mt-2">
        Progress: <span class="accent-color">60%</span>
      </p>
    </div>

    <!-- Main content grid with Kanban + Recent Activities -->
    <div class="flex space-x-4">
      <!-- Left content (Kanban + Chat) -->
      <div class="space-y-4 flex-1">
        <!-- Kanban Board -->
        <div class="card p-3 w-full">
          <h3 class="text-base font-semibold mb-3">Tasks</h3>
          <div class="grid grid-cols-4 gap-3">
            <!-- To-do -->
            <div>
              <h4 class="text-sm font-medium mb-2">To-do</h4>
              <div class="card p-2 bg-gray-100">
                <p class="text-sm">Design UI</p>
                <p class="text-xs text-gray-500">Due: 20/03</p>
              </div>
            </div>
            <!-- In Progress -->
            <div>
              <h4 class="text-sm font-medium mb-2">In Progress</h4>
              <div class="card p-2 bg-yellow-50">
                <p class="text-sm">Code Backend</p>
                <p class="text-xs text-gray-500">Due: 22/03</p>
              </div>
            </div>
            <!-- Review -->
            <div>
              <h4 class="text-sm font-medium mb-2">Review</h4>
              <div class="card p-2 bg-blue-50">
                <p class="text-sm">Test Features</p>
                <p class="text-xs text-gray-500">Due: 25/03</p>
              </div>
            </div>
            <!-- Done -->
            <div>
              <h4 class="text-sm font-medium mb-2">Done</h4>
              <div class="card p-2 bg-green-50">
                <p class="text-sm">Setup DB</p>
                <p class="text-xs text-gray-500">Completed</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Chat Panel -->
        <div class="card p-3 w-full">
          <h3 class="text-base font-semibold mb-3">
            Chat - replace this with github changes or keep (the /chat different
            will display all the damn chat)
          </h3>
          <div class="h-40 overflow-y-auto bg-gray-100 p-2 rounded mb-3">
            <div class="flex justify-start mb-2">
              <p class="text-sm bg-white p-2 rounded">
                Hey, how's the backend going?
              </p>
            </div>
            <div class="flex justify-end mb-2">
              <p
                class="text-sm p-2 rounded text-white"
                style="background-color: #6b48ff"
              >
                Good, almost done!
              </p>
            </div>
          </div>
          <div class="flex space-x-2">
            <input
              type="text"
              class="flex-1 p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              placeholder="Type a message..."
            />
            <button class="btn">Send</button>
            <button class="btn-secondary">
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
                  d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"
                ></path>
              </svg>
            </button>
          </div>
          <button class="btn-secondary mt-2">Start Video Call</button>
        </div>
      </div>

      <!-- Right sidebar - Recent Activities -->
      <div class="w-72">
        <div class="card p-4 sticky top-4">
          <h3 class="text-base font-semibold mb-4">Recent Activities</h3>
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

              <div class="relative pl-8">
                <!-- Timeline dot -->
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

              <!-- Yesterday's activities -->
              <div class="relative pl-8">
                <!-- Timeline dot -->
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
                <!-- Timeline dot -->
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
              </div>
            </div>
          </div>

          <!-- View all button -->
          <button class="btn-secondary w-full mt-4 text-sm">
            View All Activities
          </button>
        </div>
      </div>
    </div>
  </div>
</main>
