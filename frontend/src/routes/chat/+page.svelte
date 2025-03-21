<script lang="ts">
  // Lưu ý component này có các javascript rất thú vị,
  //  xem thêm ở phần /design để implement trong tương lai.
</script>

<main class="flex-1 flex flex-col ml-64 pr-4 pl-4 h-[calc(100vh-20px)] pt-4">
  <div class="flex space-x-4 flex-1">
    <!-- Chat Rooms List -->
    <div class="card p-3 w-64 flex flex-col">
      <div class="flex items-center justify-between mb-3">
        <h3 class="text-base font-semibold">Rooms</h3>
        <button class="text-[#6b48ff] hover:text-[#5a3de6]">
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
              d="M12 6v6m0 0v6m0-6h6m-6 0H6"
            ></path>
          </svg>
        </button>
      </div>

      <!-- Search Box -->
      <div class="relative mb-3">
        <input
          type="text"
          placeholder="Search rooms..."
          class="w-full p-2 pl-8 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-[#6b48ff]"
        />
        <svg
          class="w-4 h-4 absolute left-2 top-2.5 text-gray-400"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          ></path>
        </svg>
      </div>

      <!-- Room List -->
      <div class="flex-1 overflow-y-auto">
        <div class="space-y-1">
          <!-- Active Chat Room -->
          <div class="chat-room active p-2 rounded">
            <div class="flex justify-between items-center">
              <div class="flex items-center">
                <div class="w-2 h-2 bg-green-500 rounded-full mr-2"></div>
                <span class="text-sm font-medium">Build a Chat App</span>
              </div>
            </div>
            <p class="text-xs text-gray-500 mt-1 truncate">
              Jane: Hey, how's the backend going?
            </p>
          </div>

          <!-- Chat Room with Unread Messages -->
          <div class="chat-room p-2 rounded">
            <div class="flex justify-between items-center">
              <div class="flex items-center">
                <div class="w-2 h-2 bg-green-500 rounded-full mr-2"></div>
                <span class="text-sm font-medium">E-commerce Dashboard</span>
              </div>
              <span class="unread-badge">3</span>
            </div>
            <p class="text-xs text-gray-500 mt-1 truncate">
              Alex: Can we discuss the UI design?
            </p>
          </div>

          <!-- Inactive Chat Room -->
          <div class="chat-room p-2 rounded">
            <div class="flex justify-between items-center">
              <div class="flex items-center">
                <div class="w-2 h-2 bg-gray-300 rounded-full mr-2"></div>
                <span class="text-sm font-medium">Portfolio Website</span>
              </div>
            </div>
            <p class="text-xs text-gray-500 mt-1 truncate">
              Maria: Great job on the landing page!
            </p>
          </div>

          <!-- More Chat Rooms -->
          <div class="chat-room p-2 rounded">
            <div class="flex justify-between items-center">
              <div class="flex items-center">
                <div class="w-2 h-2 bg-green-500 rounded-full mr-2"></div>
                <span class="text-sm font-medium">Mobile App Project</span>
              </div>
            </div>
            <p class="text-xs text-gray-500 mt-1 truncate">
              John: Let's review the wireframes tomorrow
            </p>
          </div>

          <div class="chat-room p-2 rounded">
            <div class="flex justify-between items-center">
              <div class="flex items-center">
                <div class="w-2 h-2 bg-green-500 rounded-full mr-2"></div>
                <span class="text-sm font-medium">AI Research Team</span>
              </div>
              <span class="unread-badge">1</span>
            </div>
            <p class="text-xs text-gray-500 mt-1 truncate">
              Emma: Check out this new paper I found
            </p>
          </div>

          <div class="chat-room p-2 rounded">
            <div class="flex justify-between items-center">
              <div class="flex items-center">
                <div class="w-2 h-2 bg-gray-300 rounded-full mr-2"></div>
                <span class="text-sm font-medium">Database Design</span>
              </div>
            </div>
            <p class="text-xs text-gray-500 mt-1 truncate">
              Robert: I've updated the schema diagram
            </p>
          </div>

          <div class="chat-room p-2 rounded">
            <div class="flex justify-between items-center">
              <div class="flex items-center">
                <div class="w-2 h-2 bg-gray-300 rounded-full mr-2"></div>
                <span class="text-sm font-medium">UI/UX Workshop</span>
              </div>
            </div>
            <p class="text-xs text-gray-500 mt-1 truncate">
              Sophie: Materials for tomorrow's session are ready
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Chat Window -->
    <!-- <div class="card p-4 flex-1"> -->
    <div class="card p-4 flex-1 flex flex-col">
      <!-- <div class="h-96 overflow-y-auto bg-gray-100 p-3 rounded mb-4"> -->
      <div class="flex-1 overflow-y-auto bg-gray-100 p-3 rounded mb-4">
        <!-- Message Bubbles -->
        <div class="flex items-start mb-3">
          <img
            class="w-8 h-8 rounded-full mr-2"
            src="https://avatars.githubusercontent.com/u/789012?v=4"
            alt="Teammate Avatar"
          />
          <div>
            <p class="text-sm font-medium">Jane (Teammate)</p>
            <p class="text-sm bg-white p-2 rounded">
              Hey, how's the backend going?
            </p>
            <p class="text-xs text-gray-500 mt-1">18/03/2025 10:30</p>
          </div>
        </div>
        <div class="flex justify-end items-start mb-3">
          <div class="text-right">
            <p class="text-sm font-medium">You</p>
            <p
              class="text-sm bg-[#6b48ff] text-white p-2 rounded"
              style="background-color: #6b48ff"
            >
              Good, almost done!
            </p>
            <p class="text-xs text-gray-500 mt-1">18/03/2025 10:32</p>
          </div>
          <img
            class="w-8 h-8 rounded-full ml-2"
            src="https://avatars.githubusercontent.com/u/123456?v=4"
            alt="User Avatar"
          />
        </div>
        <div class="flex items-start mb-3">
          <img
            class="w-8 h-8 rounded-full mr-2"
            src="https://avatars.githubusercontent.com/u/345678?v=4"
            alt="Business Avatar"
          />
          <div>
            <p class="text-sm font-medium">TechCorp (Business)</p>
            <img
              class="w-32 rounded"
              src="https://via.placeholder.com/150"
              alt="Uploaded Image"
            />
            <p class="text-xs text-gray-500 mt-1">18/03/2025 10:35</p>
          </div>
        </div>
      </div>
      <!-- Input Area -->
      <div class="flex space-x-2">
        <input
          type="text"
          class="flex-1 p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
          placeholder="Type a message..."
        />
        <button class="btn">Send</button>
        <label class="btn-secondary flex items-center cursor-pointer">
          <svg
            class="w-5 h-5 mr-1"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 16v4h4l12-12-4-4L4 16zm12-8l-4-4"
            ></path>
          </svg>
          <input type="file" class="hidden" accept="image/*" />
        </label>
      </div>
      <button
        class="btn-secondary mt-2 w-full flex items-center justify-center"
      >
        <svg
          class="w-5 h-5 mr-1"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M15 10l4.5-4.5M9 10l-4.5-4.5M12 3v3m0 12v3m-9-9h3m12 0h3"
          ></path>
        </svg>
        Start Video Call
      </button>
    </div>

    <!-- Team Members -->
    <div class="card p-4 w-64 flex flex-col">
      <h3 class="text-base font-semibold mb-3">Team Members</h3>
      <div class="space-y-3">
        <div class="flex items-center">
          <img
            class="w-8 h-8 rounded-full mr-2"
            src="https://avatars.githubusercontent.com/u/123456?v=4"
            alt="User Avatar"
          />
          <div>
            <p class="text-sm font-medium">John Doe</p>
            <p class="text-xs text-gray-500">You | Student</p>
          </div>
        </div>
        <div class="flex items-center">
          <img
            class="w-8 h-8 rounded-full mr-2"
            src="https://avatars.githubusercontent.com/u/789012?v=4"
            alt="Teammate Avatar"
          />
          <div>
            <p class="text-sm font-medium">Jane</p>
            <p class="text-xs text-gray-500">Teammate | Student</p>
          </div>
        </div>
        <div class="flex items-center">
          <img
            class="w-8 h-8 rounded-full mr-2"
            src="https://avatars.githubusercontent.com/u/345678?v=4"
            alt="Business Avatar"
          />
          <div>
            <p class="text-sm font-medium">TechCorp</p>
            <p class="text-xs text-gray-500">Business</p>
          </div>
        </div>
      </div>

      <hr class="my-3 border-gray-200" />

      <!-- Project Files -->
      <h3 class="text-base font-semibold mb-3">Shared Files</h3>
      <div class="space-y-2 flex-1 overflow-y-auto">
        <div class="p-2 bg-gray-50 rounded text-sm flex items-center">
          <svg
            class="w-4 h-4 text-blue-500 mr-2"
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
          <span class="truncate">Project_spec.pdf</span>
        </div>
        <div class="p-2 bg-gray-50 rounded text-sm flex items-center">
          <svg
            class="w-4 h-4 text-green-500 mr-2"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
            ></path>
          </svg>
          <span class="truncate">budget.xlsx</span>
        </div>
        <div class="p-2 bg-gray-50 rounded text-sm flex items-center">
          <svg
            class="w-4 h-4 text-[#6b48ff] mr-2"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
            ></path>
          </svg>
          <span class="truncate">mockup.png</span>
        </div>
      </div>
    </div>
  </div>
</main>
