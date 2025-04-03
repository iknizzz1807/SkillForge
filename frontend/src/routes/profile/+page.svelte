<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();
  let editProfileModalOpen: boolean = $state(false);
  let generatePortfolioModalOpen: boolean = $state(false);
  let addJobModalOpen: boolean = $state(false);

  // Toggle this to switch between student and business profiles
  let role: "student" | "business" = $state("student");

  let porfolioType: string = $state("");

  const toggleRole = () => {
    role = role === "student" ? "business" : "student";
  };

  const toggleEditProfileModal = () => {
    editProfileModalOpen = !editProfileModalOpen;
  };

  const toggleGeneratePorfolioModal = () => {
    generatePortfolioModalOpen = !generatePortfolioModalOpen;
  };

  const toggleAddJobModal = () => {
    addJobModalOpen = !addJobModalOpen;
  };

  // Badges data for gamification
  const studentBadges = [
    {
      name: "First Project",
      icon: "🏆",
      description: "Completed your first project",
      achieved: true,
      date: "10/02/2025",
    },
    {
      name: "Quick Learner",
      icon: "🚀",
      description: "Completed 5 skills assessments",
      achieved: true,
      date: "15/02/2025",
    },
    {
      name: "Team Player",
      icon: "👥",
      description: "Collaborated with 3+ students",
      achieved: true,
      date: "01/03/2025",
    },
    {
      name: "Code Master",
      icon: "💻",
      description: "Wrote 1000+ lines of code",
      achieved: false,
      date: "",
    },
    {
      name: "Perfect Score",
      icon: "⭐",
      description: "Received 5.0 rating on a project",
      achieved: false,
      date: "",
    },
  ];

  // Calculate XP and level for gamification
  const studentXP = 1250;
  const currentLevel = Math.floor(studentXP / 300) + 1;
  const xpForNextLevel = currentLevel * 300;
  const xpProgress = ((studentXP % 300) / 300) * 100;
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
              value={role === "student"
                ? "Student | Web Developer"
                : "Business | Tech Solutions Inc."}
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
              {role === "student"
                ? "Passionate web developer with experience in modern JavaScript frameworks."
                : "Innovative tech company focused on creating cutting-edge web and mobile applications for businesses."}</textarea
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
  <div id="portfolioModal" class="modal">
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

{#if addJobModalOpen}
  <!-- Add Job Modal -->
  <div id="addJobModal" class="modal">
    <div class="bg-white rounded-lg w-1/2 max-w-2xl shadow-lg">
      <div class="p-4 border-b">
        <h3 class="text-lg font-semibold">Add New Project</h3>
      </div>
      <div class="p-4">
        <form>
          <div class="mb-4">
            <label class="block text-sm font-medium mb-1" for="projectTitle"
              >Project Title</label
            >
            <input
              type="text"
              id="projectTitle"
              placeholder="e.g. E-commerce Website Redesign"
              class="w-full p-2 border border-gray-300 rounded"
            />
          </div>

          <div class="mb-4">
            <label
              class="block text-sm font-medium mb-1"
              for="projectDescription">Project Description</label
            >
            <textarea
              id="projectDescription"
              rows="3"
              placeholder="Describe what students will work on and what they'll learn"
              class="w-full p-2 border border-gray-300 rounded"
            ></textarea>
          </div>

          <div class="grid grid-cols-2 gap-4 mb-4">
            <div>
              <label class="block text-sm font-medium mb-1" for="startDate"
                >Start Date</label
              >
              <input
                type="date"
                id="startDate"
                class="w-full p-2 border border-gray-300 rounded"
              />
            </div>
            <div>
              <label class="block text-sm font-medium mb-1" for="endDate"
                >End Date</label
              >
              <input
                type="date"
                id="endDate"
                class="w-full p-2 border border-gray-300 rounded"
              />
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4 mb-4">
            <div>
              <label class="block text-sm font-medium mb-1" for="maxStudents"
                >Max Students</label
              >
              <input
                type="number"
                id="maxStudents"
                value="4"
                min="1"
                class="w-full p-2 border border-gray-300 rounded"
              />
            </div>
            <div>
              <label class="block text-sm font-medium mb-1" for="projectBudget"
                >Budget (optional)</label
              >
              <input
                type="number"
                id="projectBudget"
                placeholder="500"
                class="w-full p-2 border border-gray-300 rounded"
              />
            </div>
          </div>

          <div class="mb-4">
            <label class="block text-sm font-medium mb-1" for="requiredSkills"
              >Required Skills</label
            >
            <select
              id="requiredSkills"
              multiple
              class="w-full p-2 border border-gray-300 rounded"
            >
              <option>JavaScript</option>
              <option>React</option>
              <option>Node.js</option>
              <option>MongoDB</option>
              <option>UI/UX Design</option>
              <option>HTML/CSS</option>
              <option>Python</option>
              <option>Django</option>
            </select>
            <p class="text-xs text-gray-500 mt-1">
              Hold Ctrl/Cmd to select multiple skills
            </p>
          </div>
        </form>
      </div>
      <div class="p-4 border-t flex justify-end space-x-2">
        <button
          id="cancelAddJobBtn"
          class="btn-secondary"
          onclick={toggleAddJobModal}>Cancel</button
        >
        <button class="btn">Create Project</button>
      </div>
    </div>
  </div>
{/if}

<main class="flex-1 pr-4 pl-4 ml-64 pt-4">
  <!-- Header with role toggle -->
  <div class="mb-4 flex justify-between items-center">
    <h1 class="text-2xl font-bold">
      {#if role === "student"}
        Student Profile
      {:else}
        Business Profile
      {/if}
    </h1>
    <button
      onclick={toggleRole}
      class="px-3 py-2 bg-purple-500 text-white rounded-lg hover:bg-purple-600 transition-colors"
    >
      Switch to {role === "student" ? "Business" : "Student"} View
    </button>
  </div>

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
            <h3 class="text-lg font-semibold">{data.name}</h3>
            <p class="text-sm text-gray-500">
              {#if role === "student"}
                {data.role} | Web Developer
              {:else}
                Business | Tech Solutions Inc.
              {/if}
            </p>
            <p class="text-sm">Email: {data.email}</p>
          </div>
          <button
            id="editProfileBtn"
            class="btn-secondary"
            onclick={toggleEditProfileModal}
          >
            Edit Profile
          </button>
        </div>

        {#if role === "student"}
          <!-- Student Level Bar -->
          <div class="mt-4 p-3 bg-gray-50 rounded-lg">
            <div class="flex justify-between items-center mb-1">
              <span class="text-sm font-medium">Level {currentLevel}</span>
              <span class="text-xs text-gray-500">{studentXP} XP</span>
            </div>
            <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
              <div
                class="bg-[#6b48ff] h-full"
                style="width: {xpProgress}%"
              ></div>
            </div>
            <div class="flex justify-between mt-1">
              <span class="text-xs text-gray-500">Current</span>
              <span class="text-xs text-gray-500"
                >{xpForNextLevel} XP for Level {currentLevel + 1}</span
              >
            </div>
          </div>
        {/if}
      </div>

      {#if role === "student"}
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

        <!-- Student Badges & Achievements -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Badges & Achievements</h3>
          <div class="grid grid-cols-3 gap-3">
            {#each studentBadges as badge}
              <div
                class="text-center p-3 rounded border {badge.achieved
                  ? 'border-[#6b48ff] bg-purple-50'
                  : 'border-gray-200 bg-gray-50 opacity-50'}"
              >
                <div class="text-2xl mb-1">{badge.icon}</div>
                <p class="text-sm font-medium">{badge.name}</p>
                <p class="text-xs text-gray-500">{badge.description}</p>
                {#if badge.achieved}
                  <p class="text-xs text-[#6b48ff] mt-1">Earned {badge.date}</p>
                {:else}
                  <p class="text-xs text-gray-400 mt-1">Locked</p>
                {/if}
              </div>
            {/each}
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
                class={`border rounded p-2 cursor-pointer hover:border-[#6b48ff] ${porfolioType === "minimal" ? "border-[#6b48ff]" : "border-gray-200"}`}
              >
                <div class="bg-gray-100 h-12 mb-1 rounded"></div>
                <p class="text-xs text-center">Minimal</p>
              </div>
            </button>

            <button onclick={() => (porfolioType = "modern")}>
              <div
                class={`border rounded p-2 cursor-pointer hover:border-[#6b48ff] ${porfolioType === "modern" ? "border-[#6b48ff]" : "border-gray-200"}`}
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
      {:else}
        <!-- Business Info -->
        <div class="card p-4">
          <div class="flex justify-between items-center mb-3">
            <h3 class="text-base font-semibold">Company Information</h3>
            <button class="text-xs text-[#6b48ff]">Edit Info</button>
          </div>

          <div class="space-y-3">
            <div>
              <p class="text-sm font-medium">Company Type</p>
              <p class="text-sm text-gray-600">Technology Solutions</p>
            </div>
            <div>
              <p class="text-sm font-medium">Founded</p>
              <p class="text-sm text-gray-600">2018</p>
            </div>
            <div>
              <p class="text-sm font-medium">Size</p>
              <p class="text-sm text-gray-600">15-50 employees</p>
            </div>
            <div>
              <p class="text-sm font-medium">Website</p>
              <a href="#" class="text-sm text-[#6b48ff] hover:underline"
                >www.techsolutions.com</a
              >
            </div>
            <div>
              <p class="text-sm font-medium">About Us</p>
              <p class="text-sm text-gray-600">
                Tech Solutions Inc. is a forward-thinking technology company
                specializing in developing innovative web and mobile
                applications. We work with businesses of all sizes to create
                custom software solutions that drive growth and efficiency.
              </p>
            </div>
          </div>
        </div>

        <!-- Business Post New Project -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Post a New Project</h3>
          <p class="text-sm text-gray-600 mb-3">
            Create a new project to find talented students to help bring your
            ideas to life.
          </p>
          <button
            id="addJobBtn"
            class="btn w-full text-sm"
            onclick={toggleAddJobModal}
          >
            Create New Project
          </button>
        </div>

        <!-- Business Engagement Stats -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Engagement Statistics</h3>
          <div class="grid grid-cols-2 gap-3">
            <div class="text-center p-2 bg-gray-50 rounded">
              <p class="text-xl font-bold text-[#6b48ff]">15</p>
              <p class="text-xs text-gray-500">Projects Posted</p>
            </div>
            <div class="text-center p-2 bg-gray-50 rounded">
              <p class="text-xl font-bold text-[#6b48ff]">42</p>
              <p class="text-xs text-gray-500">Students Engaged</p>
            </div>
            <div class="text-center p-2 bg-gray-50 rounded">
              <p class="text-xl font-bold text-[#6b48ff]">12</p>
              <p class="text-xs text-gray-500">Completed Projects</p>
            </div>
            <div class="text-center p-2 bg-gray-50 rounded">
              <p class="text-xl font-bold text-[#6b48ff]">4.8</p>
              <p class="text-xs text-gray-500">Avg. Student Rating</p>
            </div>
          </div>
        </div>
      {/if}
    </div>

    <!-- Right Column -->
    <div class="space-y-4">
      {#if role === "student"}
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
              <span
                class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded"
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
              <span
                class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded"
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
              <span
                class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded"
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

        <!-- Learning Path -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Your Learning Path</h3>
          <div class="relative pb-12">
            <!-- Timeline line -->
            <div class="absolute left-2.5 top-0 h-full w-0.5 bg-gray-200"></div>

            <!-- Timeline items -->
            <div class="relative pl-8 pb-6">
              <div class="absolute left-0 rounded-full bg-green-500 p-1">
                <div
                  class="h-4 w-4 rounded-full border-2 border-white bg-green-500"
                ></div>
              </div>
              <div class="bg-green-50 p-3 rounded-lg border border-green-100">
                <p class="text-sm font-medium">
                  Completed: Front-End Development
                </p>
                <p class="text-xs text-gray-500">
                  5 projects • 3 certificates earned
                </p>
              </div>
            </div>

            <div class="relative pl-8 pb-6">
              <div class="absolute left-0 rounded-full bg-[#6b48ff] p-1">
                <div
                  class="h-4 w-4 rounded-full border-2 border-white bg-[#6b48ff]"
                ></div>
              </div>
              <div class="bg-purple-50 p-3 rounded-lg border border-purple-100">
                <p class="text-sm font-medium">
                  In Progress: Back-End Development
                </p>
                <p class="text-xs text-gray-500">
                  2/5 projects completed • 65% complete
                </p>
              </div>
            </div>

            <div class="relative pl-8">
              <div class="absolute left-0 rounded-full bg-gray-300 p-1">
                <div
                  class="h-4 w-4 rounded-full border-2 border-white bg-gray-300"
                ></div>
              </div>
              <div class="bg-gray-50 p-3 rounded-lg border border-gray-200">
                <p class="text-sm font-medium">Next: Full-Stack Integration</p>
                <p class="text-xs text-gray-500">
                  Unlock after Back-End completion
                </p>
              </div>
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
      {:else}
        <!-- Business Active Projects -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Active Projects</h3>
          <div class="space-y-3">
            <div class="p-3 bg-gray-50 rounded border border-gray-200">
              <div class="flex justify-between mb-1">
                <p class="text-sm font-medium">Mobile Banking App</p>
                <span
                  class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded"
                  >In Progress</span
                >
              </div>
              <p class="text-xs text-gray-500 mb-2">
                Skills: React Native, Node.js | Due: 25/05/2025
              </p>
              <div class="flex justify-between items-center text-xs">
                <div>Students: <span class="font-medium">4/5</span></div>
                <div>Budget: <span class="font-medium">$750</span></div>
                <div>Progress: <span class="font-medium">65%</span></div>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-2">
                <div
                  class="bg-blue-500 h-1.5 rounded-full"
                  style="width: 65%"
                ></div>
              </div>
            </div>

            <div class="p-3 bg-gray-50 rounded border border-gray-200">
              <div class="flex justify-between mb-1">
                <p class="text-sm font-medium">E-commerce Dashboard</p>
                <span
                  class="text-xs bg-yellow-100 text-yellow-800 px-2 py-1 rounded"
                  >Just Started</span
                >
              </div>
              <p class="text-xs text-gray-500 mb-2">
                Skills: React, MongoDB, Express | Due: 10/06/2025
              </p>
              <div class="flex justify-between items-center text-xs">
                <div>Students: <span class="font-medium">3/4</span></div>
                <div>Budget: <span class="font-medium">$600</span></div>
                <div>Progress: <span class="font-medium">15%</span></div>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-2">
                <div
                  class="bg-yellow-500 h-1.5 rounded-full"
                  style="width: 15%"
                ></div>
              </div>
            </div>

            <div class="p-3 bg-gray-50 rounded border border-gray-200">
              <div class="flex justify-between mb-1">
                <p class="text-sm font-medium">AI Content Generator</p>
                <span
                  class="text-xs bg-purple-100 text-purple-800 px-2 py-1 rounded"
                  >Recruiting</span
                >
              </div>
              <p class="text-xs text-gray-500 mb-2">
                Skills: Python, ML/AI, JavaScript | Due: 30/07/2025
              </p>
              <div class="flex justify-between items-center text-xs">
                <div>Students: <span class="font-medium">2/6</span></div>
                <div>Budget: <span class="font-medium">$1,200</span></div>
                <div>Progress: <span class="font-medium">0%</span></div>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-1.5 mt-2">
                <div
                  class="bg-purple-500 h-1.5 rounded-full"
                  style="width: 0%"
                ></div>
              </div>
            </div>
          </div>
        </div>

        <!-- Business Top Talent Pool -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Top Talent Pool</h3>
          <div class="space-y-3">
            <div
              class="flex items-center p-2 bg-gray-50 rounded border border-gray-200"
            >
              <img
                src="https://randomuser.me/api/portraits/women/44.jpg"
                alt="Student"
                class="w-10 h-10 rounded-full mr-3"
              />
              <div class="flex-1">
                <p class="text-sm font-medium">Sarah Johnson</p>
                <p class="text-xs text-gray-500">React, Node.js • 4.9/5 ⭐</p>
              </div>
              <div>
                <button
                  class="text-xs bg-[#6b48ff] text-white px-2 py-1 rounded"
                  >Invite</button
                >
              </div>
            </div>

            <div
              class="flex items-center p-2 bg-gray-50 rounded border border-gray-200"
            >
              <img
                src="https://randomuser.me/api/portraits/men/32.jpg"
                alt="Student"
                class="w-10 h-10 rounded-full mr-3"
              />
              <div class="flex-1">
                <p class="text-sm font-medium">David Chen</p>
                <p class="text-xs text-gray-500">Python, ML/AI • 4.8/5 ⭐</p>
              </div>
              <div>
                <button
                  class="text-xs bg-[#6b48ff] text-white px-2 py-1 rounded"
                  >Invite</button
                >
              </div>
            </div>

            <div
              class="flex items-center p-2 bg-gray-50 rounded border border-gray-200"
            >
              <img
                src="https://randomuser.me/api/portraits/women/68.jpg"
                alt="Student"
                class="w-10 h-10 rounded-full mr-3"
              />
              <div class="flex-1">
                <p class="text-sm font-medium">Maria Garcia</p>
                <p class="text-xs text-gray-500">UI/UX, React • 4.7/5 ⭐</p>
              </div>
              <div>
                <button
                  class="text-xs bg-[#6b48ff] text-white px-2 py-1 rounded"
                  >Invite</button
                >
              </div>
            </div>
          </div>
          <div class="text-center mt-3">
            <a href="#" class="text-sm text-[#6b48ff] hover:underline"
              >View All Talent</a
            >
          </div>
        </div>

        <!-- Business Completed Projects -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Completed Projects</h3>
          <div class="space-y-3">
            <div
              class="flex justify-between items-center p-2 bg-gray-50 rounded border border-gray-200"
            >
              <div>
                <p class="text-sm font-medium">Company Website Redesign</p>
                <p class="text-xs text-gray-500">
                  Completed: 10/03/2025 • 5 students
                </p>
              </div>
              <div class="flex items-center">
                <span
                  class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded mr-2"
                  >Completed</span
                >
                <a href="#" class="text-xs text-[#6b48ff] hover:underline"
                  >View</a
                >
              </div>
            </div>

            <div
              class="flex justify-between items-center p-2 bg-gray-50 rounded border border-gray-200"
            >
              <div>
                <p class="text-sm font-medium">CRM Integration</p>
                <p class="text-xs text-gray-500">
                  Completed: 25/02/2025 • 3 students
                </p>
              </div>
              <div class="flex items-center">
                <span
                  class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded mr-2"
                  >Completed</span
                >
                <a href="#" class="text-xs text-[#6b48ff] hover:underline"
                  >View</a
                >
              </div>
            </div>

            <div
              class="flex justify-between items-center p-2 bg-gray-50 rounded border border-gray-200"
            >
              <div>
                <p class="text-sm font-medium">Payment Gateway API</p>
                <p class="text-xs text-gray-500">
                  Completed: 15/01/2025 • 4 students
                </p>
              </div>
              <div class="flex items-center">
                <span
                  class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded mr-2"
                  >Completed</span
                >
                <a href="#" class="text-xs text-[#6b48ff] hover:underline"
                  >View</a
                >
              </div>
            </div>
          </div>
          <div class="text-center mt-3">
            <a href="#" class="text-sm text-[#6b48ff] hover:underline"
              >View All Projects</a
            >
          </div>
        </div>

        <!-- Business Ratings & Reviews -->
        <div class="card p-4">
          <h3 class="text-base font-semibold mb-3">Student Feedback</h3>
          <div class="flex justify-center items-center py-3">
            <div class="text-center mr-6">
              <p class="text-3xl font-bold text-[#6b48ff]">4.8</p>
              <div class="flex text-yellow-400 text-lg">
                <span>★</span><span>★</span><span>★</span><span>★</span><span
                  >★</span
                >
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
              <div class="flex items-center text-xs">
                <span class="w-12">2 ★</span>
                <div class="w-32 bg-gray-200 rounded-full h-2 mx-2">
                  <div
                    class="bg-yellow-400 h-2 rounded-full"
                    style="width: 0%"
                  ></div>
                </div>
                <span>0%</span>
              </div>
              <div class="flex items-center text-xs">
                <span class="w-12">1 ★</span>
                <div class="w-32 bg-gray-200 rounded-full h-2 mx-2">
                  <div
                    class="bg-yellow-400 h-2 rounded-full"
                    style="width: 0%"
                  ></div>
                </div>
                <span>0%</span>
              </div>
            </div>
          </div>
        </div>
      {/if}
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
