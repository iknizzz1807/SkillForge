<script lang="ts">
</script>

<main class="flex-1 flex flex-col ml-64 pr-4 pl-4 h-[calc(100vh-20px)] pt-4">
  <div class="flex space-x-4 flex-1">
    <!-- Chat Rooms List -->
    <div class="card p-3 w-64 flex flex-col">
      <div class="flex items-center justify-between mb-3">
        <h3 class="text-base font-semibold">Rooms</h3>
        <button class="text-[#6b48ff] hover:text-[#5a3de6]"> </button>
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
          <div
            class="chat-room active p-2 rounded cursor-pointer hover:bg-gray-100 min-h-[48px] flex items-center"
          >
            <span class="text-sm font-medium">Build a Chat App</span>
          </div>

          <!-- Other Chat Rooms -->
          <div
            class="chat-room p-2 rounded cursor-pointer hover:bg-gray-100 min-h-[48px] flex items-center"
          >
            <span class="text-sm font-medium">E-commerce Dashboard</span>
          </div>

          <div
            class="chat-room p-2 rounded cursor-pointer hover:bg-gray-100 min-h-[48px] flex items-center"
          >
            <span class="text-sm font-medium">Portfolio Website</span>
          </div>

          <div
            class="chat-room p-2 rounded cursor-pointer hover:bg-gray-100 min-h-[48px] flex items-center"
          >
            <span class="text-sm font-medium">Mobile App Project</span>
          </div>

          <div
            class="chat-room p-2 rounded cursor-pointer hover:bg-gray-100 min-h-[48px] flex items-center"
          >
            <span class="text-sm font-medium">AI Research Team</span>
          </div>

          <div
            class="chat-room p-2 rounded cursor-pointer hover:bg-gray-100 min-h-[48px] flex items-center"
          >
            <span class="text-sm font-medium">Database Design</span>
          </div>

          <div
            class="chat-room p-2 rounded cursor-pointer hover:bg-gray-100 min-h-[48px] flex items-center"
          >
            <span class="text-sm font-medium">UI/UX Workshop</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Chat Window -->
    <div class="card p-4 flex-1 flex flex-col">
      <!-- No room selected state -->
      <div
        class="flex-1 flex items-center justify-center bg-gray-100 p-3 rounded mb-4"
      >
        <div class="text-center p-6">
          <svg
            class="w-16 h-16 mx-auto text-gray-400 mb-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"
            ></path>
          </svg>
          <h3 class="text-lg font-medium text-gray-700 mb-2">
            No conversation selected
          </h3>
          <p class="text-gray-500 mb-4">
            Choose a room from the list to start chatting
          </p>
        </div>
      </div>

      <!-- Message Input Area (disabled when no room selected) -->
      <div class="flex space-x-2">
        <input
          type="text"
          class="flex-1 p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff] disabled:bg-gray-100 disabled:cursor-not-allowed"
          placeholder="Select a room to start messaging..."
          disabled
        />
        <button class="btn opacity-50 cursor-not-allowed" disabled>Send</button>
        <label
          class="btn-secondary flex items-center cursor-not-allowed opacity-50"
        >
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
              d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13"
            ></path>
          </svg>
          <input type="file" class="hidden" accept="image/*" disabled />
        </label>
      </div>
      <button
        class="btn-secondary mt-2 w-full flex items-center justify-center opacity-50 cursor-not-allowed"
        disabled
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
            d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"
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

<style>
  .chat-room.active {
    background-color: rgba(107, 72, 255, 0.1);
    border-left: 3px solid #6b48ff;
  }

  .chat-room.active:hover {
    background-color: rgba(107, 72, 255, 0.15);
  }

  .btn {
    background-color: #6b48ff;
    color: white;
    padding: 0.5rem 1rem;
    border-radius: 0.375rem;
    font-weight: 500;
    transition-property: background-color, border-color, color, fill, stroke;
    transition-duration: 150ms;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }

  .btn:hover {
    background-color: #5a3cd9;
  }

  .btn-secondary {
    background-color: white;
    color: #4b5563;
    padding: 0.5rem 1rem;
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

  .card {
    background-color: white;
    border-radius: 0.5rem;
    box-shadow:
      0 1px 3px 0 rgba(0, 0, 0, 0.1),
      0 1px 2px 0 rgba(0, 0, 0, 0.06);
  }
</style>
