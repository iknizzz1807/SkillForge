<script lang="ts">
  import type { PageData } from "./$types";

  let { data }: { data: PageData } = $props();
  const userData = data.userData;
  let editProfileModalOpen: boolean = $state(false);
  let generatePortfolioModalOpen: boolean = $state(false);

  // Form data state
  let name: string = $state(data.name || "");
  let email: string = $state(data.email || "");
  let title: string = $state(data.title || "");
  let selectedFile: File | null = $state(null);
  let previewUrl: string | null = $state(null);
  let isSubmitting: boolean = $state(false);
  let updateSuccess: boolean = $state(false);
  let errorMessage: string | null = $state(null);

  let isSubmittingBusinessInfo: boolean = $state(false);
  let businessInfoUpdateSuccess: boolean = $state(false);
  let businessInfoErrorMessage: string | null = $state(null);

  // States for business info
  // Consider change the default init states to empty or something to announce business that they need to update their info
  let editBusinessInfoModalOpen: boolean = $state(false);
  let companyType: string = $state(data.businessInfo.companyType);
  let founded: string = $state(data.businessInfo.founded);
  let companySize: string = $state(data.businessInfo.companySize);
  let website: string = $state(data.businessInfo.website);
  let aboutUs: string = $state(data.businessInfo.aboutUs);

  // Toggle this to switch between student and business profiles
  let role: string | undefined = $state(data.role);
  let porfolioType: string = $state("");

  const toggleRole = () => {
    role = role === "student" ? "business" : "student";
  };

  const titleOptions = {
    student: [
      "Software Engineer",
      "AI Engineer",
      "Data Scientist",
      "Frontend Developer",
      "Backend Developer",
      "Full Stack Developer",
      "Mobile Developer",
      "DevOps Engineer",
      "Cyber Security Engineer",
      "UI/UX Designer",
      "QA Engineer",
      "Machine Learning Engineer",
      "Blockchain Developer",
      "Cloud Engineer",
      "Game Developer",
    ],
    business: [
      "IT Consultant",
      "IT Outsourcing Manager",
      "CTO",
      "CEO",
      "Project Manager",
      "Product Manager",
      "IT Director",
      "Software Architect",
      "Technical Lead",
      "Head of Engineering",
      "VP of Engineering",
      "IT Manager",
      "Digital Transformation Lead",
      "Technology Strategist",
      "Innovation Officer",
    ],
  };

  // Add alongside other toggle functions
  const toggleEditBusinessInfoModal = () => {
    // Reset form state when opening modal
    if (!editBusinessInfoModalOpen) {
      // Reset to current values
      companyType = data.businessInfo.companyType;
      founded = data.businessInfo.founded;
      companySize = data.businessInfo.companySize;
      website = data.businessInfo.website;
      aboutUs = data.businessInfo.aboutUs;

      // Reset status states
      isSubmittingBusinessInfo = false;
      businessInfoUpdateSuccess = false;
      businessInfoErrorMessage = null;
    }
    editBusinessInfoModalOpen = !editBusinessInfoModalOpen;
  };

  const toggleEditProfileModal = () => {
    // Reset form state when opening modal
    if (!editProfileModalOpen) {
      name = data.name || "";
      email = data.email || "";
      title = data.title || "";
      selectedFile = null;
      previewUrl = null;
      updateSuccess = false;
      errorMessage = null;
    }
    editProfileModalOpen = !editProfileModalOpen;
  };

  const toggleGeneratePorfolioModal = () => {
    generatePortfolioModalOpen = !generatePortfolioModalOpen;
  };

  // Handle file selection for avatar
  const handleFileChange = (event: Event) => {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files[0]) {
      selectedFile = input.files[0];

      // Create preview URL
      const reader = new FileReader();
      reader.onload = (e) => {
        previewUrl = e.target?.result as string;
      };
      reader.readAsDataURL(selectedFile);
    }
  };

  // Handle profile update submission
  const handleSubmit = async () => {
    isSubmitting = true;
    errorMessage = null;
    updateSuccess = false;

    try {
      const formData = new FormData();
      formData.append("name", name);
      formData.append("email", email);
      formData.append("title", title);

      if (selectedFile) {
        formData.append("avatar", selectedFile);
      }

      const response = await fetch("/api/profile", {
        method: "PUT",
        body: formData,
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to update profile");
      }

      const result = await response.json();

      // Update local data
      data.name = result.name;
      data.email = result.email;
      if (result.avatar_name) {
        data.avatarUrl = "/api/avatars";
      }

      updateSuccess = true;
      setTimeout(() => {
        // Close modal after successful update
        if (updateSuccess) {
          // toggleEditProfileModal();
          // invalidateAll();
          window.location.reload();
        }
      }, 1500);
    } catch (error) {
      console.error("Error updating profile:", error);
      errorMessage =
        error instanceof Error ? error.message : "An unexpected error occurred";
    } finally {
      isSubmitting = false;
    }
  };

  const handleSaveBusinessInfo = async () => {
    isSubmittingBusinessInfo = true;
    businessInfoErrorMessage = null;
    businessInfoUpdateSuccess = false;

    const formData = new FormData();
    formData.append("companyType", companyType);
    formData.append("founded", founded);
    formData.append("companySize", companySize);
    formData.append("website", website);
    formData.append("aboutUs", aboutUs);

    try {
      const response = await fetch("/api/business-info", {
        method: "PUT",
        body: formData,
      });

      const result = await response.json();

      if (response.ok) {
        // Cập nhật state thành công
        businessInfoUpdateSuccess = true;
        // Đợi một khoảng thời gian rồi reload trang
        setTimeout(() => {
          if (businessInfoUpdateSuccess) {
            window.location.reload();
          }
        }, 1000);
      } else {
        // Xử lý lỗi
        businessInfoErrorMessage =
          result.error || "Failed to update business information";
      }
    } catch (error) {
      console.error("Failed to submit form:", error);
      businessInfoErrorMessage = "An unexpected error occurred";
    } finally {
      isSubmittingBusinessInfo = false;
    }
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

{#if editBusinessInfoModalOpen}
  <div id="editBusinessInfoModal" class="modal">
    <div
      class="bg-white rounded-lg w-full max-w-3xl shadow-xl mx-4"
      style="max-height: 80vh;"
    >
      <div
        class="p-4 sticky top-0 bg-white z-10 flex justify-between items-center rounded-t-lg"
      >
        <h3 class="text-lg font-semibold">Edit Company Information</h3>
        <button
          onclick={toggleEditBusinessInfoModal}
          class="text-gray-500 hover:text-gray-700 transition-colors"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
              clip-rule="evenodd"
            />
          </svg>
        </button>
      </div>

      <div class="overflow-y-auto p-5" style="max-height: calc(80vh - 132px);">
        <form>
          <div class="mb-4">
            <label class="block text-sm font-medium mb-1" for="companyType"
              >Company Type</label
            >
            <input
              type="text"
              id="companyType"
              bind:value={companyType}
              placeholder="e.g. IT Consulting, Software Development, E-commerce"
              class="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none"
            />
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
            <div>
              <label class="block text-sm font-medium mb-1" for="founded"
                >Founded</label
              >
              <input
                type="text"
                id="founded"
                bind:value={founded}
                placeholder="e.g. 2018"
                class="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none"
              />
            </div>
            <div>
              <label class="block text-sm font-medium mb-1" for="companySize"
                >Company Size</label
              >
              <input
                type="text"
                id="companySize"
                bind:value={companySize}
                placeholder="e.g. 10-50 employees"
                class="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none"
              />
            </div>
          </div>

          <div class="mb-4">
            <label class="block text-sm font-medium mb-1" for="website"
              >Website</label
            >
            <input
              type="text"
              id="website"
              bind:value={website}
              placeholder="e.g. www.yourcompany.com"
              class="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none"
            />
          </div>

          <div class="mb-4">
            <label class="block text-sm font-medium mb-1" for="aboutUs"
              >About Us</label
            >
            <textarea
              id="aboutUs"
              bind:value={aboutUs}
              rows="4"
              placeholder="Describe your company, mission, and what sets you apart..."
              class="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none"
            ></textarea>
          </div>
        </form>
      </div>

      <!-- In your business info modal, update the Save Changes button -->
      <div
        class="p-4 border-t sticky bottom-0 bg-white z-10 flex justify-end space-x-3 rounded-b-lg"
      >
        {#if businessInfoUpdateSuccess}
          <div class="flex-1 text-green-600 flex items-center">
            <svg
              class="w-5 h-5 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M5 13l4 4L19 7"
              ></path>
            </svg>
            Business information updated successfully! Reloading the page...
          </div>
        {/if}

        {#if businessInfoErrorMessage}
          <div class="flex-1 text-red-600 flex items-center">
            <svg
              class="w-5 h-5 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
              ></path>
            </svg>
            {businessInfoErrorMessage}
          </div>
        {/if}

        <button
          class="btn-secondary px-4"
          onclick={toggleEditBusinessInfoModal}
          disabled={isSubmittingBusinessInfo}
        >
          Cancel
        </button>

        <button
          class="btn px-4"
          onclick={handleSaveBusinessInfo}
          disabled={isSubmittingBusinessInfo || businessInfoUpdateSuccess}
        >
          {#if isSubmittingBusinessInfo}
            <div class="flex items-center">
              <svg
                class="animate-spin -ml-1 mr-2 h-4 w-4 text-white"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle
                  class="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  stroke-width="4"
                ></circle>
                <path
                  class="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                ></path>
              </svg>
              Saving...
            </div>
          {:else if businessInfoUpdateSuccess}
            Saved
          {:else}
            Save Changes
          {/if}
        </button>
      </div>
    </div>
  </div>
{/if}

{#if editProfileModalOpen}
  <div id="editProfileModal" class="modal">
    <div
      class="bg-white rounded-lg w-full max-w-3xl shadow-xl mx-4"
      style="max-height: 80vh;"
    >
      <div
        class="p-4 sticky top-0 bg-white z-10 flex justify-between items-center rounded-t-lg"
      >
        <h3 class="text-lg font-semibold">Edit Profile</h3>
        <button
          onclick={toggleEditProfileModal}
          class="text-gray-500 hover:text-gray-700 transition-colors"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
              clip-rule="evenodd"
            />
          </svg>
        </button>
      </div>

      <div class="overflow-y-auto p-5" style="max-height: calc(80vh - 132px);">
        <form>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
            <div>
              <label class="block text-sm font-medium mb-1" for="name"
                >Name</label
              >
              <input
                type="text"
                id="name"
                bind:value={name}
                class="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none"
              />
            </div>
            <div>
              <label class="block text-sm font-medium mb-1" for="email"
                >Email</label
              >
              <input
                type="email"
                id="email"
                bind:value={email}
                class="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none"
              />
            </div>
          </div>

          <div class="mb-4">
            <label class="block text-sm font-medium mb-1" for="title"
              >Title
            </label>
            <select
              id="title"
              name="title"
              bind:value={title}
              required
              class="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-[#6b48ff] focus:border-transparent outline-none appearance-none bg-white"
            >
              <option value="" disabled>Select your title</option>
              {#if role === "student"}
                {#each titleOptions.student as option}
                  <option value={option}>{option}</option>
                {/each}
              {:else if role === "business"}
                {#each titleOptions.business as option}
                  <option value={option}>{option}</option>
                {/each}
              {/if}
            </select>
            <p class="text-xs text-gray-500 mt-1">
              {role === "student"
                ? "Select your professional title or area of expertise"
                : "Select your company's primary business function"}
            </p>
          </div>

          <!-- In the edit profile modal form, update the profile picture section -->
          <div class="mb-4">
            <label class="block text-sm font-medium mb-1">Profile Picture</label
            >
            <div class="flex items-center space-x-4">
              <img
                src={previewUrl || data.avatarUrl}
                alt="Current avatar"
                class="w-16 h-16 rounded-full object-cover border-2 border-gray-200"
              />
              <div class="flex flex-col space-y-2 items-center">
                <label class="btn-secondary text-sm py-1 px-3 cursor-pointer">
                  Change Photo
                  <input
                    type="file"
                    accept="image/jpeg,image/png,image/gif"
                    class="hidden"
                    onchange={handleFileChange}
                  />
                </label>
                <p class="text-xs text-gray-500">Recommended: 400x400px</p>
              </div>
            </div>
          </div>
        </form>
      </div>

      <div
        class="p-4 border-t sticky bottom-0 bg-white z-10 flex justify-end space-x-3 rounded-b-lg"
      >
        {#if updateSuccess}
          <div class="flex-1 text-green-600 flex items-center">
            <svg
              class="w-5 h-5 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M5 13l4 4L19 7"
              ></path>
            </svg>
            Profile updated successfully! Reloading the page...
          </div>
        {/if}

        {#if errorMessage}
          <div class="flex-1 text-red-600 flex items-center">
            <svg
              class="w-5 h-5 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
              ></path>
            </svg>
            {errorMessage}
          </div>
        {/if}

        <button
          class="btn-secondary px-4"
          onclick={toggleEditProfileModal}
          disabled={isSubmitting}
        >
          Cancel
        </button>

        <button
          class="btn px-4"
          onclick={handleSubmit}
          disabled={isSubmitting || updateSuccess}
        >
          {#if isSubmitting}
            <div class="flex items-center">
              <svg
                class="animate-spin -ml-1 mr-2 h-4 w-4 text-white"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle
                  class="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  stroke-width="4"
                ></circle>
                <path
                  class="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                ></path>
              </svg>
              Saving...
            </div>
          {:else if updateSuccess}
            Saved
          {:else}
            Save Changes
          {/if}
        </button>
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
    <!-- <button
      onclick={toggleRole}
      class="px-3 py-2 bg-purple-500 text-white rounded-lg hover:bg-purple-600 transition-colors"
    >
      Switch to {role === "student" ? "Business" : "Student"} View
    </button> -->
  </div>

  <!-- Two Column Layout -->
  <div class="grid grid-cols-2 gap-4">
    <!-- Left Column -->
    <div class="space-y-4">
      <!-- Profile Card -->
      <div class="card p-4">
        <div class="flex items-center space-x-4">
          <div class="relative">
            <img
              class="w-16 h-16 rounded-full object-cover border-2 border-[#6b48ff]"
              src={data.avatarUrl}
              alt="User Avatar"
            />
            <div
              class="absolute -bottom-1 -right-1 bg-white rounded-full p-0.5 shadow-sm"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4 text-green-500"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>
          </div>
          <div class="flex-1">
            <h3 class="text-lg font-semibold">{data.name}</h3>
            <p class="text-sm text-gray-500 flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4 mr-1 text-gray-400"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M6 6V5a3 3 0 013-3h2a3 3 0 013 3v1h2a2 2 0 012 2v3.57A22.952 22.952 0 0110 13a22.95 22.95 0 01-8-1.43V8a2 2 0 012-2h2zm2-1a1 1 0 011-1h2a1 1 0 011 1v1H8V5zm1 5a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1z"
                  clip-rule="evenodd"
                />
                <path
                  d="M2 13.692V16a2 2 0 002 2h12a2 2 0 002-2v-2.308A24.974 24.974 0 0110 15c-2.796 0-5.487-.46-8-1.308z"
                />
              </svg>
              {#if role === "student"}
                Student {" | " + data.title}
              {:else}
                {data.title}
              {/if}
            </p>
            <p class="text-sm flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4 mr-1 text-gray-400"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z"
                />
                <path
                  d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z"
                />
              </svg>
              {data.email}
            </p>
          </div>
          <button
            id="editProfileBtn"
            class="btn-secondary flex items-center"
            onclick={toggleEditProfileModal}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4 mr-1"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z"
              />
            </svg>
            Edit Profile
          </button>
        </div>

        {#if role === "student"}
          <!-- Student Level Bar with improved visual appeal -->
          <div class="mt-4 p-3 bg-gray-50 rounded-lg border border-gray-100">
            <div class="flex justify-between items-center mb-1">
              <div class="flex items-center">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4 mr-1 text-[#6b48ff]"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path d="M13.5 3a1.5 1.5 0 013 0v7.5a1.5 1.5 0 01-3 0V3z" />
                  <path d="M3.5 3a1.5 1.5 0 013 0v7.5a1.5 1.5 0 01-3 0V3z" />
                  <path
                    d="M3.5 10.5a1.5 1.5 0 013 0v6a1.5 1.5 0 01-3 0v-6zM13.5 10.5a1.5 1.5 0 013 0v6a1.5 1.5 0 01-3 0v-6z"
                  />
                </svg>
                <span class="text-sm font-medium">Level {currentLevel}</span>
              </div>
              <span
                class="text-xs bg-[#6b48ff] bg-opacity-10 text-white px-2 py-1 rounded-full"
                >{studentXP} XP</span
              >
            </div>
            <div class="h-2 bg-gray-200 rounded-full overflow-hidden">
              <div
                class="bg-[#6b48ff] h-full"
                style="width: {xpProgress}%"
              ></div>
            </div>
            <div class="flex justify-between mt-1 text-xs text-gray-500">
              <div class="flex items-center">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-3 w-3 mr-1"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-8.707l-3-3a1 1 0 00-1.414 0l-3 3a1 1 0 001.414 1.414L9 9.414V13a1 1 0 102 0V9.414l1.293 1.293a1 1 0 001.414-1.414z"
                    clip-rule="evenodd"
                  />
                </svg>
                Current: Level {currentLevel}
              </div>
              <div class="flex items-center">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-3 w-3 mr-1"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                    clip-rule="evenodd"
                  />
                </svg>
                Next: {xpForNextLevel} XP for Level {currentLevel + 1}
              </div>
            </div>
          </div>
        {/if}
      </div>

      {#if role === "student"}
        <!-- Detailed Skills with Progress -->
        <div class="card p-4">
          <div class="flex justify-between items-center mb-3">
            <h3 class="text-base font-semibold">Skills & Expertise</h3>
            <!-- <button class="text-xs text-[#6b48ff]">+ Add Skill</button> -->
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
          <div class="flex justify-between items-center mb-4">
            <div class="flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5 mr-2 text-[#6b48ff]"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M4 4a2 2 0 012-2h8a2 2 0 012 2v12a1 1 0 110 2h-3a1 1 0 01-1-1v-2a1 1 0 00-1-1H9a1 1 0 00-1 1v2a1 1 0 01-1 1H4a1 1 0 110-2V4zm3 1h2v2H7V5zm2 4H7v2h2V9zm2-4h2v2h-2V5zm2 4h-2v2h2V9z"
                  clip-rule="evenodd"
                />
              </svg>
              <h3 class="text-base font-semibold">Company Information</h3>
            </div>
            <button
              class="text-sm text-[#6b48ff] hover:text-[#5a3dd3] flex items-center"
              onclick={toggleEditBusinessInfoModal}
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4 mr-1"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z"
                />
              </svg>
              Edit Info
            </button>
          </div>

          <div class="space-y-3">
            <div class="flex items-center">
              <div class="bg-purple-100 p-1.5 rounded-lg mr-3">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4 text-[#6b48ff]"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M4 4a2 2 0 012-2h8a2 2 0 012 2v12a1 1 0 110 2h-3a1 1 0 01-1-1v-2a1 1 0 00-1-1H9a1 1 0 00-1 1v2a1 1 0 01-1 1H4a1 1 0 110-2V4zm3 1h2v2H7V5zm2 4H7v2h2V9zm2-4h2v2h-2V5zm2 4h-2v2h2V9z"
                    clip-rule="evenodd"
                  />
                </svg>
              </div>
              <div>
                <p class="text-sm font-medium">Company Type</p>
                <p class="text-sm text-gray-600">
                  {companyType || "Not provided"}
                </p>
              </div>
            </div>
            <div class="flex items-center">
              <div class="bg-blue-100 p-1.5 rounded-lg mr-3">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4 text-blue-600"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z"
                    clip-rule="evenodd"
                  />
                </svg>
              </div>
              <div>
                <p class="text-sm font-medium">Founded</p>
                <p class="text-sm text-gray-600">{founded || "Not provided"}</p>
              </div>
            </div>
            <div class="flex items-center">
              <div class="bg-green-100 p-1.5 rounded-lg mr-3">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4 text-green-600"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z"
                  />
                </svg>
              </div>
              <div>
                <p class="text-sm font-medium">Size</p>
                <p class="text-sm text-gray-600">
                  {companySize || "Not provided"}
                </p>
              </div>
            </div>
            <div class="flex items-center">
              <div class="bg-red-100 p-1.5 rounded-lg mr-3">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4 text-red-600"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M4.083 9h1.946c.089-1.546.383-2.97.837-4.118A6.004 6.004 0 004.083 9zM10 2a8 8 0 100 16 8 8 0 000-16zm0 2c-.076 0-.232.032-.465.262-.238.234-.497.623-.737 1.182-.389.907-.673 2.142-.766 3.556h3.936c-.093-1.414-.377-2.649-.766-3.556-.24-.56-.5-.948-.737-1.182C10.232 4.032 10.076 4 10 4zm3.971 5c-.089-1.546-.383-2.97-.837-4.118A6.004 6.004 0 0115.917 9h-1.946zm-2.003 2H8.032c.093 1.414.377 2.649.766 3.556.24.56.5.948.737 1.182.233.23.389.262.465.262.076 0 .232-.032.465-.262.238-.234.498-.623.737-1.182.389-.907.673-2.142.766-3.556zm1.166 4.118c.454-1.147.748-2.572.837-4.118h1.946a6.004 6.004 0 01-2.783 4.118zm-6.268 0C6.412 13.97 6.118 12.546 6.03 11H4.083a6.004 6.004 0 002.783 4.118z"
                    clip-rule="evenodd"
                  />
                </svg>
              </div>
              <div>
                <p class="text-sm font-medium">Website</p>
                {#if website}
                  <a
                    href={website.startsWith("http")
                      ? website
                      : `https://${website}`}
                    target="_blank"
                    class="text-sm text-[#6b48ff] hover:underline flex items-center"
                  >
                    {website}
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-3 w-3 ml-1"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        d="M11 3a1 1 0 100 2h2.586l-6.293 6.293a1 1 0 101.414 1.414L15 6.414V9a1 1 0 102 0V4a1 1 0 00-1-1h-5z"
                      />
                      <path
                        d="M5 5a2 2 0 00-2 2v8a2 2 0 002 2h8a2 2 0 002-2v-3a1 1 0 10-2 0v3H5V7h3a1 1 0 000-2H5z"
                      />
                    </svg>
                  </a>
                {:else}
                  <p class="text-sm text-gray-600">Not provided</p>
                {/if}
              </div>
            </div>
            <div class="flex items-start">
              <div class="bg-yellow-100 p-1.5 rounded-lg mr-3 mt-1">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4 text-yellow-600"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                    clip-rule="evenodd"
                  />
                </svg>
              </div>
              <div class="flex-1">
                <p class="text-sm font-medium">About Us</p>
                <p
                  class="text-sm text-gray-600 mt-1 bg-gray-50 p-2 rounded-lg border border-gray-100"
                >
                  {aboutUs || "Not provided"}
                </p>
              </div>
            </div>
          </div>
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
      {/if}
    </div>

    <!-- Right Column -->
    <div class="space-y-4">
      {#if role === "student"}
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
              </div>
              <p class="text-xs text-gray-500 mb-2">
                Skills: React Native, Node.js | Due: 25/05/2025
              </p>
              <div class="flex justify-between items-center text-xs">
                <div>Students: <span class="font-medium">4/5</span></div>
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
              </div>
              <p class="text-xs text-gray-500 mb-2">
                Skills: React, MongoDB, Express | Due: 10/06/2025
              </p>
              <div class="flex justify-between items-center text-xs">
                <div>Students: <span class="font-medium">3/4</span></div>
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
              </div>
              <p class="text-xs text-gray-500 mb-2">
                Skills: Python, ML/AI, JavaScript | Due: 30/07/2025
              </p>
              <div class="flex justify-between items-center text-xs">
                <div>Students: <span class="font-medium">2/6</span></div>
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
                  >View</button
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
                  >View</button
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
                  >View</button
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
