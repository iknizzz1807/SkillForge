<script lang="ts">
  import { enhance } from "$app/forms";
  import type { ActionData, PageData } from "./$types";
  import { goto } from "$app/navigation";

  let { form, data }: { form: ActionData; data: PageData } = $props();
  let selectedRole = $state<string | null>(null); // Bắt đầu với null
  let avatarFile: File | null = $state(null);
  let avatarPreview: string | null = $state(null);
  let avatarError: string | null = $state(null);

  // State cho việc loading và success
  let isSubmitting = $state(false);
  let registerSuccess = $state(false);

  // Danh sách title theo role
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

  // Allowed file extensions (matching backend validation)
  const allowedImageTypes = [".jpg", ".jpeg", ".png", ".gif"];

  // Handle file selection with validation
  function handleFileChange(event: Event) {
    const input = event.target as HTMLInputElement;
    avatarError = null; // Clear previous errors

    if (input.files && input.files[0]) {
      const file = input.files[0];
      const fileName = file.name.toLowerCase();
      const fileExt = fileName.substring(fileName.lastIndexOf("."));

      // Validate file extension
      if (!allowedImageTypes.includes(fileExt)) {
        avatarError =
          "Invalid file type. Only JPG, JPEG, PNG, and GIF are allowed.";
        input.value = ""; // Clear the file input
        avatarFile = null;
        avatarPreview = null;
        return;
      }

      // If valid format, proceed with preview
      avatarFile = file;

      // Create a preview URL
      const reader = new FileReader();
      reader.onload = (e) => {
        avatarPreview = e.target?.result as string;
      };
      reader.readAsDataURL(avatarFile);
    }
  }

  // Handle role change
  function handleRoleChange(event: Event) {
    const input = event.target as HTMLInputElement;
    selectedRole = input.value;
  }

  $effect(() => {
    if (form?.success) {
      registerSuccess = true;
      // Wait 1.5s then redirect
      setTimeout(() => {
        goto("/login");
      }, 1500);
    }
  });
</script>

<svelte:head>
  <title>Register to SkillForge</title>
</svelte:head>

<!-- Full height container -->
<div class="min-h-screen flex items-center justify-center bg-gray-50">
  <main class="w-full max-w-md p-4 flex items-center justify-center">
    <div class="card p-6 w-full shadow-md bg-white rounded-lg">
      <h2 class="text-2xl font-semibold mb-2 text-center">
        Register to SKILLFORGE
      </h2>
      <p class="text-sm text-gray-600 text-center mb-6">
        Create an account and start your journey
      </p>

      <!-- Registration Form -->
      <form
        class="space-y-4"
        method="POST"
        use:enhance={() => {
          isSubmitting = true;
          registerSuccess = false;

          return async ({ result, update }) => {
            // Wait for the response and update the form
            await update();

            // If registration was successful
            if (result.type === "success") {
              registerSuccess = true;
            } else {
              // If there was an error, stop loading state
              isSubmitting = false;
            }
          };
        }}
        enctype="multipart/form-data"
      >
        <!-- Avatar and Name in single row -->
        <div class="flex items-center gap-3">
          <!-- Avatar Upload -->
          <div class="flex-shrink-0">
            <div class="relative">
              {#if avatarPreview}
                <img
                  src={avatarPreview}
                  alt="Avatar preview"
                  class="w-16 h-16 rounded-full object-cover border-2 border-[#6b48ff]"
                />
              {:else}
                <div
                  class="w-16 h-16 rounded-full bg-gray-200 flex items-center justify-center text-gray-500"
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
                      d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                    />
                  </svg>
                </div>
              {/if}
              <label
                class="absolute bottom-0 right-0 bg-[#6b48ff] text-white rounded-full p-1 cursor-pointer hover:bg-[#5a3bd7] shadow-md"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-3 w-3"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z"
                  />
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M15 13a3 3 0 11-6 0 3 3 0 016 0z"
                  />
                </svg>
                <input
                  type="file"
                  name="avatar"
                  accept="image/jpeg,image/jpg,image/png,image/gif"
                  class="sr-only"
                  onchange={handleFileChange}
                />
              </label>
            </div>
          </div>

          <!-- Name field - Takes up remaining space -->
          <div class="flex-grow">
            <label class="block text-sm font-medium mb-1" for="name"
              >Full Name <span class="text-red-500">*</span></label
            >
            <input
              type="text"
              id="name"
              name="name"
              required
              class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              placeholder="Your full name"
              value={form?.formData?.name || ""}
            />
          </div>
        </div>

        <!-- Display avatar error if any -->
        {#if avatarError}
          <p class="text-red-500 text-sm">{avatarError}</p>
        {/if}

        <!-- Email field -->
        <div>
          <label class="block text-sm font-medium mb-1" for="email"
            >Email <span class="text-red-500">*</span></label
          >
          <input
            type="email"
            id="email"
            name="email"
            required
            class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
            placeholder="you@example.com"
            value={form?.formData?.email || ""}
          />
        </div>

        <!-- Account Type - Already has red asterisk -->
        <div>
          <label class="block text-sm font-medium mb-1"
            >Account Type <span class="text-red-500">*</span></label
          >
          <div class="grid grid-cols-2 gap-3">
            <label class="role-option">
              <input
                type="radio"
                name="role"
                value="student"
                class="sr-only"
                checked={selectedRole === "student"}
                onchange={handleRoleChange}
                required
              />
              <div
                class="border border-gray-300 rounded p-3 hover:border-[#6b48ff] cursor-pointer flex items-center justify-center role-card"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5 mr-2"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"
                  />
                </svg>
                <span class="font-medium">Student</span>
              </div>
            </label>
            <label class="role-option">
              <input
                type="radio"
                name="role"
                value="business"
                class="sr-only"
                checked={selectedRole === "business"}
                onchange={handleRoleChange}
              />
              <div
                class="border border-gray-300 rounded p-3 hover:border-[#6b48ff] cursor-pointer flex items-center justify-center role-card"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5 mr-2"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                  />
                </svg>
                <span class="font-medium">Business</span>
              </div>
            </label>
          </div>
        </div>

        <!-- Professional Title field - Only show after role selection -->
        {#if selectedRole}
          <div>
            <label class="block text-sm font-medium mb-1" for="title"
              >Professional Title <span class="text-red-500">*</span></label
            >
            <select
              id="title"
              name="title"
              required
              class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff] appearance-none bg-white"
            >
              <option value="" disabled selected>Select your title</option>
              {#if selectedRole === "student"}
                {#each titleOptions.student as option}
                  <option value={option}>{option}</option>
                {/each}
              {:else if selectedRole === "business"}
                {#each titleOptions.business as option}
                  <option value={option}>{option}</option>
                {/each}
              {/if}
            </select>
          </div>
        {/if}

        <!-- Website Field (only for Business) - Not required -->
        {#if selectedRole === "business"}
          <div>
            <label class="block text-sm font-medium mb-1" for="website"
              >Company Website</label
            >
            <input
              type="url"
              id="website"
              name="website"
              class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              placeholder="https://yourcompany.com"
              value={form?.formData?.website || ""}
            />
          </div>
        {/if}

        <!-- Password Fields -->
        <div>
          <label class="block text-sm font-medium mb-1" for="password"
            >Password <span class="text-red-500">*</span></label
          >
          <input
            type="password"
            id="password"
            name="password"
            required
            class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
            placeholder="••••••••••••••••"
          />
        </div>

        <div>
          <label class="block text-sm font-medium mb-1" for="confirmPassword"
            >Confirm Password <span class="text-red-500">*</span></label
          >
          <input
            type="password"
            id="confirmPassword"
            name="confirmPassword"
            required
            class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
            placeholder="••••••••••••••••"
          />
        </div>

        {#if form?.error && !form.success}
          <div class="text-red-400 text-sm">{form.message}</div>
        {/if}

        {#if form?.success}
          <div class="text-green-400 text-sm">
            Registration successful, redirecting you to Login page
          </div>
        {/if}

        <button
          type="submit"
          class="btn w-full flex items-center justify-center"
          disabled={!!avatarError ||
            !selectedRole ||
            isSubmitting ||
            registerSuccess}
        >
          {#if isSubmitting && !registerSuccess}
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
          {:else if registerSuccess}
            Account Created
          {:else}
            Create Account
          {/if}
        </button>

        <p class="text-sm text-center mt-4">
          Already have an account?
          <a href="/login" class="text-[#6b48ff] hover:underline">Login</a>
        </p>
      </form>
    </div>
  </main>
</div>

<style>
  .role-option input:checked + .role-card {
    border-color: #6b48ff;
    border-width: 2px;
    background-color: rgba(107, 72, 255, 0.05);
  }

  /* Styling cho dropdown */
  select {
    background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3e%3cpath stroke='%236B7280' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='M6 8l4 4 4-4'/%3e%3c/svg%3e");
    background-position: right 0.5rem center;
    background-repeat: no-repeat;
    background-size: 1.5em 1.5em;
    padding-right: 2.5rem;
  }

  /* Card style */
  .card {
    border-radius: 0.5rem;
  }
</style>
