<script lang="ts">
  import { enhance } from "$app/forms";
  import type { ActionData, PageData } from "./$types";

  let { form, data }: { form: ActionData; data: PageData } = $props();
</script>

<div class="login-container">
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
</div>

<style>
  .login-container {
    max-width: 400px;
    margin: 0 auto;
    padding: 2rem;
  }

  .form-group {
    margin-bottom: 1rem;
  }

  label {
    display: block;
    margin-bottom: 0.5rem;
  }

  input {
    width: 100%;
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 4px;
  }

  button {
    width: 100%;
    padding: 0.75rem;
    background-color: #4a6cf7;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    margin-top: 1rem;
  }

  button:hover {
    background-color: #3a5cd7;
  }

  .error-message {
    background-color: #ffebee;
    color: #c62828;
    padding: 0.75rem;
    border-radius: 4px;
    margin-bottom: 1rem;
  }

  .signup-link {
    margin-top: 1rem;
    text-align: center;
  }
</style>
