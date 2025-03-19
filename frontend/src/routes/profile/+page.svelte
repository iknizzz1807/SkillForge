<script lang="ts">
  let editProfileModalOpen: boolean = $state(false);
  let generatePortfolioModalOpen: boolean = $state(false);

  let porfolioType: string = $state("");

  const toggleEditProfileModal = () => {
    editProfileModalOpen = !editProfileModalOpen;
  };

  const toggleGeneratePorfolioModal = () => {
    generatePortfolioModalOpen = !generatePortfolioModalOpen;
  };
</script>

{#if editProfileModalOpen}
  <div id="editProfileModal" class="modal">
    <div class="bg-white rounded-lg w-1/2 max-w-2xl shadow-lg">
      <div class="p-4 border-b">
        <h3 class="text-lg font-semibold">Edit Profile</h3>
      </div>
      <div class="p-4">
        <form>
          <div class="grid grid-cols-2 gap-4 mb-4">
            <div>
              <label class="block text-sm font-medium mb-1" for="firstName"
                >First Name</label
              >
              <input
                type="text"
                id="firstName"
                value="John"
                class="w-full p-2 border border-gray-300 rounded"
              />
            </div>
            <div>
              <label class="block text-sm font-medium mb-1" for="lastName"
                >Last Name</label
              >
              <input
                type="text"
                id="lastName"
                value="Doe"
                class="w-full p-2 border border-gray-300 rounded"
              />
            </div>
          </div>

          <div class="mb-4">
            <label class="block text-sm font-medium mb-1" for="email"
              >Email</label
            >
            <input
              type="email"
              id="email"
              value="john.doe@example.com"
              class="w-full p-2 border border-gray-300 rounded"
            />
          </div>

          <div class="mb-4">
            <label class="block text-sm font-medium mb-1" for="title"
              >Title</label
            >
            <input
              type="text"
              id="title"
              value="Student | Web Developer"
              class="w-full p-2 border border-gray-300 rounded"
            />
          </div>

          <div class="mb-4">
            <label class="block text-sm font-medium mb-1" for="bio">Bio</label>
            <textarea
              id="bio"
              rows="3"
              class="w-full p-2 border border-gray-300 rounded"
            >
              Passionate web developer with experience in modern JavaScript
              frameworks.</textarea
            >
          </div>

          <div class="mb-4">
            <label class="block text-sm font-medium mb-1" for="location"
              >Location</label
            >
            <input
              type="text"
              id="location"
              value="San Francisco, CA"
              class="w-full p-2 border border-gray-300 rounded"
            />
          </div>

          <div class="mb-4">
            <label class="block text-sm font-medium mb-1">Profile Picture</label
            >
            <div class="flex items-center space-x-4">
              <img
                src="https://avatars.githubusercontent.com/u/123456?v=4"
                alt="Current avatar"
                class="w-16 h-16 rounded-full"
              />
              <button type="button" class="btn-secondary text-sm">
                Change Photo
              </button>
            </div>
          </div>
        </form>
      </div>
      <div class="p-4 border-t flex justify-end space-x-2">
        <button
          id="cancelEditBtn"
          class="btn-secondary"
          onclick={toggleEditProfileModal}>Cancel</button
        >
        <button class="btn">Save Changes</button>
      </div>
    </div>
  </div>
{/if}

{#if generatePortfolioModalOpen}
  <!-- Portfolio Generation Modal -->
  <div id="portfolioModal" class="modal hidden">
    <div class="bg-white rounded-lg w-1/2 max-w-2xl shadow-lg text-center">
      <div class="p-6">
        <svg
          class="w-16 h-16 text-[#6b48ff] mx-auto mb-4"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
          ></path>
        </svg>
        <h3 class="text-xl font-semibold mb-2">Portfolio Generated!</h3>
        <p class="text-gray-600 mb-4">
          Your portfolio website has been generated successfully.
        </p>
        <div class="bg-gray-100 p-4 rounded mb-4 text-left">
          <p class="text-sm font-mono break-all">
            https://skillforge.example.com/portfolio/johndoe
          </p>
        </div>
        <div class="flex space-x-2 justify-center">
          <button
            id="closePortfolioBtn"
            class="btn-secondary"
            onclick={toggleGeneratePorfolioModal}>Close</button
          >
          <button class="btn flex items-center">
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
                d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
              ></path>
            </svg>
            Visit Site
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}

<main class="flex-1 pr-4 pl-4 ml-56">
  <!-- Two Column Layout -->
  <div class="grid grid-cols-2 gap-4">
    <!-- Left Column -->
    <div class="space-y-4">
      <!-- Profile Card -->
      <div class="card p-4">
        <div class="flex items-center space-x-4">
          <img
            class="w-16 h-16 rounded-full"
            src="https://avatars.githubusercontent.com/u/123456?v=4"
            alt="User Avatar"
          />
          <div class="flex-1">
            <h3 class="text-lg font-semibold">John Doe</h3>
            <p class="text-sm text-gray-500">Student | Web Developer</p>
            <p class="text-sm">Email: john.doe@example.com</p>
          </div>
          <button
            id="editProfileBtn"
            class="btn-secondary"
            onclick={toggleEditProfileModal}
          >
            Edit Profile
          </button>
        </div>
      </div>

      <!-- Detailed Skills with Progress -->
      <div class="card p-4">
        <div class="flex justify-between items-center mb-3">
          <h3 class="text-base font-semibold">Skills & Expertise</h3>
          <button class="text-xs text-[#6b48ff]">+ Add Skill</button>
        </div>

        <div class="space-y-4">
          <div>
            <div class="flex justify-between mb-1">
              <span class="text-sm font-medium">React</span>
              <span class="text-xs text-[#6b48ff]">Advanced</span>
            </div>
            <div class="progress-bar">
              <div class="progress-fill" style="width: 85%"></div>
            </div>
          </div>

          <div>
            <div class="flex justify-between mb-1">
              <span class="text-sm font-medium">Node.js</span>
              <span class="text-xs text-[#6b48ff]">Intermediate</span>
            </div>
            <div class="progress-bar">
              <div class="progress-fill" style="width: 65%"></div>
            </div>
          </div>

          <div>
            <div class="flex justify-between mb-1">
              <span class="text-sm font-medium">JavaScript</span>
              <span class="text-xs text-[#6b48ff]">Advanced</span>
            </div>
            <div class="progress-bar">
              <div class="progress-fill" style="width: 90%"></div>
            </div>
          </div>

          <div>
            <div class="flex justify-between mb-1">
              <span class="text-sm font-medium">MongoDB</span>
              <span class="text-xs text-[#6b48ff]">Intermediate</span>
            </div>
            <div class="progress-bar">
              <div class="progress-fill" style="width: 60%"></div>
            </div>
          </div>

          <div>
            <div class="flex justify-between mb-1">
              <span class="text-sm font-medium">Python</span>
              <span class="text-xs text-[#6b48ff]">Beginner</span>
            </div>
            <div class="progress-bar">
              <div class="progress-fill" style="width: 30%"></div>
            </div>
          </div>

          <div>
            <div class="flex justify-between mb-1">
              <span class="text-sm font-medium">UI/UX Design</span>
              <span class="text-xs text-[#6b48ff]">Intermediate</span>
            </div>
            <div class="progress-bar">
              <div class="progress-fill" style="width: 70%"></div>
            </div>
          </div>
        </div>
      </div>

      <!-- Portfolio Generator -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">Portfolio Generator</h3>
        <p class="text-sm text-gray-600 mb-3">
          Create a professional portfolio website in seconds based on your
          profile data.
        </p>
        <div class="grid grid-cols-3 gap-2 mb-3">
          <button onclick={() => (porfolioType = "minimal")}>
            <div
              class={`border  rounded p-2 cursor-pointer hover:border-[#6b48ff] ${porfolioType === "minimal" ? "border-[#6b48ff]" : "border-gray-200"}`}
            >
              <div class="bg-gray-100 h-12 mb-1 rounded"></div>
              <p class="text-xs text-center">Minimal</p>
            </div>
          </button>

          <button onclick={() => (porfolioType = "modern")}>
            <div
              class={`border  rounded p-2 cursor-pointer hover:border-[#6b48ff] ${porfolioType === "modern" ? "border-[#6b48ff]" : "border-gray-200"}`}
            >
              <div class="bg-gray-100 h-12 mb-1 rounded"></div>
              <p class="text-xs text-center">Modern</p>
            </div>
          </button>

          <button onclick={() => (porfolioType = "creative")}>
            <div
              class={`border rounded p-2 cursor-pointer hover:border-[#6b48ff] ${porfolioType === "creative" ? "border-[#6b48ff]" : "border-gray-200"}`}
            >
              <div class="bg-gray-100 h-12 mb-1 rounded"></div>
              <p class="text-xs text-center">Creative</p>
            </div>
          </button>
        </div>
        <button
          id="generatePortfolioBtn"
          class="btn w-full text-sm"
          onclick={toggleGeneratePorfolioModal}
        >
          Generate Portfolio
        </button>
      </div>
    </div>

    <!-- Right Column -->
    <div class="space-y-4">
      <!-- Project History -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">Project History</h3>
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <div>
              <p class="text-sm font-medium">Build a Chat App</p>
              <p class="text-xs text-gray-500">
                Skills: React, Node.js | Completed: 18/03/2025
              </p>
            </div>
            <span class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded"
              >Completed</span
            >
          </div>
          <div class="flex justify-between items-center">
            <div>
              <p class="text-sm font-medium">Portfolio Website</p>
              <p class="text-xs text-gray-500">
                Skills: HTML, CSS, JS | Completed: 10/03/2025
              </p>
            </div>
            <span class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded"
              >Completed</span
            >
          </div>
          <div class="flex justify-between items-center">
            <div>
              <p class="text-sm font-medium">E-commerce Dashboard</p>
              <p class="text-xs text-gray-500">
                Skills: React, MongoDB | Completed: 05/03/2025
              </p>
            </div>
            <span class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded"
              >Completed</span
            >
          </div>
        </div>
      </div>

      <!-- Certificates -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">Certificates</h3>
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <div>
              <p class="text-sm font-medium">React Development</p>
              <p class="text-xs text-gray-500">
                Issued: 18/03/2025 by TechCorp
              </p>
            </div>
            <a href="#" class="text-sm text-[#6b48ff] hover:underline"
              >Download</a
            >
          </div>
          <div class="flex justify-between items-center">
            <div>
              <p class="text-sm font-medium">Web Development Basics</p>
              <p class="text-xs text-gray-500">
                Issued: 10/03/2025 by SKILLFORGE
              </p>
            </div>
            <a href="#" class="text-sm text-[#6b48ff] hover:underline"
              >Download</a
            >
          </div>
          <div class="flex justify-between items-center">
            <div>
              <p class="text-sm font-medium">MongoDB for Developers</p>
              <p class="text-xs text-gray-500">
                Issued: 05/02/2025 by MongoDB University
              </p>
            </div>
            <a href="#" class="text-sm text-[#6b48ff] hover:underline"
              >Download</a
            >
          </div>
        </div>
      </div>

      <!-- Stats Card -->
      <div class="card p-4">
        <h3 class="text-base font-semibold mb-3">Profile Stats</h3>
        <div class="grid grid-cols-2 gap-3">
          <div class="text-center p-2 bg-gray-50 rounded">
            <p class="text-xl font-bold text-[#6b48ff]">12</p>
            <p class="text-xs text-gray-500">Completed Projects</p>
          </div>
          <div class="text-center p-2 bg-gray-50 rounded">
            <p class="text-xl font-bold text-[#6b48ff]">8</p>
            <p class="text-xs text-gray-500">Certificates</p>
          </div>
          <div class="text-center p-2 bg-gray-50 rounded">
            <p class="text-xl font-bold text-[#6b48ff]">98%</p>
            <p class="text-xs text-gray-500">On-time Completion</p>
          </div>
          <div class="text-center p-2 bg-gray-50 rounded">
            <p class="text-xl font-bold text-[#6b48ff]">4.9</p>
            <p class="text-xs text-gray-500">Client Rating</p>
          </div>
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
  .skill-level {
    position: absolute;
    right: 0;
    top: 0;
    font-size: 10px;
    padding: 2px 4px;
    border-radius: 0 4px 0 4px;
    background-color: rgba(255, 255, 255, 0.3);
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
</style>
