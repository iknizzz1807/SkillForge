<script lang="ts">
  import { enhance } from "$app/forms";
  import type { ActionData, PageData } from "./$types";

  let { form, data }: { form: ActionData; data: PageData } = $props();
</script>

<!-- <div class="login-container">
  <h1>Login to SkillForge</h1>

  {#if form?.error}
    <div class="error-message">
      {form.message}
    </div>
  {/if}

  <form method="POST" use:enhance>
    <input
      type="hidden"
      name="redirectTo"
      value={form?.redirectTo || data.redirectTo}
    />

    <div class="form-group">
      <label for="email">Email</label>
      <input
        type="email"
        id="email"
        name="email"
        required
        value={form?.email || ""}
      />
    </div>

    <div class="form-group">
      <label for="password">Password</label>
      <input type="password" id="password" name="password" required />
    </div>

    <button type="submit">Log In</button>
  </form>

  <div class="signup-link">
    Don't have an account? <a href="/register">Sign up</a>
  </div>
</div> -->

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
      <form class="space-y-4" method="POST" use:enhance>
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
          />
        </div>
        <div class="flex items-center justify-between">
          <label class="flex items-center text-sm">
            <input type="checkbox" class="mr-2" />
            Remember me
          </label>
          <a
            href="/forgot-password"
            class="text-sm text-[#6b48ff] hover:underline">Forgot Password?</a
          >
        </div>
        {#if form?.error}
          <div class="text-red-400 text-1xl">{form.message}</div>
        {/if}
        {#if form?.success}
          <div class="text-green-400 text-1xl">
            {"Successfully, redirecting"}
          </div>
        {/if}
        <button type="submit" class="btn w-full">Login</button>
      </form>

      <!-- OAuth Buttons -->
      <div class="mt-6 space-y-3">
        <button class="btn-secondary w-full flex items-center justify-center">
          <svg
            class="w-5 h-5 mr-2"
            fill="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M12.48 10.92v3.28h3.36c-.14 1-.78 1.84-1.62 2.36v1.96h2.62c1.54-1.42 2.42-3.52 2.42-5.96 0-1.42-.34-2.76-.94-3.94h-5.84zM12 4.36c1.62 0 3.06.58 4.2 1.52l3.14-3.14C17.36 1.06 14.92 0 12 0 7.66 0 3.94 2.62 2.34 6.36l3.62 2.82C6.9 6.82 9.24 4.36 12 4.36zM5.64 10.92l-3.62 2.82C3.06 17.38 6.78 20 11.12 20c2.92 0 5.36-1.06 7.34-2.74v-3.28H12v-3.26h6.48c.14 1 .78 1.84 1.62 2.36v1.96c-2.28 1.96-5.36 3.06-8.98 3.06-4.34 0-8.06-2.62-9.66-6.36h3.18z"
            />
          </svg>
          Login with Google
        </button>
        <button class="btn-secondary w-full flex items-center justify-center">
          <svg
            class="w-5 h-5 mr-2"
            fill="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M20 2H4c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-1 17H5V5h14v14zm-7-13c-1.66 0-3 1.34-3 3s1.34 3 3 3 3-1.34 3-3-1.34-3-3-3zm5 11H8v-1c0-1.1.9-2 2-2h4c1.1 0 2 .9 2 2v1z"
            />
          </svg>
          Login with LinkedIn
        </button>
      </div>

      <!-- Signup Link -->
      <p class="text-sm text-center mt-4">
        Don't have an account?
        <a href="/register" class="text-[#6b48ff] hover:underline">Sign up</a>
      </p>
    </div>
  </main>
</div>
