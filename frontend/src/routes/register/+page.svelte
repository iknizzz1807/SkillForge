<script lang="ts">
  import { enhance } from "$app/forms";
  import type { ActionData, PageData } from "./$types";
  import { goto } from "$app/navigation";

  let { form, data }: { form: ActionData; data: PageData } = $props();
  const role = data.role;
  let selectedRole = $state(role);
  let avatarFile: File | null = $state(null);
  let avatarPreview: string | null = $state(null);

  // Handle file selection
  function handleFileChange(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files[0]) {
      avatarFile = input.files[0];

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
      // Wait 1s then redirect
      setTimeout(() => {
        goto("/login");
      }, 1500);
    }
  });
</script>

<svelte:head>
  <title>Register to SkillForge</title>
</svelte:head>

<div class="min-h-screen flex">
  <main class="flex-1 p-4 flex items-center justify-center">
    <div class="card p-6 w-full max-w-md">
      <h2 class="text-2xl font-semibold mb-2 text-center">
        Register to SKILLFORGE
      </h2>
      <p class="text-sm text-gray-600 text-center mb-6">
        Create your account and start connecting
      </p>

      <!-- Registration Form -->
      <form
        class="space-y-4"
        method="POST"
        use:enhance
        enctype="multipart/form-data"
      >
        <!-- Avatar Upload -->
        <div class="flex flex-col items-center mb-4">
          <div class="relative mb-2">
            {#if avatarPreview}
              <img
                src={avatarPreview}
                alt="Avatar preview"
                class="w-24 h-24 rounded-full object-cover border-2 border-[#6b48ff]"
              />
            {:else}
              <div
                class="w-24 h-24 rounded-full bg-gray-200 flex items-center justify-center text-gray-500"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-12 w-12"
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
                class="h-5 w-5"
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
                accept="image/*"
                class="sr-only"
                onchange={handleFileChange}
              />
            </label>
          </div>
          <span class="text-sm text-gray-500">Upload profile picture</span>
        </div>

        <div>
          <label class="block text-sm font-medium mb-1" for="name"
            >Full Name</label
          >
          <input
            type="text"
            id="name"
            name="name"
            required
            class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
            placeholder="Nguyen My Thong"
            value={form?.formData?.name || ""}
          />
        </div>

        <div>
          <label class="block text-sm font-medium mb-1" for="email">Email</label
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

        <div>
          <label class="block text-sm font-medium mb-1">Account Type</label>
          <div class="grid grid-cols-2 gap-3">
            <label class="role-option">
              <input
                type="radio"
                name="role"
                value="student"
                class="sr-only"
                checked={selectedRole === "student"}
                onchange={handleRoleChange}
              />
              <div
                class="border border-gray-300 rounded p-3 hover:border-[#6b48ff] cursor-pointer flex flex-col items-center role-card"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-6 w-6 mb-1"
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
                <span class="text-xs text-gray-500">Learn and collaborate</span>
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
                class="border border-gray-300 rounded p-3 hover:border-[#6b48ff] cursor-pointer flex flex-col items-center role-card"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-6 w-6 mb-1"
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
                <span class="text-xs text-gray-500">Recruit and mentor</span>
              </div>
            </label>
          </div>
        </div>

        <!-- Website Field (only for Business) -->
        {#if selectedRole === "business"}
          <div class="pt-2">
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

        <div>
          <label class="block text-sm font-medium mb-1" for="password"
            >Password</label
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
            >Confirm Password</label
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
          <div class="text-red-400 text-1xl">{form.message}</div>
        {/if}

        {#if form?.success}
          <div class="text-green-400 text-1xl">
            Registration successful, redirecting you to Login page
          </div>
        {/if}

        <button type="submit" class="btn w-full">Register</button>
      </form>

      <p class="text-sm text-center mt-4">
        Already have an account?
        <a href="/login" class="text-[#6b48ff] hover:underline">Login</a>
      </p>
    </div>
  </main>
</div>

<style>
  .role-option input:checked + .role-card {
    border-color: #6b48ff;
    border-width: 2px;
    background-color: rgba(107, 72, 255, 0.05);
  }
</style>
