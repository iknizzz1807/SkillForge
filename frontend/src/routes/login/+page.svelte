<script lang="ts">
  import { enhance } from "$app/forms";
  import { goto } from "$app/navigation";
  import type { ActionData, PageData } from "./$types";

  let { form, data }: { form: ActionData; data: PageData } = $props();
  let isLoggingIn = $state(false);
  let loginSuccess = $state(false);
</script>

<svelte:head>
  <title>Login to SkillForge</title>
</svelte:head>

<div class="min-h-screen flex">
  <main class="flex-1 p-4 flex items-center justify-center">
    <div class="card p-6 w-full max-w-md">
      <h2 class="text-2xl font-semibold mb-2 text-center">
        Login to SKILLFORGE
      </h2>
      <p class="text-sm text-gray-600 text-center mb-6">
        Access your account and start connecting
      </p>

      <!-- Login Form -->
      <form
        class="space-y-4"
        method="POST"
        use:enhance={() => {
          // Set loading state when form is submitted
          isLoggingIn = true;
          loginSuccess = false;

          return async ({ result, update }) => {
            // Wait for the response and update the form
            await update();

            // If login was successful
            if (result.type === "success") {
              loginSuccess = true;
            } else {
              // If there was an error, stop loading state
              isLoggingIn = false;
            }
          };
        }}
      >
        <input
          type="hidden"
          name="redirectTo"
          value={form?.redirectTo || data.redirectTo}
        />
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
            value={form?.email || ""}
            disabled={isLoggingIn || loginSuccess}
          />
        </div>
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
            disabled={isLoggingIn || loginSuccess}
          />
        </div>
        <div class="flex items-center justify-between">
          <label class="flex items-center text-sm">
            <input
              type="checkbox"
              class="mr-2"
              disabled={isLoggingIn || loginSuccess}
            />
            Remember me
          </label>
          <a href="/login" class="text-sm text-[#6b48ff] hover:underline"
            >Forgot Password?</a
          >
        </div>

        {#if loginSuccess}
          <div
            class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded relative"
            role="alert"
          >
            <div class="flex items-center">
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
              <span>Login successful! Redirecting you...</span>
            </div>
          </div>
        {:else if form?.error}
          <div class="error">Invalid credentials</div>
        {/if}

        <button
          type="submit"
          class="btn w-full flex items-center justify-center"
          disabled={isLoggingIn || loginSuccess}
        >
          {#if isLoggingIn}
            <svg
              class="animate-spin -ml-1 mr-2 h-5 w-5 text-white"
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
          {:else if loginSuccess}
            Logged In
          {:else}
            Login
          {/if}
        </button>
      </form>

      <!-- Signup Link -->
      <p class="text-sm text-center mt-4">
        Don't have an account?
        <a href="/register" class="text-[#6b48ff] hover:underline">Sign up</a>
      </p>
    </div>
  </main>
</div>
